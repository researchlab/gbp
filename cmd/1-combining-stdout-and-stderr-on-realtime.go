package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("/bin/bash", "-c", "ls /")

	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("stdout pipe error", err)
		return
	}
	cmd.Stderr = cmd.Stdout

	done := make(chan bool, 1)
	go func() {
		stdoutReader := bufio.NewReader(cmdReader)
		for {
			line, err := stdoutReader.ReadString('\n')
			if err != nil {
				if err != io.EOF {
					fmt.Println("read error", err)
				}
				break
			}
			fmt.Print("msg: ", line)
		}
		//done <- true
		close(done)
	}()
	if err := cmd.Start(); err != nil {
		log.Fatalf("could not start cmd: %v", err)
	}
	err = cmd.Wait()
	if err != nil {
		log.Fatalf("could not wait for cmd: %v", err)
	}
	fmt.Println("wait finished")
	// 需要注意 cmd.Wait() 在 done 之前到来
	// 所以需要在scanner 之后主动发送done
	<-done
}

/* output

msg: Applications
msg: Library
msg: System
msg: Users
msg: Volumes
msg: bin
msg: cores
msg: dev
msg: etc
msg: home
msg: opt
msg: private
msg: sbin
msg: tmp
wait finished
msg: usr
msg: var
*/
