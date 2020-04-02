package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
	. "github.com/suda/go-gooey/pkg/objects"
	. "github.com/suda/go-gooey/pkg/property"
	. "github.com/suda/go-gooey/pkg/widgets"
)

func main() {
	gtk.Init(nil)

	counter := NewStringProperty()

	window := Window{
		Title: "Hello bindings!",
		Destroy: func() {
			gtk.MainQuit()
		},
		DefaultSize: &Size{Width: 400, Height: 200},
		CSS: &Css{
			Path: "assets/css/style.css",
		},
		Children: []Widgetable{
			&Box{
				Orientation: gtk.ORIENTATION_VERTICAL,
				Children: []Widgetable{
					&Label{
						Text: *counter,
					},
					&Button{
						Label: "Hello!",
						Clicked: func() {
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
