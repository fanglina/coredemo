package gin

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
)

type IResponse interface {
	//Json输出
	IJson(obj interface{}) IResponse

	//Jsonp输出
	IJsonp(obj interface{}) IResponse

	//xml输出
	IXml(obj interface{}) IResponse

	//html输出
	IHtml(template string, obj interface{}) IResponse

	//text输出
	IText(format string, values ...interface{}) IResponse

	//重定向
	IRedirect(path string) IResponse

	//Head
	ISetHeader(key string, value string) IResponse

	//cookie
	ISetCookie(key string, value string, maxAge int, path, domain string, secure, httpOnly bool) IResponse

	//设置状态码
	ISetStatus(code int) IResponse

	//设置200
	ISetOkStatus() IResponse
}

func (ctx *Context) IJsonp(obj interface{}) IResponse {
	//获取请求参数callback
	callback := ctx.Query("callback")
	ctx.ISetHeader("Context-Type", "application/javascript")

	//输出前端进行字符过滤，防止xss攻击
	callback = template.JSEscapeString(callback)

	//输出函数名
	_, err := ctx.Writer.Write([]byte(callback))
	if err != nil {
		return ctx
	}

	//输出左括号
	_, err = ctx.Writer.Write([]byte("("))
	if err != nil {
		return ctx
	}

	//输出参数
	ret, err := json.Marshal(obj)
	if err != nil {
		return ctx
	}
	_, err = ctx.Writer.Write(ret)
	if err != nil {
		return ctx
	}

	//输出右括号
	_, err = ctx.Writer.Write([]byte(")"))
	if err != nil {
		return ctx
	}

	return ctx
}

func (ctx *Context) ISetHeader(key, val string) IResponse {
	ctx.Writer.Header().Set(key, val)
	return ctx
}

func (ctx *Context) IXml(obj interface{}) IResponse {
	byt, err := json.Marshal(obj)
	if err != nil {
		return ctx.ISetStatus(http.StatusInternalServerError)
	}
	ctx.ISetHeader("Context-Type", "application/html")
	ctx.Writer.Write(byt)
	return ctx
}

func (ctx *Context) IHtml(file string, obj interface{}) IResponse {
	//读取模板文件，创建templates实例
	t, err := template.New("output").ParseFiles(file)
	if err != nil {
		return ctx
	}
	//执行excecute插入obj
	err = t.Execute(ctx.Writer, obj)
	if err != nil {
		return ctx
	}
	ctx.ISetHeader("Context-Type", "application/html")
	return ctx
}

func (ctx *Context) IText(format string, values ...interface{}) IResponse {
	out := fmt.Sprintf(format, values...)
	ctx.ISetHeader("Context-Type", "application/text")
	ctx.Writer.Write([]byte(out))
	return ctx
}

func (ctx *Context) IRedirect(path string) IResponse {
	http.Redirect(ctx.Writer, ctx.Request, path, http.StatusMovedPermanently)
	return ctx
}

func (ctx *Context) ISetCookie(key string, value string, maxAge int, path, domain string, secure, httpOnly bool) IResponse {
	if path == "" {
		path = "/"
	}
	http.SetCookie(ctx.Writer, &http.Cookie{
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

func (ctx *Context) ISetStatus(code int) IResponse {
	ctx.Writer.WriteHeader(code)
	return ctx
}

func (ctx *Context) ISetOkStatus() IResponse {
	ctx.Writer.WriteHeader(http.StatusOK)
	return ctx
}

func (ctx *Context) IJson(obj interface{}) IResponse {
	byt, err := json.Marshal(obj)
	if err != nil {
		ctx.ISetStatus(http.StatusInternalServerError)
		return ctx
	}
	ctx.ISetHeader("Context-Type", "application/json")
	ctx.Writer.Write(byt)
	return nil
}
