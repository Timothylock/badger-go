package counter

import (
	_ "embed"
	"github.com/timothylock/badger-go/internal/ui"
	"image/color"
	"machine"
	"strconv"
	"time"

	"github.com/timothylock/badger-go/internal/apps"

	"tinygo.org/x/drivers/uc8151"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

const appName = "Counter"

//go:embed icon.png
var appIcon []byte

type Counter struct {
	display uc8151.Device
}

func NewCounterApp(display *uc8151.Device) apps.Application {
	return &Counter{display: *display}
}

func (h *Counter) GetAppConfig() apps.AppConfig {
	return apps.AppConfig{
		Name: appName,
		Icon: appIcon,
	}
}

func (h *Counter) Run() error {
	i := 0
	lastI := 0
	refresh := true

	h.display.ClearBuffer()
	ui.TopNavBar(&h.display, apps.OSName, h.GetAppConfig().Name, apps.OSVersion)
	ui.BottomNavBar(&h.display, "[a] Back", "", "")

	// Initialize Buttons
	btnA := machine.BUTTON_A
	btnUp := machine.BUTTON_UP
	btnDown := machine.BUTTON_DOWN
	btnA.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	btnUp.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	btnDown.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	// Draw buttons
	tinyfont.WriteLine(&h.display, &freesans.Regular9pt7b, 270, 40, "(+)", color.RGBA{R: 1, G: 1, B: 1, A: 255})
	tinyfont.WriteLine(&h.display, &freesans.Regular9pt7b, 272, 100, "(-)", color.RGBA{R: 1, G: 1, B: 1, A: 255})

	for {
		if btnA.Get() {
			return nil
		} else if btnUp.Get() {
			lastI = i
			i++
			refresh = true
		} else if btnDown.Get() {
			lastI = i
			i--
			refresh = true
		}

		if refresh {
			// Write the last number in white to remove it from the screen. More efficient than redrawing all the elements
			tinyfont.WriteLine(&h.display, &freesans.Bold24pt7b, 40, 80, "Count: "+strconv.Itoa(lastI), color.RGBA{R: 0, G: 0, B: 0, A: 255})
			// Write the new number
			tinyfont.WriteLine(&h.display, &freesans.Bold24pt7b, 40, 80, "Count: "+strconv.Itoa(i), color.RGBA{R: 1, G: 1, B: 1, A: 255})

			// Show the buffer on the screen
			h.display.Display()
			refresh = false
		}

		time.Sleep(200 * time.Millisecond)
	}
}
