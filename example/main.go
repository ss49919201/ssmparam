package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/s-beats/ssmparm"
)

func main() {
	sess := session.New()
	client := ssmparm.New(sess, &aws.Config{Region: aws.String("ap-northeast-1")})

	fmt.Println(client.GetParametersByPath("/s-beats/"))
	fmt.Println(client.GetParameter("name"))
}
