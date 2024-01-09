package user

import (
	"context"

	"github.com/diegom0ta/grpc-server/pb"
)

type Server struct {
	pb.UserServer
}

func (s *Server) GetById(ctx context.Context, req *pb.GetByIdRequest) (*pb.GetByIdResponse, error) {
	return &pb.GetByIdResponse{
		Id:   req.Id,
		Name: "Diego Mota",
	}, nil
}
