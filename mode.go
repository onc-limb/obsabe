package obsabe

import (
	"io"
	"os"
)

const (
	debugCode = iota
	releaseCode
	testCode
)

var DefaultWriter io.Writer = os.Stdout
var DefaultErrorWriter io.Writer = os.Stderr
var ginMode int32 = debugCode
