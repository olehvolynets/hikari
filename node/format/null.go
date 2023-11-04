package format

import "github.com/olehvolynets/sylphy/render"

type NullFormatter struct {
	r *render.Renderer
}

func (f *NullFormatter) Format(v any) {
	f.r.Print(PrettyNull)
}
