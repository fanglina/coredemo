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
	handlers []ControllerHandler
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

func (c *Core) Use(handlers ...ControllerHandler)  {
	c.handlers = append(c.handlers, handlers...)
}

//匹配GET方法，增加路由规则
func (c *Core) Get(url string, handler ...ControllerHandler)  {
	allHandlers := append(c.handlers, handler... )
	if err := c.router["GET"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router error:", err)
	}
}

// 匹配POST方法，增加路由规则
func (c *Core) Post(url string, handler ...ControllerHandler)  {
	allHandlers := append(c.handlers, handler... )
	if err := c.router["POST"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router error:", err)
	}
}

//匹配Put 方法，增加路由规则
func (c *Core) Put(url string, handler ...ControllerHandler)  {
	allHandlers := append(c.handlers, handler... )
	if err := c.router["PUT"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router error:", err)
	}
}

func (c *Core) Delete(url string, handler ...ControllerHandler)  {
	allHandlers := append(c.handlers, handler... )
	if err := c.router["DELETE"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router error:", err)
	}
}

func (c *Core) Group(prefix string) IGroup  {
	return NewGroup(c, prefix)
}

func (c *Core) FindRouteRequest(request *http.Request) *node  {
	//uri 和method 全部转换成大写，保证大小写不敏感
	uri := request.URL.Path
	method := request.Method
	upperMethod := strings.ToUpper(method)

	//查找第一层map
	if methodHandlers, ok := c.router[upperMethod];ok {
		return methodHandlers.root.mathNode(uri)
	}
	return nil
}

//所有请求进入这个函数，这个函数负责路由分发
func (c *Core) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("core.serverHTTP")
	ctx := NewContext(r, w)

	//寻找路由
	node := c.FindRouteRequest(r)
	if node == nil {
		//如果没有找到，这里打印日志
		ctx.SetStatus(404).Json("Not Found")
		return
	}

	ctx.SetHandlers(node.handlers)

	// 设置路由参数
	params := node.parseParamsFromEndNode(r.URL.Path)
	log.Println(params, r.URL.Path)
	ctx.SetParams(params)

	//调用路由函数
	if err := ctx.Next(); err != nil {
		ctx.SetStatus(500).Json( "inner error")
		return
	}
}