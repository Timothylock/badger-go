package apps

import "embed"

const OSName = "BadgerGoOS"
const OSVersion = "v0.0.1"

// Application is an interface that represents an application. It should be implemented by any type that wants to be
// considered an app.
type Application interface {
	// GetAppConfig should return the configuration of the app. This includes the name and icon of the app.
	// It is used by the home screen to display the app.
	GetAppConfig() AppConfig

	// Run should execute the app. It should not return until the app has finished executing.
	// Full control of the hardware is handed over to the app when it is invoked.
	// For example, the Run should not return until the user hits the back/exit button.
	//
	// The config parameter is the configuration that the app can use to fetch/store data.
	Run() error
}

type AppConfig struct {
	Name string
	Icon embed.FS
}
