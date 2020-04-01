package widgets

import (
	"github.com/gotk3/gotk3/gtk"
)

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
