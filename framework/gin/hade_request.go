package gin

import (
	"github.com/spf13/cast"
	"mime/multipart"
)

//代表请求的函数
type IRequest interface {
	// 请求地址url中带的参数
	// 形如: foo.com?a=1&b=bar&c[]=bar
	DefaultQueryInt(key string, def int) (int , bool)
	DefaultQueryInt64(key string, def int64) (int64, bool)
	DefaultQueryFloat64(key string, def float64) (float64, bool)
	DefaultQueryFloat32(key string, def float32) (float32, bool)
	DefaultQueryString(key string, def string) (string, bool)
	DefaultQueryStringSlice(key string, def []string) ([]string, bool)
	DefaultQueryBool(key string, def bool) (bool, bool)
	DefaultQuery(key string) interface{}

	// 路由匹配中带的参数
	// 形如 /book/:id
	DefaultParamInt(key string, def int) (int, bool)
	DefaultParamInt64(key string, def int64) (int64, bool)
	DefaultParamFloat64(key string, def float64) (float64, bool)
	DefaultParamFloat32(key string, def float32) (float32, bool)
	DefaultParamBool(key string, def bool) (bool, bool)
	DefaultParamString(key string, def string) (string, bool)
	DefaultParam(key string) interface{}

	//form表单带的数据
	DefaultFormInt(key string, def int) (int, bool)
	DefaultFormInt64(key string, def int64) (int64, bool)
	DefaultFormFloat64(key string, def float64) (float64, bool)
	DefaultFormFloat32(key string, def float32 ) (float32, bool)
	DefaultFormBool(key string, def bool) (bool, bool)
	DefaultFormString(key string, def string) (string, bool)
	DefaultFormStringSlice(key string, def []string) ([]string, bool)
	DefaultFormFile(key string) (*multipart.FileHeader, error)
	DefaultForm(key string) interface{}
}

func (ctx *Context) QueryAll() map[string][]string  {
	ctx.initQueryCache()
	return map[string][]string(ctx.queryCache)
}

func (ctx *Context) DefaultQueryInt(key string , def int) (int, bool) {
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

func (ctx *Context) DefaultQueryInt64(key string , def int64) (int64, bool) {
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

func (ctx *Context) DefaultQueryFloat32(key string , def float32) (float32, bool) {
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

func (ctx *Context) DefaultQueryBool(key string , def bool) (bool, bool) {
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

func (ctx *Context) DefaultQueryString(key string, def string) (string, bool)   {
	params := ctx.QueryAll()
	if vals, ok := params[key];ok {
		l := len(vals)
		if l > 0 {
			return vals[l-1], true
		}
	}
	return def, false
}

func (ctx *Context) DefaultQueryStringSlice(key string, def []string) ([]string, bool)  {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		return vals, true
	}
	return def, false
}

// 获取路由参数
func (ctx *Context) HadeParam(key string) interface{} {
	if val, ok := ctx.Params.Get(key); ok {
		return val
	}
	return nil
}


func (ctx *Context) DefaultParamInt(key string, def int) (int, bool)  {
	if v := ctx.HadeParam(key); v != nil {
		return cast.ToInt(v), true
	}
	return def, false
}

func (ctx *Context) DefaultParamInt64(key string, def int64) (int64, bool)  {
	if v := ctx.HadeParam(key); v != nil {
		return cast.ToInt64(v), true
	}
	return def, false
}

func (ctx *Context) DefaultParamFloat64(key string, def float64) (float64, bool)  {
	if v := ctx.HadeParam(key); v != nil {
		return cast.ToFloat64(v), true
	}
	return def, false
}

func (ctx *Context) DefaultParamFloat32(key string, def float32) (float32, bool)  {
	if v := ctx.HadeParam(key); v != nil {
		return cast.ToFloat32(v), true
	}
	return def, false
}

func (ctx *Context) DefaultParamString(key string, def string) (string, bool)  {
	if v := ctx.HadeParam(key); v != nil {
		return cast.ToString(v), true
	}
	return def, false
}


func (ctx *Context) DefaultParamBool(key string, def bool) (bool, bool)  {
	if v := ctx.HadeParam(key); v != nil {
		return cast.ToBool(v), true
	}
	return def, false
}

func (ctx *Context) DefaultParam(key string) interface{}  {
	if ctx.params != nil {
		if v := ctx.HadeParam(key); v != nil {
			return v
		}
	}

	return nil
}



func (ctx *Context) FormAll() map[string][]string  {
	ctx.initFormCache()
	return map[string][]string(ctx.formCache)
}



func (ctx *Context) DefaultFormInt(key string, def int) (int, bool)  {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultFormInt64(key string, def int64) (int64, bool)  {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt64(vals[0]), true
		}
	}
	return def, false
}


func (ctx *Context) DefaultFormFloat64(key string, def float64) (float64, bool)  {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat64(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultFormFloat32(key string, def float32) (float32, bool)  {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat32(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultFormBool(key string, def bool) (bool, bool)  {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToBool(vals[0]), true
		}
	}
	return def, false
}

func (ctx *Context) DefaultFormString(key string, def string) (string, bool)  {
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

func (ctx *Context) DefaultForm(key string) interface{} {
	params := ctx.FormAll()
	if vals, ok := params[key];ok{
		return vals[0]
	}
	return nil
}
