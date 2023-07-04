package server

import (
	"context"
	"time"

	"github.com/deka-microservices/go-bcrypt-service/pkg/service"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type BcryptRPCServer struct {
	service.UnimplementedBcryptServiceServer

	cost int
}

func NewBcryptRPCServer(cost int) *BcryptRPCServer {
	return &BcryptRPCServer{
		cost: cost,
	}
}

func (server *BcryptRPCServer) HashPassword(ctx context.Context, req *service.HashRequest) (*service.HashResponse, error) {
	password := req.GetPassword()

	start_time := time.Now()

	defer func() {
		elapsed := time.Since(start_time)
		log.Info().Any("elapsed_us", elapsed.Microseconds()).Msg("hash performance")
	}()

	hash, err := bcrypt.GenerateFromPassword([]byte(password), server.cost)
	if err != nil {
		log.Error().Err(err).Msg("bcrypt hash error")
		return &service.HashResponse{}, err
	}

	return &service.HashResponse{Hash: string(hash)}, nil
}
