package server

import (
	"github.com/binarycp/go-grpc/chat"
	"github.com/binarycp/gutils/strs"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

type Local struct {
	Remote
}

func NewLocal(port int) *Local {
	local := &Local{Remote{
		Port: port,
	}}
	for port < 63000 {
		if local.listenTcp() == nil {
			break
		}
		local.Port++
	}

	return local
}

func (l *Local) listenTcp() error {
	addr := strs.Concat("127.0.0.1:", strconv.Itoa(l.Port))
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return err
	}
	l.listen, err = net.ListenTCP("tcp", tcpAddr)
	return err
}

func (l *Local) Run() error {
	if l.listen == nil {
		return nil
	}

	server := grpc.NewServer()

	s := chat.Server{}
	chat.RegisterChatServiceServer(server, &s)
	err := server.Serve(l.listen)
	if err != nil {
		return err
	}
	return nil
}

func (l *Local) Close() error {
	return l.listen.Close()
}
