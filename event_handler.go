package hikari

import (
	"fmt"

	"github.com/fatih/color"

	"github.com/olehvolynets/hikari/config"
)

var (
	numberFormat Colorizer = color.New(color.FgHiRed)
	stringFormat Colorizer = color.New(color.FgYellow)
	boolFormat   Colorizer = color.New(color.FgBlue)
	nullFormat   Colorizer = color.New(color.FgYellow, color.Bold)
)

type Entry map[string]any

type Colorizer = *color.Color

type Handler interface {
	Render(*Context, Entry)
}
type EventHandler struct {
	Event    config.Event
	Handlers []Handler
}

func NewEventHandler(evt config.Event) *EventHandler {
	handler := EventHandler{
		Event:    evt,
		Handlers: make([]Handler, len(evt.Scheme)),
	}

	for idx, schemeItem := range evt.Scheme {
		if schemeItem.Literal == "" {
			handler.Handlers[idx] = &AttributeHandler{
				Key:       schemeItem.Name,
				Colorizer: schemeItem.ToColor(),
			}
		} else {
			handler.Handlers[idx] = &LiteralHandler{
				Literal:   schemeItem.Literal,
				Colorizer: schemeItem.ToColor(),
			}
		}
	}

	return &handler
}

func (h *EventHandler) Render(ctx *Context, val Entry) {
	if len(val) == 0 {
		return
	}

	for _, attrHandler := range h.Handlers {
		attrHandler.Render(ctx, val)
	}

	fmt.Fprintln(ctx.W)
}

var DefaultEventHandler = &EventHandler{}
