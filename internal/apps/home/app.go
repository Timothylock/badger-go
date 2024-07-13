package home

import (
	"fmt"
	"machine"
	"time"
	"tinygo.org/x/drivers/uc8151"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/proggy"

	"github.com/timothylock/badger-go/internal/apps"
	"github.com/timothylock/badger-go/internal/ui"
)

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
		Icon: nil,
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
			tinyfont.WriteLine(&h.display, &proggy.TinySZ8pt7b, int16(45-(len(cfg.Name)*3)), 100, cfg.Name, ui.ColourBlack())
			cfg = curPage[1].GetAppConfig()
			tinyfont.WriteLine(&h.display, &proggy.TinySZ8pt7b, int16(146-(len(cfg.Name)*3)), 100, cfg.Name, ui.ColourBlack())
			cfg = curPage[2].GetAppConfig()
			tinyfont.WriteLine(&h.display, &proggy.TinySZ8pt7b, int16(255-(len(cfg.Name)*3)), 100, cfg.Name, ui.ColourBlack())

			ui.TopNavBar(&h.display, apps.OSName, h.GetAppConfig().Name, apps.OSVersion)
			h.display.Display()
			h.display.WaitUntilIdle()
			refresh = false
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func (h *Home) launchApp(curPage []apps.Application, i int, btn machine.Pin) error {
	// Make sure there is an app in the slot
	if len(curPage) <= 0 {
		return nil
	}

	// Wait until the user lets go of the button
	for btn.Get() {
		time.Sleep(50 * time.Millisecond)
	}

	h.display.ClearBuffer()
	h.display.WaitUntilIdle()
	err := curPage[i].Run()
	h.display.ClearBuffer()
	h.display.WaitUntilIdle()
	return err
}

func waitForAllButtonsToLetGo() {
	for machine.BUTTON_A.Get() || machine.BUTTON_B.Get() || machine.BUTTON_C.Get() {
		time.Sleep(50 * time.Millisecond)
	}
}
