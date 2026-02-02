package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

var Log *slog.Logger

type PrettyHandler struct {
	slog.Handler
	level slog.Level
}

func (h *PrettyHandler) Handle(ctx context.Context, r slog.Record) error {
	// Emoji por nível
	var emoji string
	switch r.Level {
	case slog.LevelDebug:
		emoji = "🔍"
	case slog.LevelInfo:
		emoji = "ℹ️ "
	case slog.LevelWarn:
		emoji = "⚠️ "
	case slog.LevelError:
		emoji = "❌"
	default:
		emoji = "  "
	}

	// Timestamp simples
	timestamp := r.Time.Format("15:04:05")

	// Monta a mensagem
	msg := fmt.Sprintf("%s [%s] %s %s", emoji, timestamp, r.Level.String(), r.Message)

	// Adiciona os atributos
	r.Attrs(func(a slog.Attr) bool {
		msg += fmt.Sprintf(" %s=%v", a.Key, a.Value)
		return true
	})

	fmt.Println(msg)
	return nil
}

func (h *PrettyHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.level
}

func Init() {
	var level slog.Level

	switch os.Getenv("LOG_LEVEL") {
	case "DEBUG":
		level = slog.LevelDebug
	case "INFO":
		level = slog.LevelInfo
	case "WARN":
		level = slog.LevelWarn
	case "ERROR":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	var handler slog.Handler

	if os.Getenv("ENV") == "production" {
		// JSON em produção
		opts := &slog.HandlerOptions{Level: level}
		handler = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		// Pretty print em desenvolvimento
		handler = &PrettyHandler{
			Handler: slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level}),
			level:   level,
		}
	}

	Log = slog.New(handler)
}

// Fatal loga erro e sai do programa
func Fatal(msg string, args ...any) {
	Log.Error(msg, args...)
	os.Exit(1)
}

// Fatalf loga erro formatado e sai do programa
func Fatalf(format string, v ...any) {
	Log.Error(fmt.Sprintf(format, v...))
	os.Exit(1)
}
