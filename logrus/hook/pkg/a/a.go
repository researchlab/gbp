package a 
import (
	log "github.com/sirupsen/logrus"
)

var aLog = log.WithField("key","aLog")
func Aoo(){
	log.Info("Aoo test")
	aLog.Info("ALog test")
}
