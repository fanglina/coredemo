package framework

import (
	"context"
	"net/http"
	"sync"
)

type Context struct {
	request        *http.Request
	responseWriter http.ResponseWriter
	ctx            context.Context
	handler        ControllerHandler

	//是否超时标记位
	hasTimeout bool
	writerMux  *sync.Mutex

	//当前请求的链条
	handlers []ControllerHandler
	index int //当前请求到哪里

	params map[string]string // url路由匹配的参数
}

func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		request:        r,
		responseWriter: w,
		ctx:            r.Context(),
		writerMux:      &sync.Mutex{},
		index: -1,
	}
}

func (ctx *Context) WriterMux() *sync.Mutex  {
	return ctx.writerMux
}

func (ctx *Context) GetRequest() *http.Request  {
	return ctx.request
}

func (ctx *Context) GetResponse() http.ResponseWriter  {
	return ctx.responseWriter
}

func (ctx *Context) HasTimeout() bool  {
	return ctx.hasTimeout
}

func (ctx *Context) SetHasTimeout() {
	ctx.hasTimeout = true
}

func (ctx *Context) BaseContext() context.Context  {
	return ctx.request.Context()
}

func (ctx *Context) Done() <- chan struct{}  {
	return ctx.BaseContext().Done()
}

func (ctx *Context) Err() error  {
	return ctx.BaseContext().Err()
}

func (ctx *Context) Value(key interface{} ) interface{}  {
	return ctx.BaseContext().Value(key)
}

func (ctx *Context) SetHandlers(handlers []ControllerHandler)  {
	ctx.handlers = handlers
}

func (ctx *Context) SetParams(params map[string]string)  {
	ctx.params = params
}

func (ctx *Context) Next() error {
	ctx.index++
	if ctx.index < len(ctx.handlers) {
		if err := ctx.handlers[ctx.index](ctx); err != nil {
			return err
		}
	}
	return nil

}



