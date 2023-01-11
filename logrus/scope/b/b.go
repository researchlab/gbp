package b

import "github.com/researchlab/gbp/logrus/scope/logs"
import "github.com/researchlab/gbp/logrus/scope/c"

// commone 
var bLog = logs.Entry()

func Boo(){
	bLog.WithField("func","boo").Info("b boo test")
	c.Coo()
}
