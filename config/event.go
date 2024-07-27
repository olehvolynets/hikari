package config

import "github.com/fatih/color"

type Event struct {
	Name    string         `yaml:"name"`
	Matcher map[string]any `yaml:"matcher"`
	Scheme  []SchemeItem   `yaml:"scheme"`
}

func (evt *Event) Match(entry map[string]any) bool {
	for key, val := range evt.Matcher {
		param, ok := entry[key]
		if !ok {
			return false
		}

		if param != val {
			return false
		}
	}

	return true
}

type SchemeItem struct {
	Property     `yaml:",inline"`
	Literal      `yaml:",inline"`
	DisplayProps `yaml:"display"`
}

func (item *SchemeItem) ToColor() *color.Color {
	var zero DisplayProps
	if item.DisplayProps == zero {
		return nil
	}

	return item.DisplayProps.ToColor()
}
