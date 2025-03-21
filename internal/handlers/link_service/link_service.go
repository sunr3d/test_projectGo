package link_service_handler

import (
	"context"

	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/types/known/emptypb"

	"link_service/internal/interfaces/services"
	pb "link_service/proto/link_service"
)

var linkTopic = "link_service"

type LinkService struct {
	pb.UnimplementedLinkServiceServer
	service services.Service
}

func New(service services.Service) *LinkService {
	return &LinkService{service: service}
}

func (ls *LinkService) GetLink(ctx context.Context, req *pb.GetLinkRequest) (*pb.GetLinkResponse, error) {
	link, err := ls.service.Find(ctx, req.GetLink())
	if err != nil {
		return &pb.GetLinkResponse{}, err
	}
	return &pb.GetLinkResponse{Link: link}, nil
}

func (ls *LinkService) InputLink(ctx context.Context, req *pb.InputLinkRequest) (*emptypb.Empty, error) {
	inputLink := services.InputLink{
		Link:      req.GetLink(),
		FakeLink:  req.GetFakeLink(),
		EraseTime: req.GetEraseTime().AsTime(),
	}

	err := ls.service.Create(ctx, inputLink)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (ls *LinkService) AddMessage(ctx context.Context, req *pb.AddMessageRequest) (*emptypb.Empty, error) {
	msg := kafka.Message{
		Topic: linkTopic,
		Key:   []byte(req.GetLink()),
		Value: []byte(req.GetFakeLink()),
	}

	err := ls.service.AddMessage(ctx, msg)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
