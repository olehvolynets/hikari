package sylphy

import "encoding/json"

type Config struct {
	// Refer to fields as $<field_name>
	Format string `json:"format"`

	// A field not mentioned in the Config.Format (default: "k: v")
	UnknownFieldFormat string `json:"unknown_field_format"`
	// How to join unknown fields (default: new line), i.e.
	//  k1: v1, k2: v2
	// or
	//  k1: v1
	//  k2: v2
	UnknownFieldSeparator string `json:"unknown_field_separator"`

	Fields map[string]FieldProp `json:"fields"`
}

type FieldProp struct {
	// Default - term white
	Fg string `json:"fg"`
	// Default - term black
	Bg string `json:"bg"`
	// Default - false
	Bold bool `json:"bold"`
	// Default - false
	Italic bool `json:"italic"`
	// Default - false
	Underline bool `json:"underline"`
	// Default - false
	Strikethrough bool `json:"strikethrough"`
	// Default - false
	Blink bool `json:"blink"`
	// Default - string
	Type string `json:"type"`
}

// TODO: use io.Reader for more general approach
func NewConfig(r []byte) (*Config, error) {
	cfg := &Config{
		Fields: make(map[string]FieldProp),
	}

	err := json.Unmarshal(r, cfg)
	if err != nil {
		return nil, err
	}

	for field, fp := range cfg.Fields {
		cfg.Fields[field] = populateFieldProp(fp)
	}

	return cfg, nil
}

func populateFieldProp(fp FieldProp) FieldProp {
	if fp.Type == "" {
		fp.Type = "string"
	}

	if fp.Fg == "" {
		fp.Fg = "white"
	}

	if fp.Bg == "" {
		fp.Bg = "black"
	}

	return fp
}
