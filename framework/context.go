package framework

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
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

func (ctx *Context) Next() error {
	ctx.index++
	if ctx.index < len(ctx.handlers) {
		if err := ctx.handlers[ctx.index](ctx); err != nil {
			return err
		}
	}
	return nil

}

func (ctx *Context) QueryAll() map[string][]string  {
	if ctx.request != nil {
		return map[string][]string(ctx.request.URL.Query())
	}
	return map[string][]string{}
}

func (ctx *Context) QueryInt(key string , def int) int {
	params := ctx.QueryAll()
	if val, ok := params[key];ok {
		l := len(val)
		if l > 0 {
			intval, err := strconv.Atoi(val[l-1])
			if err != nil {
				return def
			}
			return intval
		}
	}
	return def
}

func (ctx *Context) QueryString(key string, def string) string  {
	params := ctx.QueryAll()
	if vals, ok := params[key];ok {
		l := len(vals)
		if l > 0 {
			return vals[l-1]
		}
	}
	return def
}

func (ctx *Context) QueryArray(key string, def []string) []string  {
	params := ctx.QueryAll()
	if vals, ok := params[key];ok{
		return vals
	}
	return def
}

func (ctx *Context) FormAll() map[string][]string  {
	if ctx.request != nil {
		return map[string][]string(ctx.request.PostForm)
	}
	return map[string][]string{}
}

func (ctx *Context) FormInt(key string, def int) int  {
	params := ctx.FormAll()
	if vals, ok := params[key];ok {
		l := len(vals)
		if l > 0 {
			intval, err := strconv.Atoi(vals[l-1])
			if err != nil {
				return  def
			}
			return intval
		}
	}
	return def
}

func (ctx *Context) FormString(key string, def string) string  {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		l := len(vals)
		if l > 0 {
			return vals[l-1]
		}
	}
	return def
}

func (ctx *Context) FormArray(key string, def []string) []string  {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		return vals
	}
	return def
}

func (ctx *Context) BindJson(obj interface{}) error  {
	if ctx.request != nil {
		body, err := ioutil.ReadAll(ctx.request.Body)
		if err != nil {
			return err
		}
		ctx.request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		err = json.Unmarshal(body, obj)
		if err != nil {
			return err
		}
	}else {
		return errors.New("ctx.request.empty")
	}
	return nil
}

func (ctx *Context) Json(status int, obj interface{}) error {
	if ctx.HasTimeout() {
		return nil
	}
	ctx.responseWriter.Header().Set("Context-Type", "application.json")
	ctx.responseWriter.WriteHeader(status)
	byt, err := json.Marshal(obj)
	if err != nil {
		ctx.responseWriter.WriteHeader(500)
		return err
	}
	ctx.responseWriter.Write(byt)
	return nil
}

func (ctx *Context) HTML(status int, obj interface{}, template string) error   {
	return nil
}

func (ctx *Context) Text(status int, obj string) error  {
	return nil
}