package render

import (
	"errors"

	"github.com/fatih/color"
)

var (
	ErrTypeMismatch = errors.New("type mismatch")
	ErrMissing      = errors.New("value missing")
)

type Handler interface {
	Handle(*Context) error
}

type HandlerBuilder func(key string, optional bool, colorizer Colorizer) Handler

type handler[T any] struct {
	Accessor[T]
	formatter Colorizer
}

type Entry map[string]any

type Colorizer = *color.Color
