package client

import (
	"context"
	"github.com/binarycp/go-grpc/chat"
	"google.golang.org/grpc"
	"log"
)

func run() {
	var (
		conn *grpc.ClientConn
		err  error
	)
	conn, err = grpc.Dial(":9999", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)
	hello, err := c.SayHello(context.Background(), &chat.Message{Body: "不记得了"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", hello.Body)
}
