// Chain of Responsibility Pattern avoids coupling the sender of a request to its receiver
// by giving more than one object a chance to handle the request. Chain the receiving
// objects and pass the request along the chain until an object handles it.
//
// Use cases:
// - When more than one object may handle a request, and the handler isn't known a priori
// - When you want to issue a request to one of several objects without specifying the receiver explicitly
// - When the set of objects that can handle a request should be specified dynamically

package behavioral

// LogLevel represents different logging levels
type LogLevel int

const (
	INFO LogLevel = iota
	DEBUG
	ERROR
)

// LogEntry represents a log message
type LogEntry struct {
	Message string
	Level   LogLevel
}

// Logger interface defines the contract for loggers
type Logger interface {
	SetNext(logger Logger)
	Log(entry LogEntry) string
}

// BaseLogger provides common functionality for all loggers
type BaseLogger struct {
	next Logger
}

func (l *BaseLogger) SetNext(logger Logger) {
	l.next = logger
}

// InfoLogger handles INFO level logs
type InfoLogger struct {
	BaseLogger
}

func NewInfoLogger() *InfoLogger {
	return &InfoLogger{}
}

func (l *InfoLogger) Log(entry LogEntry) string {
	if entry.Level == INFO {
		return "Info: " + entry.Message
	}
	if l.next != nil {
		return l.next.Log(entry)
	}
	return ""
}

// DebugLogger handles DEBUG level logs
type DebugLogger struct {
	BaseLogger
}

func NewDebugLogger() *DebugLogger {
	return &DebugLogger{}
}

func (l *DebugLogger) Log(entry LogEntry) string {
	if entry.Level == DEBUG {
		return "Debug: " + entry.Message
	}
	if l.next != nil {
		return l.next.Log(entry)
	}
	return ""
}

// ErrorLogger handles ERROR level logs
type ErrorLogger struct {
	BaseLogger
}

func NewErrorLogger() *ErrorLogger {
	return &ErrorLogger{}
}

func (l *ErrorLogger) Log(entry LogEntry) string {
	if entry.Level == ERROR {
		return "Error: " + entry.Message
	}
	if l.next != nil {
		return l.next.Log(entry)
	}
	return ""
}

// LoggerChain sets up the chain of responsibility
type LoggerChain struct {
	firstLogger Logger
}

func NewLoggerChain() *LoggerChain {
	infoLogger := NewInfoLogger()
	debugLogger := NewDebugLogger()
	errorLogger := NewErrorLogger()

	infoLogger.SetNext(debugLogger)
	debugLogger.SetNext(errorLogger)

	return &LoggerChain{
		firstLogger: infoLogger,
	}
}

func (c *LoggerChain) Log(entry LogEntry) string {
	return c.firstLogger.Log(entry)
}
