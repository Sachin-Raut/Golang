package main

import (
	bff "08-App-Using-Go-And-gRPC/M-APIS/m-bff/bff"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	grpcAddressHighScore := flag.String("address-m-highscore", ":50051", "The gRPC server for highscore")

	grpcAddressGameEngine := flag.String("address-m-game-engine", ":60051", "The gRPC server for game engine")

	flag.Parse()

	gameClient, err := bff.NewGrpcGameServiceClient(*grpcAddressHighScore)

	if err != nil {
		fmt.Println("error while creating a client for m-highscore")
		return
	}

	gameEngineClient, err := bff.NewGrpcGameEngineServiceClient(*grpcAddressGameEngine)

	if err != nil {
		fmt.Println("error while creating a client for m-game-engine")
		return
	}

	gr := bff.NewGameResource(gameClient, gameEngineClient)

	router := mux.NewRouter()
	router.HandleFunc("/getHighscore", gr.GetHighScore()).Methods("GET")
	router.HandleFunc("/setHighScore", gr.SetHighScore()).Methods("POST")
	router.HandleFunc("/getSize", gr.GetSize()).Methods("GET")
	router.HandleFunc("/setScore", gr.SetScore()).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))

}
