package main

import (
	"bufio"
	"context"
	"fmt"
	"os/exec"
)

func main() {
	cancel, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	cmd0 := exec.CommandContext(cancel, "cmd", "/c", "go run ./son/main.go")

	//创建一个能获取此命令输出的管道
	stdout, err := cmd0.StdoutPipe()
	//stdin, err := cmd0.StdinPipe()
	if err != nil {
		fmt.Printf("Error: Can not obtain the stdout pipe for command No.0: %s\n", err)
		return
	}
	if err := cmd0.Start(); err != nil {
		fmt.Printf("Error: The command No.0 can not be startup: %s\n", err)
		return
	}

	buf := bufio.NewReader(stdout)
	readString, err := buf.ReadString('$')
	fmt.Println(readString, err)
}
