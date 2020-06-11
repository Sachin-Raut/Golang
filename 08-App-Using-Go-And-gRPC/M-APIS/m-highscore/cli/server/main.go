package main

import (
	grpcSetup "08-App-Using-Go-And-gRPC/M-APIS/m-highscore/internal/server/grpc"
	"flag"
	"fmt"
)

func main() {

	var addressPtr = flag.String("address", ":50051", "address where you can connect with service")
	flag.Parse()
	s := grpcSetup.NewServer(*addressPtr)

	err := s.ListenAndServe()

	if err != nil {
		fmt.Println("Failed to start gRPC server")
	}
	fmt.Println("gRPC Started")
}
