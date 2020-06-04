package main

import (
	"fmt"
	"log"
)

func main() {
	in := make(chan int, 2) // CALL runtime.makechan(SB)
	in <- 1                 // CALL runtime.chansend1(SB)
	close(in)               // CALL runtime.closechan(SB)
	_, ok := <-in           // CALL runtime.chanrecv2(SB)

	// LEAQ runtime.staticbytes(SB), DX
	// LEAQ go.string.*+794(SB), AX
	// CALL log.Printf(SB)
	log.Printf("ok:%v", ok)
	// LEAQ runtime/internal/sys.DefaultGoroot.str+376(SB), CX
	// LEAQ runtime.staticbytes(SB), CX
	// CALL fmt.Println(SB)
	fmt.Println("ok:", ok)
}
