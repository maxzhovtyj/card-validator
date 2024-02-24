package handler

import (
	"context"
	"github.com/maxzhovtyj/card-validator/internal/service"
	pb "github.com/maxzhovtyj/card-validator/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var _ pb.CardServiceServer = (*CardServer)(nil)

type CardServer struct {
	logger  *zap.SugaredLogger
	service service.Card
	pb.UnimplementedCardServiceServer
}

func NewServer(srv *grpc.Server, service service.Card, logger *zap.SugaredLogger) {
	cardSrv := &CardServer{service: service, logger: logger}
	pb.RegisterCardServiceServer(srv, cardSrv)
	reflection.Register(srv)
}

func (srv *CardServer) Validate(ctx context.Context, req *pb.ValidateCardRequest) (*pb.ValidateCardResponse, error) {
	card := &pb.Card{
		Number:          req.Card.Number,
		ExpirationMonth: req.Card.ExpirationMonth,
		ExpirationYear:  req.Card.ExpirationYear,
	}

	err := srv.service.Validate(card)
	if err != nil {
		resp := &pb.ValidateCardResponse{
			Valid: false,
			Error: &pb.Error{
				Code:    "001",
				Message: err.Error(),
			},
		}

		return resp, nil
	}

	return &pb.ValidateCardResponse{Valid: true}, nil
}
