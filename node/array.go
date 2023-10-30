package node

import (
	"bytes"
	"fmt"

	sch "github.com/olehvolynets/sylphy/scheme"
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
		fmt.Fprint(buf, colorizeValue(item, false))

		if idx < len(s)-1 {
			operatorColor.Fprint(buf, inlineSeparator)
		}
	}

	operatorColor.Fprint(buf, "]")
}

func formatMultilineSlice(buf *bytes.Buffer, s []any) {
	operatorColor.Fprint(buf, "[\n")

	for idx, item := range s {
		fmt.Fprintf(buf, "\t%s", colorizeValue(item, true))

		if idx < len(s)-1 {
			operatorColor.Fprint(buf, multilineSeparator)
		}
	}

	operatorColor.Fprint(buf, "\n]")
}
