package link_server

import (
	"link_service/internal/interfaces/services"
	pb "link_service/proto/link_service"
)

type LinkServer struct {
	pb.UnimplementedLinkServiceServer
	service services.Service
}

func New() {

}

func (s *LinkServer) GetLink() {

}
