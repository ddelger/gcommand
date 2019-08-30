package logging

import (
	"fmt"
	"io"
	"log"

	"github.com/ddelger/glog"
	"gopkg.in/natefinch/lumberjack.v2"
)

func Initialize(level glog.Level, options *Options) io.WriteCloser {
	l := &lumberjack.Logger{
		Filename:   options.Filename,
		MaxAge:     options.MaxAge,
		MaxSize:    options.MaxSize,
		MaxBackups: options.MaxBackups,
	}

	loggers := map[glog.Level]*log.Logger{
		glog.LevelError: log.New(l, fmt.Sprintf("[%s] ", glog.LevelError.String()), glog.FlagsDateTimeFile),
		glog.LevelWarn:  log.New(l, fmt.Sprintf("[%s] ", glog.LevelWarn.String()), glog.FlagsDateTimeFile),
		glog.LevelInfo:  log.New(l, fmt.Sprintf("[%s] ", glog.LevelInfo.String()), glog.FlagsDateTime),
		glog.LevelDebug: log.New(l, fmt.Sprintf("[%s] ", glog.LevelDebug.String()), glog.FlagsDateTime),
	}

	wc := glog.MultiLogger(glog.StandardLogger(level), glog.WriterCloserLogger(level, l, loggers))
	glog.Loggers(wc)

	return wc
}
