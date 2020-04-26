package main

import (
	"context"
	"flag"
	"io"
	"log"
	"time"

	pb "hands-on-microservices/08gRPCComplexExample/proto"
	"google.golang.org/grpc"
)

const (
	address       = "localhost:50051"
	defaultNumber = uint32(1)
)

func main() {
	typ := flag.Int("type", 1, "Type of Request.")
	playerID := flag.Int("playerId", 2, "playerID for request type one.")
	flag.Parse()

	if *typ > 3 || *typ < 1 {
		log.Fatal("Invalid type.")
	}

	//Start connection
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	//Create client with that connection
	c := pb.NewWTAClient(conn)

	switch *typ {
	case 1:
		//We create a context with timeout of one second
		ctx, _ := context.WithTimeout(context.Background(), time.Second)
		res, err := c.GetPlayerWithHighestRanking(ctx, &pb.PlayerIdRequest{PlayerId: uint32(*playerID)})
		if err != nil {
			log.Printf("Error On Request 1: %s \n", err.Error())
		} else {
			log.Printf("Recieved: %+v", res)
		}
	case 2:
		ctx := context.Background()
		stream, err := c.GetRankingsByPlayerId(ctx, &pb.PlayerIdRequest{PlayerId: uint32(*playerID)})
		if err != nil {
			log.Printf("Error On Request 2: %s \n", err.Error())
			break
		}
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				log.Println("Stream ended.")
				break
			}
			if err != nil {
				log.Printf("Error On Request 2: %s \n", err.Error())
			}
			log.Printf("Recieved: %+v\n", in)
		}
	case 3:
		ctx := context.Background()
		stream, err := c.GetPlayers(ctx)
		if err != nil {
			log.Printf("Error On Request 3: %s \n", err.Error())
			break
		}
		for i := 1; i < 4; i++ {
			time.Sleep(100 * time.Millisecond)
			err = stream.Send(&pb.PlayerIdRequest{PlayerId: uint32(i)})
			if err != nil {
				log.Printf("Error On Request 3: %s \n", err.Error())
				break
			}
		}
		res, err := stream.CloseAndRecv()
		if err != nil {
			log.Printf("Error On Request 3: %s \n", err.Error())
			break
		}
		log.Printf("Recieved: %+v", res)
	}
}
