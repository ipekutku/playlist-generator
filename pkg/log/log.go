package log

import "go.uber.org/zap"

// Interface for Logger operations, mainly for testing purposes
type Logger interface {
	Info(msg string, keysAndValues ...interface{})
	Error(msg string, keysAndValues ...interface{})
	Sync()
}

// Implementation of Logger
// This is a wrapper around zap.Logger
type LoggerImpl struct {
	logger *zap.Logger
}

// Asserting LoggerImpl implements Logger
var _ Logger = (*LoggerImpl)(nil)

// Constructor for LoggerImpl
func NewLogger() *LoggerImpl {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	return &LoggerImpl{
		logger: logger,
	}
}

// Implementation of Info method in Logger
func (l *LoggerImpl) Info(msg string, keysAndValues ...interface{}) {
	l.logger.Sugar().Infow(msg, keysAndValues...)
}

// Implementation of Error method in Logger
func (l *LoggerImpl) Error(msg string, keysAndValues ...interface{}) {
	l.logger.Sugar().Errorw(msg, keysAndValues...)
}
func (l *LoggerImpl) Sync() {
	l.logger.Sync()
}
