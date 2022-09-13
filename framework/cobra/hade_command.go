package cobra

import (
	"fmt"
	"github.com/gohade/hade/framework"
)

// SetContainer 设置服务容器
func (c *Command) SetContainer(container framework.Container) {
	fmt.Println("SetContainer",container)
	c.container = container
}

// GetContainer 获取服务容器
func (c *Command) GetContainer() framework.Container  {
	return c.Root().container
}
