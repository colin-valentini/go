package main

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

func newGetGreetingRequest() *pb.GreetingRequest {
	return &pb.GreetingRequest{Id: rand.Int31()} //nolint:gosec
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

	log.Printf("Calling HelloWorld.GetGreeting...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	greeting, err := client.GetGreeting(ctx, newGetGreetingRequest())
	if err != nil {
		log.Fatalf("client.GetGreeting failed: %v", err)
	}
	log.Printf("Got greeting \"%s\" from server", greeting.GetMessage())

	log.Println("Client main completed successfully")
}
