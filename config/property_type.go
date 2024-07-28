package config

import (
	"strings"

	"gopkg.in/yaml.v3"
)

type PropertyType string

const (
	ArrayType    PropertyType = "array"
	BoolType     PropertyType = "bool"
	DateTimeType PropertyType = "datetime"
	DateType     PropertyType = "date"
	TimeType     PropertyType = "time"
	DurationType PropertyType = "duration"
	EnumType     PropertyType = "enum"
	MapType      PropertyType = "map"
	NumberType   PropertyType = "number"
	StringType   PropertyType = "string"
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
	case "datetime":
		*s = DateTimeType
	case "date":
		*s = DateType
	case "time":
		*s = TimeType
	case "duration":
		*s = DurationType
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
