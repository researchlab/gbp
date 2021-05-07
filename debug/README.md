- [golang代码调试](#golang代码调试)
  - [1.delve 安装](#1delve-安装)
  - [2.本地调试](#2本地调试)
    - [2.1 debug 方式](#21-debug-方式)
    - [2.2 attach 方式](#22-attach-方式)
  - [3.远程调试](#3远程调试)
  - [4.设置断点方式](#4设置断点方式)
    - [4.1 通过函数名](#41-通过函数名)
    - [4.2 通过文件+行号](#42-通过文件行号)
  - [5.调试过程](#5调试过程)
  - [6.命令说明](#6命令说明)
    - [6.1 常用命令](#61-常用命令)
    - [6.2 dlv 命令](#62-dlv-命令)
    - [6.3 dlv 调试会话命令](#63-dlv-调试会话命令)
  - [7.命令示例](#7命令示例)
    - [7.1 查看命令帮助](#71-查看命令帮助)
    - [7.2 高频命令](#72-高频命令)
      - [7.2.1  查找源码文件](#721--查找源码文件)
      - [7.2.2 查看源代码](#722-查看源代码)
      - [7.1.3 断点相关的操作](#713-断点相关的操作)
      - [7.1.4 执行相关的操作](#714-执行相关的操作)
      - [7.1.5  for循环调试利器 frame](#715--for循环调试利器-frame)
      - [7.1.6 coredump 排查利器](#716-coredump-排查利器)
    - [7.3 低频命令](#73-低频命令)
      - [7.3.1 线程相关操作](#731-线程相关操作)
      - [7.3.2 协程相关操作](#732-协程相关操作)
  - [附录](#附录)
    - [1.delve 调试容器化](#1delve-调试容器化)
    - [2.在vscode中借助delve 调试](#2在vscode中借助delve-调试)
    - [3.在goland 中借助delve 调试](#3在goland-中借助delve-调试)
    - [4.在vscode中借助delve 调试 k8s 应用](#4在vscode中借助delve-调试-k8s-应用)

### golang代码调试

gdb 和 delve 是两个可用来调试golang 代码的主要工具; 本文记录通过delve方式调试golang代码的过程;



#### 1.delve 安装

```shell
➜  ~ go get -u github.com/derekparker/delve/cmd/dlv
➜  ~ dlv version
Delve Debugger
Version: 1.6.0
Build: $Id: 8cc9751909843dd55a46e8ea2a561544f70db34d $
```



#### 2.本地调试

可通过debug 和attach 两种方式使用本地调试;

##### 2.1 debug 方式

```shell
# cd 到项目根目录下
➜  ~ dlv debug  path/main.go 
```

##### 2.2 attach 方式

对于已经运行的程序进程, 可以通过attach 程序进程pid 来进行调试

```shell
# 查询程序进程id 
ps aux |grep xxxx 

# attach pid 
dlv attach 29260 

```



#### 3.远程调试

:exclamation:**注意**

```diff
- 需要在远程服务器和本地都安装dlv工具;
- 本机dlv 通过输入quit 结束调试,  会让你选择是否关闭调试进程, 选择N ,不要关闭远程的调试进程PID ;
- 如果本地没有dlv connect,  直接kill掉远程服务器上的dlv, 会导致调试进程PID 直接退出， 所以要谨慎; 
```

在远程服务器上

```shell
# 查询目标进程PID
ps aux |grep {目标进程名称}
# attach 到目标进程 并开启远程调试端口8181 
dlv attach $PID --headless --api-version=2 --log --listen=:8181 
```

在本地调试端

```shell
# 连接到远程服务器上的dlv进程, 进入本地调试
dlv connect 远程服务器ip:8181 
```



#### 4.设置断点方式

[测试代码](https://github.com/researchlab/gbp/blob/master/debug/02-dlv-tester.go)

##### 4.1 通过函数名

```shell
# 通过 b 包名.函数名 设置断点 
b main.main 
b main.worker 
```

##### 4.2 通过文件+行号

```shell
➜  dlv debug 02-dlv-tester.go
Type 'help' for list of commands.
(dlv) b ./02-dlv-tester.go:22
Breakpoint 1 (enabled) set at 0x10c488f for main.main() ./02-dlv-tester.go:22
(dlv)
```



#### 5.调试过程

1、dlv debug main.go #debug一个main.go程序

2、break（b）main.main #在main包里的main函数入口打断点

3、continue（c） #继续运行，直到断点处停止

4、next（n） #单步运行

5、locals #打印local variables

6、print（p） #打印一个变量或者表达式

7、quit (q)  # Would you like to kill the process? [Y/n] n



#### 6.命令说明

##### 6.1 常用命令

```shell
# 常用操作
break 	      设置一个断点 
clear         删除断点 
clearall      删除所有的断点 
continue      运行到断点或程序终止 
next 	      跳到下一行 
on            在遇到断点时执行一个命令 
set           更改变量的值 
trace         设置跟踪点 
step          单步执行程序 
thread        切换到指定的线程 

# 打印命令
args 	      打印函数参数 
breakpoints   打印激活的断点信息 
funcs 	      打印函数列表 
goroutine     显示或更改当前goroutine 
goroutines    列出程序的全部goroutines 
list 	      显示源代码 
locals 	      打印局部变量 
print         评估表达式 
source        执行包含delve命令列表的文件 
sources       打印源文件列表 
stack         打印堆栈跟踪 
threads       打印每一个跟踪线程的信息 
types         打印类型列表 
vars          打印某个包内的(全局)变量 
```

##### 6.2 dlv 命令

```shell
➜  dlv --help
Delve 是Golang程序源代码级调试器

Delve 通过控制进程的执行, 评估变量及提供线程/goroutine状态, CPU寄存器状态等信息, 使你可以与程序进行交互

这个工具的目标是为调试golang程序提供简单而强大的接口

Pass flags to the program you are debugging using `--`, for example:

`dlv exec ./hello -- server --config conf/config.toml`

Usage:
  dlv [command]

Available Commands:
  attach      连接到正在运行的进程并开始调试.
  connect     连接到headless(headless server which means running a server without monitor and keyboard, just like a remote server without a separate monitor)调试服务器.
  core        Examine a core dump.
  debug       编译并调试当前目录下的主包或指定包
  exec        执行预编译的二进制文件，并开始调试会话.
  help        Help about any command
  run         Deprecated command. Use 'debug' instead.
  test        Compile test binary and begin debugging program.
  trace       Compile and begin tracing program.
  version     Prints version.

Flags:
      --accept-multiclient   Allows a headless server to accept multiple client connections. Note that the server API is not reentrant（可重入） and clients will have to coordinate.
      --api-version int      Selects API version when headless. (default 1)
      --backend string       Backend selection:
	default		Uses lldb on macOS, native everywhere else.
	native		Native backend.
	lldb		Uses lldb-server or debugserver.
	rr		Uses mozilla rr (https://github.com/mozilla/rr).
 (default "default")
      --build-flags string   Build flags, to be passed to the compiler.
      --headless             'Run debug server only, in headless mode'.
      --init string          Init file, executed by the terminal client.
  -l, --listen string        Debugging server listen address. (default "localhost:0")
      --log                  Enable debugging server logging.
      --log-output string    指定调试日志输出内容，多个内容通过逗号分割, 可能的值如下:
					debugger	'Log debugger commands'
					gdbwire		Log connection to gdbserial backend
					lldbout		Copy output from debugserver/lldb to standard output
					debuglineerr	Log recoverable errors reading .debug_line
					rpc		'Log all RPC messages'
					fncall		Log function call protocol
					minidump	Log minidump loading
Defaults to "debugger" when logging is enabled with --log.
      --wd string            Working directory for running the program. (default ".")

Use "dlv [command] --help" for more information about a command.
```

core 即对一个golang 的core dump 文件进行回溯

```shell
注：什么是core dump
文件？当一份代码编译后运行一段时间会发生崩溃，但是又很难定位错误时，较原始的办法是不停地在一些关键代码上报日志；而一个更方便的方法则是通过设置环境变量GOTRACEBACK=crash
，生成一份进程运行直到崩溃时详细信息的快照，然后对这个快照进行回溯。
```

##### 6.3 dlv 调试会话命令

```shell
Type 'help' for list of commands.
(dlv) help
The following commands are available:

Running the program:
    call ------------------------ Resumes process, injecting a function call (EXPERIMENTAL!!!)
    continue (alias: c) --------- Run until breakpoint or program termination.
    next (alias: n) ------------- Step over to next source line.
    rebuild --------------------- Rebuild the target executable and restarts it. It does not work if the executable was not built by delve.
    restart (alias: r) ---------- Restart process.
    step (alias: s) ------------- Single step through program.
    step-instruction (alias: si)  Single step a single cpu instruction.
    stepout (alias: so) --------- Step out of the current function.('跳出当前函数, 当前函数中的断点也会被忽略')

Manipulating breakpoints:
    break (alias: b) ------- Sets a breakpoint.
    breakpoints (alias: bp)  Print out info for active breakpoints.
    clear ------------------ Deletes breakpoint.
    clearall --------------- Deletes multiple breakpoints.
    condition (alias: cond)  Set breakpoint condition.
    on --------------------- Executes a command when a breakpoint is hit.
    toggle ----------------- Toggles on or off a breakpoint.('开启或关闭某个断点')
    trace (alias: t) ------- Set tracepoint.

Viewing program variables and memory:
    args ----------------- Print function arguments.('打印函数参数')
    display -------------- Print value of an expression every time the program stops.
    examinemem (alias: x)  Examine memory:
    locals --------------- Print local variables.
    print (alias: p) ----- Evaluate an expression.('计算一个表达式')
    regs ----------------- Print contents of CPU registers.('打印CPU寄存器内容')
    set ------------------ Changes the value of a variable.('改变某个变量的值')
    vars ----------------- Print package variables.
    whatis --------------- Prints type of an expression.

Listing and switching between threads and goroutines:
    goroutine (alias: gr) -- Shows or changes current goroutine
    goroutines (alias: grs)  List program goroutines.
    thread (alias: tr) ----- Switch to the specified thread.
    threads ---------------- Print out info for every traced thread.

Viewing the call stack and selecting frames:
    deferred --------- Executes command in the context of a deferred call.
    down ------------- Move the current frame down.('向下移动一帧')
    frame ------------ Set the current frame, or execute command on a different frame.('设置当前帧，或在某一个帧上执行命令')
    stack (alias: bt)  Print stack trace.('打印堆栈信息, 可根据堆栈帧号调试')
    up --------------- Move the current frame up.('向上移动一帧')

Other commands:
    config --------------------- Changes configuration parameters.
    disassemble (alias: disass)  Disassembler.
    dump ----------------------- Creates a core dump from the current process state('core dump')
    edit (alias: ed) ----------- Open where you are in $DELVE_EDITOR or $EDITOR
    exit (alias: quit | q) ----- Exit the debugger.
    funcs ---------------------- Print list of functions.
    help (alias: h) ------------ Prints the help message.
    libraries ------------------ List loaded dynamic libraries
    list (alias: ls | l) ------- Show source code.
    source --------------------- Executes a file containing a list of delve commands
    sources -------------------- Print list of source files.
    types ---------------------- Print list of types

Type help followed by a command for full documentation.
(dlv)
```



#### 7.命令示例

[测试代码](https://github.com/researchlab/gbp/blob/master/debug/03-dlv-http-server.go)

##### 7.1 查看命令帮助

delve 提供了非常多实用的调试命令，**如何查看某一条命令的使用说明很重要**，如下

```shell
# 查看所有命令说明
(dlv) help
The following commands are available:
    help (alias: h) ------------- Prints the help message.
    list (alias: ls | l) -------- Show source code.
Type help followed by a command for full documentation.

# 查看具体某一条命令的使用说明  help [command]
(dlv) help list
Show source code.

	[goroutine <n>] [frame <m>] list [<linespec>]

Show source around current point or provided linespec.
(dlv)
```

##### 7.2 高频命令

高频命令，是调试过程中使用频率高，对众多场景能提供有用信息的命令; 

###### 7.2.1  查找源码文件

**命令说明**

`sources` 命令用来打印 源码文件路径， 可通过添加正则表达式过滤到合适的源码文件路径

**使用场景**

- 设置断点，在线上场景，包路径很深，则可以直接通过查找具体的源码路径加行号设置断点;
- 查看源代码，通过list命令查看源代码时, 需要提供编译时的源码绝对路径;

**使用示例**

```shell
# 查看sources 使用说明
(dlv) help sources
Print list of source files.

	sources [<regex>]

If regex is specified only the source files matching it will be returned.

# sources 正则表达式  过滤到合适的源码文件路径
(dlv) sources 03
/Users/lihong/workbench/dev/src/github.com/researchlab/gbp/debug/03-dlv-http-server.go
(dlv)
```

###### 7.2.2 查看源代码

**命令说明**

> list (alias: ls | l) -------- Show source code.

`list` 查看源代码， 默认显示指定**具体源码文件路径**的**具体行号**的上下5行源代码； 必须指定具体行号；

**使用场景**

- 查看某个具体文件，某个函数的具体源代码；

**使用示例**

```shell
# list 命令使用说明
(dlv) help list
Show source code.

	[goroutine <n>] [frame <m>] list [<linespec>]

Show source around current point or provided linespec.

# 查找具体源码文件编译路径
(dlv) sources 03
/Users/lihong/workbench/dev/src/github.com/researchlab/gbp/debug/03-dlv-http-server.go

# 方式一 ls 源代码编译绝对路径:行号  显示指定行号上下5行源代码
(dlv) ls /Users/lihong/workbench/dev/src/github.com/researchlab/gbp/debug/03-dlv-http-server.go:15
Showing /Users/lihong/workbench/dev/src/github.com/researchlab/gbp/debug/03-dlv-http-server.go:15 (PC: 0x135984f)
  10:		http.HandleFunc("/hi", hi)
  11:		fmt.Println("running on port:8081")
  12:		http.ListenAndServe(":8081", nil)
  13:	}
  14:
  15:	func hi(w http.ResponseWriter, r *http.Request) {
  16:		doHi(w)
  17:	}
  18:
  19:	func doHi(w http.ResponseWriter) {
  20:		hostName, _ := os.Hostname()
(dlv)
# 方式二 shows source code by package and function name
(dlv) list main.main
Showing /Users/lihong/workbench/dev/src/github.com/researchlab/gbp/debug/03-dlv-http-server.go:10 (PC: 0x135a023)
   5:		"net/http"
   6:		"os"
   7:		"strconv"
   8:	)
   9:
  10:	func main() {
  11:		http.HandleFunc("/hi", hi)
  12:		func01()
  13:		func02()
  14:		http.HandleFunc("/sum", sum)
  15:		fmt.Println("running on port:8081")
(dlv)
```

###### 7.1.3 断点相关的操作

**命令说明**

与断点相关的操作有

1. break/breakpoints 设置/查看断点;
2. clear/clearall 清除某一个(或所有)断点;
3. condition 设置条件断点;
4. on 命中断点时执行命令;



1.1**设置/查看断点**

> break (alias: b) ------------ Sets a breakpoint.

`break`  设置断点, 可以通过包名.函数名 设置断点, 也可以通过指定具体源码文件+具体行号 设置断点;

> breakpoints (alias: bp) ----- Print out info for active breakpoints.

1.2主要的断点设置方式

> delve 工具已经内置了panic 的断点设置;

1.2.0 设置断点方式说明

```shell
(dlv) help break
Sets a breakpoint.
  # break [断点名称]  断点具体位置
	break [name] <linespec>
	
# b 断点具体位置
# 这里没有指定断点名称, dlv默认会以数字递增的形式为断点命名, 如下文的1 就是默认的断点名称;
(dlv) b main.main
Breakpoint 1 set at 0x1359773 for main.main() ./03-dlv-http-server.go:9
# b 断点名称  断点具体位置
# 这里自定义了断点名称为 hiInterface 
(dlv) b hiInterface main.hi
Breakpoint hiInterface set at 0x135984f for main.hi() ./03-dlv-http-server.go:15
# 查看当前断点列表
(dlv) breakpoints
...
Breakpoint 1 at 0x1359773 for main.main() ./03-dlv-http-server.go:9 (0)
Breakpoint hiInterface at 0x135984f for main.hi() ./03-dlv-http-server.go:15 (0)
```

1.2.1通过  **包名.函数名**,  命令格式:  **b 包名.函数名**

```shell
(dlv) b main.main
Breakpoint 2 set at 0x1359773 for main.main() ./03-dlv-http-server.go:9
```

1.2.2通过 **具体文件+行号 **,  命令格式:  **b 具体源文件路径:行号**

```shell
# 通过绝对路径设置断点
(dlv) b /Users/lihong/workbench/dev/src/github.com/researchlab/gbp/debug/03-dlv-http-server.go:15
Breakpoint 3 set at 0x135984f for main.hi() ./03-dlv-http-server.go:15
# 注意 这里通过相对路径设置断点，但是不要加./ 前缀
(dlv) b 03-dlv-http-server.go:16
Breakpoint 7 set at 0x135985d for main.hi() ./03-dlv-http-server.go:16
```

1.2.3通过 **基于上文**, 命令格式: **b  +/-offset**

```shell
(dlv) c
> main.main() ./03-dlv-http-server.go:9 (hits goroutine(1):1 total:1) (PC: 0x1359773)
  ....
=>   9:	func main() {
    10:		http.HandleFunc("/hi", hi)
    11:		fmt.Println("running on port:8081")
  ....

(dlv) b +2  # 基于上文执行位置, 向后编译2行处加断点;
Breakpoint 4 set at 0x13597a6 for main.main() ./03-dlv-http-server.go:11
(dlv) c
> main.hi() ./03-dlv-http-server.go:15 (hits goroutine(19):1 total:1) (PC: 0x135984f)
    10:		http.HandleFunc("/hi", hi)
    11:		fmt.Println("running on port:8081")
    12:		http.ListenAndServe(":8081", nil)
    13:	}
    14:
=>  15:	func hi(w http.ResponseWriter, r *http.Request) {
  ....
(dlv) b -3 # 基于上文执行位置, 向前偏移3行处加断点;
Breakpoint 6 set at 0x1359808 for main.main() ./03-dlv-http-server.go:12 
```

1.2.3 通过 正则表达式 设置断点,  所有符合正则表达式的位置都会被设置断点， 不常用;

2.**清除某一个(或所有)断点**

> clear ----------------------- Deletes breakpoint.

> clearall -------------------- Deletes multiple breakpoints.

```shell
# 清除具体的断点
(dlv) clear hiInterface
Breakpoint hiInterface cleared at 0x135984f for main.hi() ./03-dlv-http-server.go:15
# 清除所有断点
(dlv) clearall
Breakpoint 1 cleared at 0x1359773 for main.main() ./03-dlv-http-server.go:9
```

3.**设置条件断点**

**命令说明**

> condition (alias: cond) ----- Set breakpoint condition.

```shell
(dlv) help condition
Set breakpoint condition.

	condition <breakpoint name or id> <boolean expression>.

Specifies that the breakpoint or tracepoint should break only if the boolean expression is true.
(dlv)
```

**使用场景**

- 在面对复杂的循环(for,while) 或者多条件并发执行的场景下, 条件断点非常有用，能帮助节省大量的时间; 

- 条件断点，并不是设置一个新断点， 而是在一个已知的断点上设置执行条件，那么当执行到这个断点处并且满足这个条件时并会触发断点

- 对同一个断点可以多次设置条件，以最后一次为准

**使用示例**

```shell
➜  debug git:(master) ✗ dlv debug 03-dlv-http-server.go
Type 'help' for list of commands.
(dlv) ls 03-dlv-http-server.go:25
Showing /Users/lihong/workbench/dev/src/github.com/researchlab/gbp/debug/03-dlv-http-server.go:25 (PC: 0x135a29c)
  20:			vals := r.URL.Query()
  21:			_num := vals.Get("num")
  22:			if len(_num) != 0 {
  23:				num, _ := strconv.Atoi(_num)
  24:				if num != 0 {
  25:					for i := 0; i <= num; i++ {
  26:						result += i
  27:					}
  28:				}
  29:			}
  30:		}
(dlv) b 03-dlv-http-server.go:26
Breakpoint 1 set at 0x135a2b5 for main.sum() ./03-dlv-http-server.go:26
# 对 断点1 设置条件  i==num-1 , 当i == num-1 时并会触发此断点
(dlv) cond 1 i==num-1
# 查看断点1 已被设置cond i == num-1 
(dlv) breakpoints
Breakpoint 1 at 0x135a2b5 for main.sum() ./03-dlv-http-server.go:26 (0)
	cond i == num-1
# 修改断点1 的条件设置
(dlv) cond 1 result==3
# 断点1 的条件以被修改为 cond result == 3 触发; 
(dlv) breakpoints
Breakpoint 1 at 0x135a2b5 for main.sum() ./03-dlv-http-server.go:26 (0)
	cond result == 3
```

4.on 命中断点时 执行命令

**命令说明**

```shell
(dlv) help on
Executes a command when a breakpoint is hit.

    on <breakpoint name or id> <command>.

Supported commands: print, stack and goroutine)
```

**使用示例**

```shell
(dlv) b /Users/user/go-learning/test.go:21 //添加断点
Breakpoint 2 set at 0x10b1368 for main.main() ./test.go:21
(dlv) on 2 print msg == "starting a gofunc" //设置当到达该断点后运行的命令，即判断msg的值是否为"starting a gofunc"
(dlv) breakpoints //查看此时的断点信息
Breakpoint runtime-fatal-throw at 0x102b3e0 for runtime.fatalthrow() /usr/local/Cellar/go/1.11.4/libexec/src/runtime/panic.go:654 (0)
Breakpoint unrecovered-panic at 0x102b450 for runtime.fatalpanic() /usr/local/Cellar/go/1.11.4/libexec/src/runtime/panic.go:681 (0)
    print runtime.curg._panic.arg
Breakpoint 2 at 0x10b1368 for main.main() ./test.go:21 (0)
    print msg == "starting a gofunc"
(dlv) c
Starting main
> main.main() ./test.go:21 (hits goroutine(1):1 total:1) (PC: 0x10b1368)
    msg == "starting a gofunc": true //'可见打印出来了，并返回值为true'
    16: func main() {
    17:     msg := "Starting main"
    18:     fmt.Println(msg)
    19:     bus := make(chan int)
    20:     msg = "starting a gofunc"
=>  21:     go counting(bus)
    22:     for count := range bus{
    23:         fmt.Println("count : ", count)
    24:     }
    25: }
```



**断点相关操作使用示例**

```shell
# break 使用说明
(dlv) help break
Sets a breakpoint.

	break [name] <linespec>

See $GOPATH/src/github.com/derekparker/delve/Documentation/cli/locspec.md for the syntax of linespec.

# 设置断点方式1 通过包名.函数名 设置断点
(dlv) b main.main
Breakpoint 2 set at 0x1359773 for main.main() ./03-dlv-http-server.go:9

# 错误示例  通过文件行号设置断点，相对路径不要加 ./ 前缀
(dlv) b ./03-dlv-http-server.go:15
Command failed: Location "./03-dlv-http-server.go:15" not found

# 注意 这里通过相对路径设置断点，但是不要加./ 前缀
(dlv) b 03-dlv-http-server.go:16
Breakpoint 7 set at 0x135985d for main.hi() ./03-dlv-http-server.go:16

# 设置断点方式2 通过源码编译时的绝对路径+具体行号 设置断点;
(dlv) b /Users/lihong/workbench/dev/src/github.com/researchlab/gbp/debug/03-dlv-http-server.go:15
Breakpoint 3 set at 0x135984f for main.hi() ./03-dlv-http-server.go:15

# 查看断点 
# 可以看到delve 工具已经内置了 panic 函数的断点;
(dlv) breakpoints
Breakpoint unrecovered-panic at 0x1039590 for runtime.fatalpanic() /usr/local/go/src/runtime/panic.go:1185 (0)
	print runtime.curg._panic.arg
Breakpoint 2 at 0x1359773 for main.main() ./03-dlv-http-server.go:9 (0)
Breakpoint 3 at 0x135984f for main.hi() ./03-dlv-http-server.go:15 (0)
(dlv) c
> main.main() ./03-dlv-http-server.go:9 (hits goroutine(1):1 total:1) (PC: 0x1359773)
     4:		"fmt"
     5:		"net/http"
     6:		"os"
     7:	)
     8:
=>   9:	func main() {
    10:		http.HandleFunc("/hi", hi)
    11:		fmt.Println("running on port:8081")
    12:		http.ListenAndServe(":8081", nil)
    13:	}
    
# 设置断点方式3  从上一次continue后 的某个偏移位置处设置断点;
# 如上次执行到第9行，偏移2行设置断点，即设置第11行位置断点;
(dlv) b +2
Breakpoint 4 set at 0x13597a6 for main.main() ./03-dlv-http-server.go:11
(dlv) c
> main.main() ./03-dlv-http-server.go:11 (hits goroutine(1):1 total:1) (PC: 0x13597a6)
     6:		"os"
     7:	)
     8:
     9:	func main() {
    10:		http.HandleFunc("/hi", hi)
=>  11:		fmt.Println("running on port:8081")
    12:		http.ListenAndServe(":8081", nil)
    13:	}
    14:
    15:	func hi(w http.ResponseWriter, r *http.Request) {
    16:		doHi(w)
(dlv)
```

###### 7.1.4 执行相关的操作

与执行到下一步的相关操作，

1. 执行到下一个断点或程序退出;
2. 执行到下一行（不进入函数内部);
3. 进入/退出函数内部；

1.**执行到下一个断点或直到程序退出**

**命令说明**

> ​    continue (alias: c) --------- Run until breakpoint or program termination.

2.**执行到下一行**(不进入函数内部)

**命令说明**

> ​    next (alias: n) ------------- Step over to next source line.

3.**进入/退出函数内部**

**命令说明**

> step (alias: s) ------------- Single step through program.
>
> stepout --------------------- Step out of the current function.



###### 7.1.5  for循环调试利器 frame

对源码进行调试时，偶尔会产生需要回顾之前运行过的代码片段的信息需求，就好比下面的斐波那契递归，在调试下已经递归了20次，然后需要查看第18次的信息，这时候就可以通过查阅frame来满足需求。

frame实际上是栈帧，即记录了每一个函数调用过程的信息帧。可以用`stack`
或者简称`bt`
来获取栈上的帧信息。

每一帧都记录了执行的函数对应的文件地址，可以据此判断哪一帧才是我们需要的。

然后使用`frame 帧编号`
的方式进入该帧，结合上面的命令查看该帧中的相应信息。

[测试代码](https://github.com/researchlab/gbp/blob/master/debug/04-fib.go)

```shell
(dlv) b ./04-fib.go:13
Breakpoint 1 (enabled) set at 0x10c2854 for main.FibIter() ./04-fib.go:13
(dlv) bp
...
Breakpoint 1 (enabled) at 0x10c2854 for main.FibIter() ./04-fib.go:13 (0)
(dlv) c
> main.FibIter() ./04-fib.go:13 (hits goroutine(21):1 total:2) (PC: 0x10c2854)
> main.FibIter() ./04-fib.go:13 (hits goroutine(18):1 total:2) (PC: 0x10c2854)
     8:	func FibIter(a, b, n int) int {
     9:		if n == 0 {
    10:			return b
    11:		}
    12:
=>  13:		return FibIter(a+b, a, n-1)
    14:	}
    15:
    16:	func Fib(n int) int {
    17:		return FibIter(1, 0, n)
    18:	}
(dlv) n
...
(dlv) stack
0  0x00000000010c2854 in main.FibIter
   at ./04-fib.go:13
1  0x00000000010c287e in main.FibIter
   at ./04-fib.go:13
2  0x00000000010c287e in main.FibIter
   at ./04-fib.go:13
3  0x00000000010c287e in main.FibIter
   at ./04-fib.go:13
4  0x00000000010c287e in main.FibIter
   at ./04-fib.go:13
5  0x00000000010c28f6 in main.Fib
   at ./04-fib.go:17
6  0x00000000010c2a4a in main.main.func1
   at ./04-fib.go:25
7  0x0000000001063a91 in runtime.goexit
   at /usr/local/go/src/runtime/asm_amd64.s:1373
(dlv) args
a = 5
b = 3
n = 96
~r3 = 0
(dlv) bt
0  0x00000000010c2854 in main.FibIter
   at ./04-fib.go:13
1  0x00000000010c287e in main.FibIter
   at ./04-fib.go:13
2  0x00000000010c287e in main.FibIter
   at ./04-fib.go:13
3  0x00000000010c287e in main.FibIter
   at ./04-fib.go:13
4  0x00000000010c287e in main.FibIter
   at ./04-fib.go:13
5  0x00000000010c28f6 in main.Fib
   at ./04-fib.go:17
6  0x00000000010c2a4a in main.main.func1
   at ./04-fib.go:25
7  0x0000000001063a91 in runtime.goexit
   at /usr/local/go/src/runtime/asm_amd64.s:1373
(dlv) frame 1
> main.FibIter() ./04-fib.go:13 (hits goroutine(25):2 total:59) (PC: 0x10c2854)
> main.FibIter() ./04-fib.go:13 (hits goroutine(18):5 total:59) (PC: 0x10c2854)
Frame 1: ./04-fib.go:13 (PC: 10c287e)
     8:	func FibIter(a, b, n int) int {
     9:		if n == 0 {
    10:			return b
    11:		}
    12:
=>  13:		return FibIter(a+b, a, n-1)
    14:	}
    15:
    16:	func Fib(n int) int {
    17:		return FibIter(1, 0, n)
    18:	}
(dlv) args
a = 3
b = 2
n = 97
~r3 = 0
(dlv) frame 3
> main.FibIter() ./04-fib.go:13 (hits goroutine(25):2 total:59) (PC: 0x10c2854)
> main.FibIter() ./04-fib.go:13 (hits goroutine(18):5 total:59) (PC: 0x10c2854)
Frame 3: ./04-fib.go:13 (PC: 10c287e)
     8:	func FibIter(a, b, n int) int {
     9:		if n == 0 {
    10:			return b
    11:		}
    12:
=>  13:		return FibIter(a+b, a, n-1)
    14:	}
    15:
    16:	func Fib(n int) int {
    17:		return FibIter(1, 0, n)
    18:	}
(dlv) args
a = 1
b = 1
n = 99
~r3 = 0
(dlv)
```

###### 7.1.6 coredump 排查利器

**命令说明**

coredump 文件是复杂故障排查利器;

`coredump`是一个包含程序意外终止时的内存快照的文件。它可以用于事后调试，以了解崩溃发生的原因以及其中涉及的变量。通过GOTRACEBACK，Go提供了一个环境变量来控制程序崩溃时产生的输出。这个变量还可以强制生成`coredump`，这样以来就使得调试成为可能。

**GOTRACEBACK**

`GOTRACEBACK`可以用于控制程序程序崩溃时输出内容的多少，它可以有以下一些取值：

- `none` 不显示goroutine的堆栈调用的trace
- `single` （默认值）显示当前goroutine的堆栈调用的trace
- `all`显示用户创建的所有的goroutine的堆栈调用的trace
- `system`显示包含运行时goroutine及其它所有goroutine的堆栈调用的trace
- `crash`类似于`system`，不同的是，它同时也会产生coredump

最后一个选项使我们能够在程序崩溃时进行调试。如果你没有得到足够的日志，或者崩溃无法重现，这可能是一个不错的选择。

**生成 coredump 文件条件说明**

1. `ulimit -c unlimited ` 修改 系统生成coredump 文件大小不受限制, 默认为0 ，表示不会生成coredump 文件，但是在macos 系统, delve 1.6.0 上实际测试 不设置这个命令不受影响

```shell
# 可以看到 此时系统充许生成的core file size 为0， 即不充许生成coredump文件
# 实际dlv dump 时 似乎不受影响;
# 下面实验保持 如下系统默认值 不修改;
➜  ulimit -a
-t: cpu time (seconds)              unlimited
-f: file size (blocks)              unlimited
-d: data seg size (kbytes)          unlimited
-s: stack size (kbytes)             8192
-c: core file size (blocks)         0
-v: address space (kbytes)          unlimited
-l: locked-in-memory size (kbytes)  unlimited
-u: processes                       2784
-n: file descriptors                256
```

2. 设置`GOTRACEBACK` 环境变量， 然后在delve 1.6.0 版本中似乎不设置这个变量没有影响;

**coredump 实验过程**

测试代码

```go
package main

import "math/rand"

func main() {
    sum := 0
    for {
        n := rand.Intn(1e6)
        sum += n
        if sum%42 == 0 {
            panic(":(")
        }
    }
}
```

执行程序

```shell
➜  go run 05-core-dump.go
panic: :(

goroutine 1 [running]:
main.main()
	/Users/lihong/workbench/dev/src/github.com/researchlab/gbp/debug/05-core-dump.go:11 +0x88
exit status 2
```

但这里的问题是，我们无法从堆栈调用的trace中判断出具体是哪个值引起的崩溃。当然了，您可以通过添加日志的方法去一步步的定位问题具体出现在哪里，但我们并不总是能知道在我们应该把日志添加在哪里。另外呢，当一个问题无法重现时，写测试用例并确保它被修复也是相当困难的。

总结一下上面的思路：在添加日志和运行应用程序之间迭代，直到它崩溃，并查看可能的原因运行后。

是否有其它的办法呢？答案是：有的。

我们可以用GOTRACEBACK=crash再运行一次，这样将有比较详细的输出了，因为我们现在已经打印出了所有的goroutines，包括运行时的。除此之外，也生成了相应的`coredump`文件

:exclamation: 注意

**Currently supports linux/amd64 and linux/arm64 core files, windows/amd64 minidumps and core files generated by Delve's 'dump' command.**

生成coredump文件 , `dlv debug 源文件` 启动debug 并生成coredump 文件

```shell
# 步骤一 启动调试器
➜  dlv debug 05-core-dump.go
Type 'help' for list of commands.
# 步骤二 设置断点
(dlv) b main.main
Breakpoint 1 (enabled) set at 0x1060913 for main.main() ./05-core-dump.go:5
# 步骤三 执行到断点
(dlv) c
> main.main() ./05-core-dump.go:5 (hits goroutine(1):1 total:1) (PC: 0x1060913)
     1:	package main
     2:
     3:	import "math/rand"
     4:
=>   5:	func main() {
     6:		sum := 0
     7:		for {
     8:			n := rand.Intn(1e6)
     9:			sum += n
    10:			if sum%42 == 0 {
# 步骤四 继续执行 触发panic
# 此时注意, 在调试文件夹目录下 会生成一个 __debug_bin 二进制文件 这个就是程序crash后会自动生成的;
# 这个 __debug_bin 二进制文件 包含了当前进程中的goroutines等信息, 在dlv core调试时需要用到
# 这个 __debug_bin  在退出当前dlv 调试会话时会自动被删除， 所以在dlv core时不要退出当前dlv 调试会话
(dlv) c
> [unrecovered-panic] runtime.fatalpanic() /usr/local/go/src/runtime/panic.go:1185 (hits goroutine(1):1 total:1) (PC: 0x102d730)
Warning: debugging optimized function
	runtime.curg._panic.arg: interface {}(string) ":("
  1180:	// fatalpanic implements an unrecoverable panic. It is like fatalthrow, except
  1181:	// that if msgs != nil, fatalpanic also prints panic messages and decrements
  1182:	// runningPanicDefers once main is blocked from exiting.
  1183:	//
  1184:	//go:nosplit
=>1185:	func fatalpanic(msgs *_panic) {
  1186:		pc := getcallerpc()
  1187:		sp := getcallersp()
  1188:		gp := getg()
  1189:		var docrash bool
  1190:		// Switch to the system stack to avoid any stack growth, which
 # 查看dump 命令使用
(dlv) help dump
Creates a core dump from the current process state
  # dump 输出dump文件名称
	dump <output file>

The core dump is always written in ELF, even on systems (windows, macOS) where this is not customary. For environments other than linux/amd64 threads and registers are dumped in a format that only Delve can read back.
# 设置coredump 文件名为 05-coredump 
(dlv) dump 05-coredump
Dumping memory 2542632960 / 2542632960...
# 注意 这个dlv 调试会话放在这里不要退出， 另外开一个新的dlv core 进行coredump 调试
(dlv)
```

新看一个dlv core 调试会话进行coredump 调试

```shell
# 新开一个dlv core  调试会话
# 命令格式 dlve core <executable_file> <core file> 
# 注意 这里的 <executable_file> 一定要用上面程序crash后自动生成的__debug_bin
# 因为只有这个__debug_bin包含了上面这个进程当时的goroutines 信息， 如果用自己go build 生成的二进制文件时没有包含这些goroutines 运行时信息的;
➜  dlv core __debug_bin 05-coredump
Type 'help' for list of commands.
# 查看进程当时的goroutines 信息;
# * Goroutine 1  这个前面的* 表示进程crash的时候正在执行的goroutine; 没有什么特殊作用;
# 通过查看这些goroutine 后面跟的源码文件，看到和问题差不多的源码文件比如这里的./05-core-dump.go:11 大概就知道 时执行这个goroutine 出现的crash
(dlv) goroutines
* Goroutine 1 - User: ./05-core-dump.go:11 main.main (0x10609a2) (thread 565295)
  Goroutine 2 - User: /usr/local/go/src/runtime/proc.go:305 runtime.gopark (0x102ff8b) [force gc (idle)]
  Goroutine 3 - User: /usr/local/go/src/runtime/proc.go:305 runtime.gopark (0x102ff8b) [GC sweep wait]
  Goroutine 4 - User: /usr/local/go/src/runtime/proc.go:305 runtime.gopark (0x102ff8b) [GC scavenge wait]
[4 goroutines]
# 切换到上面定位的可能(一般就是)引发crash的goroutine 1 
(dlv) goroutine 1
Switched from 1 to 1 (thread 565295)
# 查看这个goroutine 的堆栈信息; 
# 一般也是找里面的源码文件是业务文件的位置;
# 如下 查看到 栈帧 2 位置有和业务相关的源码文件  at ./05-core-dump.go:11
(dlv) bt
0  0x000000000102d730 in runtime.fatalpanic
   at /usr/local/go/src/runtime/panic.go:1185
1  0x000000000102d151 in runtime.gopanic
   at /usr/local/go/src/runtime/panic.go:1060
2  0x00000000010609a2 in main.main
   at ./05-core-dump.go:11
3  0x000000000102fbd8 in runtime.main
   at /usr/local/go/src/runtime/proc.go:203
4  0x000000000105a911 in runtime.goexit
   at /usr/local/go/src/runtime/asm_amd64.s:1373
# 切换到 栈帧 2 
# 可以看到 程序crash 的位置在 11行  直接panic 导致的; 
(dlv) frame 2
> runtime.fatalpanic() /usr/local/go/src/runtime/panic.go:1185 (PC: 0x102d730)
Warning: debugging optimized function
Frame 2: ./05-core-dump.go:11 (PC: 10609a2)
     6:		sum := 0
     7:		for {
     8:			n := rand.Intn(1e6)
     9:			sum += n
    10:			if sum%42 == 0 {
=>  11:				panic(":(")
    12:			}
    13:		}
    14:	}
# 查看当前栈帧 局部变量的值; 
(dlv) locals
sum = 5705994
n = 203300
(dlv)
```

至此 整个coredump 排查故障过程就结束了; 

**清理工作**

结束后 关闭第一个dlv debug 调试会话 __debug_bin 会被自动删除， 生成的05-coredump文件需要手动删除; 

**coredump 实验02**

下面是一个相对更为复杂的coredump 排查示例; 

[测试代码](https://github.com/researchlab/gbp/blob/master/debug/06-core-dump-nil.go)

调试过程遇到crash, 生成__debug_bin 文件

```shell
# 启动dlv debug 调试会话 
➜  dlv debug 06-core-dump-nil.go
Type 'help' for list of commands.
# 设置断点
(dlv) b main.main
Breakpoint 1 (enabled) set at 0x10c27b3 for main.main() ./06-core-dump-nil.go:10
# 执行并触发crash 
(dlv) c
> main.main() ./06-core-dump-nil.go:10 (hits goroutine(1):1 total:1) (PC: 0x10c27b3)
     5:		"fmt"
     6:		"time"
     7:		"unsafe"
     8:	)
     9:
=>  10:	func main() {
    11:		for i := 1; i <= 10; i++ {
    12:			go func(gid int) {
    13:				n := 0
    14:				for {
    15:					fmt.Println(time.Now().Format("2006-01-02 15:04:05"), gid, n)
(dlv) c
1
1
myfun2
# 触发crash 
> main.myfun3() ./06-core-dump-nil.go:48 (PC: 0x10c2aaf)
Command failed: bad access
# 生成 coredump 文件
(dlv) dump 06-coredump
Dumping memory 2546270208 / 2546270208...
# 其实此时 可以通过ls 查看当前源代码执行位置
# 可以看到crash 的代码位置为 48行
(dlv) ls
> main.myfun3() ./06-core-dump-nil.go:48 (PC: 0x10c2aaf)
    43:	}
    44:
    45:	func myfun3() {
    46:		var p uintptr = 0
    47:		arr := (*int)(unsafe.Pointer(p))
=>  48:		*arr = 1
    49:		fmt.Println(*arr)
    50:	}
# 查看此时的局部变量，发现试图将值赋值给nil 因而造成了panic 
(dlv) locals
p = 0
arr = *int nil
(dlv)
```

通过上面其实可以查看到问题是 给nil 赋值 造成的， 但是我们也可以通过coredump 来查看问题， 

```shell
# 同样 新开一个 dlv core 调试会话
➜  dlv core __debug_bin 06-coredump
Type 'help' for list of commands.
# 查看所有goroutines, 可以看到只有goroutine 29 是业务代码文件相关的;
(dlv) goroutines
  Goroutine 1 - User: /usr/local/go/src/runtime/time.go:198 time.Sleep (0x10531fa) [sleep]
  Goroutine 2 - User: /usr/local/go/src/runtime/proc.go:305 runtime.gopark (0x103623b) [force gc (idle)]
  Goroutine 3 - User: /usr/local/go/src/runtime/proc.go:305 runtime.gopark (0x103623b) [GC sweep wait]
  Goroutine 4 - User: /usr/local/go/src/runtime/proc.go:305 runtime.gopark (0x103623b) [GC scavenge wait]
  Goroutine 17 - User: /usr/local/go/src/runtime/proc.go:305 runtime.gopark (0x103623b) [finalizer wait]
  Goroutine 18 - User: /usr/local/go/src/runtime/sys_darwin.go:64 syscall.syscall (0x1051fd2)
  Goroutine 19 - User: /usr/local/go/src/runtime/sema.go:71 sync.runtime_SemacquireMutex (0x1045437) [semacquire]
  Goroutine 20 - User: /usr/local/go/src/runtime/sema.go:71 sync.runtime_SemacquireMutex (0x1045437) [semacquire]
  Goroutine 21 - User: /usr/local/go/src/runtime/sema.go:71 sync.runtime_SemacquireMutex (0x1045437) [semacquire]
  Goroutine 22 - User: /usr/local/go/src/sync/mutex.go:124 sync.(*Mutex).lockSlow (0x107a913) (thread 588116)
  Goroutine 23 - User: /usr/local/go/src/runtime/sema.go:71 sync.runtime_SemacquireMutex (0x1045437) [semacquire]
  Goroutine 24 - User: /usr/local/go/src/runtime/sema.go:71 sync.runtime_SemacquireMutex (0x1045437) [semacquire]
  Goroutine 25 - User: /usr/local/go/src/runtime/sema.go:71 sync.runtime_SemacquireMutex (0x1045437) [semacquire]
  Goroutine 26 - User: /usr/local/go/src/runtime/sema.go:71 sync.runtime_SemacquireMutex (0x1045437) [semacquire]
  Goroutine 27 - User: /usr/local/go/src/runtime/sema.go:71 sync.runtime_SemacquireMutex (0x1045437) [semacquire]
  Goroutine 29 - User: ./06-core-dump-nil.go:48 main.myfun3 (0x10c2aaf) (thread 588009)
[16 goroutines]
# 切换到 goroutine 29 
(dlv) goroutine 29
Switched from 0 to 29 (thread 588009)
# 查看 goroutine 29 下面的堆栈信息; 
# 可以看到与业务代码相关的栈帧由0 和1，此时一般从上往下排查，即先从栈帧0开始排查；
(dlv) bt
0  0x00000000010c2aaf in main.myfun3
   at ./06-core-dump-nil.go:48
1  0x00000000010c2a54 in main.myfun2
   at ./06-core-dump-nil.go:42
2  0x0000000001063c61 in runtime.goexit
   at /usr/local/go/src/runtime/asm_amd64.s:1373
# 切换到栈帧0 
(dlv) frame 0
> main.myfun3() ./06-core-dump-nil.go:48 (PC: 0x10c2aaf)
Frame 0: ./06-core-dump-nil.go:48 (PC: 10c2aaf)
    43:	}
    44:
    45:	func myfun3() {
    46:		var p uintptr = 0
    47:		arr := (*int)(unsafe.Pointer(p))
=>  48:		*arr = 1
    49:		fmt.Println(*arr)
    50:	}
# 查看局部变量值，可知 原因是 不能为Nil变量赋值1
(dlv) locals
p = 0
arr = *int nil
(dlv)
```



##### 7.3 低频命令

###### 7.3.1 线程相关操作

golang 的goroutine 非常方便且资源开销非常低，通常一个(主)线程足够了，多线程场景之一是调用shell/python 脚本; 

**命令说明**

> thread (alias: tr) ---------- Switch to the specified thread.
> threads --------------------- Print out info for every traced thread.

**使用示例**

```shell
(dlv) help threads
Print out info for every traced thread.
(dlv) help thread
Switch to the specified thread.

	thread <id>
# 查看当前进程的所有线程;
(dlv) threads
* Thread 1520324 at :0
# 线程切换;
(dlv) thread 1520324
Switched from 1520324 to 1520324
```

###### 7.3.2 协程相关操作

协程是golang 并发的常见操作，但是协程由golang自己负责调度，所以一般也用的少

**命令说明**

> goroutine ------------------- Shows or changes current goroutine
> goroutines ------------------ List program goroutines.

**使用示例**

```shell
(dlv) help goroutine
Shows or changes current goroutine

	goroutine
	goroutine <id>
	goroutine <id> <command>

Called without arguments it will show information about the current goroutine.
Called with a single argument it will switch to the specified goroutine.
Called with more arguments it will execute a command on the specified goroutine.
(dlv) help goroutines
List program goroutines.

	goroutines [-u (default: user location)|-r (runtime location)|-g (go statement location)|-s (start location)] [ -t (stack trace)]

Print out info for every goroutine. The flag controls what information is shown along with each goroutine:

	-u	displays location of topmost stackframe in user code
	-r	displays location of topmost stackframe (including frames inside private runtime functions)
	-g	displays location of go instruction that created the goroutine
	-s	displays location of the start function
	-t	displays stack trace of goroutine

If no flag is specified the default is -u.
```



#### 附录

##### 1.delve 调试容器化

[Dockerized golang  with delve debugger](https://github.com/researchlab/gbp/tree/master/debug/docker)

##### 2.在vscode中借助delve 调试

[debug with vscode](https://github.com/researchlab/gbp/tree/master/debug/debug-with-vscode)

##### 3.在goland 中借助delve 调试

[debug with goland](https://github.com/researchlab/gbp/tree/master/debug/debug-with-goland)

##### 4.在vscode中借助delve 调试 k8s 应用

[debug with vscode for k8s application](https://github.com/researchlab/gbp/tree/master/debug/debug-with-vscode-k8s)


