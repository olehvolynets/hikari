package render

type NumberHandler struct {
	handler[float64]
}

func NewNumberHandler(key string, optional bool, formatter Colorizer) Handler {
	return &NumberHandler{
		handler: handler[float64]{
			Accessor: Accessor[float64]{
				Key:      key,
				Optional: optional,
			},
			formatter: formatter,
		},
	}
}

var _ HandlerBuilder = NewNumberHandler

func (h *NumberHandler) Handle(ctx *Context) error {
	val, skip, err := h.Get(ctx)
	if skip {
		return nil
	}
	if err != nil {
		return err
	}

	h.formatter.Fprint(ctx.W, val)

	return nil
}
