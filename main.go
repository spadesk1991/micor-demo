package main

import (
	"context"

	"log"

	micro "github.com/micro/go-micro"
	proto "github.com/spadesk1991/micor-demo/proto"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, res *proto.HelloResponse) error {
	res.Greeting = "我的名字是：" + req.Namme
	return nil
}

func main() {
	service := micro.NewService(micro.Name("greeter"))
	service.Init()
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))
	err := service.Run()
	if err != nil {
		log.Fatal(err)
	}
}
