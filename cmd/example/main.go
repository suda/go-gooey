package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
	. "github.com/suda/go-gooey/pkg/objects"
	. "github.com/suda/go-gooey/pkg/property"
	. "github.com/suda/go-gooey/pkg/widgets"
)

// func RunSpec(params ...interface{}) {
func RunSpec(w Window) {
	win, err := w.Widget()
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}

	win.ShowAll()
}

func main() {
	gtk.Init(nil)

	counter := NewStringProperty()

	RunSpec(Window{
		Title: "Hello bindings!",
		Destroy: func() {
			gtk.MainQuit()
		},
		DefaultSize: Size{400, 200},
		CSS: &Css{
			Path: "assets/style.css",
		},
		Children: []Widgetable{
			&Box{
				Orientation: gtk.ORIENTATION_VERTICAL,
				Children: []Widgetable{
					&Label{
						Text: counter,
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
	})

	gtk.Main()
}
