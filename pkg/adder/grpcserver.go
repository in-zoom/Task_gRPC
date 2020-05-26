package adder

import (
	"Tes_gRPC/pkg/api"
	"context"
)

type GRPCServer struct{}

func (s *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	elements := make(map[string]string)
	elements["Name"] = "Ivan"
	elements["Address"] = "Moscow"
	elements["Email"] = "Ivan@mail.ru"

	if elements["Name"] != req.Name {
		return &api.AddResponse{Result: false}, nil
	}
	if elements["Address"] != req.Address {
		return &api.AddResponse{Result: false}, nil
	}
	if elements["Email"] != req.Email {
		return &api.AddResponse{Result: false}, nil
	}
	return &api.AddResponse{Result: true}, nil
}
