
//documents -> GoWorkspace -> src -> hands-on-microservices -> 08gRPC -> server -> main.go

/* 

//this creates greet.pb.go file
$ protoc greet.proto --go_out=plugins=grpc:.

*/
package main

import (
	"log"
	"net"

	pb "hands-on-microservices/08gRPCComplexExample/proto"

	"hands-on-microservices/08gRPCComplexExample/server/wtaserver"
	"hands-on-microservices/08gRPCComplexExample/server/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	myS := &WTAServer.MyWTAServer {
		Repo: repository.NewWTARepository(),
	}
	defer myS.Repo.CloseWTARepository()

	pb.RegisterWTAServer(s, myS)

	reflection.Register(s)

	log.Println("Starting Server.")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
