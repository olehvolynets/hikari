package render

type BoolHandler struct {
	handler[bool]
}

func NewBoolHandler(key string, optional bool, formatter Colorizer) Handler {
	return &BoolHandler{
		handler: handler[bool]{
			Accessor: Accessor[bool]{
				Key:      key,
				Optional: optional,
			},
			formatter: formatter,
		},
	}
}

var _ HandlerBuilder = NewBoolHandler

func (h *BoolHandler) Handle(ctx *Context) error {
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
