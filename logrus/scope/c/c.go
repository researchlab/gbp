package c 
import "github.com/researchlab/gbp/logrus/scope/logs"
var cLog = logs.Entry().WithField("scope","c")

func Coo(){
	cLog.Info("Coo test")
}
