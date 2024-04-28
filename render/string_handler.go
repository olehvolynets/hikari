package render

type StringHandler struct {
	handler[string]
}

func NewStringHandler(key string, optional bool, formatter Colorizer) Handler {
	return &StringHandler{
		handler: handler[string]{
			Accessor: Accessor[string]{
				Key:      key,
				Optional: optional,
			},
			formatter: formatter,
		},
	}
}

var _ HandlerBuilder = NewStringHandler

func (h *StringHandler) Handle(ctx *Context) error {
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
