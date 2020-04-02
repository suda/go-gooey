package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
	o "github.com/suda/go-gooey/pkg/objects"
	p "github.com/suda/go-gooey/pkg/properties"
	w "github.com/suda/go-gooey/pkg/widgets"
)

func main() {
	gtk.Init(nil)

	// Define a StringProperty
	counter := p.NewStringProperty()

	window := o.Window{
		Title: "Hello bindings!",
		Destroy: func() {
			gtk.MainQuit()
		},
		DefaultSize: &o.Size{Width: 400, Height: 200},
		CSS: &o.Css{
			Path: "assets/css/style.css",
		},
		Children: []w.Widgetable{
			&w.Box{
				Orientation: gtk.ORIENTATION_VERTICAL,
				Children: []w.Widgetable{
					&w.Label{
						// Assign the property to Text
						Text: *counter,
					},
					&w.Button{
						Label: "Hello!",
						Clicked: func() {
							// Modify the property value
							counter.Set(counter.Value + " 👋")
						},
					},
				},
			},
		},
	}

	err := window.Open()
	if err != nil {
		log.Fatal("Unable to open window:", err)
	}

	gtk.Main()
}
