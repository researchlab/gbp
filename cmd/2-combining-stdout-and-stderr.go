package main

import (
	"fmt"
	"os/exec"
)

func main() {

	cmd := exec.Command("/bin/bash", "-c", "ls /")
	res, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(res))
		panic(err)
	}

	fmt.Println(string(res))
	cmd = exec.Command("/bin/bash", "-c", "ls /notfile")
	res, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(res))
		panic(err)
	}
	fmt.Println(string(res))
}
