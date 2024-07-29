package config

import (
	"errors"
	"reflect"

	"gopkg.in/yaml.v3"
)

type Event struct {
	Name    string       `yaml:"name"`
	Matcher Matchers     `yaml:"matcher"`
	Scheme  []SchemeItem `yaml:"scheme"`
}

func (evt *Event) Match(entry map[string]any) bool {
	return evt.Matcher.Match(entry)
}

type Matcher struct {
	key     string
	value   any
	Present bool `yaml:"present"`
	Filled  bool `yaml:"filled"`
}

type Matchers struct {
	ms []Matcher
}

func (container *Matchers) UnmarshalYAML(node *yaml.Node) error {
	if node.Kind != yaml.MappingNode {
		return errors.New("unsupported matcher format")
	}

	for i := 0; i < len(node.Content); i++ {
		keyNode := node.Content[i]
		valueNode := node.Content[i+1]
		i++

		matcher := Matcher{
			key: keyNode.Value,
		}

		switch valueNode.Kind {
		case yaml.ScalarNode:
			matcher.value = valueNode.Value
		case yaml.MappingNode:
			if err := valueNode.Decode(&matcher); err != nil {
				return err
			}
		}

		container.ms = append(container.ms, matcher)
	}

	return nil
}

func (matchers *Matchers) Match(entry map[string]any) bool {
	for _, m := range matchers.ms {
		param, ok := entry[m.key]
		if !ok {
			return false
		}

		if m.Present || m.Filled {
			if m.Filled {
				if reflect.ValueOf(param).IsZero() {
					return false
				}

				continue
			}

			continue
		}

		if param != m.value {
			return false
		}
	}

	return true
}
