package render

type NullHandler struct {
	// FIX: nil type
	handler[bool]
}

func NewNullHandler(key string, optional bool, colorizer Colorizer) Handler {
	return &NullHandler{
		handler: handler[bool]{
			Accessor: Accessor[bool]{
				Key:      key,
				Optional: optional,
			},
			formatter: colorizer,
		},
	}
}

func (h *NullHandler) Handle(ctx *Context) error {
	return nil
}
