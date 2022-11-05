package ssmparam

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewParameterCollection(t *testing.T) {
	got := NewParameterCollection([]*Parameter{{name: "name", value: "value"}})
	assert.Equal(t, &ParameterCollection{parmeters: []*Parameter{{name: "name", value: "value"}}}, got)
}

func TestFirstByName(t *testing.T) {
	collection := &ParameterCollection{parmeters: []*Parameter{
		{name: "name", value: "value"},
		{name: "name", value: "value-copy"},
		{name: "name2", value: "value2"},
	}}

	assert.Equal(t, &Parameter{name: "name", value: "value"}, collection.FirstByName("name"))
	assert.Equal(t, &Parameter{name: "name2", value: "value2"}, collection.FirstByName("name2"))
	assert.Nil(t, collection.FirstByName("nil"))
}

func TestPrameters(t *testing.T) {
	collection := &ParameterCollection{parmeters: []*Parameter{
		{name: "name", value: "value"},
		{name: "name", value: "value-copy"},
		{name: "name2", value: "value2"},
	}}

	assert.Equal(
		t,
		[]*Parameter{
			{name: "name", value: "value"},
			{name: "name", value: "value-copy"},
			{name: "name2", value: "value2"},
		},
		collection.Parameters(),
	)
}

func TestParameterCollection_Filter(t *testing.T) {
	collection := &ParameterCollection{parmeters: []*Parameter{
		{name: "name", value: "value"},
		{name: "name", value: "value-copy"},
		{name: "name2", value: "value2"},
	}}

	callback := func(p *Parameter) bool {
		return p.name == "name"
	}

	assert.Equal(
		t,
		[]*Parameter{{name: "name", value: "value"}, {name: "name", value: "value-copy"}},
		collection.Filter(callback),
	)
}
