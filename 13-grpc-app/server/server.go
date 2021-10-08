package main

import (
	"context"
	"grpc-app/proto"
	"io"
	"log"
	"net"

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
