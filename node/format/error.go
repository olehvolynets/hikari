package format

import (
	"github.com/fatih/color"

	"github.com/olehvolynets/sylphy/render"
)

var errColor = color.New(color.BgHiRed)

type ErrorFormatter struct {
	r *render.Renderer
}

func (f *ErrorFormatter) Format(v any) {
	errColor.Fprintf(f.r, "wtf type? - %v", v)
}
