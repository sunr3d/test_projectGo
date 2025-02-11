package link_service

import (
	"context"
	"link_service/internal/interfaces/services"
	pb "link_service/proto/link_service"
)

type LinkService struct {
	pb.UnimplementedLinkServiceServer
	service services.Service
}

func NewLinkService(service services.Service) *LinkService {
	return &LinkService{service: service}
}

func (ls *LinkService) GetLink(ctx context.Context, req *pb.GetLinkRequest) (*pb.GetLinkResponse, error) {
	// Реализация метода по получению данных о fakeLink из БД
	return &pb.GetLinkResponse{Link: "PLACEHOLDER"}, nil
}

func (ls *LinkService) InputLink(ctx context.Context, req *pb.InputLinkRequest) (*pb.InputLinkResponse, error) {
	inputLink := services.InputLink{
		Link:      req.Link,
		FakeLink:  req.FakeLink,
		EraseTime: req.EraseTime.AsTime(),
	}

	if err := ls.service.Create(ctx, inputLink); err != nil {
		return &pb.InputLinkResponse{
			Success: false,
			Message: err.Error(),
			Id:      "",
		}, err
	}
	return &pb.InputLinkResponse{
		Success: true,
		Message: "Link successfully added.",
		Id:      "PLACEHOLDER_ID",
	}, nil
}
