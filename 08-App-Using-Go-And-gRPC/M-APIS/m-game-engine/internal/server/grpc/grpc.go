package grpc

import (
	"08-App-Using-Go-And-gRPC/M-APIS/m-game-engine/internal/server/logic"
	pbgameengine "08-App-Using-Go-And-gRPC/M-APIS/m-game-engine/v1"
	"context"
	"fmt"
	"net"

	"github.com/pkg/errors"

	"google.golang.org/grpc"
)

//Grpc is
type Grpc struct {
	address string
	srv     *grpc.Server
}

//NewServer is
func NewServer(address string) *Grpc {
	return &Grpc{
		address: address,
	}
}

//ListenAndServe is
func (g *Grpc) ListenAndServe() error {
	lis, err := net.Listen("tcp", g.address)

	if err != nil {
		return errors.Wrap(err, "failed to open tcp port")
	}

	serverOpts := []grpc.ServerOption{}

	g.srv = grpc.NewServer(serverOpts...)

	pbgameengine.RegisterGameEngineServer(g.srv, g)

	fmt.Println("Starting gRPC server :", g.address)

	err = g.srv.Serve(lis)

	if err != nil {
		return errors.Wrap(err, "Failed to start gRPC for m-game-engine")
	}
	return nil
}

//GetSize is
func (g *Grpc) GetSize(ctx context.Context, input *pbgameengine.GetSizeRequest) (*pbgameengine.GetSizeResponse, error) {
	fmt.Println("GetSize is called")
	return &pbgameengine.GetSizeResponse{
		Size: logic.GetSize(),
	}, nil
}

//SetScore is
func (g *Grpc) SetScore(ctx context.Context, input *pbgameengine.SetScoreRequest) (*pbgameengine.SetScoreResponse, error) {
	fmt.Println("SetScore is called")
	set := logic.SetScore(input.Score)
	return &pbgameengine.SetScoreResponse{
		Set: set,
	}, nil
}
