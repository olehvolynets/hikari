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
		for i := 0; i < len(node.Content); i += 2 {
			keyNode := node.Content[i]

			if keyNode.Value == "literal" {
				valueNode := node.Content[i+1]
				d.Literal = valueNode.Value
				break
			}
		}

		node.Decode(&d.DisplayProps)
	default:
		return errors.New("unsupported value for prefix/postfix, must be a literal or an object")
	}

	return nil
}
