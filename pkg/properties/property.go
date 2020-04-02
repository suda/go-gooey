package properties

import (
	"github.com/reactivex/rxgo/v2"
)

// Property is a observable type, used to bind values to components
type Property struct {
	Channel chan rxgo.Item
	Value   interface{}
}

// NewProperty creates a new Property instance
func NewProperty() *Property {
	p := new(Property)
	p.Channel = make(chan rxgo.Item)
	return p
}

// Set updates the value of the property and notifies about the change
func (p *Property) Set(v interface{}) {
	p.Value = v
	go func() {
		p.Channel <- rxgo.Of(p.Value)
	}()
}

// StringProperty is a Property where the Value can only be a string
type StringProperty struct {
	Property
	Value string
}

// NewStringProperty creates a new StringProperty instance
func NewStringProperty() *StringProperty {
	p := new(StringProperty)
	p.Property = *NewProperty()
	return p
}

// Set updates the value of the property and notifies about the change
func (p *StringProperty) Set(v string) {
	p.Value = v
	p.Property.Set(v)
}
