package client

import (
	"context"
	"errors"
	"net"
	"reflect"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	pb "github.com/wwgberlin/grpc/salute"
	"google.golang.org/grpc"
)

type serveFunc func(ctx context.Context, in *pb.SaluteRequest) (*pb.SaluteResponse, error)

func (f serveFunc) Salute(ctx context.Context, in *pb.SaluteRequest) (*pb.SaluteResponse, error) {
	return f(ctx, in)
}

func TestSend_Success(t *testing.T) {
	const address = ":9090"
	req := pb.SaluteRequest{
		Saluters:    []string{"we", "the", "people"},
		Salutee:     "You",
		VideoUrl:    "some video",
		LogoUrl:     "some image",
		Url:         "some url",
		Description: "You know why",
		Timestamp:   &timestamp.Timestamp{Seconds: 1},
	}
	resp := pb.SaluteResponse{Body: "hello world", ResponseType: pb.ResponseType_text}

	lis, err := net.Listen("tcp", address)
	require.NoError(t, err)

	fn := func(ctx context.Context, in *pb.SaluteRequest) (*pb.SaluteResponse, error) {
		if req.Salutee != in.Salutee ||
			!reflect.DeepEqual(req.Saluters, in.Saluters) ||
			req.LogoUrl != in.LogoUrl ||
			req.VideoUrl != in.VideoUrl ||
			req.Url != in.Url ||
			req.Description != in.Description ||
			req.Timestamp.Seconds != in.Timestamp.Seconds {
			t.Fatal("Requests were expected to be equal")
		}

		return &resp, nil
	}

	//serve
	s := grpc.NewServer()
	pb.RegisterSaluterServer(s, serveFunc(fn))

	ready := make(chan struct{})
	go func() {
		close(ready)
		s.Serve(lis)
	}()

	<-ready

	//test the client
	r, err := Send(address, &req, context.Background())
	require.NoError(t, err)

	if !reflect.DeepEqual(*r, resp) {
		t.Fatalf("Responses were expected to be equal")
	}

	//release resources
	lis.Close()
	s.Stop()
}

func TestSend_Error(t *testing.T) {
	const address = ":9090"
	const retErr = "some error"

	lis, err := net.Listen("tcp", address)
	require.NoError(t, err)

	//fail the operation
	fn := func(ctx context.Context, in *pb.SaluteRequest) (*pb.SaluteResponse, error) {
		return nil, errors.New(retErr)
	}

	//serve
	s := grpc.NewServer()
	pb.RegisterSaluterServer(s, serveFunc(fn))

	ready := make(chan struct{})
	go func() {
		close(ready)
		s.Serve(lis)
	}()

	<-ready

	//test the client
	_, err = Send(address, &pb.SaluteRequest{}, context.Background())
	require.Error(t, err)

	//release resources
	lis.Close()
	s.Stop()
}

func TestSend_ContextCancellation(t *testing.T) {
	const address = ":9090"
	const retErr = "some error"

	lis, err := net.Listen("tcp", address)
	require.NoError(t, err)

	//fail the operation
	fn := func(ctx context.Context, in *pb.SaluteRequest) (*pb.SaluteResponse, error) {
		return nil, errors.New(retErr)
	}

	//serve
	s := grpc.NewServer()
	pb.RegisterSaluterServer(s, serveFunc(fn))

	ready := make(chan struct{})
	go func() {
		close(ready)
		s.Serve(lis)
	}()

	<-ready

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	//test the client
	_, err = Send(address, &pb.SaluteRequest{}, ctx)
	require.Error(t, err)

	//release resources
	lis.Close()
	s.Stop()
}

func TestSend_ContextTimeout(t *testing.T) {
	const address = ":9090"

	const retErr = "some error"

	lis, err := net.Listen("tcp", address)
	require.NoError(t, err)

	//fail the operation
	var called bool
	fn := func(ctx context.Context, in *pb.SaluteRequest) (*pb.SaluteResponse, error) {
		called = true
		time.Sleep(time.Second)
		return nil, errors.New(retErr)
	}

	//serve
	s := grpc.NewServer()
	pb.RegisterSaluterServer(s, serveFunc(fn))

	ready := make(chan struct{})
	go func() {
		close(ready)
		s.Serve(lis)
	}()

	<-ready

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//test the client
	_, err = Send(address, &pb.SaluteRequest{}, ctx)
	require.Error(t, err)
	assert.True(t, called)

	//release resources
	lis.Close()
	s.Stop()
}
