package ircformat

import (
	"bytes"
	"fmt"
)

// See https://github.com/myano/jenni/wiki/IRC-String-Formatting

const (
	escapeBold = '\x02'
	escapeColor = '\x03'
	escapeItalic = '\x1D'
	escapeUnderline = '\x1F'
	escapeInvert = '\x16'
	escapeReset = '\x0F'
)

type Color int

const (
	ColorWhite Color = iota
	ColorBlack
	ColorBlue // Navy
	ColorGreen
	ColorRed
	ColorBrown // Maroon
	ColorPurple
	ColorOrange // Olive
	ColorYellow
	ColorLightGreen // Lime
	ColorTeal // Green/blue cyan
	ColorLightCyan // Cyan/aqua
	ColorLightBlue // Royal
	ColorPink // Light purple/fuchsia
	ColorGrey
	ColorLightGrey // Silver
	ColorDefault = 99
)

type Buffer struct {
	Buffer bytes.Buffer
	lastEscape rune
}

func (b *Buffer) Append(s string) *Buffer {
	b.Buffer.WriteString(s)
	return b
}

func (b *Buffer) appendEscape(esc rune) *Buffer {
	b.lastEscape = esc
	b.Buffer.WriteRune(esc)
	return b
}

func (b *Buffer) Bold() *Buffer {
	return b.appendEscape(escapeBold)
}

func (b *Buffer) Color(foreground, background Color) *Buffer {
	b.appendEscape(escapeColor)
	b.Buffer.WriteString(fmt.Sprintf("%02d,%02d", foreground, background))
	return b
}

func (b *Buffer) Italic() *Buffer {
	return b.appendEscape(escapeItalic)
}

func (b *Buffer) Underline() *Buffer {
	return b.appendEscape(escapeUnderline)
}

func (b *Buffer) Invert() *Buffer {
	return b.appendEscape(escapeInvert)
}

func (b *Buffer) String() string {
	if b.lastEscape != escapeReset {
		b.appendEscape(escapeReset)
	}
	return b.Buffer.String()
}
