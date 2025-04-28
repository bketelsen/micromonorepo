package main

import (
	"context"

	"github.com/micro/monorepo/protos/helloworld"
	"go-micro.dev/v5"
)

type Say struct{}

func (s *Say) Hello(ctx context.Context, req *helloworld.HelloRequest, rsp *helloworld.HelloResponse) error {
	rsp.Msg = "Hello " + req.Name
	return nil
}

func main() {

	// create service
	service := micro.New("helloworld")
	// register handler
	helloworld.RegisterSayHandler(service.Server(), &Say{})
	// init service
	service.Init()
	// run service
	service.Run()
}
