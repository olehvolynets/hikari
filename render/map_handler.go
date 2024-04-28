package render

type MapHandler struct {
	handler[map[string]any]
}

func NewMapHandler(key string, optional bool, colorizer Colorizer) Handler {
	return &MapHandler{
		handler: handler[map[string]any]{
			Accessor: Accessor[map[string]any]{
				Key:      key,
				Optional: optional,
			},
			formatter: colorizer,
		},
	}
}

func (h *MapHandler) Handle(ctx *Context) error {
	val, skip, err := h.Get(ctx)
	if skip {
		return nil
	}
	if err != nil {
		return err
	}

	h.formatter.Fprint(ctx.W, "{")
	for k, v := range val {
		// TODO:use custom separator
		h.formatter.Fprint(ctx.W, k, ": ", v)
	}
	h.formatter.Fprint(ctx.W, "}")

	return nil
}
