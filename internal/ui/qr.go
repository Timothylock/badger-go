package ui

import (
	"fmt"

	qrcode "github.com/skip2/go-qrcode"
	"tinygo.org/x/drivers/uc8151"
	"tinygo.org/x/tinydraw"
)

func DrawQR(display *uc8151.Device, msg string, x, y, w, h int16) error {
	qr, err := qrcode.New(msg, qrcode.Medium)
	if err != nil {
		return err
	}

	qrBytes := qr.Bitmap()
	size := int16(len(qrBytes))

	factor := int16(int(h) / len(qrBytes))

	bx := (w - size*factor) / 2
	by := (h - size*factor) / 2

	fmt.Println(size)
	for i := int16(0); i < size; i++ {
		for j := int16(0); j < size; j++ {
			if qrBytes[i][j] {
				tinydraw.FilledRectangle(display, x+bx+j*factor, y+by+i*factor, factor, factor, ColourBlack())
			} else {
				tinydraw.FilledRectangle(display, x+bx+j*factor, y+by+i*factor, factor, factor, ColourWhite())
			}
		}
	}

	return nil
}
