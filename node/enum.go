package node

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/olehvolynets/sylphy/scheme"
)

type EnumNode struct {
	AttrName  string
	Optional  bool
	Colorizer *color.Color

	Prefix         string
	PrefixInherit  bool
	Postfix        string
	PostfixInherit bool

	Values map[string]*LiteralNode
}

func NewEnumNode(si *scheme.SchemeItem) *EnumNode {
	if si.Type != "string" {
		panic("sylphy: only enum strings are supported")
	}

	en := &EnumNode{
		AttrName:  si.Name,
		Optional:  si.Optional,
		Colorizer: si.ToColor(),
		Values:    make(map[string]*LiteralNode),
	}

	if si.Prefix != nil {
		en.Prefix = si.Prefix.Literal
		en.PrefixInherit = si.Prefix.Inherit
	}

	if si.Postfix != nil {
		en.Postfix = si.Postfix.Literal
		en.PostfixInherit = si.Postfix.Inherit
	}

	for _, variant := range si.Enum {
		if variant.Value == "" {
			panic("sylphy: enum variant must have value")
		}

		c := variant.ToColor()
		lit := c.Sprint(variant.Value)

		en.Values[variant.Value] = &LiteralNode{Literal: lit, Colorizer: c}
	}

	return en
}

func (e *EnumNode) Print(entry map[string]any) (int, error) {
	var written int
	var resultErr error = nil
	var err error
	var val any
	var strVal string
	var ok bool
	var variant *LiteralNode
	var i int

	if !e.PrefixInherit {
		i, err = e.Colorizer.Print(e.Prefix)
		if err != nil {
			resultErr = err
			goto exit
		}

		written += i
	}

	val, ok = entry[e.AttrName]
	if !ok {
		if e.Optional {
			goto exit
		}

		i, err = fmt.Printf("%s=<ERRNOVAL>", e.AttrName)
		if err != nil {
			resultErr = err
			goto exit
		}

		written += i
	}

	strVal, ok = val.(string)
	if !ok {
		i, err = fmt.Print("<enum value is not a string>")
		if err != nil {
			resultErr = err
			goto exit
		}

		written += i
		goto exit
	}

	variant, ok = e.Values[strVal]
	if !ok {
		i, err = fmt.Printf("%s=<ERRUNKNVAL> - %s", e.AttrName, strVal)
		if err != nil {
			resultErr = err
			goto exit
		}

		written += i
	}

	if e.PrefixInherit {
		variant.Colorizer.Print(e.Prefix)
	}

	i, err = variant.Print(entry)
	if err != nil {
		resultErr = err
		goto exit
	}
	written += i

	if !e.PostfixInherit {
		i, err = e.Colorizer.Print(e.Postfix)
		if err != nil {
			resultErr = err
			goto exit
		}

		written += i
	} else {
		i, err = variant.Colorizer.Print(e.Postfix)
		if err != nil {
			resultErr = err
			goto exit
		}
		written += i
	}

exit:
	return written, resultErr
}
