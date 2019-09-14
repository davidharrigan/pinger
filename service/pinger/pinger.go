package pinger

import (
	"context"
	"log"

	pb "github.com/davidharrigan/pinger/grpc/protos"
)

// Pinger repesents our grpc service that can respond to a simple ping request
type Pinger struct{}

// Ping repliese with single ping response
func (s *Pinger) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PingResponse, error) {
	log.Println("received")
	return &pb.PingResponse{
		Payload: []byte(`pong`),
	}, nil
}

// PingStream replies a stream of ping responses
func (s *Pinger) PingStream(in *pb.PingRequest, stream pb.Pinger_PingStreamServer) error {
	log.Println("received")
	for i := 0; i < int(in.Count); i++ {
		log.Println("streaming...")
		resp := &pb.PingResponse{
			Payload: []byte(`pong`),
		}
		if err := stream.Send(resp); err != nil {
			return err
		}
	}
	return nil
}
