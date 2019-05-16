package client

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	_ "google.golang.org/grpc"

	pb "github.com/wwgberlin/grpc/salute"
)

const address = "localhost:8081"

//Send takes an address to dial to, a request and a context.
//it opens an (insecure) connection to the server using the address and the context

//it creates a saluter client with the connection created

//it salutes (calls Salute on the client) giving the context and request args.

//Follow example: https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_client/main.go
func Send(address string, req *pb.SaluteRequest, ctx context.Context) (*pb.SaluteResponse, error) {
	//implement here!
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	c := pb.NewSaluterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Salute(ctx, &pb.SaluteRequest{Name: "Salute"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	//Don't forget to close the connection on exit!
	defer conn.Close()

}
