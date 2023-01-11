package main

import (
	"fmt"
	"io"
	"os"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func main() {
	Logger.Error("error....")
	Logger.Info("info....")
	Logger.Warn("warn....")
}

func init() {
	initLogger()
}

func initLogger() {
	// 注意: 这里要是绝对路径, 否则 roatelogs WithLinkName  后的日志为空
	logFile := "/tmp/00.log"
	if !Exists(logFile) {
		file, err := os.Create(logFile)
		defer file.Close()
		if err != nil {
			fmt.Println("mkdir logPath err!", err)
			return
		}
	}
	encoder := initEncoder()

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(getWriter(logFile)), zap.LevelEnablerFunc(func(zapcore.Level) bool { return true })),
	)
	Logger = zap.New(core, zap.AddCaller())
}

//初始化Encoder
func initEncoder() zapcore.Encoder {
	//return zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
	return zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		TimeKey:     "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		CallerKey:    "file",
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})
}

//日志文件切割
func getWriter(filename string) io.Writer {
	// 保存30天内的日志，每24小时(整点)分割一次日志
	hook, err := rotatelogs.New(
		filename+".%Y%m%d",
		// 注意: filename这里要是绝对路径, 否则 roatelogs WithLinkName  后的日志为空
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*30),
		rotatelogs.WithRotationTime(time.Hour*24),
	)

	if err != nil {
		panic(err)
	}
	return hook
}

//查看文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
