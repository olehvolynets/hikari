package sylphy

import (
	"io"
	"strings"
)

type Context struct {
	Entry

	IndentLevel int
	IndentChar  string

	W io.Writer
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
