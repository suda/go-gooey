package widgets

import (
	"github.com/gotk3/gotk3/gtk"
)

// Box is a representation of GTK's GtkBox
type Box struct {
	Orientation gtk.Orientation
	Spacing     int
	Children    []Widgetable
	Childful
}

// Widget returns GTK's IWidget representation of the Box component
func (b *Box) Widget() (gtk.IWidget, error) {
	box, err := gtk.BoxNew(b.Orientation, b.Spacing)
	if err != nil {
		return nil, err
	}

	b.AddChildrenToParent(&box.Container, &b.Children)

	return box, nil
}
