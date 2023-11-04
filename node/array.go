package node

import (
	"fmt"

	"github.com/olehvolynets/sylphy/node/format"
	"github.com/olehvolynets/sylphy/render"
	"github.com/olehvolynets/sylphy/scheme"
)

type ArrayNode struct {
	baseT

	formatter format.Formatter
}

func NewArrayNode(sch *scheme.SchemeItem, r *render.Renderer) *ArrayNode {
	rendCfg := render.Config{
		ElementSeparator: ",",
		KeyValSeparator:  ": ",
		LineSeparator:    " ",
		IndentChar:       "",
	}

	if sch.Multiline {
		rendCfg.LineSeparator = "\n"
		rendCfg.IndentChar = "\t"
	}

	arrayRenderer := r.WithConfig(rendCfg)

	return &ArrayNode{
		baseT: baseT{
			AttrName: sch.Name,
			Optional: sch.Optional,
		},
		formatter: format.NewArrayFormatter(arrayRenderer),
	}
}

func (n *ArrayNode) Print(entry map[string]any, r *render.Renderer) (int, error) {
	val, ok := n.findSelf(entry)
	if !ok {
		return r.Print(val)
	}

	arrVal, ok := val.([]any)
	if !ok {
		return r.Print(fmt.Errorf("%s - ERRNOTARR - %v", n.AttrName, val))
	}

	n.formatter.Format(arrVal)

	return 0, nil
}
