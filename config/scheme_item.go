package config

type SchemeItem struct {
	Name    string `yaml:"name"`
	Literal string `yaml:"literal"`

	As       string       `yaml:"as"`
	Type     PropertyType `yaml:"type"`
	Skip     bool         `yaml:"skip"`
	Optional bool         `yaml:"optional"`

	Variants map[string]EnumVariant `yaml:"variants"`

	Prefix  *Decorator `yaml:"prefix"`
	Postfix *Decorator `yaml:"postfix"`

	DisplayProps `yaml:",inline"`
}

type EnumVariant struct {
	Literal string `yaml:"literal"`

	Prefix  *Decorator `yaml:"prefix"`
	Postfix *Decorator `yaml:"postfix"`

	DisplayProps `yaml:",inline"`
}
