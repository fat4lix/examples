package main

import (
	"github.com/murlokswarm/app"
	_ "github.com/murlokswarm/mac"
)

func main() {
	app.OnLaunch = func() {
		appMenu := &AppMainMenu{} // Creates the AppMainMenu component.
		app.Menu().Mount(appMenu) // Mounts the AppMainMenu component into the application menu context.

		// Creates a window context.
		win := app.NewWindow(app.Window{
			Title:          "Hello World",
			Width:          1280,
			Height:         720,
			TitlebarHidden: true,
		})

		hello := &Hello{} // Creates a Hello component.
		win.Mount(hello)  // Mounts the Hello component into the window context.
	}

	app.Run()
}
