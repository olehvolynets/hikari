package render

import (
	"github.com/fatih/color"
)

type DisplayProps struct {
	Fg            string `yaml:"fg"`
	Bg            string `yaml:"bg"`
	Bold          bool   `yaml:"bold"`
	Italic        bool   `yaml:"italic"`
	Underline     bool   `yaml:"underline"`
	Strikethrough bool   `yaml:"strikethrough"`
	Blink         bool   `yaml:"blink"`
}

func (dp *DisplayProps) ToColor() Colorizer {
	c := color.New()

	if fg, ok := FgColors[dp.Fg]; ok {
		c.Add(fg)
	}
	if bg, ok := BgColors[dp.Bg]; ok {
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

	return Colorizer(c)
}
