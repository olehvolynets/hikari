package sylphy

import (
	"fmt"

	"github.com/olehvolynets/sylphy/node"
	"github.com/olehvolynets/sylphy/scheme"
)

type Sylphy struct {
	sequence []Node
}

type Node interface {
	Print(entry map[string]any) (int, error)
}

func New(sc *scheme.Scheme) *Sylphy {
	s := &Sylphy{}

	for _, item := range sc.Items {
		var n Node

		if item.Literal != "" {
			n = node.NewLiteralNode(item)
		} else if len(item.Enum) > 0 {
			n = node.NewEnumNode(item)
		} else {
			n = newAttrNode(item)
		}

		s.sequence = append(s.sequence, n)
	}

	return s
}

func newAttrNode(scheme *scheme.SchemeItem) Node {
	switch scheme.Type {
	case "array":
		return node.NewArrayNode(scheme)
	case "object":
		return node.NewObjectNode(scheme)
	default:
		return node.NewAttributeNode(scheme)
	}
}

func (s *Sylphy) Print(entry map[string]any) {
	for _, n := range s.sequence {
		_, err := n.Print(entry)
		if err != nil {
			fmt.Println("sylphy: output err %w", err)
		}
	}

	fmt.Println()
}
