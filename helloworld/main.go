package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/option"
	"google.golang.org/grpc"

	greeter "github.com/idlebot/monorepo/hellogrpc/greeter/v1"
	greetercli "github.com/idlebot/monorepo/hellogrpc/greeter/v1/client"
)

func main() {
	ctx := context.Background()
	conn, err := grpc.Dial("localhost:5432", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	opt := option.WithGRPCConn(conn)
	c, err := greetercli.NewGreeterClient(ctx, opt)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	req := &greeter.HelloRequest{
		Name: "Go",
	}
	resp, err := c.Hello(ctx, req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Message)
}
