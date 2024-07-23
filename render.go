package sylphy

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/fatih/color"

	"github.com/olehvolynets/sylphy/config"
)

var (
	ErrTypeMismatch = errors.New("type mismatch")
	ErrMissing      = errors.New("value missing")
	ErrNil          = errors.New("is nil")
)

var (
	numberFormat Colorizer = color.New(color.FgHiRed)
	stringFormat Colorizer = color.New(color.FgYellow)
	boolFormat   Colorizer = color.New(color.FgMagenta)
	nullFormat   Colorizer = color.New(color.FgYellow, color.Bold)
)

type Entry map[string]any

type Colorizer = *color.Color

type Handler struct {
	Event config.Event
}

func (h *Handler) Render(ctx *Context, val Entry) error {
	if len(val) == 0 {
		return nil
	}

	for k, v := range val {
		fmt.Fprintf(
			ctx.W,
			"%s%s: ",
			ctx.CurrentIndent(),
			k,
		)
		h.render(ctx, reflect.ValueOf(v))
		fmt.Fprintln(ctx.W)
	}

	fmt.Fprintln(ctx.W)

	return nil
}

func (h *Handler) render(ctx *Context, val reflect.Value) {
	switch val.Kind() {
	case reflect.Invalid:
		fmt.Fprint(ctx.W, "<Invalid>")
	case reflect.Interface:
		if val.IsNil() {
			h.renderNull(ctx)
		} else {
			h.render(ctx, val.Elem())
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		h.renderNumber(ctx, val.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		h.renderNumber(ctx, val.Uint())
	case reflect.Float32, reflect.Float64:
		h.renderNumber(ctx, val.Float())
	case reflect.Bool:
		boolFormat.Fprint(ctx.W, val.Bool())
	case reflect.String:
		stringFormat.Fprintf(ctx.W, "%q", val.String())
	case reflect.Array, reflect.Slice:
		h.renderArray(ctx, val)
	case reflect.Map:
		h.renderMap(ctx, val)
	default:
		if val.IsNil() {
			h.renderNull(ctx)
		} else {
			fmt.Fprint(ctx.W, val.Kind())
		}
	}
}

func (h *Handler) renderNull(ctx *Context) {
	nullFormat.Fprint(ctx.W, "null")
}

func (h *Handler) renderNumber(ctx *Context, v any) {
	numberFormat.Fprint(ctx.W, v)
}

func (h *Handler) renderArray(ctx *Context, val reflect.Value) {
	fmt.Fprint(ctx.W, "[")
	ctx.Indent()
	if val.Len() > 0 {
		h.render(ctx, val.Index(0))

		if val.Len() > 1 {
			for i := 1; i < val.Len(); i++ {
				fmt.Fprint(ctx.W, ", ")
				h.render(ctx, val.Index(i))
			}
		}
	}
	ctx.Dedent()
	fmt.Fprint(ctx.W, "]")
}

func (h *Handler) renderMap(ctx *Context, val reflect.Value) {
	fmt.Fprintln(ctx.W, "{")
	ctx.Indent()
	for _, key := range val.MapKeys() {
		fmt.Fprintf(
			ctx.W,
			"%s%s: ",
			ctx.CurrentIndent(),
			key.String(),
		)
		h.render(ctx, val.MapIndex(key))
		fmt.Fprintln(ctx.W)
	}
	ctx.Dedent()
	fmt.Fprint(ctx.W, ctx.CurrentIndent(), "}")
}
