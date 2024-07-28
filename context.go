package hikari

import (
	"io"
	"strings"
)

type Context struct {
	IndentLevel int
	IndentChar  string

	W io.Writer

	HandledAttributes []string
}

func (ctx *Context) Indent() {
	ctx.IndentLevel += 1
}

func (ctx *Context) Dedent() {
	ctx.IndentLevel -= 1
}

func (ctx *Context) CurrentIndent() string {
	return strings.Repeat(ctx.IndentChar, ctx.IndentLevel)
}

func (ctx *Context) AddHandled(key string) {
	ctx.HandledAttributes = append(ctx.HandledAttributes, key)
}
