package a1 
import "github.com/researchlab/gbp/logrus/scope/a"
import "github.com/sirupsen/logrus"
func Foo(){
	a.ALog.Info("a1 foo test")	
	a.ALog.WithFields(logrus.Fields{
		"key":"a1",
		"size":10,
	}).Error("with fields test for a1")
}
