package format

import "github.com/olehvolynets/sylphy/render"

type AnyFormatter struct {
	r *render.Renderer
}

func (f *AnyFormatter) Format(v any) {}
