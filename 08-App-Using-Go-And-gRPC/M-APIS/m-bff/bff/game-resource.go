package bff

import (
	pbgameengine "08-App-Using-Go-And-gRPC/M-APIS/m-game-engine/v1"
	pbhighscore "08-App-Using-Go-And-gRPC/M-APIS/m-highscore/v1"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"google.golang.org/grpc"
)

//gameResource is
type gameResource struct {
	gameClient       pbhighscore.GameClient
	gameEngineClient pbgameengine.GameEngineClient
}

//Score is
type Score struct {
	Highscore float64 `json:"highscore"`
}

//Error is
type Error struct {
	Message string `json:"message"`
}

//Size is
type Size struct {
	Size float64 `json:"size"`
}

func sendSuccess(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}

func respondWithError(w http.ResponseWriter, status int, error Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}

//NewGameResource is
func NewGameResource(gameClient pbhighscore.GameClient, gameEngineClient pbgameengine.GameEngineClient) *gameResource {
	return &gameResource{
		gameClient:       gameClient,
		gameEngineClient: gameEngineClient,
	}
}

//NewGrpcGameServiceClient is
func NewGrpcGameServiceClient(serverAdd string) (pbhighscore.GameClient, error) {
	conn, err := grpc.Dial(serverAdd, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Failed to dial", err)
		return nil, err
	}

	fmt.Println("Successfully connected :", serverAdd)

	if conn == nil {
		fmt.Println("m-highscore connection is nil in bff")
	}

	client := pbhighscore.NewGameClient(conn)
	return client, nil
}

//NewGrpcGameEngineServiceClient is
func NewGrpcGameEngineServiceClient(serverAdd string) (pbgameengine.GameEngineClient, error) {
	conn, err := grpc.Dial(serverAdd, grpc.WithInsecure())

	if err != nil {
		fmt.Println("failed to dial", err)
		return nil, err
	}

	fmt.Println("Successfully connected :", serverAdd)

	if conn == nil {
		fmt.Println("m-game-engine connection is nil in m-bff")
	}

	client := pbgameengine.NewGameEngineClient(conn)

	return client, nil
}

func (gr *gameResource) SetHighScore() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("SetHighScore() called")

		var score Score
		var error Error

		json.NewDecoder(r.Body).Decode(&score)
		fmt.Println("Score", score)

		//check if user provides empty value
		if score.Highscore == 0 {
			error.Message = "Highscore can't be empty"
			respondWithError(w, http.StatusBadRequest, error)
			return
		}

		gr.gameClient.SetHighScore(context.Background(), &pbhighscore.SetHighScoreRequest{
			HighScore: score.Highscore,
		})
		w.Write([]byte("Successfully set highscore"))
	}
}

func (gr *gameResource) GetHighScore() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GetHighScore() called")

		var score Score

		highscoreResponse, err := gr.gameClient.GetHighScore(context.Background(), &pbhighscore.GetHighScoreRequest{})

		if err != nil {
			fmt.Println("Error while getting highscore -", err)
			return
		}

		// hsString := strconv.FormatFloat(highscoreResponse.HighScore, 'e', -1, 64)

		score.Highscore = highscoreResponse.HighScore //hsString

		w.Header().Set("Content-Type", "application/json")
		sendSuccess(w, score)
	}
}

func (gr *gameResource) GetSize() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GetSize() called")
		var size Size

		sizeResponse, err := gr.gameEngineClient.GetSize(context.Background(), &pbgameengine.GetSizeRequest{})

		if err != nil {
			fmt.Println("Error while getting size")
			return
		}

		size.Size = sizeResponse.GetSize()

		w.Header().Set("Content-Type", "application/json")
		sendSuccess(w, size)
	}
}

func (gr *gameResource) SetScore() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("SetScore() called")

		var score Score

		json.NewDecoder(r.Body).Decode(&score)

		_, err := gr.gameEngineClient.SetScore(context.Background(), &pbgameengine.SetScoreRequest{
			Score: score.Highscore,
		})
		if err != nil {
			fmt.Println("Error while setting score")
			return
		}
		w.Write([]byte("Successfully set score"))
	}
}
