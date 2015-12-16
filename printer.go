package uiutil

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// Out is the default out
var Out = os.Stdout

// Component in the interface that UI components need to implement
type Component interface {
	Bytes() []byte
}

// Printer represents a buffered container for components that can be flushed
type Printer struct {
	out   io.Writer
	comps []Component
}

// NewPrinter returns a pointer to a new printer object
func NewPrinter(out io.Writer) *Printer {
	if out == nil {
		out = Out
	}
	return &Printer{out: out}
}

// Add adds the components to the printer
func (p *Printer) Add(c Component) *Printer {
	p.comps = append(p.comps, c)
	return p
}

// AddTitle Adds a Title to the printer
func (p *Printer) AddTitle(title string) *Printer {
	return p.Add(&Title{text: title})
}

// Bytes returns the formmated string of the output
func (p *Printer) Bytes() []byte {
	var buf bytes.Buffer
	for _, c := range p.comps {
		buf.Write(c.Bytes())
		buf.Write([]byte{'\n'})
	}
	return buf.Bytes()
}

// Flush prints the output to the writer and clears the buffer
func (p *Printer) Flush() error {
	_, err := fmt.Fprintln(p.out, p)
	if err != nil {
		return err
	}
	p.comps = nil
	return nil
}
