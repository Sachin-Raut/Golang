package main

import (
	pbhighscore "08-App-Using-Go-And-gRPC/M-APIS/m-highscore/v1"
	"context"
	"flag"
	"fmt"

	"google.golang.org/grpc"
)

func main() {
	var addressPtr = flag.String("address", ":50051", "address to connect")
	flag.Parse()

	conn, err := grpc.Dial(*addressPtr, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Failed to load gRPC service")
	}
	defer conn.Close()

	c := pbhighscore.NewGameClient(conn)

	if c != nil {
		fmt.Println("Client nil")
	}

	r, err := c.GetHighScore(context.Background(), &pbhighscore.GetHighScoreRequest{})

	if err != nil {
		fmt.Println("Failed to get a response")
	}

	if r != nil {
		fmt.Println("Highscore =", r.GetHighScore())
	} else {
		fmt.Println("Could not get highscore")
	}
}
