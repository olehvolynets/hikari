package hikari

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/olehvolynets/hikari/config"
)

type Hikari struct {
	sink          io.Writer
	eventHandlers []*EventHandler
}

func NewHikari(out io.Writer, cfg *config.Config) (*Hikari, error) {
	s := Hikari{
		sink:          out,
		eventHandlers: make([]*EventHandler, len(cfg.Events)),
	}

	for idx, evt := range cfg.Events {
		s.eventHandlers[idx] = NewEventHandler(evt)
	}

	return &s, nil
}

func (app *Hikari) Start(r io.Reader) error {
	// Has to be bufio.Reader to be able to read till the end of the line.
	buffReader := bufio.NewReader(r)

	decoder := json.NewDecoder(buffReader)

	for {
		entry := make(Entry)
		var handler *EventHandler

		if err := decoder.Decode(&entry); errors.Is(err, io.EOF) {
			// All readers are exhausted at this point.
			break
		} else if err != nil {
			// json.Decoder internally buffers 4kB read from the reader
			// so to process input need to aggregate remaining buffered data
			// with the primary reader.
			mr := io.MultiReader(decoder.Buffered(), buffReader)
			buffReader = bufio.NewReader(mr)

			line, err := buffReader.ReadString('\n')
			if err != nil {
				return err
			}

			fmt.Fprint(app.sink, line)
			// Recreate decoder after processing the line that caused
			//   the previous one to fail. It doesn't mean that the next line
			//   will be parsed successfully, but in that case it will
			//   be handled in the same maner.
			decoder = json.NewDecoder(buffReader)
		} else {
			// No decoding errors branch.
			handler = app.MatchEvent(entry)
		}

		if handler != nil {
			ctx := Context{W: app.sink, IndentChar: "\t"}

			handler.Render(&ctx, entry)
		}
	}

	return nil
}

func (app *Hikari) MatchEvent(entry Entry) *EventHandler {
	for _, handler := range app.eventHandlers {
		if handler.Event.Match(entry) {
			return handler
		}
	}

	return nil
}
