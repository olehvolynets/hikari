package sylphy

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"

	"github.com/olehvolynets/sylphy/config"
)

type Sylphy struct {
	sink          io.Writer
	eventHandlers []Handler
}

func NewSylphy(out io.Writer, cfg *config.Config) (*Sylphy, error) {
	s := Sylphy{
		sink:          out,
		eventHandlers: make([]Handler, len(cfg.Events)),
	}

	for idx, evt := range cfg.Events {
		s.eventHandlers[idx] = Handler{Event: evt}
	}

	return &s, nil
}

func (app *Sylphy) Start(r io.Reader) error {
	// Has to be bufio.Reader to be able to read till the end of the line.
	buffReader := bufio.NewReader(r)

	decoder := json.NewDecoder(buffReader)

	for {
		entry := make(Entry)
		var handler *Handler

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

			fmt.Fprintln(app.sink, line)
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
			ctx := Context{W: app.sink, Entry: entry, IndentChar: "\t"}

			if err := handler.Render(&ctx, entry); err != nil {
				slog.Error(err.Error())
			}
		}
	}

	return nil
}

func (app *Sylphy) MatchEvent(entry Entry) *Handler {
	for _, handler := range app.eventHandlers {
		if handler.Event.Match(entry) {
			return &handler
		}
	}

	return nil
}
