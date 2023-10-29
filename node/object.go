package node

import (
	"fmt"

	sch "github.com/olehvolynets/sylphy/scheme"
)

type ObjectNode struct{}

func NewObjectNode(scheme *sch.SchemeItem) *ObjectNode {
	fmt.Printf("%s - ObjectNode not implemented\n", scheme.Name)

	return &ObjectNode{}
}

func (n *ObjectNode) Print(entry map[string]any) (int, error) {
	return fmt.Print("{ ObjectNode }")
}
