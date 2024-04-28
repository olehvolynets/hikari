package config

import (
	"strings"

	"gopkg.in/yaml.v3"
)

type SpecialProperty string

const (
	rest SpecialProperty = "rest"
)

func (s *SpecialProperty) UnmarshalYAML(node *yaml.Node) error {
	switch strings.ToLower(node.Value) {
	case "rest":
		*s = rest
	default:
		return &ErrUnknownSpecialProperty{Value: node.Value}
	}

	return nil
}

type ErrUnknownSpecialProperty struct {
	Value string
}

func (e *ErrUnknownSpecialProperty) Error() string {
	return "unknown SpecialProperty: " + e.Value
}
