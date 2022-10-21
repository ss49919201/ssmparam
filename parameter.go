package ssmparam

// Contains name and value.
type Parameter struct {
	name  string
	value string
}

// Name returns parameter's name.
func (p *Parameter) Name() string {
	return p.name
}

// Value returns parameter's value.
func (p *Parameter) Value() string {
	return p.value
}

// NewParmeter returns new Parameter.
func NewParameter(name, value string) *Parameter {
	return &Parameter{
		name:  name,
		value: value,
	}
}

// Contains slice of Prameter.
type ParameterCollection struct {
	parmeters []*Parameter
}

// NewParameter returns new ParameterCollection.
func NewParameterCollection(parameters []*Parameter) *ParameterCollection {
	return &ParameterCollection{parameters}
}

// Parameter returns Parameter by name.
func (p *ParameterCollection) Parameter(name string) *Parameter {
	for _, v := range p.parmeters {
		if v.name == name {
			return v
		}
	}
	return nil
}

// Parameters returns all Parameter.
func (p *ParameterCollection) Parameters() []*Parameter {
	return p.parmeters
}
