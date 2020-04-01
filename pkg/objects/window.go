package objects

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	. "github.com/suda/go-gooey/pkg/widgets"
)

type Size struct {
	Width  int
	Height int
}

type Window struct {
	Title       string
	Destroy     interface{}
	DefaultSize Size
	Children    []Widgetable
	CSS         *Css
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

	if w.CSS != nil {
		provider, err := w.CSS.Provider()
		if err == nil {
			screen, _ := gdk.ScreenGetDefault()
			gtk.AddProviderForScreen(screen, provider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)
		}
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
