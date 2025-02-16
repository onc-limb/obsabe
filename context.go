package obsabe

import "net/http"

type Context struct {
	writermem http.ResponseWriter
	Request   *http.Request
}
