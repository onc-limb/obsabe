package obsabe

import (
	"fmt"
	"strings"
	"sync/atomic"
)

var DebugPrintFunc func(format string, values ...interface{})

func IsDebugging() bool {
	return atomic.LoadInt32(&ginMode) == debugCode
}

func debugPrint(format string, values ...any) {
	if !IsDebugging() {
		return
	}

	if DebugPrintFunc != nil {
		DebugPrintFunc(format, values...)
		return
	}

	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	fmt.Fprintf(DefaultWriter, "[GIN-debug] "+format, values...)
}

func DebugPrintError(err error) {
	if err != nil && IsDebugging() {
		fmt.Fprintf(DefaultErrorWriter, "[GIN-debug] [ERROR] %v\n", err)
	}
}
