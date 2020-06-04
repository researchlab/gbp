# golang defer tips

A defer statement pushes a function call onto a list. The list of saved calls is executed after the surrounding function returns. Defer is commonly used to simplify functions that perform various clean-up actions.
> -- https://blog.golang.org/defer-panic-and-recover 

> defer表达式将一个函数调用保存在列表中，当包裹defer的函数"返回"后，列表中的调用会被执行。defer通常用于清理收尾工作。

#### 0-1-unamed-return
```
➜  go run 0-1-unamed-return.go
defer1: 1
defer2: 2
return: 0
```

```
func a() int {
	var i int
	defer func() {
		i++
		fmt.Println("defer2:", i) // 打印结果为 defer: 2
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i) // 打印结果为 defer: 1
	}()
	return i
}
```
> 函数a()  返回列表中只声明了返回的数据类型，并没声明变量, return 返回时首先将返回值i 赋值给返回变量，此处a()没有，golang语言会在此时先为其声明一个返回变量假如为j , 则return 先将i的值赋值给j(注意这里是值传递操作，而非引用操作), 然后交给defer 做清理动作，当defer 清理完成后 返回值j 才会给下一步代码调用; 
> 这里可以看到defer清理过程改变的是变量i, 而j不会被改变;

#### 0-2-named-return
```
➜  go run 0-2-named-return.go
defer1: 1
defer2: 2
return: 2
```

```
func b() (i int) {
	defer func() {
		i++
		fmt.Println("defer2:", i) // 打印结果为 defer: 2
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i) // 打印结果为 defer: 1
	}()
	return i // 或者直接 return 效果相同
}
```
> 函数b()返回列表中声明了返回变量为i, 所以在return赋值之后，defer中的修改也是在修改这个返回变量i, 所以最终i的值是两次defer修改之后的值;

#### 0-3-ptr-return
```
➜  go run 0-3-ptr-return.go
c defer1: 1
c defer2: 2
c return: 2
```
```
func c() *int {
	var i int
	defer func() {
		i++
		fmt.Println("c defer2:", i) // 打印结果为 c defer: 2
	}()
	defer func() {
		i++
		fmt.Println("c defer1:", i) // 打印结果为 c defer: 1
	}()
	return &i
}
```
> 函数c() 返回列表中也没有声明变量，与函数a()不同的是 函数c()返回的是地址引用值, 所以当return之后defer的清理操作修改的还是返回变量指向的地址空间中的值，所以最终返回变量的值也因为地址空间值被defer修改而改变;

#### 0-4-compilation-analysis
```
func a() int {
	var i int
	defer func() {
		i++
	}()
	return i
}
```
编译成汇编指令
```
go build -gcflags "-N -l" -ldflags=-compressdwarf=false -o 0-4-compilation-analysis.out 0-4-compilation-analysis.go
go tool objdump -s "main.a" 0-4-compilation-analysis.out >  0-4-compilation-analysis.S
```
> go tool objdump -s 参数"main.a" 表示指抓取namespace main 下a函数执行的汇编指令

```
  0-4-compilation-analysis.go:12	0x109d026		e805faf8ff		CALL runtime.deferprocStack(SB)		
  0-4-compilation-analysis.go:12	0x109d02b		85c0			TESTL AX, AX				
  0-4-compilation-analysis.go:12	0x109d02d		751c			JNE 0x109d04b				
  0-4-compilation-analysis.go:12	0x109d02f		eb00			JMP 0x109d031				
  0-4-compilation-analysis.go:15	0x109d031		488b442408		MOVQ 0x8(SP), AX			
  0-4-compilation-analysis.go:15	0x109d036		4889442470		MOVQ AX, 0x70(SP)			
  0-4-compilation-analysis.go:15	0x109d03b		90			NOPL					
  0-4-compilation-analysis.go:15	0x109d03c		e84f02f9ff		CALL runtime.deferreturn(SB)		
  0-4-compilation-analysis.go:15	0x109d041		488b6c2460		MOVQ 0x60(SP), BP			
  0-4-compilation-analysis.go:15	0x109d046		4883c468		ADDQ $0x68, SP				
  0-4-compilation-analysis.go:15	0x109d04a		c3			RET					
  0-4-compilation-analysis.go:12	0x109d04b		90			NOPL					
  0-4-compilation-analysis.go:12	0x109d04c		e83f02f9ff		CALL runtime.deferreturn(SB)		
  0-4-compilation-analysis.go:12	0x109d051		488b6c2460		MOVQ 0x60(SP), BP			
  0-4-compilation-analysis.go:12	0x109d056		4883c468		ADDQ $0x68, SP				
  0-4-compilation-analysis.go:12	0x109d05a		c3			RET					
  0-4-compilation-analysis.go:10	0x109d05b		e820c3fbff		CALL runtime.morestack_noctxt(SB)	
  0-4-compilation-analysis.go:10	0x109d060		e96bffffff		JMP main.a(SB)				
```
> 上述0-4-compilation-analysis.go:12 是文件名:行号,  可以清晰的看到代码中的某一行对应的汇编指令集;
> 0-4-compilation-analysis.go第12行defer func(){}  汇编指令集进行了CALL runtime.deferprocStack(SB)操作, 将defer 入栈;
>  第15行调用CALL runtime.deferreturn(SB) 执行return ,紧接着回到第12行 执行之前入栈的defer 函数； 当前main.a()函数调用结束，打印 CALL runtime.morestack_noctxt(SB) 表示当前main.a()上下文栈中以全部出栈，即最后跳转到第10行;

#### 1-1-unamed-return
```
➜  go run 1-1-unamed-return.go
golang
panic: unexpected EOF

goroutine 1 [running]:
main.main()
	/Users/lihong/workbench/dev/src/github.com/researchlab/gbp/defer/1-1-unamed-return.go:17 +0x280
exit status 2
```
```
func gzipFlush(data []byte) bytes.Buffer {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	defer w.Close()
	w.Write(data)
	w.Flush()
	return b
}
```
> flush() 返回的不是基本数据类型，而是结构体类型,  同样return 将b的值赋值给一个新的bytes.Buffer类型变量(注意这里是值传递而非引用), 之后虽然defer w.Close() 关闭了gzip.Writer类型,但此时defer影响的对象是b变量，而之前值传递操作的返回值变量不受此时的defer影响，所以导致最终在ioutil.ReaddAll(b)时无法正常读取到`End-of-File`而触发panic: unexpected EOF; 

#### 1-2-named-return
```
func flush(data []byte) (b bytes.Buffer) {
	w := gzip.NewWriter(&b)
	defer w.Close()
	w.Write(data)
	w.Flush()
	return b
}
```
> 有名返回， defer修改直接影响到返回值，所以最终ioutil.RealAll()可以正确读取到EOF
#### 1-3-ptr-return
```
func flush(data []byte) *bytes.Buffer {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	defer w.Close()
	w.Write(data)
	w.Flush()
	return &b
}
```
> 指针返回值，return之后的defer操作修改也能直接影响返回值变量，因为是引用操作;

#### 1-4-no-defer
```
func flush(data []byte) bytes.Buffer {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	w.Write(data)
	w.Flush()
	w.Close()
	return b
}
```
> 如果不使用defer, 只要在return之前 记得w.Close() 自然也不存在ioutil.ReadAll()读取不到EOF;

#### 1-5-determined-at-declaration
```
func calc(index string, a, b int) int {
    ret := a + b
    fmt.Println(index, a, b, ret)
    return ret
}

func main() {
    a := 1
    b := 2
    defer calc("1", a, calc("10", a, b))
    a = 0
    defer calc("2", a, calc("20", a, b))
    b = 1
}
```
> defer的参数在声明时即被确定下来, defer calc("1", a, calc("10", a, b)) 的第3个参数会在调用runtime.deferproc时确定，并不会在延时调用时才会被计算。

> defer 与 return 在一起的执行流程
> > step 1 : 在defer表达式的地方，会调用runtime.deferproc(size int32, fn *funcval)保存延时调用，注意这里保存了延时调用的参数
> > step 2 : 在return时，先将返回值保存起来
> > step 3 : 按FILO顺序调用runtime.deferreturn，即延时调用
> > step 4 : RET指令
