package channel

import (
	"fmt"
	"log"
	"reflect"
	"time"
)

var (
	errSendCloseChannel = "send on closed channel"
)

func CloseWriteChanException() (err error) {
	defer func() {
		r := recover()
		if r != nil {
			if excpt := fmt.Sprintf("%v", r); errSendCloseChannel != excpt {
				err = fmt.Errorf("UnkownError: %v", r)
			} else {
				log.Printf("RecoveryException: %v - %s\n", reflect.TypeOf(r), excpt)
			}
		}
	}()
	in := make(chan int, 2)
	go func() {
		time.Sleep(time.Second)
		close(in)
	}()

	for i := 0; i < 2; i++ {
		in <- i
		time.Sleep(time.Second)
	}
	return nil
}

func CloseReadChanException() (err error) {
	defer func() {
		r := recover()
		if r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	in := make(chan int, 2)
	go func() {
		time.Sleep(time.Second)
		close(in)
	}()

	for out := range in {
		_ = out
	}
	return nil
}
