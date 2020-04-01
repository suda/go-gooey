package objects

import (
	"github.com/gotk3/gotk3/gtk"
)

type Css struct {
	Path string
}

func (c *Css) Provider() (*gtk.CssProvider, error) {
	css, err := gtk.CssProviderNew()
	if err != nil {
		return nil, err
	}

	if c.Path != "" {
		css.LoadFromPath(c.Path)
	}

	return css, nil
}
