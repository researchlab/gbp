package main

import (
	"github.com/researchlab/gbp/logrus/hook/logs"
	"github.com/researchlab/gbp/logrus/hook/pkg/a"
	log "github.com/sirupsen/logrus"
)

func main() {
	logs.InitLog("debug")
	log.Info("main start")
	a.Aoo()
}
