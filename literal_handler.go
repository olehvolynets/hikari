package hikari

import "fmt"

type LiteralHandler struct {
	Literal string
	Colorizer
}

func (h *LiteralHandler) Render(ctx *Context, _ Entry) {
	if h.Colorizer == nil {
		fmt.Fprint(ctx.W, h.Literal)
	} else {
		h.Colorizer.Fprint(ctx.W, h.Literal)
	}
}
