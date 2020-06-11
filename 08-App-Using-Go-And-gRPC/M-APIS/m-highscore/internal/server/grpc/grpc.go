package grpc

import (
	pbhighscore "08-App-Using-Go-And-gRPC/M-APIS/m-highscore/v1"
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

//HighScore is
var HighScore = 999.0

//SetHighScore is
func (g *Grpc) SetHighScore(ctx context.Context, input *pbhighscore.SetHighScoreRequest) (*pbhighscore.SetHighScoreResponse, error) {
	fmt.Println("SetHighScore called")
	return &pbhighscore.SetHighScoreResponse{
		Set: true,
	}, nil
}

//GetHighScore is
func (g *Grpc) GetHighScore(ctx context.Context, input *pbhighscore.GetHighScoreRequest) (*pbhighscore.GetHighScoreResponse, error) {
	return &pbhighscore.GetHighScoreResponse{
		HighScore: HighScore,
	}, nil
}

//ListenAndServe is
func (g *Grpc) ListenAndServe() error {
	lis, err := net.Listen("tcp", g.address)

	if err != nil {
		return errors.Wrap(err, "failed to open TCP port")
	}

	serverOpts := []grpc.ServerOption{}

	g.srv = grpc.NewServer(serverOpts...)

	pbhighscore.RegisterGameServer(g.srv, g)
	fmt.Println("Starting gRPC server on port", g.address)

	err = g.srv.Serve(lis)
	if err != nil {
		return errors.Wrap(err, "Failed to start gRPC server")
	}
	return nil
}

//NewServer is
func NewServer(address string) *Grpc {
	return &Grpc{
		address: address,
	}
}
