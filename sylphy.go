package sylphy

import (
	"log"

	"github.com/olehvolynets/sylphy/node"
	"github.com/olehvolynets/sylphy/render"
	"github.com/olehvolynets/sylphy/scheme"
)

type Sylphy struct {
	renderer *render.Renderer
	sequence []node.Base
}

func New(sc *scheme.Scheme) *Sylphy {
	s := &Sylphy{
		renderer: render.NewRenderer(),
	}

	for _, item := range sc.Items {
		var n node.Base

		if item.Literal != "" {
			n = node.NewLiteralNode(item)
		} else if len(item.Enum) > 0 {
			n = node.NewEnumNode(item)
		} else {
			n = newAttrNode(item, s.renderer)
		}

		s.sequence = append(s.sequence, n)
	}

	return s
}

func newAttrNode(scheme *scheme.SchemeItem, r *render.Renderer) node.Base {
	switch scheme.Type {
	case "array":
		return node.NewArrayNode(scheme, r)
	case "object":
		return node.NewObjectNode(scheme, r)
	default:
		return node.NewAttributeNode(scheme)
	}
}

func (s *Sylphy) Print(entry map[string]any) {
	for _, n := range s.sequence {
		_, err := n.Print(entry, s.renderer)
		if err != nil {
			log.Println("sylphy: output err %w", err)
		}
	}

	s.renderer.Endl()
}
