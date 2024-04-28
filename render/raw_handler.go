package render

import "fmt"

type RawHandler struct {
	Value string
}

func (h *RawHandler) Handle(ctx *Context) error {
	fmt.Fprint(ctx.W, h.Value)

	return nil
}
