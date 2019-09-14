package config

import (
	"fmt"
	"log"
	"math/rand"
	"net"

	"google.golang.org/grpc"

	pb "github.com/davidharrigan/pinger/grpc/protos"
	"github.com/davidharrigan/pinger/service/pinger"
)

// ServerConfig stores info of grpc server
type ServerConfig struct {
	Address string
	Port    int
}

func randomPort() int {
	min := 51000
	max := 55000
	return rand.Intn(max-min) + min
}

// PingerConfig stores info of pinger server config
var PingerConfig = ServerConfig{
	Address: "localhost",
	Port:    50000,
}

// Pinger returns a new pinger server for local testing
func Pinger() *grpc.Server {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", PingerConfig.Address, PingerConfig.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterPingerServer(s, &pinger.Pinger{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return s
}
