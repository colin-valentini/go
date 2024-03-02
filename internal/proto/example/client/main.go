package main

// This package is a simple demonstration of gRPC client usage of the service
// declared in internal/proto/example/example.proto, using stub code generated
// by the proto compiler.

import (
	"context"
	"flag"
	"log"
	"math/rand"
	"time"

	pb "github.com/colin-valentini/go/internal/proto/example"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var serverAddr = flag.String("addr", "localhost:50051", "The server address in the format of host:port")

func getGreeting(client pb.HelloWorldClient, req *pb.GreetingRequest) *pb.GreetingResponse {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	greeting, err := client.GetGreeting(ctx, req)
	if err != nil {
		log.Fatalf("client.GetGreeting failed: %v", err)
	}
	return greeting
}

func main() {
	flag.Parse()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	log.Printf("Dialing HelloWorld server at %s...", *serverAddr)
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	log.Printf("Successfully connected to HelloWorld server")
	defer conn.Close()
	client := pb.NewHelloWorldClient(conn)

	log.Printf("Calling HelloWorld.GetGreeting with random ID...")
	greeting := getGreeting(client, &pb.GreetingRequest{
		Id: rand.Int31(), //nolint:gosec
	})
	log.Printf("Got greeting \"%s\" from server", greeting.GetMessage())

	log.Printf("Calling HelloWorld.GetGreeting with special ID...")
	canned_greeting := getGreeting(client, &pb.GreetingRequest{
		Id:       69420,
		Language: pb.Language_LANGUAGE_ITALIAN, //nolint:gosec
	})
	log.Printf("Got greeting \"%s\" from server", canned_greeting.GetMessage())

	log.Println("Client main completed successfully")
}
