package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
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

type Widgetable interface {
	Widget() (gtk.IWidget, error)
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

type Label struct {
	Text string
}

func (l *Label) Widget() (gtk.IWidget, error) {
	lbl, err := gtk.LabelNew(l.Text)
	if err != nil {
		return nil, err
	}
	return lbl, nil
}

type Button struct {
	Label   string
	Clicked interface{}
}

func (b *Button) Widget() (gtk.IWidget, error) {
	btn, err := gtk.ButtonNewWithLabel(b.Label)
	if err != nil {
		return nil, err
	}

	if b.Clicked != nil {
		btn.Connect("clicked", b.Clicked)
	}

	return btn, nil
}

type Box struct {
	Orientation gtk.Orientation
	Spacing     int
	Children    []Widgetable
}

func (b *Box) Widget() (gtk.IWidget, error) {
	box, err := gtk.BoxNew(b.Orientation, b.Spacing)
	if err != nil {
		return nil, err
	}

	if b.Children != nil {
		for _, child := range b.Children {
			widget, err := child.Widget()
			if err != nil {
				return nil, err
			}
			box.Add(widget)
		}
	}

	return box, nil
}

func main() {
	gtk.Init(nil)

	RunSpec(Window{
		Title: "Simple Example",
		Destroy: func() {
			gtk.MainQuit()
		},
		DefaultSize: Size{800, 600},
		Children: []Widgetable{
			&Box{
				Orientation: gtk.ORIENTATION_VERTICAL,
				Spacing:     20,
				Children: []Widgetable{
					&Label{
						Text: "Hello, gotk3ui!",
					},
					&Button{
						Label: "Click me!",
						Clicked: func() {
							log.Println("Hi!")
						},
					},
				},
			},
		},
	})

	gtk.Main()
}
