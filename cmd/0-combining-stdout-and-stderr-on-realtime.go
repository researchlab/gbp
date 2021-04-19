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
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatalf("could not get stderr pipe: %v", err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("could not get stdout pipe: %v", err)
	}
	done := make(chan bool, 1)
	go func() {
		merged := io.MultiReader(stderr, stdout)
		scanner := bufio.NewScanner(merged)
		for scanner.Scan() {
			msg := scanner.Text()
			fmt.Printf("msg: %s\n", msg)
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
	// 需要注意 cmd.Wait() 在 done 之前到来
	// 所以需要在scanner 之后主动发送done
	<-done
}
