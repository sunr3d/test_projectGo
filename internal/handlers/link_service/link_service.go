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
	// Реализация метода по добавлению данных в БД
	id := "0" // PLACEHOLDER
	return &pb.InputLinkResponse{Success: true, Message: "Link successfully added.", Id: id}, nil
}
