package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// func NewZapLogger() *zap.SugaredLogger {
// 	cfg := zap.Config{
// 		Encoding:    "json",                              // json hoặc console
// 		Level:       zap.NewAtomicLevelAt(zap.InfoLevel), // InfoLevel cho phép log ở cả 3 level
// 		OutputPaths: []string{"stderr"},
// 		// Cấu hình phần logging
// 		EncoderConfig: zapcore.EncoderConfig{
// 			MessageKey:   "message",
// 			LevelKey:     "level",
// 			CallerKey:    "caller",
// 			TimeKey:      "time",
// 			EncodeTime:   CustomTimeEncoder,                // format hiển thị thời gian log
// 			EncodeCaller: zapcore.FullCallerEncoder,        // lấy dòng code bắt đầu log
// 			EncodeLevel:  zapcore.CapitalColorLevelEncoder, // format cách hiển thị level
// 		},
// 	}

// 	logger, _ := cfg.Build() // build ra logger từ config
// 	return logger.Sugar()
// }

type Logger *zap.SugaredLogger

func NewZapLogger() *zap.SugaredLogger {
	writer := getLogWriter()
	encoder := getEncoder()

	core := zapcore.NewCore(encoder, writer, zapcore.DebugLevel)

	logger := zap.New(core)
	return logger.Sugar()
}

func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("./logs/app.log") // Tạo file
	return zapcore.AddSync(file)           // Ghi log vào file
}

func getEncoder() zapcore.Encoder {
	// NewConsoleEncoder hoặc NewJSONEncoder
	return zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:   "message",
		LevelKey:     "level",
		CallerKey:    "caller",
		TimeKey:      "time",
		EncodeTime:   CustomTimeEncoder,           // format hiển thị thời gian log
		EncodeCaller: zapcore.ShortCallerEncoder,  // lấy dòng code bắt đầu log
		EncodeLevel:  zapcore.CapitalLevelEncoder, // format cách hiển thị level
	})
}

func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}
