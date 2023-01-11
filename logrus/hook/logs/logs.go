package logs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"runtime"

	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
)

const logDir = "./data"
const logPath = "hook"

type MyJSONFormatter struct {
	Time  string `json:"time"`
	File  string `json:"file"`
	Line  int    `json:"line"`
	Level string `json:"level"`
	Info  string `json:"info"`
	Msg   string `json:"msg"`
}

func (f *MyJSONFormatter) Format(entry *log.Entry) ([]byte, error) {
	logrusJF := &(log.JSONFormatter{})
	logrusJF.TimestampFormat = "2006-01-02 15:04:05.000"
	bytes, _ := logrusJF.Format(entry)
	json.Unmarshal(bytes, &f)
	if _, file, no, ok := runtime.Caller(8); ok {
		f.File = file
		f.Line = no
	}
	str := fmt.Sprintf("[%s] %s %s:%d %s\n", f.Level, f.Time, f.File, f.Line, f.Msg)
	return []byte(str), nil
}

func InitLog(level string) (err error) {
	err = setAppLog(level)
	return
}

func DefaultJSONFormatter()*log.JSONFormatter{
	return &log.JSONFormatter{
		TimestampFormat:"2006-01-02 15:04:05",
		//PrettyPrint:true, 表示Json展开打印
	}
}
func setAppLog(level string) (err error) {
	switch level {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	default:
		log.Fatal("log conf only allow [debug, info,warn,error], please check your confguire")
	}

	//os.MkdirAll(logDir+"/"+logPath, os.ModeDir)
	//var path = logDir + "/" + logPath + "/" + logPath + ".log"
	var path = logDir + "/" + logPath + ".log"
	log.SetOutput(ioutil.Discard)
	log.AddHook(lfshook.NewHook(
		lfshook.PathMap{
			log.InfoLevel:  path,
			log.DebugLevel: path,
			log.WarnLevel:  path,
			log.ErrorLevel: path,
			log.FatalLevel: path,
			log.PanicLevel: path,
		},
		//&MyJSONFormatter{}))
		DefaultJSONFormatter()))
	return nil
}
