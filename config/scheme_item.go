package config

type SchemeItem struct {
	Name    string `yaml:"name"`
	Literal string `yaml:"literal"`

	As       string       `yaml:"as"`
	Type     PropertyType `yaml:"type"`
	Skip     bool         `yaml:"skip"`
	Optional bool         `yaml:"optional"`

	Variants []EnumVariant `yaml:"variants"`

	Prefix  *Decorator `yaml:"prefix"`
	Postfix *Decorator `yaml:"postfix"`

	DisplayProps `yaml:",inline"`
}

type EnumVariant struct {
	Value   string `yaml:"value"`
	Replace string `yaml:"replace"`

	Min float64 `yaml:"min"`
	Max float64 `yaml:"max"`

	Prefix  *Decorator `yaml:"prefix"`
	Postfix *Decorator `yaml:"postfix"`

	DisplayProps `yaml:",inline"`
}
