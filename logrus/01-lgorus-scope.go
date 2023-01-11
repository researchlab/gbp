package main

import (
	"fmt"
	"os"

	"github.com/researchlab/gbp/logrus/scope/a/a1"
	"github.com/researchlab/gbp/logrus/scope/a/a2"
	"github.com/researchlab/gbp/logrus/scope/b"
	"github.com/researchlab/gbp/logrus/scope/logs"
)

var mLog = logs.Entry().WithField("scope", "main")

func main() {
	logFile, err := os.OpenFile("./data/scope.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeExclusive)
	if err != nil {
		fmt.Println("create file scope.log failed:", err) // 所以这一句日志也会丢失， 应该用fmt
	}
	defer logFile.Close()
	logs.Init(logFile)
	mLog.Info("main start")
	a1.Foo()
	b.Boo()
	a2.Doo()
}
