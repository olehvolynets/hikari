package config

type Property struct {
	Name string       `yaml:"name"`
	Type PropertyType `yaml:"type"`
}
