package home

import (
	"embed"
	"fmt"
	_ "github.com/nfnt/resize"
	"image/color"
	"image/png"
	"machine"
	"time"
	"tinygo.org/x/drivers/uc8151"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/proggy"

	"github.com/timothylock/badger-go/internal/apps"
	"github.com/timothylock/badger-go/internal/ui"
)

//go:embed icon.png
var appIconFS embed.FS

// Home represents the home screen
type Home struct {
	display      uc8151.Device
	applications []apps.Application
}

// NewHome creates a new home screen
func NewHome(display *uc8151.Device) Home {
	return Home{display: *display}
}

// AddApp adds an application to the home screen
func (h *Home) AddApp(app apps.Application) error {
	h.applications = append(h.applications, app)
	return nil
}

func (h *Home) GetAppConfig() apps.AppConfig {
	return apps.AppConfig{
		Name: "Home",
		Icon: appIconFS,
	}
}

// Run should execute the app. It should not return until the app has finished executing.
func (h *Home) Run() error {
	h.display.ClearBuffer()
	h.display.WaitUntilIdle()
	refresh := true

	// Initialize Buttons
	btnA := machine.BUTTON_A
	btnB := machine.BUTTON_B
	btnC := machine.BUTTON_C
	btnA.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	btnB.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	btnC.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	for {
		curPage := h.applications[0:3]

		if btnA.Get() {
			err := h.launchApp(curPage, 0, btnA)
			if err != nil {
				return fmt.Errorf("%s has crashed: %v", curPage[0].GetAppConfig().Name, err)
			}
			waitForAllButtonsToLetGo()
			refresh = true
		} else if btnB.Get() {
			err := h.launchApp(curPage, 1, btnB)
			if err != nil {
				return fmt.Errorf("%s has crashed: %v", curPage[1].GetAppConfig().Name, err)
			}
			waitForAllButtonsToLetGo()
			refresh = true
		} else if btnC.Get() {
			err := h.launchApp(curPage, 2, btnC)
			if err != nil {
				return fmt.Errorf("%s has crashed: %v", curPage[2].GetAppConfig().Name, err)
			}
			waitForAllButtonsToLetGo()
			refresh = true
		}

		if refresh {
			// Draw Apps
			cfg := curPage[0].GetAppConfig()
			err := drawIconResized(&h.display, 18, 40, 60, 60, cfg.Icon)
			if err != nil {
				return err
			}
			tinyfont.WriteLine(&h.display, &proggy.TinySZ8pt7b, int16(48-(len(cfg.Name)*3)), 110, cfg.Name, ui.ColourBlack())

			cfg = curPage[1].GetAppConfig()
			err = drawIconResized(&h.display, 118, 40, 60, 60, cfg.Icon)
			if err != nil {
				return err
			}
			tinyfont.WriteLine(&h.display, &proggy.TinySZ8pt7b, int16(149-(len(cfg.Name)*3)), 110, cfg.Name, ui.ColourBlack())

			cfg = curPage[2].GetAppConfig()
			err = drawIconResized(&h.display, 218, 40, 60, 60, cfg.Icon)
			if err != nil {
				return err
			}
			tinyfont.WriteLine(&h.display, &proggy.TinySZ8pt7b, int16(248-(len(cfg.Name)*3)), 110, cfg.Name, ui.ColourBlack())

			ui.TopNavBar(&h.display, apps.OSName, h.GetAppConfig().Name, apps.OSVersion)
			h.display.Display()
			h.display.WaitUntilIdle()
			refresh = false
		}

		time.Sleep(200 * time.Millisecond)
	}
}

func (h *Home) launchApp(curPage []apps.Application, i int, btn machine.Pin) error {
	// Make sure there is an app in the slot
	if len(curPage) <= 0 {
		return nil
	}

	// TODO: INTERRUPTS

	// Wait until the user lets go of the button
	for btn.Get() {
		time.Sleep(50 * time.Millisecond)
	}

	h.display.ClearBuffer()
	h.display.ClearDisplay()
	h.display.WaitUntilIdle()
	err := curPage[i].Run()
	h.display.ClearBuffer()
	h.display.ClearDisplay()
	h.display.WaitUntilIdle()
	return err
}

func waitForAllButtonsToLetGo() {
	for machine.BUTTON_A.Get() || machine.BUTTON_B.Get() || machine.BUTTON_C.Get() {
		time.Sleep(50 * time.Millisecond)
	}
}

func drawIconResized(display *uc8151.Device, x, y, w, h int16, iconFS embed.FS) error {
	iconReader, err := iconFS.Open("icon.png")
	if err != nil {
		return err
	}
	defer iconReader.Close()

	img, err := png.Decode(iconReader)
	if err != nil {
		return err
	}

	//newImage := resize.Resize(uint(w), uint(h), img, resize.Lanczos3)

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			r, g, b, a := img.At(j, i).RGBA()
			if a <= 5 {
				continue
			}
			display.SetPixel(int16(j)+x, int16(i)+y, color.RGBA{uint8(r), uint8(g), uint8(b), uint8(255)})
		}
	}

	return nil
}
