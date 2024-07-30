package hikari

import (
	"fmt"
	"reflect"
	"time"

	"github.com/olehvolynets/hikari/config"
)

type AttributeHandler struct {
	Key      string
	Skip     bool
	Optional bool
	Type     config.PropertyType
	Prefix   *Decorator
	Postfix  *Decorator

	RefHandler *ReferenceHandler

	Colorizer
}

type Decorator struct {
	Literal string
	Colorizer
}

func NewDecorator(d *config.Decorator) *Decorator {
	if d == nil {
		return nil
	}

	return &Decorator{
		Literal:   d.Literal,
		Colorizer: d.ToColor(),
	}
}

func (h *AttributeHandler) Render(ctx *Context, val Entry) {
	ctx.AddHandled(h.Key)

	if h.Skip {
		return
	}
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

	if h.RefHandler == nil {
		h.renderDecorator(ctx, h.Prefix)
		h.render(ctx, reflect.ValueOf(h.typeConvert(v)))
		h.renderDecorator(ctx, h.Postfix)
	} else {
		h.renderDecorator(ctx, h.Prefix)

		h.RefHandler.Render(ctx, v)
		h.renderDecorator(ctx, h.Postfix)
	}
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
		h.Colorizer.Fprint(ctx.W, val)
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
	if val.Len() == 0 {
		fmt.Fprintln(ctx.W, "[]")
		return
	}

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
	if val.Len() == 0 {
		fmt.Fprintln(ctx.W, "{}")
		return
	}

	fmt.Fprintln(ctx.W, "{")
	ctx.Indent()
	for _, key := range val.MapKeys() {
		mapKeyFormat.Fprintf(
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

func (h *AttributeHandler) typeConvert(v any) any {
	switch h.Type {
	case config.DurationType:
		vFloat, ok := v.(float64)
		if ok {
			return time.Duration(vFloat).String()
		} else {
			return v
		}
	case config.DateType, config.TimeType, config.DateTimeType:
		dateString, ok := v.(string)
		if !ok {
			return v
		}

		date, err := time.Parse(time.RFC3339, dateString)
		if err != nil {
			return v
		}

		switch h.Type {
		case config.DateType:
			return date.Format(time.DateOnly)
		case config.TimeType:
			return date.Format(time.TimeOnly)
		default:
			return date.Format(time.DateTime)
		}
	default:
		return v
	}
}
