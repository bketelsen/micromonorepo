package main

import (
	"context"

	"github.com/micro/monorepo/protos/helloworld"
	"github.com/micro/monorepo/protos/stringsvc"
	"go-micro.dev/v5"
)

type Say struct {
	tools stringsvc.ToolsService
}

func NewSay(tools stringsvc.ToolsService) *Say {
	return &Say{
		tools: tools,
	}
}

func (s *Say) Hello(ctx context.Context, req *helloworld.HelloRequest, rsp *helloworld.HelloResponse) error {
	rsp.Msg = "Hello " + req.Name
	return nil
}
func (s *Say) StrongHello(ctx context.Context, req *helloworld.HelloRequest, rsp *helloworld.HelloResponse) error {
	trsp, err := s.tools.Upper(ctx, &stringsvc.UpperRequest{Value: req.Name})
	if err != nil {
		return err
	}
	rsp.Msg = "Hello " + trsp.Value
	return nil
}

func main() {

	// create service
	service := micro.New("helloworld")
	hd := NewSay(stringsvc.NewToolsService("stringsvc", service.Client()))
	// register handler
	helloworld.RegisterSayHandler(service.Server(), hd)
	// init service
	service.Init()
	// run service
	service.Run()
}
