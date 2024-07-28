package config

type Type struct {
	Name    string       `yaml:"name"`
	Type    PropertyType `yaml:"type"`
	Prefix  *Decorator   `yaml:"prefix"`
	Postfix *Decorator   `yaml:"postfix"`

	DisplayProps `yaml:",inline"`
}
