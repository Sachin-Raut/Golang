package main

import (
	grpcSetup "08-App-Using-Go-And-gRPC/M-APIS/m-game-engine/internal/server/grpc"
	"flag"
	"fmt"
)

func main() {
	var addressPtr = flag.String("address", ":60051", "address to connect to port")
	flag.Parse()
	s := grpcSetup.NewServer(*addressPtr)

	err := s.ListenAndServe()

	if err != nil {
		fmt.Println("failed to start gRPC server")
	}
}
