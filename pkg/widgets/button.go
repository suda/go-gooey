package widgets

import (
	"github.com/gotk3/gotk3/gtk"
)

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
