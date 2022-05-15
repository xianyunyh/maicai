package main

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

type LogOptions struct {
	output []io.Writer
	level  string
}
type LogOption func(*LogOptions)

func withOutput(out io.Writer) LogOption {
	return func(lo *LogOptions) {
		lo.output = append(lo.output, out)
	}
}

func withLogLevel(level string) LogOption {
	return func(lo *LogOptions) {
		lo.level = level
	}
}

func InitLogger(args ...LogOption) {
	opts := &LogOptions{
		level:  "trace",
		output: []io.Writer{os.Stdout},
	}
	for _, f := range args {
		f(opts)
	}
	log.SetOutput(io.MultiWriter(opts.output...))
	l, err := logrus.ParseLevel(opts.level)
	if err != nil {
		l = logrus.TraceLevel
	}
	log.SetLevel(l)
}
