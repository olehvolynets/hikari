package format

import "github.com/olehvolynets/sylphy/render"

type DatetimeFormatter struct {
	r *render.Renderer
}

func NewDatetimeFormatter() *DatetimeFormatter {
	return &DatetimeFormatter{}
}

func (f *DatetimeFormatter) Format() {
	f.r.Print("datetime")
}
