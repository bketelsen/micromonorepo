package main

import (
	"context"
	"strings"

	"github.com/micro/monorepo/protos/stringsvc"
	"go-micro.dev/v5"
)

type Utils struct{}

func (s *Utils) Upper(ctx context.Context, req *stringsvc.UpperRequest, rsp *stringsvc.UpperResponse) error {
	rsp.Value = strings.ToUpper(req.Value)
	return nil
}

func main() {

	// create service
	service := micro.New("stringsvc")
	// register handler
	stringsvc.RegisterToolsHandler(service.Server(), &Utils{})
	// init service
	service.Init()
	// run service
	service.Run()
}
