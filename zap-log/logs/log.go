package logs

import (
	"fmt"
	"runtime"
)

func Error(err string) {
	var code, funcName string
	pc, codePath, codeLine, ok := runtime.Caller(1)
	if !ok {
		code = "-"
		funcName = "_"
	} else {
		code = fmt.Sprintf("%s:%d", codePath, codeLine)
		funcName = runtime.FuncForPC(pc).Name()
	}
	level := "ERROR"
	msg := err
	pathLine := code
	funcName = funcName
	logMsg := fmt.Sprintf("<%s>:<%s>:<%s>:<%s>", level, msg, pathLine, funcName)
	fmt.Println(logMsg)
}
