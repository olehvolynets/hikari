package node

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/olehvolynets/sylphy/scheme"
)

type LiteralNode struct {
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

func (l *LiteralNode) Print(entry map[string]any) (int, error) {
	return fmt.Print(l.Literal)
}
