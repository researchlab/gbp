package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(3)
	err := make(chan error, 3)
	go func() {
		defer wg.Done()
		err <- nil
	}()
	go func() {
		defer wg.Done()
		//time.Sleep(time.Second)
		err <- nil
	}()
	go func() {
		defer wg.Done()
		//err <- fmt.Errorf("error go")
		err <- nil
	}()
	go func() {
		wg.Wait()
		fmt.Println("wait")
		//time.Sleep(time.Second)
		close(err)
		_, ok := <-err
		if ok {
			fmt.Println("close err")
			close(err)
		}
	}()

	for e := range err {
		if e != nil {
			_, ok := <-err
			if ok {
				fmt.Println("close age")
				close(err)
			}
			fmt.Println("close", ok)
			return
		}
		fmt.Println(e)
	}
}
