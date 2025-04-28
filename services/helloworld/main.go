package main

import (
	"context"

	proto "github.com/micro/monorepo/protos"
	"go-micro.dev/v5"
)

type Say struct{}

func (s *Say) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Msg = "Hello " + req.Name
	return nil
}

func main() {

	// create service
	service := micro.New("helloworld")
	// register handler
	proto.RegisterSayHandler(service.Server(), &Say{})
	// init service
	service.Init()
	// run service
	service.Run()
}
