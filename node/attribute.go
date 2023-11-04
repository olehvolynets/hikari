package node

import (
	"fmt"
	"log"
	"time"

	"github.com/fatih/color"

	"github.com/olehvolynets/sylphy/render"
	"github.com/olehvolynets/sylphy/scheme"
)

type AttributeNode struct {
	baseT

	FormatStr string
	Colorizer *color.Color
	Prefix    *LiteralNode
	Postfix   *LiteralNode
	Preproc   func(string) string

	// formatter format.Formatter
}

func NewAttributeNode(si *scheme.SchemeItem) *AttributeNode {
	node := &AttributeNode{
		baseT: baseT{
			AttrName: si.Name,
			Optional: si.Optional,
		},
		FormatStr: si.FormatStr(),
		Colorizer: si.ToColor(),
	}

	node.Prefix = node.surrounder(si.Prefix)
	node.Postfix = node.surrounder(si.Postfix)

	if si.Type == "datetime" {
		node.Preproc = func(in string) string {
			t, err := time.Parse(si.SrcFormat, in)
			if err != nil {
				log.Println(err)
			}
			return t.Format(si.DstFormat)
		}
	}

	return node
}

func (n *AttributeNode) Print(entry map[string]any, r *render.Renderer) (int, error) {
	if n.Prefix != nil {
		n.Prefix.Print(entry, r)
	}
	defer func() {
		if n.Postfix != nil {
			n.Postfix.Print(entry, r)
		}
	}()

	value, ok := n.findSelf(entry)
	if !ok {
		return r.Print(value)
	}

	if n.Preproc != nil {
		strVal, ok := value.(string)
		if !ok {
			panic(fmt.Errorf("AttributeNode.Preproc called with non-string (%T) argument %v", value, value))
		}

		value = n.Preproc(strVal)
	}

	return n.Colorizer.Fprintf(r, n.FormatStr, value)
}

func (n *AttributeNode) surrounder(d *scheme.Decorator) *LiteralNode {
	if d == nil {
		return nil
	}

	postfixColor := n.Colorizer
	if !d.Inherit {
		postfixColor = d.ToColor()
	}

	literal := postfixColor.Sprint(d.Literal)

	return &LiteralNode{Literal: literal}
}
