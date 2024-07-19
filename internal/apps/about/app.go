package about

import (
	"embed"
	_ "embed"
	"github.com/timothylock/badger-go/internal/apps"
	"github.com/timothylock/badger-go/internal/ui"
	"machine"
	"time"

	"tinygo.org/x/drivers/uc8151"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

const appName = "About"

//go:embed icon.png
var appIconFS embed.FS

type About struct {
	display uc8151.Device
}

func NewAboutApp(display *uc8151.Device) apps.Application {
	return &About{display: *display}
}

func (h *About) GetAppConfig() apps.AppConfig {
	return apps.AppConfig{
		Name: appName,
		Icon: appIconFS,
	}
}

func (h *About) Run() error {
	h.display.ClearBuffer()
	h.display.WaitUntilIdle()
	ui.TopNavBar(&h.display, apps.OSName, h.GetAppConfig().Name, apps.OSVersion)
	ui.BottomNavBar(&h.display, "[a] Back", "", "")

	// Initialize Buttons
	btnA := machine.BUTTON_A
	btnA.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	// Draw buttons
	tinyfont.WriteLine(&h.display, &freesans.Bold12pt7b, 5, 50, apps.OSName, ui.ColourBlack())
	tinyfont.WriteLine(&h.display, &freesans.Bold9pt7b, 5, 70, apps.OSVersion, ui.ColourBlack())
	tinyfont.WriteLine(&h.display, &freesans.Regular9pt7b, 5, 100, "Created by @timothylock (Github)", ui.ColourBlack())

	h.display.Display()
	h.display.WaitUntilIdle()
	for {
		if btnA.Get() {
			return nil
		}

		time.Sleep(100 * time.Millisecond)
	}
}
