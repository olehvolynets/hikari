package format

import "github.com/olehvolynets/sylphy/render"

type BoolFormatter struct {
	r *render.Renderer
}

func (f *BoolFormatter) Format(v any) {
	BoolColor.Fprint(f.r, v)
}
