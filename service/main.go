package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/davidharrigan/pinger/grpc/protos"
	"github.com/davidharrigan/pinger/service/pinger"
)

const port = ":50051"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterPingerServer(s, &pinger.Pinger{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
