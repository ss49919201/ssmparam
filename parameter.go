package ssmparam

type Parameter struct {
	name  string
	value string
}

func (p *Parameter) Name() string {
	return p.name
}

func (p *Parameter) Value() string {
	return p.value
}

func NewParmeter(name, value string) *Parameter {
	return &Parameter{
		name:  name,
		value: value,
	}
}

type ParameterCollection struct {
	parmeters []*Parameter
}

func NewParameterCollection(parameters []*Parameter) *ParameterCollection {
	return &ParameterCollection{parameters}
}

func (p *ParameterCollection) Parameter(name string) *Parameter {
	for _, v := range p.parmeters {
		if v.name == name {
			return v
		}
	}
	return nil
}

func (p *ParameterCollection) Parameters() []*Parameter {
	return p.Parameters()
}
