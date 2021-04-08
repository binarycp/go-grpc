package server

import (
	"testing"
	"time"
)

func Test_run(t *testing.T) {
	remote := NewRemote(9999)
	go remote.Run()
	time.Sleep(time.Second * 2)
	t.Log(remote.Close())
}

func Test_local(t *testing.T) {
	remote := NewLocal(9999)
	t.Log(remote.Port)
	go remote.Run()

	time.Sleep(time.Second)
	local := NewLocal(remote.Port)
	t.Log(local.Port)
}
