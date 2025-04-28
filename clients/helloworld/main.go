package main

import (
	"context"

	proto "github.com/micro/monorepo/protos/helloworld"
	"go-micro.dev/v5"
)

func main() {
	helloService := micro.New("helloworld")
	helloService.Init()

	sayClient := proto.NewSayService("helloworld", helloService.Client())

	resp, err := sayClient.Hello(context.Background(), &proto.HelloRequest{Name: "John"})
	if err != nil {
		panic(err)
	}
	println(resp.Msg)
}
