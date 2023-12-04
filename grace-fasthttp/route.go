package main

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func getRouter() *fasthttprouter.Router {
	router := fasthttprouter.New()

	router.GET("/", Index)

	return router
}

func Index(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("ok")
}
