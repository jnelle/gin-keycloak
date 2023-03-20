package ginkeycloak

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// The following code was copied from
// https://stackoverflow.com/a/70974263/20329236

var zapLog *zap.Logger

func init() {
	var err error
	config := zap.NewProductionConfig()
	enccoderConfig := zap.NewProductionEncoderConfig()
	zapcore.TimeEncoderOfLayout(time.RFC3339)
	enccoderConfig.StacktraceKey = "" // to hide stacktrace info
	config.EncoderConfig = enccoderConfig

	zapLog, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}
