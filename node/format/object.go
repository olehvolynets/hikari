package format

import (
	"reflect"
	"slices"

	"github.com/olehvolynets/sylphy/render"
)

type ObjectFormatter struct {
	r *render.Renderer
}

func NewObjectFormatter(r *render.Renderer) *ObjectFormatter {
	return &ObjectFormatter{r: r}
}

func (f *ObjectFormatter) Format(v any) {
	val, ok := v.(map[string]any)
	if !ok {
		panic("wtf?")
	}

	keys := make([]string, 0, len(val))
	for key := range val {
		keys = append(keys, key)
	}
	slices.Sort(keys)

	keyFormatter := &StringFormatter{f.r}

	f.r.WriteOperator("{")
	f.r.Endl()
	f.r.Indent()

	for idx, key := range keys {
		f.r.WriteOffset()

		keyFormatter.Format(key)
		f.r.WriteKvSeparator()

		item := val[key]

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
	f.r.WriteOperator("}")
}
