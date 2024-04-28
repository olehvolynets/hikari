package render

type ArrayHandler struct {
	handler[[]any]
}

func NewArrayHandler(key string, optional bool, colorizer Colorizer) Handler {
	return &ArrayHandler{
		handler: handler[[]any]{
			Accessor: Accessor[[]any]{
				Key:      key,
				Optional: optional,
			},
			formatter: colorizer,
		},
	}
}

func (h *ArrayHandler) Handle(ctx *Context) error {
	value, skip, err := h.Get(ctx)
	if skip {
		return nil
	}
	if err != nil {
		return err
	}

	if len(value) == 0 {
		return nil
	}

	h.formatter.Fprint(ctx.W, "[")
	h.printElements(ctx, value)
	h.formatter.Fprint(ctx.W, "]")

	return nil
}

func (h *ArrayHandler) printElements(ctx *Context, value []any) {
	h.formatter.Fprint(ctx.W, value[0])

	for i := 1; i < len(value); i++ {
		// TODO:use custom separator
		h.formatter.Fprint(ctx.W, ", ", value[i])
	}
}
