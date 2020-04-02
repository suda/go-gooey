package widgets

import (
	"github.com/gotk3/gotk3/gtk"
)

// Button is a representation of GTK's GtkButton
type Button struct {
	Label   string
	Clicked interface{}
}

// Widget returns GTK's IWidget representation of the Button component
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
