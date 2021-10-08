package main

import (
	"context"
	"grpc-app/proto"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedAppServiceServer
}

func (s *server) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	x := req.GetX()
	y := req.GetY()
	log.Printf("[Add] Processing %d and %d\n", x, y)
	result := x + y
	response := &proto.AddResponse{
		Result: result,
	}
	return response, nil
}

func (s *server) Average(stream proto.AppService_AverageServer) error {
	var sum int32
	var count int32
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			res := &proto.AverageResponse{
				Result: sum / count,
			}
			log.Printf("[Average] Sending response")
			stream.SendAndClose(res)
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("[Average] received %d\n", req.GetNo())
		sum += req.GetNo()
		count++
	}
	return nil
}

func (s *server) GeneratePrime(req *proto.PrimeRequest, stream proto.AppService_GeneratePrimeServer) error {
	start := req.GetStart()
	end := req.GetEnd()
	log.Printf("[server] GeneratePrime request received from %d to %d\n", start, end)
	for i := start; i <= end; i++ {
		if isPrime(i) {
			time.Sleep(time.Millisecond * 500)
			log.Printf("Sending %d\n", i)
			res := &proto.PrimeResponse{
				No: i,
			}
			stream.Send(res)
		}
	}
	return nil
}

func isPrime(no int32) bool {
	if no < 2 {
		return false
	}
	for i := int32(2); i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterAppServiceServer(grpcServer, &server{})
	e := grpcServer.Serve(listener)
	if e != nil {
		log.Fatalln(e)
	}
}
