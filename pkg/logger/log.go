package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Writer zapcore.WriteSyncer = os.Stdout

func Init(output zapcore.WriteSyncer) {
	logger := newLogger(output)
	zap.ReplaceGlobals(logger)
}

func Logger() *zap.Logger {
	return zap.L()
}

func newLogger(writer zapcore.WriteSyncer) *zap.Logger {
	atom := zap.NewAtomicLevel()

	encodeCfg := zap.NewProductionEncoderConfig()
	encodeCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	bl := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encodeCfg),
		zapcore.Lock(writer),
		atom,
	))

	l := bl.With(zap.String("out", "stdout"))
	return l
}
