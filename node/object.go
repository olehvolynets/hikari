package node

import (
	"fmt"

	"github.com/olehvolynets/sylphy/node/format"
	"github.com/olehvolynets/sylphy/render"
	"github.com/olehvolynets/sylphy/scheme"
)

type ObjectNode struct {
	baseT

	formatter format.Formatter
}

func NewObjectNode(sch *scheme.SchemeItem, r *render.Renderer) *ObjectNode {
	rendCfg := render.Config{ElementSeparator: ",", KeyValSeparator: ": "}
	if sch.Multiline {
		rendCfg.LineSeparator = "\n"
		rendCfg.IndentChar = "\t"
	} else {
		rendCfg.LineSeparator = " "
		rendCfg.IndentChar = ""
	}

	mapRenderer := r.WithConfig(rendCfg)

	return &ObjectNode{
		baseT: baseT{
			AttrName: sch.Name,
			Optional: sch.Optional,
		},
		formatter: format.NewObjectFormatter(mapRenderer),
	}
}

func (n *ObjectNode) Print(entry map[string]any, r *render.Renderer) (int, error) {
	val, ok := n.findSelf(entry)
	if !ok {
		return r.Print(val)
	}

	mapVal, ok := val.(map[string]any)
	if !ok {
		return r.Print(fmt.Errorf("%s - ERRNOTMAP - %v", n.AttrName, val))
	}

	n.formatter.Format(mapVal)

	return 0, nil
}
