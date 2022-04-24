package services

import (
	"context"
	"google.golang.org/grpc"
)

type ProdService struct {
}

func (p *ProdService) GetProdStock(
	ctx context.Context,
	in *ProdRequest,
	opts ...grpc.CallOption) (response *ProdResponse, err error) {
	return &ProdResponse{
		ProdStock: 1,
	}, nil
}
