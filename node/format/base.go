package format

import (
	"github.com/fatih/color"
)

type Formatter interface {
	Format(any)
}

var (
	OperatorColor = color.New(color.FgHiWhite)
	NumberColor   = color.New(color.FgYellow)
	BoolColor     = color.New(color.FgMagenta, color.Bold)
	StringColor   = color.New(color.FgGreen)
	NullColor     = color.New(color.FgYellow, color.Bold)
)

var PrettyNull = NullColor.Sprint("null")
