package main

import (
	pbgameengine "08-App-Using-Go-And-gRPC/M-APIS/m-game-engine/v1"
	"flag"
	"fmt"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var addressPtr = flag.String("address", ":60051", "address to connect to m-game-engine")
	flag.Parse()

	conn, err := grpc.Dial(*addressPtr, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Failed to dial m-game-engine gRPC service")
	}

	defer conn.Close()

	timeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	c := pbgameengine.NewGameEngineClient(conn)

	if c == nil {
		fmt.Println("Client nil")
	}

	r, err := c.GetSize(timeout, &pbgameengine.GetSizeRequest{})

	if err != nil {
		fmt.Println("Failed to get response from address :", *addressPtr)
	}

	if r != nil {
		fmt.Println("Size =", r.GetSize())
	} else {
		fmt.Println("Couldn't get size")
	}
}
