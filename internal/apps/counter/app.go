package counter

import (
	"embed"
	_ "embed"
	"machine"
	"strconv"
	"time"

	"github.com/timothylock/badger-go/internal/apps"
	"github.com/timothylock/badger-go/internal/ui"

	"tinygo.org/x/drivers/uc8151"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

const appName = "Counter"

//go:embed icon.png
var appIconFS embed.FS

type Counter struct {
	display uc8151.Device
}

func NewCounterApp(display *uc8151.Device) apps.Application {
	return &Counter{display: *display}
}

func (h *Counter) GetAppConfig() apps.AppConfig {
	return apps.AppConfig{
		Name: appName,
		Icon: appIconFS,
	}
}

func (h *Counter) Run() error {
	i := 0
	lastI := 0
	refresh := true

	// Initialize Buttons
	btnA := machine.BUTTON_A
	btnUp := machine.BUTTON_UP
	btnDown := machine.BUTTON_DOWN
	btnA.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	btnUp.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	btnDown.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

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
			h.display.ClearBuffer()
			h.display.WaitUntilIdle()
			ui.TopNavBar(&h.display, apps.OSName, h.GetAppConfig().Name, apps.OSVersion)
			ui.BottomNavBar(&h.display, "[a] Back", "", "")

			// Draw buttons
			tinyfont.WriteLine(&h.display, &freesans.Regular9pt7b, 270, 40, "(+)", ui.ColourBlack())
			tinyfont.WriteLine(&h.display, &freesans.Regular9pt7b, 272, 100, "(-)", ui.ColourBlack())

			// Write the last number in white to remove it from the screen. More efficient than redrawing all the elements
			tinyfont.WriteLine(&h.display, &freesans.Bold24pt7b, 40, 80, "Count: "+strconv.Itoa(lastI), ui.ColourWhite())
			// Write the new number
			tinyfont.WriteLine(&h.display, &freesans.Bold24pt7b, 40, 80, "Count: "+strconv.Itoa(i), ui.ColourBlack())

			// Show the buffer on the screen
			h.display.Display()
			h.display.WaitUntilIdle()
			refresh = false
		}

		time.Sleep(200 * time.Millisecond)
	}
}
