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
		if item.Literal != "" {
			s.sequence = append(s.sequence, node.NewLiteralNode(item))
			continue
		}

		if len(item.Enum) != 0 {
			s.sequence = append(s.sequence, node.NewEnumNode(item))
			continue
		}

		s.sequence = append(s.sequence, node.NewAttributeNode(item))
	}

	return s
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
