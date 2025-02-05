package logger

import (
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"
	"time"

	sloglogstash "github.com/samber/slog-logstash/v2"
	slogmulti "github.com/samber/slog-multi"
)

const (
	Default_file_channel = "default"
)

type LoggerHandler struct {
	fileName string
	osFile   *os.File
	logger   *slog.Logger
}

func NewLoggerHandler(channel_name string) *LoggerHandler {
	conn, err := net.Dial("tcp", "localhost:5000")

	if err != nil {
		log.Fatal("Failed to connect to syslog server", err)
	}
	logstashHandler := sloglogstash.Option{
		Level: slog.LevelDebug,
		Conn:  conn,
	}.NewLogstashHandler()

	fileName := getFullFilePath(channel_name)
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatal("Failed to open log file", err)
	}

	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}

	handler := slog.NewJSONHandler(file, opts)

	logger := slog.New(slogmulti.Fanout(
		logstashHandler,
		handler,
		slog.NewJSONHandler(os.Stdout, nil),
	))

	return &LoggerHandler{
		fileName: fileName,
		osFile:   file,
		logger:   logger,
	}
}

func (h *LoggerHandler) GetLogger() *slog.Logger {
	return h.logger
}

func (h *LoggerHandler) Close() {
	h.osFile.Close()
}

func getFullFilePath(channel_name string) string {
	if channel_name == Default_file_channel {
		return fmt.Sprintf("./logs/%s-logs%s.log", os.Getenv("APP_NAME"), time.Now().Format("2006.01.02"))
	}

	return fmt.Sprintf("./logs/%s/%s-logs%s.log", channel_name, os.Getenv("APP_NAME"), time.Now().Format("2006.01.02"))
}
