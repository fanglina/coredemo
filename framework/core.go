package framework

import (
	"log"
	"net/http"
)

//类型检查
var _ http.Handler = (*Core)(nil)

//框架核心结构
type Core struct {
	router map[string]ControllerHandler
}

//初始化核心接口
func NewCore() *Core {
	return &Core{router: map[string]ControllerHandler{}}
}

func (c *Core) Get(url string, handler ControllerHandler)  {
	c.router[url] = handler
}

//框架核心结构要实现的handler
func (c *Core) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("core.serverHTTP")
	ctx := NewContext(r, w)

	//一个简单的路由选择器，这里直接写死为测试路由foo
	router := c.router["foo"]
	if router == nil {
		return
	}
	log.Println("core.router")
	router(ctx)
}