package a2 
import "github.com/sirupsen/logrus"

var a2Log = logrus.WithField("scope","a2Log")
func Doo(){
		a2Log.Info("Doo test")
		Eoo()
}

func Eoo(){
	a2Log.Info("Eoo test")
}
