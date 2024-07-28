package hikari

import (
	"fmt"
	"reflect"
)

type AttributeHandler struct {
	Key      string
	Optional bool
	Prefix   *Decorator
	Postfix  *Decorator
	Colorizer
}

type Decorator struct {
	Literal string
	Colorizer
}

func (h *AttributeHandler) Render(ctx *Context, val Entry) {
	if len(val) == 0 {
		return
	}

	v, ok := val[h.Key]
	if !ok {
		if h.Optional {
			return
		}

		h.renderDecorator(ctx, h.Prefix)
		h.renderNull(ctx)
		h.renderDecorator(ctx, h.Postfix)

		return
	}

	h.renderDecorator(ctx, h.Prefix)
	h.render(ctx, reflect.ValueOf(v))
	h.renderDecorator(ctx, h.Postfix)
}

func (h *AttributeHandler) render(ctx *Context, val reflect.Value) {
	switch val.Kind() {
	case reflect.Invalid:
		h.renderNull(ctx)
		// if val.Interface() == nil {
		// }
		// h.Colorizer.Fprint(ctx.W, "<Invalid>")
		// fmt.Fprint(ctx.W, "<Invalid>")
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
		h.renderBool(ctx, val.Bool())
	case reflect.String:
		h.renderString(ctx, val.String())
	case reflect.Array, reflect.Slice:
		h.renderArray(ctx, val)
	case reflect.Map:
		h.renderMap(ctx, val)
	default:
		if val.IsNil() {
			h.renderNull(ctx)
		} else {
			h.Colorizer.Fprint(ctx.W, val.Kind())
		}
	}
}

func (h *AttributeHandler) renderNull(ctx *Context) {
	nullFormat.Fprint(ctx.W, "null")
}

func (h *AttributeHandler) renderNumber(ctx *Context, v any) {
	if h.Colorizer == nil {
		numberFormat.Fprintf(ctx.W, "%v", v)
	} else {
		h.Colorizer.Fprint(ctx.W, v)
	}
}

func (h *AttributeHandler) renderString(ctx *Context, v any) {
	if h.Colorizer == nil {
		stringFormat.Fprintf(ctx.W, "%s", v)
	} else {
		h.Colorizer.Fprintf(ctx.W, "%s", v)
	}
}

func (h *AttributeHandler) renderBool(ctx *Context, v any) {
	if h.Colorizer == nil {
		boolFormat.Fprintf(ctx.W, "%s", v)
	} else {
		h.Colorizer.Fprintf(ctx.W, "%s", v)
	}
}

func (h *AttributeHandler) renderArray(ctx *Context, val reflect.Value) {
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

func (h *AttributeHandler) renderMap(ctx *Context, val reflect.Value) {
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

func (h *AttributeHandler) renderDecorator(ctx *Context, d *Decorator) {
	if d == nil {
		return
	}

	switch {
	case d.Colorizer != nil:
		d.Colorizer.Fprint(ctx.W, d.Literal)
	case h.Colorizer != nil:
		h.Colorizer.Fprint(ctx.W, d.Literal)
	default:
		fmt.Fprint(ctx.W, d.Literal)
	}
}