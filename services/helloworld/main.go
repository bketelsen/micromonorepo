package main

import (
	"context"

	"go-micro.dev/v5"
)

type Say struct{}

type Request struct {
	Name string
}

type Response struct {
	Message string
}

func (s *Say) Hello(ctx context.Context, req *Request, rsp *Response) error {
	rsp.Message = "Hello " + req.Name
	return nil
}

func main() {

	// create service
	service := micro.New("helloworld")
	// register handler
	service.Handle(new(Say))
	// init service
	service.Init()
	// run service
	service.Run()
}
