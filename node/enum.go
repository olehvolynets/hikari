package node

import (
	"log"

	"github.com/fatih/color"

	"github.com/olehvolynets/sylphy/render"
	"github.com/olehvolynets/sylphy/scheme"
)

type EnumNode struct {
	baseT

	Colorizer *color.Color
	Values    map[string]*LiteralNode
}

func NewEnumNode(si *scheme.SchemeItem) *EnumNode {
	if si.Type != "string" {
		panic("sylphy: only enum strings are supported")
	}

	enumColor := si.ToColor()
	en := &EnumNode{
		baseT: baseT{
			AttrName: si.Name,
			Optional: si.Optional,
		},
		Colorizer: enumColor,
		Values:    make(map[string]*LiteralNode),
	}

	for _, variant := range si.Enum {
		if variant.Value == "" {
			panic("sylphy: enum variant must have value")
		}

		variantColor := variant.ToColor()
		prefixColor, postfixColor := enumColor, enumColor
		if si.Prefix.Inherit {
			prefixColor = variantColor
		}
		if si.Postfix.Inherit {
			postfixColor = variantColor
		}

		prefix := prefixColor.Sprint(si.Prefix.Literal)
		postfix := postfixColor.Sprint(si.Postfix.Literal)

		lit := prefix + variantColor.Sprint(variant.Value) + postfix

		en.Values[variant.Value] = &LiteralNode{Literal: lit, Colorizer: variantColor}
	}

	return en
}

func (e *EnumNode) Print(entry map[string]any, r *render.Renderer) (int, error) {
	val, shouldPrint := e.findSelf(entry)
	if !shouldPrint {
		return 0, nil
	}

	switch v := val.(type) {
	case string:
		variant, ok := e.Values[v]
		if !ok {
			r.Printf("%s=<ERRUNKNVAL> - %s", e.AttrName, v)
		}

		variant.Print(entry, r)
	case error:
		log.Println(v)
		return 0, nil
	default:
		log.Println("<enum value is not a string>")
		return 0, nil
	}

	return 0, nil
}
