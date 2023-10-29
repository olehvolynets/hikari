package node

import (
	"fmt"

	"github.com/fatih/color"
)

const (
	inlineSeparator    = ", "
	multilineSeparator = ",\n"
)

var (
	operatorColor = color.New(color.FgHiWhite)
	numberColor   = color.New(color.FgYellow)
	boolColor     = color.New(color.FgMagenta, color.Bold)
	stringColor   = color.New(color.FgGreen)
	nullColor     = color.New(color.FgYellow, color.Bold)
)

var (
	prettyNull         = nullColor.Sprint("null")
	prettyColon        = operatorColor.Sprint(":")
	prettyLeftBrace    = operatorColor.Sprint("{")
	prettyRightBrace   = operatorColor.Sprint("}")
	prettyLeftBracket  = operatorColor.Sprint("[")
	prettyRightBracket = operatorColor.Sprint("]")
)

func colorizeValue(val any) string {
	switch val.(type) {
	case int64:
		return numberColor.Sprintf("%d", val)
	case float64:
		return numberColor.Sprintf("%f", val)
	case string:
		return stringColor.Sprintf("%q", val)
	case bool:
		return boolColor.Sprintf("%t", val)
	case nil:
		return prettyNull
	default:
		return fmt.Sprintf("%v", val)
	}
}
