package server

import (
	"github.com/binarycp/go-grpc/chat"
	"github.com/binarycp/gutils/strs"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

func run(port int) error {
	listen, err := net.Listen("tcp", strs.Concat("127.0.0.1:", strconv.Itoa(port)))
	if listen != nil {
		defer listen.Close()
	}

	if err != nil {
		return err
	}

	server := grpc.NewServer()

	s := chat.Server{}
	chat.RegisterChatServiceServer(server, &s)
	err = server.Serve(listen)
	if err != nil {
		return err
	}
	return nil
}
