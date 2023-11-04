package node

import (
	"github.com/fatih/color"

	"github.com/olehvolynets/sylphy/render"
	"github.com/olehvolynets/sylphy/scheme"
)

type LiteralNode struct {
	baseT

	Literal   string
	Colorizer *color.Color
}

func NewLiteralNode(si *scheme.SchemeItem) *LiteralNode {
	c := si.ToColor()

	return &LiteralNode{
		Literal:   c.Sprint(si.Literal),
		Colorizer: c,
	}
}

func (l *LiteralNode) Print(entry map[string]any, r *render.Renderer) (int, error) {
	return r.Print(l.Literal)
}
