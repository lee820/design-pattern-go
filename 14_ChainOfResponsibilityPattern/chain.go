package chain

import "fmt"

const (
	//StandardLogLevel 标准日志等级
	StandardLogLevel = iota
	//InfoLogLevel info日志等级
	InfoLogLevel
	//ErrorLogLevel 错误日志等级
	ErrorLogLevel
)

//BaseLogger 日志接口
type BaseLogger interface {
	PrintLog(level int, message string)
	Write(message string)
}

//StandardLogger 标准日志类
type StandardLogger struct {
	Level      int
	NextLogger BaseLogger
}

//NewStandardLogger 实例化标准日志类
func NewStandardLogger() *StandardLogger {
	return &StandardLogger{
		Level:      StandardLogLevel,
		NextLogger: nil,
	}
}

//Write 标准日志类写日志方法
func (sl *StandardLogger) Write(message string) {
	fmt.Printf("Standard Logger out: %s.\n", message)
}

//PrintLog 标准日志类输入日志，并且流向下一个对象方法
func (sl *StandardLogger) PrintLog(level int, message string) {
	if sl.Level == level {
		sl.Write(message)
	}
	if sl.NextLogger != nil {
		sl.NextLogger.PrintLog(level, message)
	}
}

//SetNextLogger 标准日志类设置下一个对象方法
func (sl *StandardLogger) SetNextLogger(logger BaseLogger) {
	sl.NextLogger = logger
}

//InfoLogger 提示日志类
type InfoLogger struct {
	Level      int
	NextLogger BaseLogger
}

//NewInfoLogger 实例化提示日志类
func NewInfoLogger() *InfoLogger {
	return &InfoLogger{
		Level:      InfoLogLevel,
		NextLogger: nil,
	}
}

//Write 提示日志类的写方法
func (infoL *InfoLogger) Write(message string) {
	fmt.Printf("Info Logger out: %s.\n", message)
}

//PrintLog 提示日志类的输入日志方法
func (infoL *InfoLogger) PrintLog(level int, message string) {
	if infoL.Level == level {
		infoL.Write(message)
	}
	if infoL.NextLogger != nil {
		infoL.NextLogger.PrintLog(level, message)
	}
}

//SetNextLogger 提示日志类设置下一个对象
func (infoL *InfoLogger) SetNextLogger(logger BaseLogger) {
	infoL.NextLogger = logger
}

//ErrorLogger 错误日志类
type ErrorLogger struct {
	Level      int
	NextLogger BaseLogger
}

//NewErrorLogger 实例化错误日志类
func NewErrorLogger() *ErrorLogger {
	return &ErrorLogger{
		Level:      ErrorLogLevel,
		NextLogger: nil,
	}
}

//Write 错误日志类写方法
func (el *ErrorLogger) Write(message string) {
	fmt.Printf("Error logger out: %s.\n", message)
}

//PrintLog 错误日志类输入日志方法
func (el *ErrorLogger) PrintLog(level int, message string) {
	if el.Level == level {
		el.Write(message)
	}
	if el.NextLogger != nil {
		el.NextLogger.PrintLog(level, message)
	}
}

//SetNextLogger 错误日志类设置下一个对象
func (el *ErrorLogger) SetNextLogger(logger BaseLogger) {
	el.NextLogger = logger
}

//GetChainOfLoggers 获取日志对象链
func GetChainOfLoggers() BaseLogger {
	errLog := NewErrorLogger()
	infoLog := NewInfoLogger()
	standardLog := NewStandardLogger()

	errLog.SetNextLogger(infoLog)
	infoLog.SetNextLogger(standardLog)

	return errLog
}
