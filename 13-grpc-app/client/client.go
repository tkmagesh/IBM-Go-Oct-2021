package main

import (
	"context"
	"grpc-app/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	client, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	clientConn := proto.NewAppServiceClient(client)
	ctx := context.Background()
	//Add operation
	req := &proto.AddRequest{
		X: 100,
		Y: 200,
	}
	res, err := clientConn.Add(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Add operation result: ", res.GetResult())
}
