package hikari

type EnumHandler struct {
	AttributeHandler `yaml:",inline"`
	Variants         map[string]EnumVariantHandler `yaml:"variants"`
}

type EnumVariantHandler struct {
	Value            string `yaml:"value"`
	AttributeHandler `yaml:",inline"`
}
