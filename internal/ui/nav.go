package ui

import (
	"tinygo.org/x/drivers/uc8151"
	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/proggy"
)

// TopNavBar creates a top navigation bar with left, center, and right text
func TopNavBar(disp *uc8151.Device, left, center, right string) {
	tinydraw.FilledRectangle(disp, 2, 2, 292, 14, ColourBlack())

	tinyfont.WriteLine(disp, &proggy.TinySZ8pt7b, 6, 12, left, ColourWhite())
	tinyfont.WriteLine(disp, &proggy.TinySZ8pt7b, int16(146-(len(center)*3)), 12, center, ColourWhite())
	tinyfont.WriteLine(disp, &proggy.TinySZ8pt7b, int16(290-(len(right)*6)), 12, right, ColourWhite())
}

// BottomNavBar creates a top navigation bar with left, center, and right text
func BottomNavBar(disp *uc8151.Device, left, center, right string) {
	tinydraw.FilledRectangle(disp, 2, 112, 292, 14, ColourGrey())

	tinyfont.WriteLine(disp, &proggy.TinySZ8pt7b, 6, 122, left, ColourWhite())
	tinyfont.WriteLine(disp, &proggy.TinySZ8pt7b, int16(146-(len(center)*3)), 122, center, ColourWhite())
	tinyfont.WriteLine(disp, &proggy.TinySZ8pt7b, int16(290-(len(right)*6)), 122, right, ColourWhite())
}
