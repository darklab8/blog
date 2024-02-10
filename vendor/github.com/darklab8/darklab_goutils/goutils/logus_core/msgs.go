package logus_core

import (
	"fmt"
	"os"
)

func (l *Logger) Debug(msg string, opts ...SlogParam) {
	args := append([]any{}, newSlogArgs(l.default_params...)...)
	args = append(args, newSlogArgs(opts...)...)
	if l.enable_file_showing {
		args = append(args, logGroupFiles())
	}
	l.logger.Debug(msg, args...)
}

func (l *Logger) Info(msg string, opts ...SlogParam) {
	args := append([]any{}, newSlogArgs(l.default_params...)...)
	args = append(args, newSlogArgs(opts...)...)
	if l.enable_file_showing {
		args = append(args, logGroupFiles())
	}
	l.logger.Info(msg, args...)
}

// Just potentially bad behavior to be aware of
func (l *Logger) Warn(msg string, opts ...SlogParam) {
	args := append([]any{}, newSlogArgs(l.default_params...)...)
	args = append(args, newSlogArgs(opts...)...)
	if l.enable_file_showing {
		args = append(args, logGroupFiles())
	}
	l.logger.Warn(msg, args...)
}

// It is bad but program can recover from it
func (l *Logger) Error(msg string, opts ...SlogParam) {
	args := append([]any{}, newSlogArgs(l.default_params...)...)
	args = append(args, newSlogArgs(opts...)...)
	if l.enable_file_showing {
		args = append(args, logGroupFiles())
	}
	l.logger.Error(msg, args...)
}

// Program is not allowed to run further with fatal
func (l *Logger) Fatal(msg string, opts ...SlogParam) {
	args := append([]any{}, newSlogArgs(l.default_params...)...)
	args = append(args, newSlogArgs(opts...)...)
	if l.enable_file_showing {
		args = append(args, logGroupFiles())
	}
	l.logger.Error(msg, args...)
	os.Exit(1)
}

func (l *Logger) Panic(msg string, opts ...SlogParam) {
	args := append([]any{}, newSlogArgs(l.default_params...)...)
	args = append(args, newSlogArgs(opts...)...)
	if l.enable_file_showing {
		args = append(args, logGroupFiles())
	}
	l.logger.Error(msg, args...)
	panic(msg)
}

func (l *Logger) CheckDebug(err error, msg string, opts ...SlogParam) bool {
	if err == nil {
		return false
	}
	args := append([]any{}, newSlogArgs(opts...)...)
	args = append(args, "error")
	args = append(args, fmt.Sprintf("%v", err))
	l.logger.Debug(msg, args...)
	return true
}

func (l *Logger) CheckWarn(err error, msg string, opts ...SlogParam) bool {
	if err == nil {
		return false
	}
	args := append([]any{}, newSlogArgs(opts...)...)
	args = append(args, "error")
	args = append(args, fmt.Sprintf("%v", err))
	l.logger.Warn(msg, args...)
	return true
}

func (l *Logger) CheckError(err error, msg string, opts ...SlogParam) bool {
	if err == nil {
		return false
	}
	args := append([]any{}, newSlogArgs(opts...)...)
	args = append(args, "error")
	args = append(args, fmt.Sprintf("%v", err))
	l.logger.Error(msg, args...)
	return true
}

// It has shorter error output in comparison to CheckPanic
func (l *Logger) CheckFatal(err error, msg string, opts ...SlogParam) {
	if err == nil {
		return
	}
	args := append([]any{}, newSlogArgs(opts...)...)
	args = append(args, "error")
	args = append(args, fmt.Sprintf("%v", err))
	l.logger.Error(msg, args...)
	os.Exit(1)
}

func (l *Logger) CheckPanic(err error, msg string, opts ...SlogParam) {
	if err == nil {
		return
	}
	args := append([]any{}, newSlogArgs(opts...)...)
	args = append(args, "error")
	args = append(args, fmt.Sprintf("%v", err))
	l.logger.Error(msg, args...)
	panic(msg)
}

func (l *Logger) Debugf(msg string, varname string, value any, opts ...SlogParam) {
	args := append([]any{}, newSlogArgs(opts...)...)
	if l.enable_file_showing {
		args = append(args, logGroupFiles())
	}
	args = append(args, varname)
	args = append(args, fmt.Sprintf("%v", value))
	l.logger.Debug(msg, args...)
}

func (l *Logger) Infof(msg string, varname string, value any, opts ...SlogParam) {
	args := append([]any{}, newSlogArgs(opts...)...)
	if l.enable_file_showing {
		args = append(args, logGroupFiles())
	}
	args = append(args, varname)
	args = append(args, fmt.Sprintf("%v", value))
	l.logger.Info(msg, args...)
}
