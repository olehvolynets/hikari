package config

import (
	"errors"

	"gopkg.in/yaml.v3"
)

type Decorator struct {
	Literal      string `yaml:"literal"`
	DisplayProps `yaml:",inline"`
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
