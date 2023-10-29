package node

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/olehvolynets/sylphy/node/internal/util"
	"github.com/olehvolynets/sylphy/scheme"
)

type AttributeNode struct {
	AttrName  string
	Optional  bool
	FormatStr string
	Colorizer *color.Color
	Prefix    *LiteralNode
	Postfix   *LiteralNode
	Preproc   func(string) string
}

func NewAttributeNode(si *scheme.SchemeItem) *AttributeNode {
	node := &AttributeNode{
		AttrName:  si.Name,
		Optional:  si.Optional,
		FormatStr: si.FormatStr(),
		Colorizer: si.ToColor(),
	}

	node.Prefix = node.surrounder(si.Prefix)
	node.Postfix = node.surrounder(si.Postfix)

	if si.Type == "datetime" {
		node.Preproc = func(in string) string {
			t, err := time.Parse(si.SrcFormat, in)
			if err != nil {
				fmt.Println(err)
			}
			return t.Format(si.DstFormat)
		}
	}

	return node
}

func (n *AttributeNode) Print(entry map[string]any) (int, error) {
	p := util.Printer{}

	if n.Prefix != nil {
		p.Print(n.Prefix.Print, entry)
	}

	p.Print(n.printSelf, entry)

	if n.Postfix != nil {
		p.Print(n.Postfix.Print, entry)
	}

	return p.Result()
}

func (n *AttributeNode) printSelf(entry map[string]any) (int, error) {
	value, ok := entry[n.AttrName]
	if !ok {
		if n.Optional {
			return 0, nil
		}

		return n.Colorizer.Printf("%s=<ERRNOVAL>", n.AttrName)
	}

	if n.Preproc != nil {
		strVal, ok := value.(string)
		if !ok {
			panic(fmt.Sprintf("AttributeNode.Preproc called with non-string (%T) argument %v", value, value))
		}

		return n.Colorizer.Printf(n.FormatStr, n.Preproc(strVal))
	} else {
		return n.Colorizer.Printf(n.FormatStr, value)
	}
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
