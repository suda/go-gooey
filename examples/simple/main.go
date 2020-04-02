package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
	o "github.com/suda/go-gooey/pkg/objects"
	w "github.com/suda/go-gooey/pkg/widgets"
)

func main() {
	// You need to init GTK before anything else
	gtk.Init(nil)

	// Declare the window
	window := o.Window{
		Title: "Hello go-gooey!",
		// Hook the destroy signal to a callback
		Destroy: func() {
			gtk.MainQuit()
		},
		// Set the default size, otherwise the window will be as big as the initial contents
		DefaultSize: &o.Size{Width: 400, Height: 200},
		// Describe the window children components
		Children: []w.Widgetable{
			&w.Label{
				Text: "👋",
			},
		},
	}

	// Try opening the window
	err := window.Open()
	if err != nil {
		log.Fatal("Unable to open window:", err)
	}

	// Start the GTK main loop
	gtk.Main()
}
