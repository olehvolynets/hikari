package node

import (
	"bytes"
	"fmt"

	sch "github.com/olehvolynets/sylphy/scheme"
	"golang.org/x/exp/slices"
)

var (
	keyValObjFormat         = "%s" + operatorColor.Sprint(":") + " %s"
	multilinKeyValObjFormat = "\t" + keyValObjFormat
)

type ObjectNode struct {
	AttrName  string
	Optional  bool
	Multiline bool
}

func NewObjectNode(scheme *sch.SchemeItem) *ObjectNode {
	return &ObjectNode{
		AttrName:  scheme.Name,
		Optional:  scheme.Optional,
		Multiline: scheme.Multiline,
	}
}

func (n *ObjectNode) Print(entry map[string]any) (int, error) {
	val, ok := entry[n.AttrName]
	if !ok {
		if n.Optional {
			return 0, nil
		}

		return fmt.Printf("%s - ERRMISS", n.AttrName)
	}

	mapVal, ok := val.(map[string]any)
	if !ok {
		return fmt.Printf("{ ERRNOTARR %s %v }", n.AttrName, val)
	}

	return fmt.Print(formatObject(mapVal, n.Multiline))
}

func formatObject(obj map[string]any, multiline bool) string {
	buf := bytes.NewBuffer([]byte{})

	keys := make([]string, 0, len(obj))
	for key := range obj {
		keys = append(keys, key)
	}
	slices.Sort(keys)

	if multiline {
		formatMultilineObject(buf, keys, obj)
	} else {
		formatInlineObject(buf, keys, obj)
	}

	return buf.String()
}

func formatInlineObject(buf *bytes.Buffer, keys []string, obj map[string]any) {
	operatorColor.Fprint(buf, prettyLeftBrace, " ")

	for idx, key := range keys {
		var strVal string

		val, ok := obj[key]
		if ok {
			strVal = colorizeValue(val)
		} else {
			strVal = prettyNull
		}

		fmt.Fprintf(buf, "%s%s %s",
			stringColor.Sprintf("%q", key),
			prettyColon,
			strVal)

		if idx < len(keys)-1 {
			operatorColor.Fprint(buf, inlineSeparator)
		}
	}

	operatorColor.Fprint(buf, " ", prettyRightBrace)
}

func formatMultilineObject(buf *bytes.Buffer, keys []string, obj map[string]any) {
	operatorColor.Fprint(buf, prettyLeftBrace, "\n")

	for idx, key := range keys {
		var strVal string

		val, ok := obj[key]
		if ok {
			strVal = colorizeValue(val)
		} else {
			strVal = prettyNull
		}

		fmt.Fprintf(buf, "\t%s%s %s",
			stringColor.Sprintf("%q", key),
			prettyColon,
			strVal)

		if idx < len(keys)-1 {
			operatorColor.Fprint(buf, multilineSeparator)
		}
	}

	operatorColor.Fprint(buf, "\n", prettyRightBrace)
}
