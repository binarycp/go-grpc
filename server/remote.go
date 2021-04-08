package server

import (
	"github.com/binarycp/go-grpc/chat"
	"github.com/binarycp/gutils/strs"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

type Remote struct {
	listen *net.TCPListener
	Port   int
	parent
}

func NewRemote(port int) *Remote {
	return &Remote{Port: port}
}

func (r *Remote) Run() error {
	addr := strs.Concat("127.0.0.1:", strconv.Itoa(r.Port))
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return err
	}

	r.listen, err = net.ListenTCP("tcp", tcpAddr)
	if r.listen != nil {
		defer r.listen.Close()
	}

	if err != nil {
		return err
	}

	server := grpc.NewServer()

	s := chat.Server{}
	chat.RegisterChatServiceServer(server, &s)
	err = server.Serve(r.listen)
	if err != nil {
		return err
	}
	return nil
}
