package framework

// IGroup 代表前缀分组
type IGroup interface {
	// 实现HttpMethod方法
	Get(string, ...ControllerHandler)
	Post(string, ...ControllerHandler)
	Put(string, ...ControllerHandler)
	Delete(string, ...ControllerHandler)

	//实现嵌套group
	Group(string) IGroup

	//嵌套中间件
	Use(middleware ...ControllerHandler)
}

// Group struct 实现了IGroup
type Group struct {
	core *Core //指向core结构
	parent *Group //指向上一个Group,如果有的话
	prefix string //这个group的通用前缀

	middlewares []ControllerHandler // 存放中间件
}

//初始化Group
func NewGroup(core *Core, prefix string) *Group  {
	return &Group{
		core:   core,
		parent: nil,
		prefix: prefix,
	}
}

func (g *Group) Get(uri string, handlers ...ControllerHandler)  {
	uri = g.getAbsolutePrefix() + uri
	allHandles := append(g.middlewares, handlers...)
	g.core.Get(uri, allHandles...)
}

func (g *Group) Post(uri string, handlers ...ControllerHandler)  {
	uri = g.getAbsolutePrefix()
	allHandles := append(g.middlewares, handlers...)
	g.core.Post(uri, allHandles...)
}

func (g *Group) Put(uri string, handlers ...ControllerHandler)  {
	uri = g.getAbsolutePrefix() + uri
	allHandles := append(g.middlewares, handlers...)
	g.core.Put(uri, allHandles...)
}

func (g *Group) Delete(uri string, handlers ...ControllerHandler)  {
	uri = g.getAbsolutePrefix()
	allHandles := append(g.middlewares, handlers...)
	g.core.Delete(uri, allHandles...)
}

func (g *Group) getAbsolutePrefix() string  {
	if g.parent == nil {
		return g.prefix
	}
	return g.parent.getAbsolutePrefix() + g.prefix
}

func (g *Group) Use(middleware ...ControllerHandler)  {
	g.middlewares = append(g.middlewares, middleware...)
}

func (g *Group) Group(uri string) IGroup  {
	cgroup := NewGroup(g.core, uri)
	cgroup.parent = g
	return cgroup
}