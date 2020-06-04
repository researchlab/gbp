# golang channel tips 

写入 关闭的channel
> 向关闭的channel中写入数据时会触发painc操作

> panic 类型: runtime.plainError

> panic 值: "send on closed channel"

读取 关闭的channel
> 往关闭的channel中读数据时, 当读到第二个返回值为false时 就表示已无数据;

> v, ok := <- chan

#### 0-1-channel-close

```
go build -gcflags "-N -l" -ldflags=-compressdwarf=false -o 0-1-channel-close.out 0-1-channel-close.go
go tool objdump -s "main.main" 0-1-channel-close.out > 0-1-channel-close.S
```

汇编分析
```
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
```

#### 0-2-send-on-closed-channel
```
func main() {
	in := make(chan bool)
	close(in)
	go func() { <-in }()
	in <- true // panic: send on closed channel
}
```

#### 0-3-read-on-closed-channel
```
func main() {
	in := make(chan int, 2)
	close(in)
	_, ok := <-in
	log.Printf("ok:%v", ok) // ok:false
}
```
