package framework

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
)

type IResponse interface {
	//Json输出
	Json(obj interface{}) IResponse

	//Jsonp输出
	Jsonp(obj interface{}) IResponse

	//xml输出
	Xml(obj interface{}) IResponse

	//html输出
	Html(template string, obj interface{}) IResponse

	//text输出
	Text(format string, values ...interface{}) IResponse

	//重定向
	Redirect(path string) IResponse

	//Head
	SetHeader(key string, value string) IResponse

	//cookie
	SetCookie(key string, value string, maxAge int, path, domain string, secure, httpOnly bool) IResponse

	//设置状态码
	SetStatus(code int) IResponse

	//设置200
	SetOkStatus() IResponse
}

func (ctx *Context) Jsonp(obj interface{}) IResponse {
	//获取请求参数callback
	callback, _ := ctx.QueryString("callback", "callback_func")
	ctx.SetHeader("Context-Type", "application/javascript")

	//输出前端进行字符过滤，防止xss攻击
	callback = template.JSEscapeString(callback)

	//输出函数名
	_, err := ctx.responseWriter.Write([]byte(callback))
	if err != nil {
		return ctx
	}

	//输出左括号
	_, err = ctx.responseWriter.Write([]byte("("))
	if err != nil {
		return ctx
	}

	//输出参数
	ret, err := json.Marshal(obj)
	if err != nil {
		return ctx
	}
	_, err = ctx.responseWriter.Write(ret)
	if err != nil {
		return ctx
	}

	//输出右括号
	_, err = ctx.responseWriter.Write([]byte(")"))
	if err != nil {
		return ctx
	}

	return ctx
}

func (ctx *Context) SetHeader(key, val string) IResponse {
	ctx.responseWriter.Header().Set(key, val)
	return ctx
}

func (ctx *Context) Xml(obj interface{}) IResponse {
	byt, err := json.Marshal(obj)
	if err != nil {
		return ctx.SetStatus(http.StatusInternalServerError)
	}
	ctx.SetHeader("Context-Type", "application/html")
	ctx.responseWriter.Write(byt)
	return ctx
}

func (ctx *Context) Html(file string, obj interface{}) IResponse {
	//读取模板文件，创建templates实例
	t, err := template.New("output").ParseFiles(file)
	if err != nil {
		return ctx
	}
	//执行excecute插入obj
	err = t.Execute(ctx.responseWriter, obj)
	if err != nil {
		return ctx
	}
	ctx.SetHeader("Context-Type", "application/html")
	return ctx
}

func (ctx *Context) Text(format string, values ...interface{}) IResponse {
	out := fmt.Sprintf(format, values...)
	ctx.SetHeader("Context-Type", "application/text")
	ctx.responseWriter.Write([]byte(out))
	return ctx
}

func (ctx *Context) Redirect(path string) IResponse {
	http.Redirect(ctx.responseWriter, ctx.request, path, http.StatusMovedPermanently)
	return ctx
}

func (ctx *Context) SetCookie(key string, value string, maxAge int, path, domain string, secure, httpOnly bool) IResponse {
	if path == "" {
		path = "/"
	}
	http.SetCookie(ctx.responseWriter, &http.Cookie{
		Name:     key,
		Value:    url.QueryEscape(value),
		Path:     path,
		Domain:   domain,
		MaxAge:   maxAge,
		Secure:   secure,
		HttpOnly: httpOnly,
		SameSite: 1,
	})
	return ctx
}

func (ctx *Context) SetStatus(code int) IResponse {
	ctx.responseWriter.WriteHeader(code)
	return ctx
}

func (ctx *Context) SetOkStatus() IResponse {
	ctx.responseWriter.WriteHeader(http.StatusOK)
	return ctx
}

func (ctx *Context) Json(obj interface{}) IResponse {
	byt, err := json.Marshal(obj)
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		return ctx
	}
	ctx.SetHeader("Context-Type", "application/json")
	ctx.responseWriter.Write(byt)
	return nil
}
