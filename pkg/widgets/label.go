package widgets

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/reactivex/rxgo/v2"
	"github.com/suda/go-gooey/pkg/property"
)

type Label struct {
	Text *property.StringProperty
}

func (l *Label) Widget() (gtk.IWidget, error) {
	lbl, err := gtk.LabelNew("")
	if err != nil {
		return nil, err
	}

	observable := rxgo.FromChannel(l.Text.Channel)
	observable.ForEach(func(v interface{}) {
		lbl.SetText(v.(string))
	}, func(err error) {}, func() {})

	return lbl, nil
}
