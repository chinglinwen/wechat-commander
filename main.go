package main

import (
	"flag"
	"log"
	"net/http"

	pb "github.com/chinglinwen/wxrobot/api"
	"google.golang.org/grpc"
)

var (
	grpcAddress = flag.String("a", "localhost:50051", "default wxrobot grpc server address")
	client      pb.ApiClient
	port        = flag.String("p", ":4000", "this api listening port")
)

func main() {
	log.Println("starting...")

	// Set up a connection to the server.
	close := conn()
	defer close()

	http.HandleFunc("/", handler)
	http.HandleFunc("/ui", cmdHandler)
	http.HandleFunc("/text", textHandler)
	log.Fatal(http.ListenAndServe(*port, nil))
}

func conn() func() error {
	conn, err := grpc.Dial(*grpcAddress, grpc.WithInsecure())
	if err != nil {
		return nil
	}
	client = pb.NewApiClient(conn)
	return conn.Close
}
