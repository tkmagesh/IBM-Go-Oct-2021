package main

import (
	"context"
	"grpc-app/proto"
	"log"
	"time"

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
	//doRequestResponse(clientConn, ctx)
	doClientStreaming(clientConn, ctx)

}

func doClientStreaming(clientConn proto.AppServiceClient, ctx context.Context) {
	nos := []int32{3, 1, 4, 2, 5}
	clientStream, err := clientConn.Average(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	for _, no := range nos {
		time.Sleep(500 * time.Millisecond)
		log.Printf("[client] sending average request for %d\n", no)
		req := &proto.AverageRequest{
			No: no,
		}
		clientStream.Send(req)
	}
	res, err := clientStream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("[client] received average result: %d\n", res.GetResult())
}

func doRequestResponse(clientConn proto.AppServiceClient, ctx context.Context) {
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
