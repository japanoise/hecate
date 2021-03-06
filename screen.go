package main

import "github.com/nsf/termbox-go"

type Screen interface {
	receiveEvents(input <-chan termbox.Event, output chan<- interface{}, quit <-chan bool)
	performLayout()
	drawScreen(style Style)
}

const (
	ABOUT_SCREEN_INDEX = iota
	PALETTE_SCREEN_INDEX
	DATA_SCREEN_INDEX
	EXIT_SCREEN_INDEX
)

func defaultScreensForFiles(files []FileInfo) []Screen {
	data_screen := DataScreen{}
	data_screen.initializeWithFiles(files)

	about_screen := AboutScreen{}
	palette_screen := PaletteScreen(0)
	screens := [...]Screen{
		&about_screen,
		&palette_screen,
		&data_screen,
	}

	return screens[:]
}

func drawBackground(bg termbox.Attribute) {
	termbox.Clear(0, bg)
}

func layoutAndDrawScreen(screen Screen, style Style) {
	screen.performLayout()
	drawBackground(style.Default_bg)
	screen.drawScreen(style)
	termbox.Flush()
}
