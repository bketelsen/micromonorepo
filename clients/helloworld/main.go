package main

import (
	"context"

	proto "github.com/micro/monorepo/protos"
	"go-micro.dev/v5"
)

func main() {
	service := micro.New("helloworld")
	service.Init()

	client := proto.NewSayService("helloworld", service.Client())

	resp, err := client.Hello(context.Background(), &proto.HelloRequest{Name: "John"})
	if err != nil {
		panic(err)
	}
	println(resp.Msg)
}
