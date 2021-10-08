package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"io"
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
	//doClientStreaming(clientConn, ctx)
	//doServerStreaming(clientConn, ctx)
	doBiDirectionalStreaming(clientConn, ctx)
}

func doBiDirectionalStreaming(client proto.AppServiceClient, ctx context.Context) {
	stream, err := client.Greet(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	requestData := []*proto.GreetRequest{
		&proto.GreetRequest{
			Greeting: &proto.Greeting{
				FirstName: "Magesh",
				LastName:  "Kuppan",
			},
		},
		&proto.GreetRequest{
			Greeting: &proto.Greeting{
				FirstName: "Suresh",
				LastName:  "Kannan",
			},
		},
		&proto.GreetRequest{
			Greeting: &proto.Greeting{
				FirstName: "Ramesh",
				LastName:  "Jayaraman",
			},
		},
		&proto.GreetRequest{
			Greeting: &proto.Greeting{
				FirstName: "Rajesh",
				LastName:  "Pandit",
			},
		},
		&proto.GreetRequest{
			Greeting: &proto.Greeting{
				FirstName: "Naresh",
				LastName:  "Kumar",
			},
		},
	}

	go func() {
		for _, req := range requestData {
			stream.Send(req)
		}
		stream.CloseSend()
	}()
	/* wg := &sync.WaitGroup{}
	wg.Add(1) */
	done := make(chan bool)
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				fmt.Println("Thats all folks!")
				break
			}
			if err != nil {
				log.Fatalln(err)
			}
			log.Println("Greet Result:", res.GetGreetMessage())
		}
		//done <- true
		close(done)
	}()
	<-done
}

func doServerStreaming(clientConn proto.AppServiceClient, ctx context.Context) {
	req := &proto.PrimeRequest{
		Start: 5,
		End:   100,
	}
	stream, err := clientConn.GeneratePrime(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			log.Println("All prime numbers received")
			return
		}
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("Prime number: ", res.GetNo())
	}
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
