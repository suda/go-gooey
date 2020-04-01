package widgets

import (
	"github.com/gotk3/gotk3/gtk"
)

type Widgetable interface {
	Widget() (gtk.IWidget, error)
}
