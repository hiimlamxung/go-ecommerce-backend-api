package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// Demo:
	// Zap có 3 mode:
	// - Example: dùng cho development
	// - Development: dùng cho development
	// - Production: dùng cho production
	// Run: go run cmd/cli/main.log.go để test thử

	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("Hello name: %s, age: %d", "Lâm", 27) // giống fmt.Printf(format, a...)

	// logger
	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("name", "Lâm"), zap.Int("age", 27)) // là kiểu key-value, chứ ko giống Println của go

	// logger := zap.NewExample()
	// logger.Info("Hello Example")

	// logger, _ = zap.NewDevelopment()
	// logger.Info("Hello Development")

	// logger, _ = zap.NewProduction()
	// logger.Info("Hello Production")

	// Code chính
	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller()) // zap.AddCaller(): thêm thông tin về file và dòng số của vị trí log

	logger.Info("Info Log", zap.Int("line", 1), zap.Int("my_custom_field", 1))
	logger.Error("Error Log", zap.Int("line", 2), zap.Int("my_custom_field", 2))
}

// Customize log
func getEncoderLog() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"                          // key thời gian log: ts => time
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder   // Chuyển đổi format thời gian log: 1748074313.7744408 => 2025-05-24T15:11:53.773+0700
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // Viết hoa level: info => INFO
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder // Chỉ lấy tên file và dòng số của vị trí log: "caller":"cli/main.log.go:22"
	return zapcore.NewJSONEncoder(encoderConfig)
}

// Ghi log vào file
func getWriterSync() zapcore.WriteSyncer {
	// Giải thích 3 tham số của os.OpenFile:
	// - ./log/log.txt: đường dẫn + tên file
	// - Cách thức mở file: os.O_RDWR là mở file để đọc và ghi
	// - Quyền hạn của file: os.ModePerm -> tương đương với 0777 (quyền đọc, ghi, thực thi cho owner, group, và others)
	file, _ := os.OpenFile("./log/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)

	syncFile := zapcore.AddSync(file)

	// Tạo một WriteSyncer để ghi log ra console (màn hình terminal)
	// os.Stderr: Là standard error stream - nơi thường dùng để xuất thông báo lỗi và log
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncFile, syncConsole)
}
