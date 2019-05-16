package main

import (
	"context"
	"testing"

	"strings"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wwgberlin/grpc/client"
	pb "github.com/wwgberlin/grpc/salute"
)

func TestServe(t *testing.T) {
	const addr = ":9091"
	path = "salute.tmpl"

	req := pb.SaluteRequest{
		Saluters:    []string{"we", "the", "people"},
		Salutee:     "You",
		VideoUrl:    "some_video",
		LogoUrl:     "some_image",
		Url:         "some_url",
		Description: "You know why",
		Timestamp:   &timestamp.Timestamp{Seconds: 1},
	}

	ready := make(chan struct{})
	go func() {
		close(ready)
		Serve(addr)
	}()
	<-ready
	resp, err := client.Send(addr, &req, context.Background())
	require.NoError(t, err)

	assert.True(t, strings.Contains(resp.Body, req.VideoUrl))
	assert.True(t, strings.Contains(resp.Body, req.LogoUrl))
	assert.True(t, strings.Contains(resp.Body, req.Description))
	assert.True(t, strings.Contains(resp.Body, req.Salutee))
	assert.True(t, strings.Contains(resp.Body, req.Url))
	assert.True(t, strings.Contains(resp.Body, ptypes.TimestampString(req.Timestamp)))
	for i := 0; i < len(req.Saluters); i++ {
		assert.True(t, strings.Contains(resp.Body, req.Saluters[i]))
	}
}
