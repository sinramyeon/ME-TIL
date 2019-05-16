package main

import (
	"bytes"
	"context"
	"errors"
	"html/template"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"

	"github.com/golang/protobuf/ptypes"
	pb "github.com/wwgberlin/grpc/salute"
)

var path = "server/salute.tmpl"

type server struct{}

//Serve starts a grpc server and registers itself as a SaluterServer
//See example at: https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_server/main.go
func Serve(addr string) error {
	//implement here!
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	server := server{}

	s := grpc.NewServer()
	pb.RegisterSaluterServer(s, pb.SaluterServer{})

	if err := s.Serve(listen); err != nil {
		panic(err)
	}

}

func main() {
	log.Fatal(Serve(":8081"))
}

func (s *server) Salute(ctx context.Context, in *pb.SaluteRequest) (*pb.SaluteResponse, error) {
	if in.ResponseType != pb.ResponseType_HTML {
		return nil, errors.New("unsupported response format")
	}
	var b bytes.Buffer
	t := template.New("salute")
	t = t.Funcs(map[string]interface{}{
		"timestamp": func() string {
			return ptypes.TimestampString(in.Timestamp)
		},
		"join": func(s []string) string {
			return strings.Join(s, ", ")
		},
	})
	if err := template.Must(t.ParseFiles(path)).Execute(&b, in); err != nil {
		return nil, err
	}
	return &pb.SaluteResponse{Body: string(b.Bytes())}, nil
}
