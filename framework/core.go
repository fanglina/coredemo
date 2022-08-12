package framework

import (
	"log"
	"net/http"
	"strings"
)

//类型检查
var _ http.Handler = (*Core)(nil)

//框架核心结构
type Core struct {
	router map[string]*Tree
}

//初始化核心接口
func NewCore() *Core {
	//初始化路由
	router := map[string]*Tree{}
	router["POST"] = NewTree()
	router["GET"] = NewTree()
	router["PUT"] = NewTree()
	router["DELETE"] = NewTree()
	return &Core{router: router}
}

//匹配GET方法，增加路由规则
func (c *Core) Get(url string, handler ControllerHandler)  {
	if err := c.router["GET"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error:", err)
	}
}

// 匹配POST方法，增加路由规则
func (c *Core) Post(url string, handler ControllerHandler)  {
	if err := c.router["POST"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error:", err)
	}
}

//匹配Put 方法，增加路由规则
func (c *Core) Put(url string, handler ControllerHandler)  {
	if err := c.router["PUT"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error:", err)
	}
}

func (c *Core) Delete(url string, handler ControllerHandler)  {
	if err := c.router["DELETE"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error:", err)
	}
}

func (c *Core) Group(prefix string) IGroup  {
	return NewGroup(c, prefix)
}

func (c *Core) FindRouteRequest(request *http.Request) ControllerHandler  {
	//uri 和method 全部转换成大写，保证大小写不敏感
	uri := request.URL.Path
	method := request.Method
	upperMethod := strings.ToUpper(method)

	//查找第一层map
	if methodHandlers, ok := c.router[upperMethod];ok {
		return methodHandlers.FindHandler(uri)
	}
	return nil
}

//所有请求进入这个函数，这个函数负责路由分发
func (c *Core) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("core.serverHTTP")
	ctx := NewContext(r, w)

	//寻找路由
	router := c.FindRouteRequest(r)
	if router == nil {
		//如果没有找到，这里打印日志
		ctx.Json(404, "Not Found")
		return
	}

	if err := router(ctx); err != nil {
		ctx.Json(500, "inner error")
		return
	}
}