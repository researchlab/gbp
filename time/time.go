package main
 
import (
    "time"
    "fmt"
)
 
func main() {
 
    go func() {
			var tt = time.Now()	
			for {
				fmt.Println("time:",tt)
				t := time.NewTimer(2 * time.Second)
				tt = <-t.C
			}
		}()
    time.Sleep(10 * time.Second)
		fmt.Println("finished")
	}
