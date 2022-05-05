package ssmparam

import (
	"errors"

	"github.com/aws/aws-sdk-go/service/ssm"
)

var (
	ErrParameterNotFound = errors.New("parameter not found")
)

func (c *Client) GetParameter(name string) (*Parameter, error) {
	input := &ssm.GetParameterInput{
		Name: &name,
	}

	op, err := c.ssmClient.GetParameter(input)
	if err != nil {
		return nil, err
	}

	if op.Parameter == nil {
		return nil, ErrParameterNotFound
	}

	return NewParmeter(*op.Parameter.Name, *op.Parameter.Value), nil
}

func (c *Client) GetParametersByPath(path string) (*ParameterCollection, error) {
	var nextToken *string
	var parmeters []*Parameter

	for {
		input := &ssm.GetParametersByPathInput{
			Path:      &path,
			NextToken: nextToken,
		}

		op, err := c.ssmClient.GetParametersByPath(input)
		if err != nil {
			return nil, err
		}

		for _, v := range op.Parameters {
			parmeters = append(parmeters, NewParmeter(*v.Name, *v.Value))
		}

		if op.NextToken == nil {
			break
		} else {
			nextToken = op.NextToken
		}
	}

	return NewParameterCollection(parmeters), nil
}
