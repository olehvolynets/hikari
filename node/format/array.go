package format

import (
	"reflect"

	"github.com/olehvolynets/sylphy/render"
)

type ArrayFormatter struct {
	r *render.Renderer
}

func NewArrayFormatter(r *render.Renderer) *ArrayFormatter {
	return &ArrayFormatter{r: r}
}

func (f *ArrayFormatter) Format(v any) {
	val, ok := v.([]any)
	if !ok {
		panic("wtf?")
	}

	f.r.WriteOperator("[")
	f.r.Endl()
	f.r.Indent()

	for idx, item := range val {
		f.r.WriteOffset()

		var formatter Formatter
		switch item.(type) {
		case int64, float64:
			formatter = &NumberFormatter{f.r}
		case string:
			formatter = &StringFormatter{f.r}
		case bool:
			formatter = &BoolFormatter{f.r}
		case nil:
			formatter = &NullFormatter{f.r}
		default:
			switch reflect.TypeOf(item).Kind() {
			case reflect.Slice:
				formatter = &ArrayFormatter{f.r}
			case reflect.Map:
				formatter = &ObjectFormatter{f.r}
			default:
				formatter = &ErrorFormatter{f.r}
			}
		}

		formatter.Format(item)

		if idx < len(val)-1 {
			f.r.WriteElSeparator()
			f.r.Endl()
		}
	}

	f.r.Outdent()
	f.r.Endl()
	f.r.WriteOffset()
	f.r.WriteOperator("]")
}
