package config

import (
	"errors"

	"gopkg.in/yaml.v3"
)

type SchemeItem struct {
	Name     string       `yaml:"name"`
	Type     PropertyType `yaml:"type"`
	Literal  string       `yaml:"literal"`
	Optional bool         `yaml:"optional"`
	Prefix   *Decorator   `yaml:"prefix"`
	Postfix  *Decorator   `yaml:"postfix"`

	DisplayProps `yaml:"display"`
}

type Decorator struct {
	Literal      string `yaml:"literal"`
	DisplayProps `yaml:"display"`
}

func (d *Decorator) UnmarshalYAML(node *yaml.Node) error {
	switch node.Kind {
	case yaml.ScalarNode:
		d.Literal = node.Value
	case yaml.MappingNode:
		for i := 0; i < len(node.Content); i++ {
			keyNode := node.Content[i]

			switch keyNode.Value {
			case "literal":
				valueNode := node.Content[i+1]
				d.Literal = valueNode.Value
			case "display":
				valueNode := node.Content[i+1]
				valueNode.Decode(&d.DisplayProps)
			}

			i++
		}
	default:
		return errors.New("unsupported value for prefix/postfix, must be a literal or an object")
	}

	return nil
}
