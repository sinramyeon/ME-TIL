package main

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/wwgberlin/grpc/client"
	pb "github.com/wwgberlin/grpc/salute"
)

var req pb.SaluteRequest

func init() {
	var err error
	dat, err := ioutil.ReadFile("message.dat")
	if err != nil {
		panic(err)
	}
	if err = proto.Unmarshal(dat, &req); err != nil {
		panic(err)
	}
}

const serverAddr = ":8081"

func main() {
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./web/public"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		resp, err := client.Send(serverAddr, &req, context.Background())
		if err != nil {
			panic(err)
		}
		w.Write([]byte(resp.Body))

	})
	http.ListenAndServe(":8080", nil)
}
