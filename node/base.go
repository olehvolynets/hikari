package node

import (
	"fmt"

	"github.com/olehvolynets/sylphy/render"
)

type Base interface {
	Print(entry map[string]any, r *render.Renderer) (int, error)
}

type baseT struct {
	AttrName string
	Optional bool
}

func (b *baseT) findSelf(entry map[string]any) (any, bool) {
	val, ok := entry[b.AttrName]
	if ok {
		return val, true
	}

	if b.Optional {
		return nil, false
	}

	return fmt.Errorf("<%s - MISSING>", b.AttrName), true
}
