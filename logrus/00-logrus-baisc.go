package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.Debug("lgorus new") // 这一句不会打印到文件中，因为此时还没有设置setOuput
	file, err := os.OpenFile("debug.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeExclusive)
	if err != nil {
		log.Fatalf("create file debug.log failed:%v", err)  // 所以这一句日志也会丢失， 应该用fmt
	}
	defer file.Close()
	log.SetLevel(logrus.DebugLevel)
	log.Formatter = &logrus.JSONFormatter{}
	log.SetOutput(file)
	log.WithField("key", "value").Info("test for withfield")
	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")
}
