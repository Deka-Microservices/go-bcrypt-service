package server

import (
	"context"
	"errors"
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

func (server *BcryptRPCServer) CheckPassword(ctx context.Context, req *service.CheckRequest) (*service.CheckResponse, error) {
	password := req.GetPassword()
	hash := req.GetHash()

	start_time := time.Now()

	defer func() {
		elapsed := time.Since(start_time)
		log.Info().Any("elapsed_us", elapsed.Microseconds()).Msg("check performance")
	}()

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	valid := err == nil
	out := &service.CheckResponse{Valid: valid}

	if !valid {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return out, nil
		}
		return out, err
	}

	return out, nil
}
