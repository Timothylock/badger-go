package apps

// App is an interface that represents an application. It should be implemented by any type that wants to be considered
// an app.
type App interface {
	// Run should execute the app. It should not return until the app has finished executing.
	// Full control of the hardware is handed over to the app when it is invoked.
	// For example, the Run should not return until the user hits the back/exit button.
	//
	// The config parameter is the configuration that the app can use to fetch/store data.
	Run() error
}
