package service

import (
	"github.com/ogataka50/APIServerGo/pb/proto"
	"golang.org/x/net/context"
)

type PingService struct {
}

func (s *PingService) Hello(ctx context.Context, req *ping.HelloReqest) (*ping.HelloResponse, error) {
	toMessage := req.GetToMessage()
	response := ping.HelloResponse{
		ResMessage: "I hear " + toMessage,
	}

	return &response, nil
}
