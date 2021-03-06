package main

import (
	"fmt"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

type Command struct {
	key         string
	description string
}

type AboutScreen struct {
	show_html bool
}

func drawCommandsAtPoint(commands []Command, x int, y int, style Style) {
	x_pos, y_pos := x, y
	longest_key_len := 1
	for _, cmd := range commands {
		if len(cmd.key) > longest_key_len {
			longest_key_len = len(cmd.key)
		}
	}
	for index, cmd := range commands {
		drawStringAtPoint(fmt.Sprintf("%[2]*[1]s", cmd.key, longest_key_len), x_pos, y_pos, style.Default_fg, style.Default_bg)
		drawStringAtPoint(cmd.description, x_pos+longest_key_len+2, y_pos, style.Default_fg, style.Default_bg)
		y_pos++
		if index > 2 && index%2 == 1 {
			y_pos++
		}
	}
}

func (screen *AboutScreen) receiveEvents(input <-chan termbox.Event, output chan<- interface{}, quit <-chan bool) {
	for {
		do_quit := false
		select {
		case event := <-input:
			if event.Key == termbox.KeyCtrlR {
				screen.show_html = !screen.show_html
				output <- ScreenIndex(ABOUT_SCREEN_INDEX)
			} else {
				screen.show_html = false
				output <- ScreenIndex(DATA_SCREEN_INDEX)
			}
		case <-quit:
			do_quit = true
		}
		if do_quit {
			break
		}
	}
}

func (screen *AboutScreen) performLayout() {
}

func (screen *AboutScreen) drawScreen(style Style) {
	default_fg := style.Default_fg
	default_bg := style.Default_bg
	width, height := termbox.Size()
	template := [...]string{
		"                              ############################",
		"                              ##The#hex#editor#from#hell##",
		"                              ############################",
		"                                      ####            #   ",
		"#### #### ########  #####      ###    #### ########   #   ",
		"#### #### ####### #########  #######  #### ########  #### ",
		"#### #### ####    #### #### #### #### #### ####     ##x#x ",
		"#### #### ####    ####      #### #### #### ####       #   ",
		"######### ####### ####      ######### #### #######   ###  ",
		"######### ####### ####      ######### #### #######  # # # ",
		"#### #### ####    ####      #### #### #### ####    #  #  #",
		"#### #### ####    ####      #### #### #### ####      # #  ",
		"#### #### ####    #### #### #### #### #### ####     #   # ",
		"#### #### ####### ######### #### #### #### ####### #     #",
		"#### #### ########  #####   #### #### #### ########       ",
	}
	commands1 := [...]Command{
		{"C-b", "left"},
		{"C-n", "down"},
		{"C-p", "up"},
		{"C-f", "right"},

		{"M-b", "left 4 bytes"},
		{"M-f", "right 4 bytes"},

		{"C-a", "line start"},
		{"C-e", "line end"},

		{"M-<", "file start"},
		{"M->", "file end"},

		{"M-g", "jump to byte"},
		{"x", "toggle hex"},
	}

	commands2 := [...]Command{
		{"t", "text mode"},
		{"p", "bit pattern mode"},
		{"i", "integer mode"},
		{"f", "float mode"},

		{"H", "shrink cursor"},
		{"L", "grow cursor"},

		{"e", "toggle endianness"},
		{"u", "toggle signedness"},

		{"a", "date decoding"},
		{"@", "set date epoch"},

		{"C-s", "search file"},
		{"n", "next match"},
	}

	commands3 := [...]Command{
		{"S", "show tabs"},
		{"W", "hide tabs"},

		{"A", "previous tab"},
		{"D", "next tab"},

		{"ctrl-t", "new tab"},
		{"ctrl-w", "close tab"},

		{"M-+", "scroll down"},
		{"M--", "scroll up"},

		{"C-v", "page down"},
		{"M-v", "page up"},

		{"C-l", "centre view"},
		{"enter", "edit mode"},
		{"?", "this screen"},
		{"q", "quit program"},
	}

	first_line := template[0]
	start_x := (width - len(first_line)) / 2
	start_y := (height-len(template)-2-len(commands2)/2*3-1)/2 + 1
	x_pos := start_x
	y_pos := start_y
	for _, line := range template {
		x_pos = start_x
		for _, runeValue := range line {
			bg := default_bg
			displayRune := ' '
			if runeValue != ' ' {
				bg = style.About_logo_bg
				if runeValue != '#' {
					displayRune = runeValue
				}
				termbox.SetCell(x_pos, y_pos, displayRune, default_fg, bg)
			}
			x_pos += runewidth.RuneWidth(displayRune)
		}
		y_pos++
	}
	x_pos = start_x
	y_pos++

	if screen.show_html {
		drawStringAtPoint("<table>", 0, y_pos, style.Default_fg, style.Default_bg)
		y_pos++
		for i := 0; i < len(commands1); i++ {
			x_pos = 0
			x_pos += drawStringAtPoint("<tr>", x_pos, y_pos, style.Default_fg, style.Default_bg)
			for _, cmd := range [...]Command{commands1[i], commands2[i], commands3[i]} {
				x_pos += drawStringAtPoint(fmt.Sprintf("<td>%s</td>", cmd.key), x_pos, y_pos, style.Default_fg, style.Default_bg)
				if cmd.description == "this screen" {
					x_pos += drawStringAtPoint(fmt.Sprintf("<td>%s</td>", "help screen"), x_pos, y_pos, style.Default_fg, style.Default_bg)
				} else {
					x_pos += drawStringAtPoint(fmt.Sprintf("<td>%s</td>", cmd.description), x_pos, y_pos, style.Default_fg, style.Default_bg)
				}
			}
			x_pos += drawStringAtPoint("</tr>", x_pos, y_pos, style.Default_fg, style.Default_bg)
			y_pos++
		}
		drawStringAtPoint("</table>", 0, y_pos, style.Default_fg, style.Default_bg)
		y_pos++
	} else {
		drawCommandsAtPoint(commands1[:], x_pos, y_pos+1, style)
		drawCommandsAtPoint(commands2[:], x_pos+19, y_pos+1, style)
		drawCommandsAtPoint(commands3[:], x_pos+42, y_pos+1, style)
	}
}
