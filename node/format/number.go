package format

import "github.com/olehvolynets/sylphy/render"

type NumberFormatter struct {
	r *render.Renderer
}

func (f *NumberFormatter) Format(v any) {
	NumberColor.Fprint(f.r, v)
}
