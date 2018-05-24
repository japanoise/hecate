// +build !windows,!linux

package main

import (
	"os"
	"syscall"

	"github.com/nsf/termbox-go"
)

func handleSpecialKeys(key termbox.Key) {
	if key == termbox.KeyCtrlZ {
		process, _ := os.FindProcess(os.Getpid())
		termbox.Close()
		process.Signal(syscall.SIGSTOP)
		termbox.Init()
	}
}

const outputMode = termbox.Output256
const inputMode = termbox.InputAlt

func defaultStyle() *Style {
	var style Style
	style.Default_bg = termbox.Attribute(1)
	style.Default_fg = termbox.Attribute(256)
	style.Rune_fg = termbox.Attribute(248)
	style.Int_fg = termbox.Attribute(154)
	style.Bit_fg = termbox.Attribute(154)
	style.Space_rune_fg = termbox.Attribute(240)
	style.Selected_option_bg = termbox.Attribute(240)
	style.Search_progress_fg = termbox.Attribute(76)

	style.Text_cursor_hex_bg = termbox.Attribute(167)
	style.Bit_cursor_hex_bg = termbox.Attribute(26)
	style.Int_cursor_hex_bg = termbox.Attribute(63)
	style.Fp_cursor_hex_bg = termbox.Attribute(127)

	style.Hilite_hex_fg = termbox.Attribute(231)
	style.Hilite_rune_fg = termbox.Attribute(256)

	style.About_logo_bg = termbox.Attribute(125)

	style.Field_editor_bg = style.Default_fg
	style.Field_editor_fg = style.Default_bg

	style.Field_editor_last_bg = style.Rune_fg
	style.Field_editor_last_fg = style.Default_fg

	style.Field_editor_invalid_bg = termbox.Attribute(125)
	style.Field_editor_invalid_fg = style.Rune_fg

	style.Space_rune = '•'
	style.Filled_bit_rune = '●'
	style.Empty_bit_rune = '○'

	return &style
}
