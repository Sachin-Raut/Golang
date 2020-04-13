
//documents -> GoWorkspace -> src -> hands-on-microservices -> 08gRPC -> server -> main.go

/* 

//this creates greet.pb.go file
$ protoc greet.proto --go_out=plugins=grpc:.

*/

package main

import (
	"errors"
	"log"
	"net"
	"time"

	"context"

	pb "hands-on-microservices/08gRPC/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct {
	//Cache of results
	cache [4294967295]uint64
}

func (s *server) Calculate(ctx context.Context, in *pb.FibonacciRequest) (*pb.FibonacciReply, error) {
	if in.Number > 4294967295 {
		return nil, errors.New("Invalid Input")
	}
	timeStart := time.Now()
	result := s.calculateFibonacci(in.Number)

	return &pb.FibonacciReply{
		Result:         result,
		ProcessingTime: uint64(time.Since(timeStart)),
	}, nil
}

func (s *server) calculateFibonacci(num uint32) uint64 {
	if num == 0 {
		return 0
	}
	if num == 1 {
		return 1
	}

	res := s.cache[num]
	if res != 0 {
		return res
	}

	res = s.calculateFibonacci(num-2) + s.calculateFibonacci(num-1)
	s.cache[num] = res
	return res
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	fibS := new(server)

	pb.RegisterFibonacciServer(s, fibS)
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
