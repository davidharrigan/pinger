package pinger

import (
	"context"
	"log"

	pb "github.com/davidharrigan/integration-test/grpc/protos"
)

// Pinger repesents our grpc service that can respond to a simple ping request
type Pinger struct{}

// Ping repliese with single ping response
func (s *Pinger) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PingResponse, error) {
	log.Printf("Received")
	return &pb.PingResponse{
		Payload: []byte(`pong`),
	}, nil
}

// PingStream replies a stream of ping responses
func (s *Pinger) PingStream(in *pb.PingRequest, stream pb.Pinger_PingStreamServer) error {
	return nil
}
