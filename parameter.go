package ssmparam

// Contains name and value.
type Parameter struct {
	name  string
	value string
}

// Name returns parameter's name.
func (p Parameter) Name() string {
	return p.name
}

// Value returns parameter's value.
func (p Parameter) Value() string {
	return p.value
}

// NewParmeter returns new Parameter.
func NewParameter(name, value string) Parameter {
	return Parameter{
		name:  name,
		value: value,
	}
}

// Contains slice of Prameter.
type ParameterCollection struct {
	parmeters []Parameter
}

// NewParameter returns new ParameterCollection.
func NewParameterCollection(parameters []Parameter) *ParameterCollection {
	return &ParameterCollection{parameters}
}

// FirstByName returns first parameter that match given name.
//
// If there is no match, returns empty.
func (p *ParameterCollection) FirstByName(name string) (Parameter, bool) {
	var parameter Parameter
	for _, v := range p.parmeters {
		if v.name == name {
			parameter = v
			return parameter, true
		}
	}
	return parameter, false
}

// Parameters returns all Parameter.
func (p *ParameterCollection) Parameters() []Parameter {
	return p.parmeters
}

// Filter returns filterd slice of Parameter.
func (p *ParameterCollection) Filter(callback func(Parameter) bool) []Parameter {
	var filtered []Parameter
	for _, v := range p.parmeters {
		if callback(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}
