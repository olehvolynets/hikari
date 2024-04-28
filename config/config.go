package config

import (
	"fmt"
	"io"

	"gopkg.in/yaml.v3"

	"github.com/olehvolynets/sylphy/render"
)

type Config struct {
	Events []Event
	Types  []Type
}

func LoadConfig(r io.Reader) (*Config, error) {
	var cfg Config
	dec := yaml.NewDecoder(r)

	if err := dec.Decode(&cfg); err != nil && err != io.EOF {
		return nil, err
	}

	return &cfg, nil
}

func (c *Config) CreateHandlers() []render.Handler {
	handlers := make([]render.Handler, 0, len(c.Events))

	for _, evt := range c.Events {
		attrHandlers := make([]render.Handler, 0, len(evt.Scheme))

		for _, item := range evt.Scheme {
			var attrHandler render.Handler

			if item.Literal.Literal != "" {
				attrHandler = render.NewLiteralHandler(item.Literal.Literal, item.ToColor())
			} else {
				builder := c.handlerBuilder(item.Type)
				attrHandler = builder(item.Name, false, item.ToColor())
			}

			attrHandlers = append(attrHandlers, attrHandler)
		}

		handlers = append(handlers, &EventHandler{
			AttributeHandlers: attrHandlers,
		})
	}

	return handlers
}

func (c *Config) handlerBuilder(t PropertyType) render.HandlerBuilder {
	switch t {
	case NumberType:
		return render.NewNumberHandler
	case StringType:
		return render.NewStringHandler
	case BoolType:
		return render.NewBoolHandler
	case ArrayType:
		return render.NewArrayHandler
	case MapType:
		return render.NewMapHandler
	default:
		panic(fmt.Sprint("unknown (yet) PropertyType - ", string(t)))
	}
}
