package customlogger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewCustomLogger Lambda用のカスタムロガーを生成
func NewCustomLogger() (*zap.Logger, error) {
	var lvl zapcore.Level
	lvl = zap.InfoLevel

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfg),
		zapcore.AddSync(os.Stdout),
		lvl,
	)
	// 呼び出し元のファイル名と行番号をログメッセージに追加
	logger := zap.New(core, zap.AddCaller())
	return logger, nil
}

// WithRequestID は、ロガーにリクエストIDをフィールドとして追加します。
func WithRequestID(logger *zap.Logger, requestID string) *zap.Logger {
	return logger.With(zap.String("request_id", requestID))
}
