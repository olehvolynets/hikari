package hikari

import (
	"reflect"

	"github.com/olehvolynets/hikari/config"
)

type ReferenceHandler struct {
	Name        string
	attrHandler AttributeHandler
}

func NewReferenceHandler(t config.Type) *ReferenceHandler {
	return &ReferenceHandler{
		Name: t.Name,
		attrHandler: AttributeHandler{
			Type:        t.Type,
			Inline:      t.Inline,
			EnumHandler: NewEnumHandler(t.Variants),
			Prefix:      NewDecorator(t.Prefix),
			Postfix:     NewDecorator(t.Postfix),
			Colorizer:   t.ToColor(),
		},
	}
}

func (h *ReferenceHandler) Render(ctx *Context, v any) {
	h.attrHandler.renderDecorator(ctx, h.attrHandler.Prefix)
	if h.attrHandler.EnumHandler == nil {
		h.attrHandler.render(ctx, reflect.ValueOf(h.attrHandler.typeConvert(v)))
	} else {
		h.attrHandler.EnumHandler.Render(ctx, reflect.ValueOf(h.attrHandler.typeConvert(v)))
	}
	h.attrHandler.renderDecorator(ctx, h.attrHandler.Postfix)
}
