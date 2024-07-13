package main

import (
	"fmt"
	"github.com/timothylock/badger-go/internal/apps/about"
	"machine"

	"tinygo.org/x/drivers/uc8151"

	"github.com/timothylock/badger-go/internal/apps"
	"github.com/timothylock/badger-go/internal/apps/counter"
	"github.com/timothylock/badger-go/internal/apps/home"
)

func main() {
	display, err := initScreen()
	if err != nil {
		fmt.Printf("Failed to initialize screen: %v", err)
		return
	}

	homeScreen := home.NewHome(display)

	applications, err := initApps(display)
	if err != nil {
		fmt.Printf("Failed to initialize applications: %v", err)
		return
	}

	for _, app := range applications {
		err := homeScreen.AddApp(app)
		if err != nil {
			fmt.Printf("Failed to add application to homeScreen screen: %v", err)
			return
		}
	}

	about.NewAboutApp(display).Run()

	err = homeScreen.Run()
	if err != nil {
		fmt.Printf("Failed to run homeScreen screen: %v", err)
		return
	}
}

func initApps(device *uc8151.Device) ([]apps.Application, error) {
	applications := []apps.Application{
		counter.NewCounterApp(device),
		counter.NewCounterApp(device),
		about.NewAboutApp(device), // TODO: Replace with actual apps
	}

	return applications, nil
}

// initScreen initializes the screen and returns a pointer to the device
func initScreen() (*uc8151.Device, error) {
	err := machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 12000000,
		SCK:       machine.EPD_SCK_PIN,
		SDO:       machine.EPD_SDO_PIN,
	})
	if err != nil {
		return nil, err
	}

	display := uc8151.New(machine.SPI0, machine.EPD_CS_PIN, machine.EPD_DC_PIN, machine.EPD_RESET_PIN, machine.EPD_BUSY_PIN)
	display.Configure(uc8151.Config{
		Rotation: uc8151.ROTATION_270,
		Speed:    uc8151.MEDIUM,
		Blocking: true,
	})

	return &display, nil
}
