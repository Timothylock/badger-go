package home

import (
	"github.com/timothylock/badger-go/internal/apps"
	"github.com/timothylock/badger-go/internal/ui"
	"time"
	"tinygo.org/x/drivers/uc8151"
)

// Home represents the home screen
type Home struct {
	display      uc8151.Device
	applications []apps.App
}

// NewHome creates a new home screen
func NewHome(display *uc8151.Device) Home {
	return Home{display: *display}
}

// AddApp adds an application to the home screen
func (h *Home) AddApp(app apps.App) error {
	h.applications = append(h.applications, app)
	return nil
}

// Run should execute the app. It should not return until the app has finished executing.
func (h *Home) Run() error {
	h.display.ClearBuffer()
	refresh := true
	for {

		if refresh {
			ui.TopNavBar(&h.display, "BadgerGoOS", "Home", "v0.0.1")
			h.display.Display()
			refresh = false
		}

		time.Sleep(200 * time.Millisecond)
	}
}
