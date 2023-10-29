package scheme

import "github.com/fatih/color"

var fgColors map[string]color.Attribute = map[string]color.Attribute{
	"white":      color.FgWhite,
	"hi-white":   color.FgHiWhite,
	"black":      color.FgBlack,
	"hi-black":   color.FgHiBlack,
	"red":        color.FgRed,
	"hi-red":     color.FgHiRed,
	"green":      color.FgGreen,
	"hi-green":   color.FgHiGreen,
	"yellow":     color.FgYellow,
	"hi-yellow":  color.FgHiYellow,
	"blue":       color.FgBlue,
	"hi-blue":    color.FgHiBlue,
	"magenta":    color.FgMagenta,
	"hi-magenta": color.FgHiMagenta,
	"cyan":       color.FgCyan,
	"hi-cyan":    color.FgHiCyan,
}

var bgColors map[string]color.Attribute = map[string]color.Attribute{
	"white":      color.BgWhite,
	"hi-white":   color.BgHiWhite,
	"black":      color.BgBlack,
	"hi-black":   color.BgHiBlack,
	"red":        color.BgRed,
	"hi-red":     color.BgHiRed,
	"green":      color.BgGreen,
	"hi-green":   color.BgHiGreen,
	"yellow":     color.BgYellow,
	"hi-yellow":  color.BgHiYellow,
	"blue":       color.BgBlue,
	"hi-blue":    color.BgHiBlue,
	"magenta":    color.BgMagenta,
	"hi-magenta": color.BgHiMagenta,
	"cyan":       color.BgCyan,
	"hi-cyan":    color.BgHiCyan,
}
