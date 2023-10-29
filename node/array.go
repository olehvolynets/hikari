package node

import (
	"bytes"
	"fmt"

	"github.com/fatih/color"
	sch "github.com/olehvolynets/sylphy/scheme"
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
)

type ArrayNode struct {
	AttrName  string
	Optional  bool
	Multiline bool
}

func NewArrayNode(scheme *sch.SchemeItem) *ArrayNode {
	return &ArrayNode{
		AttrName:  scheme.Name,
		Optional:  scheme.Optional,
		Multiline: scheme.Multiline,
	}
}

func (n *ArrayNode) Print(entry map[string]any) (int, error) {
	val, ok := entry[n.AttrName]
	if !ok {
		if n.Optional {
			return 0, nil
		}

		return fmt.Printf("%s - ERRMISS", n.AttrName)
	}

	arrVal, ok := val.([]any)
	if !ok {
		return fmt.Printf("[%v - ERRNOTARR]", val)
	}

	return fmt.Print(formatSlice(arrVal, n.Multiline))
}

func formatSlice(s []any, multiline bool) string {
	buf := bytes.NewBuffer([]byte{})

	if multiline {
		formatMultilineSlice(buf, s)
	} else {
		formatInlineSlice(buf, s)
	}

	return buf.String()
}

func formatInlineSlice(buf *bytes.Buffer, s []any) {
	operatorColor.Fprint(buf, "[")

	for idx, item := range s {
		switch item.(type) {
		case int64:
			numberColor.Fprintf(buf, "%d", item)
		case float64:
			numberColor.Fprintf(buf, "%f", item)
		case string:
			stringColor.Fprintf(buf, "%q", item)
		case bool:
			boolColor.Fprintf(buf, "%t", item)
		default:
			fmt.Fprintf(buf, "%v", item)
		}

		if idx < len(s)-1 {
			operatorColor.Fprint(buf, inlineSeparator)
		}
	}

	operatorColor.Fprint(buf, "]")
}

func formatMultilineSlice(buf *bytes.Buffer, s []any) {
	operatorColor.Fprint(buf, "[\n")

	for idx, item := range s {
		switch item.(type) {
		case int64:
			numberColor.Fprintf(buf, "\t%d", item)
		case float64:
			numberColor.Fprintf(buf, "\t%f", item)
		case string:
			stringColor.Fprintf(buf, "\t%q", item)
		case bool:
			boolColor.Fprintf(buf, "\t%t", item)
		default:
			fmt.Fprintf(buf, "\t%v", item)
		}

		if idx < len(s)-1 {
			operatorColor.Fprint(buf, multilineSeparator)
		}
	}

	operatorColor.Fprint(buf, "\n]")
}
