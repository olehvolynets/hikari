package render

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/fatih/color"
)

var operatorColor = color.New(color.FgHiWhite)

type Renderer struct {
	Cfg  Config
	Sink io.Writer
}

func NewRenderer() *Renderer {
	return &Renderer{
		Cfg: Config{
			IndentChar:       "\t",
			ElementSeparator: ",",
			LineSeparator:    "\n",
		},
		Sink: os.Stdout,
	}
}

func (r *Renderer) WithConfig(c Config) *Renderer {
	return &Renderer{
		Cfg:  c,
		Sink: r.Sink,
	}
}

func (r *Renderer) Indent() {
	r.Cfg.IndentLevel += 1
}

func (r *Renderer) Outdent() {
	if r.Cfg.IndentLevel > 0 {
		r.Cfg.IndentLevel -= 1
	}
}

func (r *Renderer) Write(b []byte) (int, error) {
	return r.Sink.Write(b)
}

func (r *Renderer) Offset() string {
	return strings.Repeat(r.Cfg.IndentChar, int(r.Cfg.IndentLevel))
}

func (r *Renderer) WriteOperator(op string) (int, error) {
	return operatorColor.Fprint(r, op)
}

func (r *Renderer) WriteOffset() (int, error) {
	return r.WriteOperator(r.Offset())
}

func (r *Renderer) WriteElSeparator() (int, error) {
	return r.WriteOperator(r.Cfg.ElementSeparator)
}

func (r *Renderer) WriteKvSeparator() (int, error) {
	return r.WriteOperator(r.Cfg.KeyValSeparator)
}

func (r *Renderer) Endl() (int, error) {
	return r.WriteOperator(r.Cfg.LineSeparator)
}

func (r *Renderer) Print(v ...any) (int, error) { return fmt.Fprint(r.Sink, v...) }

func (r *Renderer) Printf(format string, v ...any) (int, error) {
	return fmt.Fprintf(r.Sink, format, v...)
}

func (r *Renderer) Println(v ...any) (int, error) {
	a, _ := r.Write([]byte(r.Offset()))
	b, _ := r.Print(v...)
	c, _ := r.Endl()

	return a + b + c, nil
}
