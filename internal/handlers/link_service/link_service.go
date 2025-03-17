package link_service_handler

import (
	"context"

	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/types/known/emptypb"

	"link_service/internal/interfaces/services"
	pb "link_service/proto/link_service"
)

type LinkService struct {
	pb.UnimplementedLinkServiceServer
	service     services.Service
	kafkaWriter *kafka.Writer
}

func New(service services.Service, kafkaWriter *kafka.Writer) *LinkService {
	return &LinkService{service: service, kafkaWriter: kafkaWriter}
}

func (ls *LinkService) GetLink(ctx context.Context, req *pb.GetLinkRequest) (*pb.GetLinkResponse, error) {
	link, err := ls.service.Find(ctx, req.Link)
	if err != nil {
		return &pb.GetLinkResponse{}, err
	}
	return &pb.GetLinkResponse{Link: link}, nil
}

func (ls *LinkService) InputLink(ctx context.Context, req *pb.InputLinkRequest) (*emptypb.Empty, error) {
	inputLink := services.InputLink{
		Link:      req.Link,
		FakeLink:  req.FakeLink,
		EraseTime: req.EraseTime.AsTime(),
	}

	err := ls.service.Create(ctx, inputLink)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (ls *LinkService) AddMessage(ctx context.Context, req *pb.AddMessageRequest) (*emptypb.Empty, error) {
	msg := kafka.Message{
		Key:   []byte(req.Link),
		Value: []byte(req.FakeLink),
	}

	err := ls.kafkaWriter.WriteMessages(ctx, msg)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
