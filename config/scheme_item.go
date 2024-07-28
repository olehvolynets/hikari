package config

import "github.com/fatih/color"

type SchemeItem struct {
	Name    string       `yaml:"name"`
	Type    PropertyType `yaml:"type"`
	Literal string       `yaml:"literal"`

	DisplayProps `yaml:"display"`
}

func (item *SchemeItem) ToColor() *color.Color {
	var zero DisplayProps
	if item.DisplayProps == zero {
		return nil
	}

	return item.DisplayProps.ToColor()
}
