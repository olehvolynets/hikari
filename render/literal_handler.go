package render

type LiteralHandler struct {
	Literal   string
	formatter Colorizer
}

func NewLiteralHandler(literal string, formatter Colorizer) Handler {
	return &LiteralHandler{
		Literal:   literal,
		formatter: formatter,
	}
}

func (h *LiteralHandler) Handle(ctx *Context) error {
	_, err := h.formatter.Fprint(ctx.W, h.Literal)

	return err
}
