package hikari

import (
	"fmt"
	"reflect"
	"slices"

	"github.com/fatih/color"

	"github.com/olehvolynets/hikari/config"
)

var (
	numberFormat Colorizer = color.New(color.FgHiRed)
	stringFormat Colorizer = color.New(color.FgWhite)
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
			attrHandler := &AttributeHandler{
				Key:       schemeItem.Name,
				Skip:      schemeItem.Skip,
				Optional:  schemeItem.Optional,
				Type:      schemeItem.Type,
				Colorizer: schemeItem.ToColor(),
			}

			if schemeItem.Prefix != nil {
				attrHandler.Prefix = &Decorator{
					Literal:   schemeItem.Prefix.Literal,
					Colorizer: schemeItem.Prefix.ToColor(),
				}
			}

			if schemeItem.Postfix != nil {
				attrHandler.Postfix = &Decorator{
					Literal:   schemeItem.Postfix.Literal,
					Colorizer: schemeItem.Postfix.ToColor(),
				}
			}

			handler.Handlers[idx] = attrHandler
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

	unhandled := make(Entry, len(val)/2)
	for k, v := range val {
		if !slices.Contains(ctx.HandledAttributes, k) {
			unhandled[k] = v
		}
	}

	DefaultEventHandler.Render(ctx, unhandled)

	fmt.Fprintln(ctx.W)
}

type defaultEventHandler struct{}

func (h *defaultEventHandler) Render(ctx *Context, val Entry) {
	if len(val) == 0 {
		return
	}

	attrHandler := AttributeHandler{}
	attrHandler.render(ctx, reflect.ValueOf(val))
	fmt.Fprintln(ctx.W)
}

var DefaultEventHandler = &defaultEventHandler{}
