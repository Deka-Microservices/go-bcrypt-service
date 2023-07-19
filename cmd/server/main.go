package main

import (
	"net"

	_ "github.com/deka-microservices/go-bcrypt-service/internal/config"
	"github.com/deka-microservices/go-bcrypt-service/internal/consts"
	"github.com/deka-microservices/go-bcrypt-service/internal/server"
	"github.com/deka-microservices/go-bcrypt-service/pkg/service"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var Version = "dev"

func main() {
	log.Info().Str("version", Version).Msg("current version")

	address := viper.GetString(consts.CONFIG_IP) + ":" + viper.GetString(consts.CONFIG_PORT)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal().Err(err).Msg("tcp listen")
	}

	s := grpc.NewServer()
	bcrypt_server := server.NewBcryptRPCServer(viper.GetInt(consts.CONFIG_COST))
	service.RegisterBcryptServiceServer(s, bcrypt_server)

	log.Info().Str("address", lis.Addr().String()).Msg("listen start")
	if err := s.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("serve failure")
	}

}
