# debug local 

```shell
# 创建项目目录
mkdir debug-local 

# 项目目录下初始化 go mod
go mod init github.com/researchlab/gbp/debug/local


# 此时项目结构
.
├── README.md
├── go.mod
├── main.go
└── utils
    ├── util.go
    └── util_test.go

# 写好代码之后， 在项目目录执行调试
dlv debug
main.go:6:2: no required module provides package github.com/researchlab/gbp/debug/debug-local/utils; to add it:
	go get github.com/researchlab/gbp/debug/debug-local/utils
exit status 1

# 报上述错误， 是有因为utils 文件还没有提交到github.com , 属于本地包

本次不先提交到github.com 而是添加本地包， 具体做法如下, 在go.mod 文件中添加如下两行  require xxx utils  和 replace xxx => ./utils 
cat go.mod
module github.com/researchlab/gbp/debug/local

go 1.17

require github.com/researchlab/gbp/debug/debug-local/utils v0.0.0
replace github.com/researchlab/gbp/debug/debug-local/utils  => ./utils

如此之后 还需要在utils 目录下执行 go mod  生存一条go.mod 文件, go mod xxx 这个xxx 随便起
go mod init github.com/researchlab/gbp/debug/local/utils

# 此时项目目录
tree -L 3
.
├── README.md
├── go.mod
├── main.go
└── utils
    ├── go.mod
    ├── util.go
    └── util_test.go

# 在项目目录执行 dlv debug 进入调试模式
```

dlv 断点调试main 包

*设置断点的两种方式*
- b 包名.函数名 
- b 文件名:行号  # 这里注意， 如果指定行号上没有有意义的代码，则设置不成功， 比如指定行号为空或者括弧等, 此外只需要指定文件名即可，不需要指定具体路径

```shell
$ dlv debug # 开始调试main包
Type 'help' for list of commands.
(dlv) b main.main # 在main.main 方法上打断点
Breakpoint 1 set at 0x10d0fd8 for main.main() ./main.go:8
(dlv) bp # 打印当前断点
Breakpoint runtime-fatal-throw at 0x1038420 for runtime.fatalthrow() /usr/local/go/src/runtime/panic.go:1162 (0)
Breakpoint unrecovered-panic at 0x10384a0 for runtime.fatalpanic() /usr/local/go/src/runtime/panic.go:1189 (0)
    print runtime.curg._panic.arg
Breakpoint 1 at 0x10d0fd8 for main.main() ./main.go:8 (0)
(dlv) c # 继续执行，直到下一个断点停止
> main.main() ./main.go:8 (hits goroutine(1):1 total:1) (PC: 0x10d0fd8)
     3: import (
     4:     "fmt"
     5:     "test/utils"
     6: )
     7:
=>   8: func main() {
     9:     for i := 0; i < 10; i++ {
    10:         fmt.Println(i)
    11:     }
    12:     fmt.Println(utils.Add(100, 2300))
    13: }
(dlv) b main.go:12 # 在main.go 文件中的12行打一个断点, 注意这里12行是utils.Add() 函数是有意义的;
Breakpoint 2 set at 0x10d10b8 for main.main() ./main.go:12
(dlv) n # 下一步，遇到函数不进入
> main.main() ./main.go:9 (PC: 0x10d0fef)
     4:     "fmt"
     5:     "test/utils"
     6: )
     7:
     8: func main() {
=>   9:     for i := 0; i < 10; i++ {
    10:         fmt.Println(i)
    11:     }
    12:     fmt.Println(utils.Add(100, 2300))
    13: }
(dlv) s # 单步执行程序，遇到函数进入
> main.main() ./main.go:10 (PC: 0x10d1007)
     5:     "test/utils"
     6: )
     7:
     8: func main() {
     9:     for i := 0; i < 10; i++ {
=>  10:         fmt.Println(i)
    11:     }
    12:     fmt.Println(utils.Add(100, 2300))
    13: }
(dlv) c
> main.main() ./main.go:12 (hits goroutine(1):1 total:1) (PC: 0x10d10b8)
     7:
     8: func main() {
     9:     for i := 0; i < 10; i++ {
    10:         fmt.Println(i)
    11:     }
=>  12:     fmt.Println(utils.Add(100, 2300))
    13: }
(dlv) c
2400
Process 53433 has exited with status 0
(dlv)
```
