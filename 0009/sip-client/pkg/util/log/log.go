package log

import (
	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
	"os"
)

func init() {
	writers := make([]io.Writer, 0)
	writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
	writers = append(writers, newRollingFile("./sip.log"))
	mw := io.MultiWriter(writers...)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = zerolog.New(mw).With().Caller().Logger()
}

func newRollingFile(file string) io.Writer {
	return &lumberjack.Logger{
		Filename:   file,
		MaxBackups: 3,
		MaxSize:    500,
		MaxAge:     30,
	}
}
