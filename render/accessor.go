package render

type Accessor[T any] struct {
	Key      string
	Optional bool
}

func (a *Accessor[T]) Get(ctx *Context) (T, bool, error) {
	var value T
	anyValue, ok := ctx.Entry[a.Key]
	if !ok {
		if a.Optional {
			return value, true, nil
		}
		return value, false, ErrMissing
	}

	value, ok = anyValue.(T)
	if !ok {
		return value, false, ErrTypeMismatch
	}

	return value, false, nil
}
