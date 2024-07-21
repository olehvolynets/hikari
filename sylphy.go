package sylphy

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/olehvolynets/sylphy/config"
	"github.com/olehvolynets/sylphy/render"
)

var ErrNilHandler = SylphyError{Err: errors.New("were about to pass nil as render.Handler")}

type Sylphy struct {
	sink io.Writer
}

func NewSylphy(out io.Writer, cfg *config.Config) (*Sylphy, error) {
	s := Sylphy{
		sink: out,
	}

	return &s, nil
}

func (app *Sylphy) Start(r io.Reader) error {
	// Has to be bufio.Reader to be able to read till the end of the line.
	buffReader := bufio.NewReader(r)

	decoder := json.NewDecoder(buffReader)

	for {
		// var handler render.Handler
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

			// handler = &render.RawHandler{Value: line}

			fmt.Fprintln(app.sink, line)
			// Recreate decoder after processing the line that caused
			//   the previous one to fail. It doesn't mean that the next line
			//   will be parsed successfully, but in that case it will
			//   be handled in the same maner.
			decoder = json.NewDecoder(buffReader)
		}
		// else {
		// No decoding errors branch.
		// TODO: match event
		// handler = app.handlers[0]
		// }

		// if handler == nil {
		// 	// Something definetely went wrong here.
		// 	return ErrNilHandler
		// }

		ctx := render.Context{W: app.sink, Entry: entry, IndentChar: "\t"}
		// if err := handler.Handle(&ctx); err != nil {
		// 	slog.Error(err.Error())
		// }
		new(render.Handler).Render(&ctx, entry)
	}

	return nil
}

// func (s *Sylphy) MatchEvent(entry map[string]any) *config.Event {
// 	// TODO: really match event
// 	return &s.cfg.Events[0]
// }

// func (s *Sylphy) createHandlers(c *config.Config) {
// 	s.handlers = make([]render.Handler, 0, len(c.Events))
//
// 	for _, evt := range c.Events {
// 		attrHandlers := make([]render.Handler, 0, len(evt.Scheme))
//
// 		for _, item := range evt.Scheme {
// 			var attrHandler render.Handler
//
// 			if item.Literal.Literal != "" {
// 				attrHandler = render.NewLiteralHandler(item.Literal.Literal, item.ToColor())
// 			} else {
// 				builder := s.handlerBuilder(item.Type)
// 				attrHandler = builder(item.Name, false, item.ToColor())
// 			}
//
// 			attrHandlers = append(attrHandlers, attrHandler)
// 		}
//
// 		s.handlers = append(s.handlers, &config.EventHandler{
// 			AttributeHandlers: attrHandlers,
// 		})
// 	}
// }

// func (s *Sylphy) handlerBuilder(t config.PropertyType) render.HandlerBuilder {
// 	switch t {
// 	case config.NumberType:
// 		return render.NewNumberHandler
// 	case config.StringType:
// 		return render.NewStringHandler
// 	case config.BoolType:
// 		return render.NewBoolHandler
// 	case config.ArrayType:
// 		return render.NewArrayHandler
// 	case config.MapType:
// 		return render.NewMapHandler
// 	default:
// 		panic(fmt.Sprint("unknown (yet) PropertyType - ", string(t)))
// 	}
// }
