package config

import (
	"github.com/olehvolynets/sylphy/render"
)

type Event struct {
	Name    string            `yaml:"name"`
	Matcher map[string]string `yaml:"matcher"`
	Scheme  []SchemeItem      `yaml:"scheme"`
}

type SchemeItem struct {
	Property            `yaml:",inline"`
	Literal             `yaml:",inline"`
	render.DisplayProps `yaml:"display"`
}

func (evt *Event) Match(entry map[string]any) bool {
	return true
}

type EventHandler struct {
	AttributeHandlers []render.Handler
}

// func (h *EventHandler) Handle(ctx *render.Context) error {
// 	slog.Debug("EventHandler", "h", *h)
// 	for _, attrHandler := range h.AttributeHandlers {
// 		err := attrHandler.Handle(ctx)
// 		if err != nil {
// 			return err
// 		}
// 	}
//
// 	return nil
// }
