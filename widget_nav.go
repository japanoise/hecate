package main

type NavigationWidget int

const (
	navStrLong      string = "Navigate: ←^B ↓^N ↑^P →^F"
	navStrShort            = "Navigate: ←^B ↓^N"
	navStrShortMore        = "↑^P →^F"
	navStrWords            = "←←←←M-b M-f→→→→"
)

func (widget NavigationWidget) sizeForLayout(layout Layout) Size {
	if layout.pressure > 1 {
		return Size{0, 0}
	}
	layouts := map[int]string{
		0: navStrLong,
		1: navStrShort,
	}
	runeCount := 0
	for _, _ = range layouts[layout.pressure] {
		runeCount++
	}
	return Size{runeCount, 2}
}

func (widget NavigationWidget) drawAtPoint(tab *DataTab, layout Layout, point Point, style Style) Size {
	fg := style.Default_fg
	bg := style.Default_bg
	x_pos := point.x
	if layout.pressure == 0 {
		x_pos += drawStringAtPoint(navStrLong, x_pos, point.y, fg, bg)
		x_pos = point.x + 10
		x_pos += drawStringAtPoint(navStrWords, x_pos, point.y+1, fg, bg)
	} else if layout.pressure == 1 {
		x_pos += drawStringAtPoint(navStrShort, x_pos, point.y, fg, bg)
		x_pos = point.x + 10
		x_pos += drawStringAtPoint(navStrShortMore, x_pos, point.y+1, fg, bg)
	}
	return Size{x_pos - point.x, 2}
}
