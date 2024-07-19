package badge

import (
	"embed"
	_ "embed"
	"machine"
	"time"

	"github.com/timothylock/badger-go/internal/apps"
	"github.com/timothylock/badger-go/internal/ui"

	"tinygo.org/x/drivers/uc8151"
	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
	"tinygo.org/x/tinyfont/proggy"
)

const appName = "Badge"

var (
	company = "Electronic Arts"
	name    = "Tim L"
	title   = "Software Engineer"
	social  = "@timothylock"
	qrCode  = "https://timlock.dev"
)

//go:embed icon.png
var appIconFS embed.FS

type Badge struct {
	display uc8151.Device
}

func NewBadgeApp(display *uc8151.Device) apps.Application {
	return &Badge{display: *display}
}

func (b *Badge) GetAppConfig() apps.AppConfig {
	return apps.AppConfig{
		Name: appName,
		Icon: appIconFS,
	}
}

func (b *Badge) Run() error {
	b.display.ClearBuffer()
	b.display.WaitUntilIdle()

	// Initialize Buttons
	btnA := machine.BUTTON_A
	btnA.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	// Draw Badge
	b.display.ClearBuffer()
	b.display.WaitUntilIdle()
	ui.DrawQR(&b.display, qrCode, 172, 0, 128, 128)
	b.drawBadge(b.display)
	b.display.Display()
	b.display.WaitUntilIdle()

	for {
		if btnA.Get() {
			return nil
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func (b *Badge) drawBadge(display uc8151.Device) {
	midW := int16(176)
	//if profileErr == nil {
	//	display.DrawBuffer(0, 0, 128, 120, []uint8(profileImg))
	//}

	tinydraw.FilledRectangle(&display, 0, 0, midW, 30, ui.ColourBlack())
	tinydraw.Line(&display, 0, 0, 295, 0, ui.ColourBlack())
	tinydraw.Line(&display, 0, 0, 0, 127, ui.ColourBlack())
	tinydraw.Line(&display, 295, 0, 295, 127, ui.ColourBlack())
	tinydraw.Line(&display, midW, 0, midW, 127, ui.ColourBlack())
	tinydraw.Line(&display, 0, 87, midW, 87, ui.ColourBlack())
	tinydraw.Line(&display, 0, 107, midW, 107, ui.ColourBlack())
	tinydraw.Line(&display, 0, 127, 295, 127, ui.ColourBlack())

	tinyfont.WriteLine(&display, &freesans.Bold9pt7b, 8, 22, company, ui.ColourWhite())

	w32, _ := tinyfont.LineWidth(&freesans.Bold24pt7b, name)
	if w32 < uint32(midW) {
		tinyfont.WriteLine(&display, &freesans.Bold24pt7b, (midW-int16(w32))/2, 74, name, ui.ColourBlack())
	} else {
		w32, _ := tinyfont.LineWidth(&freesans.Bold18pt7b, name)
		if w32 < uint32(midW) {
			tinyfont.WriteLine(&display, &freesans.Bold18pt7b, (midW-int16(w32))/2, 74, name, ui.ColourBlack())
		} else {
			w32, _ := tinyfont.LineWidth(&freesans.Bold12pt7b, name)
			if w32 < uint32(midW) {
				tinyfont.WriteLine(&display, &freesans.Bold12pt7b, (midW-int16(w32))/2, 74, name, ui.ColourBlack())
			} else {
				w32, _ := tinyfont.LineWidth(&freesans.Bold9pt7b, name)
				tinyfont.WriteLine(&display, &freesans.Bold9pt7b, (midW-int16(w32))/2, 74, name, ui.ColourBlack())
			}
		}
	}

	w32, _ = tinyfont.LineWidth(&freesans.Regular9pt7b, title)
	tinyfont.WriteLine(&display, &freesans.Regular9pt7b, (midW-int16(w32))/2, 102, title, ui.ColourBlack())

	w32, _ = tinyfont.LineWidth(&freesans.Regular9pt7b, social)
	if w32 < uint32(midW) {
		tinyfont.WriteLine(&display, &freesans.Regular9pt7b, (midW-int16(w32))/2, 122, social, ui.ColourBlack())
	} else {
		w32, _ := tinyfont.LineWidth(&proggy.TinySZ8pt7b, social)
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, (midW-int16(w32))/2, 120, social, ui.ColourBlack())
	}
}
