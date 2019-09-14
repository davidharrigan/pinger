// +build integration

package pinger

import (
	"context"
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"

	pb "github.com/davidharrigan/integration-test/grpc/protos"
	"github.com/davidharrigan/integration-test/testing/config"
)

func TestPinger(t *testing.T) {

	type expectation struct {
		out *pb.PingResponse
		err error
	}

	tcs := map[string]struct {
		in       *pb.PingRequest
		expected expectation
	}{
		"ok": {
			in: &pb.PingRequest{},
			expected: expectation{
				out: &pb.PingResponse{
					Payload: []byte(`pong`),
				},
			},
		},
	}

	for scenario, tc := range tcs {
		t.Run(scenario, func(t *testing.T) {
			assert := assert.New(t)

			address := fmt.Sprintf("%s:%d", config.PingerConfig.Address, config.PingerConfig.Port)
			conn, err := grpc.Dial(address, grpc.WithInsecure())
			assert.Nil(err)
			defer conn.Close()

			c := pb.NewPingerClient(conn)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			out, err := c.Ping(ctx, tc.in)
			assert.Nil(err)

			if tc.expected.err == nil {
				assert.Nil(err)
				assert.Equal(tc.expected.out, out)
			} else {
				assert.Nil(out)
				assert.Equal(tc.expected.err, err)
			}
		})
	}
}

func TestPingerStream(t *testing.T) {

	type expectation struct {
		count int
		out   *pb.PingResponse
		err   error
	}

	tcs := map[string]struct {
		in       *pb.PingRequest
		expected expectation
	}{
		"ok": {
			in: &pb.PingRequest{Count: 5},
			expected: expectation{
				count: 5,
				out: &pb.PingResponse{
					Payload: []byte(`pong`),
				},
			},
		},
	}

	for scenario, tc := range tcs {
		t.Run(scenario, func(t *testing.T) {
			assert := assert.New(t)

			address := fmt.Sprintf("%s:%d", config.PingerConfig.Address, config.PingerConfig.Port)
			conn, err := grpc.Dial(address, grpc.WithInsecure())
			if !assert.Nil(err) {
				return
			}
			defer conn.Close()

			c := pb.NewPingerClient(conn)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			stream, err := c.PingStream(ctx, tc.in)
			if !assert.Nil(err) {
				return
			}

			for i := 0; i < tc.expected.count; i++ {
				out, err := stream.Recv()
				if tc.expected.err == nil {
					assert.Nil(err)
					assert.Equal(tc.expected.out, out)
				} else {
					assert.Nil(out)
					assert.Equal(tc.expected.err, err)
				}
			}

			_, err = stream.Recv()
			assert.Equal(io.EOF, err)
		})
	}
}
