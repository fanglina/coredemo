package framework

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"github.com/spf13/cast"
	"io/ioutil"
	"mime/multipart"
)

const defaultMultipartMemory = 32 << 20 //32M

//代表请求的函数
type IRequest interface {
	// 请求地址url中带的参数
	// 形如: foo.com?a=1&b=bar&c[]=bar
	QueryInt(key string, def int) (int , bool)
	QueryInt64(key string, def int64) (int64, bool)
	QueryFloat64(key string, def float64) (float64, bool)
	QueryFloat32(key string, def float32) (float32, bool)
	QueryString(key string, def string) (string, bool)
	QueryStringSlice(key string, def []string) ([]string, bool)
	QueryBool(key string, def bool) (bool, bool)
	Query(key string) interface{}

	// 路由匹配中带的参数
	// 形如 /book/:id
	ParamInt(key string, def int) (int, bool)
	ParamInt64(key string, def int64) (int64, bool)
	ParamFloat64(key string, def float64) (float64, bool)
	ParamFloat32(key string, def float32) (float32, bool)
	ParamBool(key string, def bool) (bool, bool)
	ParamString(key string, def string) (string, bool)
	Param(key string) interface{}

	//form表单带的数据
	FormInt(key string, def int) (int, bool)
	FormInt64(key string, def int64) (int64, bool)
	FormFloat64(key string, def float64) (float64, bool)
	FormFloat32(key string, def float32 ) (float32, bool)
	FormBool(key string, def bool) (bool, bool)
	FormString(key string, def string) (string, bool)
	FormStringSlice(key string, def []string) ([]string, bool)
	FormFile(key string) (*multipart.FileHeader, error)
	Form(key string) interface{}

	//json Body
	BindJson(obj interface{})

	//xml
	BindXml(obj interface{})

	//其他格式
	GetRawData(key string) ([]byte, error)

	//基础信息
	Uri() string
	Method() string
	Host() string
	ClientIp() string

	//Header
	Headers() map[string][]string
	Header(key string) (string , bool)

	//Cookie
	Cookies() map[string]string
	Cookie(key string) (string, bool)
}

func (ctx *Context) QueryAll() map[string][]string  {
	if ctx.request != nil {
		return map[string][]string(ctx.request.URL.Query())
	}
	return map[string][]string{}
}

func (ctx *Context) QueryInt(key string , def int) (int, bool) {
	params := ctx.QueryAll()
	if val, ok := params[key];ok {
		l := len(val)
		if l > 0 {
			intval := cast.ToInt(val[l-1])
			return intval, true
		}
	}
	return def, false
}

func (ctx *Context) QueryInt64(key string , def int64) (int64, bool) {
	params := ctx.QueryAll()
	if val, ok := params[key];ok {
		l := len(val)
		if l > 0 {
			intval := cast.ToInt64(val[l-1])
			return intval, true
		}
	}
	return def, false
}

func (ctx *Context) QueryFloat64(key string , def float64) (float64 , bool) {
	params := ctx.QueryAll()
	if val, ok := params[key];ok {
		l := len(val)
		if l > 0 {
			intval := cast.ToFloat64(val[l-1])
			return intval, true
		}
	}
	return def, false
}

func (ctx *Context) QueryFloat32(key string , def float32) (float32, bool) {
	params := ctx.QueryAll()
	if val, ok := params[key];ok {
		l := len(val)
		if l > 0 {
			intval := cast.ToFloat32(val[l-1])
			return intval, true
		}
	}
	return def, false
}

func (ctx *Context) QueryBool(key string , def bool) (bool, bool) {
	params := ctx.QueryAll()
	if val, ok := params[key];ok {
		l := len(val)
		if l > 0 {
			intval := cast.ToBool(val[l-1])
			return intval, true
		}
	}
	return def, false
}

func (ctx *Context) QueryString(key string, def string) (string, bool)   {
	params := ctx.QueryAll()
	if vals, ok := params[key];ok {
		l := len(vals)
		if l > 0 {
			return vals[l-1], true
		}
	}
	return def, false
}

func (ctx *Context) QueryStringSlice(key string, def []string) ([]string, bool)  {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		return vals, true
	}
	return def, false
}

func (ctx *Context) Query(key string, def interface{}) (interface{}, bool)  {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		return vals[0], true
	}
	return def, false
}

func (ctx *Context) ParamInt(key string, def int) (int, bool)  {
	if v, ok := ctx.params[key];ok{
		return cast.ToInt(v), true
	}
	return def, false
}

func (ctx *Context) ParamInt64(key string, def int64) (int64, bool)  {
	if v, ok := ctx.params[key];ok{
		return cast.ToInt64(v), true
	}
	return def, false
}

func (ctx *Context) ParamFloat64(key string, def float64) (float64, bool)  {
	if v, ok := ctx.params[key];ok{
		return cast.ToFloat64(v), true
	}
	return def, false
}

func (ctx *Context) ParamFloat32(key string, def float32) (float32, bool)  {
	if v, ok := ctx.params[key];ok{
		return cast.ToFloat32(v), true
	}
	return def, false
}

func (ctx *Context) ParamString(key string, def string) (string, bool)  {
	if v, ok := ctx.params[key];ok{
		return v, true
	}
	return def, false
}


func (ctx *Context) ParamBool(key string, def bool) (bool, bool)  {
	if v, ok := ctx.params[key];ok{
		return cast.ToBool(v), true
	}
	return def, false
}

func (ctx *Context) Param(key string) interface{}  {
	if ctx.params != nil {
		if v, ok := ctx.params[key];ok{
			return v
		}
	}

	return nil
}



func (ctx *Context) FormAll() map[string][]string  {
	if ctx.request != nil {
		return map[string][]string(ctx.request.PostForm)
	}
	return map[string][]string{}
}



func (ctx *Context) FormInt(key string, def int) (int, bool)  {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) FormInt64(key string, def int64) (int64, bool)  {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt64(vals[0]), true
		}
	}
	return def, false
}


func (ctx *Context) FormFloat64(key string, def float64) (float64, bool)  {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat64(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) FormFloat32(key string, def float32) (float32, bool)  {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat32(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) FormBool(key string, def bool) (bool, bool)  {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToBool(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) FormString(key string, def string) (string, bool)  {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return vals[0], true
		}
	}
	return def, false
}

func (ctx *Context) FormStringSlice(key string, def []string) ( []string , bool)  {
	params := ctx.FormAll()
	if vals, ok := params[key];ok{
		return vals, true
	}
	return def, false
}



func (ctx *Context) FormFile(key string) (*multipart.FileHeader, error)  {
	if ctx.request.MultipartForm == nil {
		if err := ctx.request.ParseMultipartForm(defaultMultipartMemory); err != nil {
			return nil, err
		}
	}
	f, fh, err := ctx.request.FormFile(key)
	if err != nil {
		return nil , err
	}
	f.Close()
	return fh, nil
}

func (ctx *Context) Form(key string) interface{} {
	params := ctx.FormAll()
	if vals, ok := params[key];ok{
		return vals[0]
	}
	return nil
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

func (ctx *Context) BindXml(obj interface{}) error  {
	if ctx.request != nil {
		body, err := ioutil.ReadAll(ctx.request.Body)
		if err != nil {
			return nil
		}
		ctx.request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		err = xml.Unmarshal(body, obj )
		if err != nil  {
			return err
		}
	}else{
		return errors.New("ctx.request.empty")
	}
	return nil
}

func (ctx *Context) GetRawData() ([]byte, error) {
	if ctx.request != nil {
		body, err := ioutil.ReadAll(ctx.request.Body)
		if err != nil {
			return nil, err
		}
		ctx.request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		return body, nil
	}
	return nil , errors.New("ctx.request.empty")
}

func (ctx *Context) Uri() string  {
	return ctx.request.RequestURI
}

func (ctx *Context) Method() string  {
	return ctx.request.Method
}

func (ctx *Context) Host() string  {
	return ctx.request.URL.Host
}

func (ctx *Context) ClientIp() string  {
	r := ctx.request
	ipAddress := r.Header.Get("X-Real-Ip")
	if ipAddress == "" {
		ipAddress = r.Header.Get("X-Forwarded-For")
	}
	if ipAddress == "" {
		ipAddress = ctx.request.RemoteAddr
	}
	return ipAddress
}

func (ctx *Context) Headers() map[string][]string  {
	return ctx.request.Header
}

func (ctx *Context) Header(key string) (string, bool)  {
	val := ctx.request.Header.Get(key)
	if  len(val) < 0 {
		return "", false
	}
	return val, true
}

func (ctx *Context) Cookies()  map[string]string {
	cookies := ctx.request.Cookies()
	ret := make(map[string]string)
	for _, cookie := range cookies {
		ret[cookie.Name] = cookie.Value
	}

	return ret
}

func (ctx *Context) Cookie(key string) (string, bool)  {
	cookies := ctx.Cookies()
	if cookie, ok := cookies[key];ok {
		return cookie, true
	}
	return "", false
}

