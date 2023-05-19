package logger

import (
	"io"
	"os"

	log "github.com/go-kit/kit/log"
)

var (
	// Logger is the shared logger instance
	Logger log.Logger
)

func init() {
	Initialize(os.Stdout)
}

// Initialize initializes the logger with the given output writer
func Initialize(out io.Writer) {
	Logger = log.NewLogfmtLogger(os.Stderr)
	Logger = log.With(Logger, "ts", log.DefaultTimestampUTC) // TODO:// change to gmt
	Logger = log.With(Logger, "caller", log.DefaultCaller)
	// Logger = log.New(out, "", log.LstdFlags|log.LUTC|log.Lshortfile)
}
