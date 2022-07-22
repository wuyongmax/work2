package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Lg *zap.Logger

func Init() {
	encoder := getEncoder()
	logF := os.Stdout
	core := zapcore.NewCore(encoder, logF, zapcore.InfoLevel)
	Lg = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(Lg)
	zap.L().Info("init successful")
	return
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}
