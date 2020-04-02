package objects

import (
	"github.com/gotk3/gotk3/gtk"
)

// Css allows setting the CSS for the parent component
type Css struct {
	Path string
}

// Provider returns GTK CssProvider that can be added to Screen or StyleContext
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
