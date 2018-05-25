// +build windows

package main

import (
	"github.com/nsf/termbox-go"
)

func handleSpecialKeys(key termbox.Key) {}

const outputMode = termbox.OutputNormal
const inputMode = termbox.InputAlt

func defaultStyle() *Style {
	var style Style
	style.Default_bg = termbox.ColorBlack
	style.Default_fg = termbox.ColorWhite
	style.Rune_fg = termbox.ColorYellow
	style.Int_fg = termbox.ColorCyan
	style.Bit_fg = termbox.ColorCyan
	style.Space_rune_fg = termbox.ColorWhite
	style.Selected_option_bg = termbox.ColorBlue
	style.Search_progress_fg = termbox.ColorCyan

	style.Text_cursor_hex_bg = termbox.ColorRed
	style.Bit_cursor_hex_bg = termbox.ColorCyan
	style.Int_cursor_hex_bg = termbox.ColorCyan
	style.Fp_cursor_hex_bg = termbox.ColorRed

	style.Hilite_hex_fg = termbox.ColorMagenta
	style.Hilite_rune_fg = termbox.ColorMagenta

	style.About_logo_bg = termbox.ColorRed

	style.Field_editor_bg = style.Default_fg
	style.Field_editor_fg = style.Default_bg

	style.Field_editor_last_bg = style.Rune_fg
	style.Field_editor_last_fg = style.Default_fg

	style.Field_editor_invalid_bg = termbox.ColorRed
	style.Field_editor_invalid_fg = style.Rune_fg

	style.Space_rune = '•'
	style.Filled_bit_rune = '1'
	style.Empty_bit_rune = '0'

	return &style
}
