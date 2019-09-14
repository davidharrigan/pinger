package main

import (
	"context"
	"log"
	"net"

	pb "github.com/davidharrigan/integration-test/grpc/protos"
	"google.golang.org/grpc"
)

const port = ":50051"

type server struct{}

func (s *server) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PingResponse, error) {
	log.Printf("Received: %v")
	return &pb.PingResponse{}, nil
}

func (s *server) PingStream(in *pb.PingRequest, stream pb.Pinger_PingStreamServer) error {
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterPingerServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
