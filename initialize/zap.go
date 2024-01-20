package initialize

import (
	"2024Winter/app/api/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

func SetupZap() {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{global.C.ZapInfo.Path}
	consoleOutput := zapcore.Lock(os.Stdout)
	fileOutput, _, err := zap.Open(global.C.ZapInfo.Path)
	if err != nil {
		panic("failed to open file: " + err.Error())
	}
	config.EncoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cores := []zapcore.Core{
		zapcore.NewCore(zapcore.NewConsoleEncoder(config.EncoderConfig), consoleOutput, config.Level),
		zapcore.NewCore(zapcore.NewJSONEncoder(config.EncoderConfig), fileOutput, config.Level),
	}

	core := zapcore.NewTee(cores...)
	global.Logger = zap.New(core)
	global.Logger.Info("zap initial success")
}
