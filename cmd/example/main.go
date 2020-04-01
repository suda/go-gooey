package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
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

type Size struct {
	Width  int
	Height int
}

type Window struct {
	Title       string
	Destroy     interface{}
	DefaultSize Size
	Children    []Widgetable
}

func (w *Window) Widget() (*gtk.Window, error) {
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		return nil, err
	}

	if w.Title != "" {
		win.SetTitle(w.Title)
	}

	if w.Destroy != nil {
		win.Connect("destroy", w.Destroy)
	}

	win.SetDefaultSize(w.DefaultSize.Width, w.DefaultSize.Height)

	if w.Children != nil {
		for _, child := range w.Children {
			widget, err := child.Widget()
			if err != nil {
				return nil, err
			}
			win.Add(widget)
		}
	}

	return win, nil
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
		Children: []Widgetable{
			&Box{
				Orientation: gtk.ORIENTATION_VERTICAL,
				Spacing:     20,
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
