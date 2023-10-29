package scheme

import (
	"github.com/fatih/color"
)

type DisplayProps struct {
	Fg            string `yaml:"fg"`
	Bg            string `yaml:"bg"`
	Type          string `yaml:"type"`
	Bold          bool   `yaml:"bold"`
	Italic        bool   `yaml:"italic"`
	Underline     bool   `yaml:"underline"`
	Strikethrough bool   `yaml:"strikethrough"`
	Blink         bool   `yaml:"blink"`

	Quotted bool `yaml:"quotted"`
}

func (dp *DisplayProps) ToColor() *color.Color {
	c := color.New()

	if fg, ok := fgColors[dp.Fg]; ok {
		c.Add(fg)
	}
	if bg, ok := bgColors[dp.Bg]; ok {
		c.Add(bg)
	}

	if dp.Bold {
		c.Add(color.Bold)
	}
	if dp.Italic {
		c.Add(color.Italic)
	}
	if dp.Underline {
		c.Add(color.Underline)
	}
	if dp.Strikethrough {
		c.Add(color.CrossedOut)
	}

	return c
}

func (dp *DisplayProps) FormatStr() string {
	switch dp.Type {
	case "string":
		if dp.Quotted {
			return "%q"
		} else {
			return "%s"
		}
	case "int":
		return "%d"
	case "float":
		return "%f"
	case "bool":
		return "%t"
	case "datetime":
		return "%s"
		// case "array":
		// case "map":
	default:
		panic("formatter: unknown type " + dp.Type)
	}
}
