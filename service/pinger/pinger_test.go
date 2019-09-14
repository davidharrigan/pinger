package pinger

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	pb "github.com/davidharrigan/pinger/grpc/protos"
)

func TestPing(t *testing.T) {
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

			p := &Pinger{}
			ctx := context.Background()

			out, err := p.Ping(ctx, tc.in)

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
