package log

// NewNoop creates a new noop logger
func NewNoop() Logger {
	return &noopLogger{}
}

type noopLogger struct{}

func (n *noopLogger) Trace(args ...any) {}

func (n *noopLogger) Debug(args ...any) {}

func (n *noopLogger) Info(args ...any) {}

func (n *noopLogger) Warn(args ...any) {}

func (n *noopLogger) Error(args ...any) {}

func (n *noopLogger) Fatal(args ...any) {}

func (n *noopLogger) Panic(args ...any) {}

func (n *noopLogger) Tracef(format string, args ...any) {}

func (n *noopLogger) Debugf(format string, args ...any) {}

func (n *noopLogger) Infof(format string, args ...any) {}

func (n *noopLogger) Warnf(format string, args ...any) {}

func (n *noopLogger) Errorf(format string, args ...any) {}

func (n *noopLogger) Fatalf(format string, args ...any) {}

func (n *noopLogger) Panicf(format string, args ...any) {}
