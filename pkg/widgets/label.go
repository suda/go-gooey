package widgets

import (
	"errors"

	"github.com/gotk3/gotk3/gtk"
	"github.com/reactivex/rxgo/v2"
	"github.com/suda/go-gooey/pkg/property"
)

// Label is a representation of GTK's GtkLabel
type Label struct {
	Text interface{}
}

// Widget returns GTK's IWidget representation of the Label component
func (l *Label) Widget() (gtk.IWidget, error) {
	lbl, err := gtk.LabelNew("")
	if err != nil {
		return nil, err
	}

	if l.Text != nil {
		if val, ok := l.Text.(property.StringProperty); ok {
			observable := rxgo.FromChannel(val.Channel)
			observable.ForEach(func(v interface{}) {
				lbl.SetText(v.(string))
			}, func(err error) {}, func() {})
		} else if val, ok := l.Text.(string); ok {
			lbl.SetText(val)
		} else {
			return nil, errors.New("Text property should be either a string or StringProperty type")
		}
	}

	return lbl, nil
}
