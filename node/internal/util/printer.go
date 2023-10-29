package util

type (
	PrinterFunc func(map[string]any) (int, error)

	Printer struct {
		written int
		err     error
	}
)

func (p *Printer) Print(f PrinterFunc, v map[string]any) {
	if p.err != nil {
		return
	}

	i, err := f(v)
	if err != nil {
		p.err = err
		return
	}

	p.written += i
}

func (p *Printer) Result() (int, error) { return p.written, p.err }
