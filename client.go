package ssmparam

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// Client is SSM client.
type Client struct {
	ssmClient *ssm.SSM
}

// New creates a client.
func New(p client.ConfigProvider, cfgs ...*aws.Config) *Client {
	client := ssm.New(p, cfgs...)
	return &Client{client}
}
