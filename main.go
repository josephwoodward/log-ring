package main

import (
	"log/slog"
	"os"

	. "github.com/josephwoodward/log-ring/ringhandler"
)

func main() {

	handler := slog.NewJSONHandler(os.Stdout, nil)
	opts := &HandlerOptions{
		TailSize: 10,
		Level:    slog.LevelInfo,
		AttrKey:  "RequestId",
	}
	log := slog.New(NewLogTailHandler(handler, opts))

	// log.With()
	log.Debug("Debug 1", "RequestId", "123")
	log.Debug("Debug 2", "RequestId", "456")

	log.Debug("Debug 2", "RequestId", "123")
	log.Debug("Debug 5", "RequestId", "456")

	log.Info("Info 1", "RequestId", "123")
	log.Info("Info 2", "RequestId", "123")

	log.Warn("Warning 1", "RequestId", "123")
	log.Error("Boom!", "RequestId", "123")
}
