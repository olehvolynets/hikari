package format

import "github.com/olehvolynets/sylphy/render"

type StringFormatter struct {
	r *render.Renderer
}

func (f *StringFormatter) Format(v any) {
	StringColor.Fprintf(f.r, "%q", v)
}
