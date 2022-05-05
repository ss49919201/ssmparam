package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/s-beats/ssmparm"
)

func main() {
	sess := session.New()
	ssmparm.New(sess)
}
