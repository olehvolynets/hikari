package sylphy

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"log/slog"

	"github.com/olehvolynets/sylphy/config"
	"github.com/olehvolynets/sylphy/render"
)

var ErrNilHandler = SylphyError{Err: errors.New("were about to pass nil as render.Handler")}

type Sylphy struct {
	sink io.Writer

	cfg *config.Config
}

func NewSylphy(out io.Writer, cfg *config.Config) (*Sylphy, error) {
	s := Sylphy{
		sink: out,
		cfg:  cfg,
	}

	return &s, nil
}

func (s *Sylphy) Start(r io.Reader) error {
	// Has to be bufio.Reader to be able to read till the end of the line.
	buffReader := bufio.NewReader(r)

	decoder := json.NewDecoder(buffReader)
	hh := s.cfg.CreateHandlers()

	for {
		var handler render.Handler
		entry := make(map[string]any)

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

			handler = &render.RawHandler{Value: line}

			// Recreate decoder after processing the line that caused
			//   the previous one to fail. It doesn't mean that the next line
			//   will be parsed successfully, but in that case it will
			//   be handled in the same maner.
			decoder = json.NewDecoder(buffReader)
		} else {
			// No decoding errors branch.
			// TODO: match event
			handler = hh[0]
		}

		if handler == nil {
			// Something definetely went wrong here.
			return ErrNilHandler
		}

		ctx := render.Context{W: s.sink, Entry: entry}
		if err := handler.Handle(&ctx); err != nil {
			slog.Error(err.Error())
		}
	}

	return nil
}

func (s *Sylphy) MatchEvent(entry map[string]any) *config.Event {
	// TODO: really match event
	return &s.cfg.Events[0]
}
