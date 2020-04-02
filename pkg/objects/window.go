package objects

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	w "github.com/suda/go-gooey/pkg/widgets"
)

// Size holds Width and Height
type Size struct {
	Width  int
	Height int
}

// Window holds a visible window
type Window struct {
	Title       string
	Destroy     interface{}
	DefaultSize *Size
	CSS         *Css
	window      *gtk.Window
	Children    []w.Widgetable
	w.Childful
}

// Create instantiates  GTK Window
func (w *Window) Create() error {
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		return err
	}
	w.window = win

	if w.Title != "" {
		win.SetTitle(w.Title)
	}

	if w.Destroy != nil {
		win.Connect("destroy", w.Destroy)
	}

	if w.CSS != nil {
		provider, err := w.CSS.Provider()
		if err != nil {
			return err
		}

		screen, _ := gdk.ScreenGetDefault()
		gtk.AddProviderForScreen(screen, provider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)
	}

	if w.DefaultSize != nil {
		win.SetDefaultSize(w.DefaultSize.Width, w.DefaultSize.Height)
	}

	w.AddChildrenToParent(&win.Container, &w.Children)

	return nil
}

// Open creates the window if needed and shows it contents
func (w *Window) Open() error {
	if w.window == nil {
		err := w.Create()

		if err != nil {
			return err
		}
	}
	w.window.ShowAll()
	return nil
}
