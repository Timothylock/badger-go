package ui

import "image/color"

func ColourWhite() color.RGBA {
	return color.RGBA{R: 0, G: 0, B: 0, A: 255}
}

func ColourBlack() color.RGBA {
	return color.RGBA{R: 1, G: 1, B: 1, A: 255}
}
