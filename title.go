package uiutil

import (
	"bytes"
)

// TitleUnderliner is the underline character for the title
var TitleUnderliner = "="

// Title is a UI component that renders a title
type Title struct {
	text string
}

func NewTitle(text string) *Title {
	return &Title{text: text}
}

// String returns the formated string of the title
func (t *Title) Bytes() []byte {
	var buf bytes.Buffer
	buf.WriteString(t.text + "\n")
	for i := 0; i < len(t.text); i++ {
		buf.Write([]byte(TitleUnderliner))
	}
	buf.WriteString("\n")
	return buf.Bytes()
}
