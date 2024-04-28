package config

import "github.com/olehvolynets/sylphy/render"

type Type struct {
	Name string       `yaml:"name"`
	Type PropertyType `yaml:"type"`

	render.DisplayProps `yaml:"display"`
}
