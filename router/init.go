package router

import (
	"minecraft-ws/config"
	"minecraft-ws/service/apihandler"

	"github.com/EricChiou/httprouter"
	"github.com/valyala/fasthttp"
)

// Init init api
func Init() {
	initConfigAPI()
}

// ListenAndServe start http server
func ListenAndServe(port string) error {
	return newFHServer().ListenAndServe(":" + port)
}

// ListenAndServeTLS start https server
func ListenAndServeTLS(port, certPath, keyPath string) error {
	return newFHServer().ListenAndServeTLS(":"+port, certPath, keyPath)
}

// SetHeader add api response header
func SetHeader(key string, value string) {
	httprouter.SetHeader(key, value)
}

func newFHServer() *fasthttp.Server {
	return &fasthttp.Server{
		Name:    config.Get().ServerName,
		Handler: httprouter.FasthttpHandler(),
	}
}

func get(path string, run func(*httprouter.Context) apihandler.ResponseEntity) {
	httprouter.Get(path, func(ctx *httprouter.Context) { run(ctx) })
}

func post(path string, run func(*httprouter.Context) apihandler.ResponseEntity) {
	httprouter.Post(path, func(ctx *httprouter.Context) { run(ctx) })
}

func put(path string, run func(*httprouter.Context) apihandler.ResponseEntity) {
	httprouter.Put(path, func(ctx *httprouter.Context) { run(ctx) })
}

func delete(path string, run func(*httprouter.Context) apihandler.ResponseEntity) {
	httprouter.Delete(path, func(ctx *httprouter.Context) { run(ctx) })
}

func patch(path string, run func(*httprouter.Context) apihandler.ResponseEntity) {
	httprouter.Patch(path, func(ctx *httprouter.Context) { run(ctx) })
}

func head(path string, run func(*httprouter.Context) apihandler.ResponseEntity) {
	httprouter.Head(path, func(ctx *httprouter.Context) { run(ctx) })
}

func options(path string, run func(*httprouter.Context) apihandler.ResponseEntity) {
	httprouter.Options(path, func(ctx *httprouter.Context) { run(ctx) })
}
