package obsabe

import (
	"net/http"
	"sync"
)

type App struct {
	pool sync.Pool
}

func New() *App {
	app := &App{}

	return app
}

func (app *App) Listen(addr ...string) (err error) {
	defer func() { DebugPrintError(err) }()
	address := resolveAddress(addr)
	debugPrint("Listening and serving HTTP on %s\n", address)
	err = http.ListenAndServe(address, app)
	return
}

func (app *App) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := app.pool.Get().(*Context)
	c.writermem.reset(w)
	c.Request = req
	c.reset()

	app.handleHTTPRequest(c)

	app.pool.Put(c)
}
