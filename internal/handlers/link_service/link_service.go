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
	return nil, nil
}

func (ls *LinkService) InputLink(ctx context.Context, req *pb.InputLinkRequest) (*pb.InputLinkResponse, error) {
	return nil, nil
}
