package chain

import "testing"

func Test(t *testing.T) {
	t.Run("Log Chain:", LoggerTest)
}

func LoggerTest(t *testing.T) {
	logChain := GetChainOfLoggers()
	logChain.PrintLog(InfoLogLevel, "This is an information")
	logChain.PrintLog(StandardLogLevel, "This is a Standard output")
	logChain.PrintLog(ErrorLogLevel, "This is an error information")
}
