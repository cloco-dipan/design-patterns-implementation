package logger

type Logger interface {
	Log(interface{})
}

func LoggerFactory(loggerType string) Logger {
	switch loggerType {
	case "console":
		return &ConsoleLogger{}

	case "file":
		return &FileLogger{}
	}
	return nil
}

var CL = LoggerFactory("console")
var FL = LoggerFactory("file")
