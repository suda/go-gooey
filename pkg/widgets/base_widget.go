package widgets

import (
	"github.com/gotk3/gotk3/gtk"
)

// Widgetable allows component to be represented as GTK Widget
type Widgetable interface {
	Widget() (gtk.IWidget, error)
}

// Childful is a component that can own children
type Childful struct {
}

// AddChildrenToParent adds passed children to the parent container
func (c *Childful) AddChildrenToParent(parent *gtk.Container, children *[]Widgetable) error {
	if (parent != nil) && (children != nil) {
		for _, child := range *children {
			widget, err := child.Widget()
			if err != nil {
				return err
			}
			parent.Add(widget)
		}
	}
	return nil
}
