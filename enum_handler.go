package hikari

import (
	"fmt"
	"reflect"

	"github.com/olehvolynets/hikari/config"
)

type EnumHandler struct {
	Variants []EnumVariantHandler
}

func NewEnumHandler(cfg []config.EnumVariant) *EnumHandler {
	handler := &EnumHandler{
		Variants: make([]EnumVariantHandler, len(cfg)),
	}

	for _, variant := range cfg {
		replaceValue := variant.Value
		if variant.Replace != "" {
			replaceValue = variant.Replace
		}

		if variant.Min > variant.Max {
			panic("hikari: enum's min can't be greater than max")
		}

		handler.Variants = append(handler.Variants, EnumVariantHandler{
			Value: variant.Value,
			Min:   variant.Min,
			Max:   variant.Max,
			LiteralHandler: LiteralHandler{
				Literal:   replaceValue,
				Colorizer: variant.ToColor(),
			},
		})
	}

	return handler
}

func (h *EnumHandler) Render(ctx *Context, val reflect.Value) {
	for _, variantHandler := range h.Variants {
		if variantHandler.Matches(val) {
			variantHandler.Render(ctx, val)
			return
		}
	}

	fmt.Fprint(ctx.W, val)
}

type EnumVariantHandler struct {
	Value string

	Min float64 `yaml:"min"`
	Max float64 `yaml:"max"`

	LiteralHandler
}

func (h *EnumVariantHandler) Render(ctx *Context, val reflect.Value) {
	if h.Min != 0 && h.Max != 0 {
		h.LiteralHandler.Colorizer.Fprint(ctx.W, numericValue(val))
	} else {
		h.LiteralHandler.Render(ctx, nil)
	}
}

func (h *EnumVariantHandler) Matches(val reflect.Value) bool {
	switch val.Kind() {
	case reflect.String:
		return val.String() == h.Value
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return h.matchNumericRange(numericValue(val))
	}

	return false
}

func (h *EnumVariantHandler) matchNumericRange(val float64) bool {
	if val >= h.Min && val <= h.Max {
		return true
	}

	return false
}

func numericValue(v reflect.Value) float64 {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(v.Uint())
	case reflect.Float32, reflect.Float64:
		return v.Float()
	default:
		panic("hikari: what????")
	}
}
