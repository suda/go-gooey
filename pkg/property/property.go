package property

import (
	"github.com/reactivex/rxgo/v2"
)

type Property struct {
	Channel chan rxgo.Item
	Value   interface{}
}

func NewProperty() *Property {
	p := new(Property)
	p.Channel = make(chan rxgo.Item)
	return p
}

func (p *Property) Set(v interface{}) {
	p.Value = v
	go func() {
		p.Channel <- rxgo.Of(p.Value)
	}()
}

type StringProperty struct {
	Property
	Value string
}

func NewStringProperty() *StringProperty {
	p := new(StringProperty)
	p.Property = *NewProperty()
	return p
}

func (p *StringProperty) Set(v string) {
	p.Value = v
	p.Property.Set(v)
}
