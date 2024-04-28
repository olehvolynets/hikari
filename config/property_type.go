package config

import (
	"strings"

	"gopkg.in/yaml.v3"
)

type PropertyType string

const (
	NumberType PropertyType = "number"
	StringType PropertyType = "string"
	BoolType   PropertyType = "bool"
	ArrayType  PropertyType = "array"
	MapType    PropertyType = "map"
	EnumType   PropertyType = "enum"
)

func (s *PropertyType) UnmarshalYAML(node *yaml.Node) error {
	switch strings.ToLower(node.Value) {
	case "number":
		*s = NumberType
	case "string":
		*s = StringType
	case "bool":
		*s = BoolType
	case "array":
		*s = ArrayType
	case "map":
		*s = MapType
	case "enum":
		*s = EnumType
	default:
		return &ErrUnknownPropertyType{Value: node.Value}
	}

	return nil
}

type ErrUnknownPropertyType struct {
	Value string
}

func (e *ErrUnknownPropertyType) Error() string {
	return "unknown PropertyType: " + e.Value
}
