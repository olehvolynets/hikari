package config

type Type struct {
	Name string       `yaml:"name"`
	Type PropertyType `yaml:"type"`

	DisplayProps `yaml:"display"`
}
