package hikari

import (
	"fmt"

	"github.com/olehvolynets/hikari/config"
)

type EnumHandler struct {
	Variants map[string]LiteralHandler
}

func NewEnumHandler(cfg map[string]config.EnumVariant) *EnumHandler {
	handler := &EnumHandler{
		Variants: make(map[string]LiteralHandler, len(cfg)),
	}

	for key, variant := range cfg {
		replaceValue := variant.Literal
		if variant.Literal == "" {
			replaceValue = key
		}

		handler.Variants[key] = LiteralHandler{
			Literal:   replaceValue,
			Colorizer: variant.ToColor(),
		}
	}

	return handler
}

func (h *EnumHandler) Render(ctx *Context, val string) {
	for k, variantHandler := range h.Variants {
		if k == val {
			variantHandler.Render(ctx, nil)
			return
		}
	}

	fmt.Fprint(ctx.W, val)
}
