package demo

import (
	"fmt"
	"github.com/gohade/hade/framework"
)

//服务提供方
type DemoServiceProvider struct {
}

// Name方法直接将服务对应的字符凭证返回
func (sp *DemoServiceProvider) Name() string  {
	return Key
}

// Register 注册初始化服务实例的方法
func (sp *DemoServiceProvider) Register(c framework.Container) framework.NewInstance {
	return NewDemoService
}

// IsDefer是否延迟实例化，设置成true, 延迟实例化
func (sp *DemoServiceProvider) IsDefer() bool  {
	return true
}

// Params 表示实例化的参数
func (sp *DemoServiceProvider) Params (c framework.Container) []interface{} {
	return []interface{}{c}
}

func (sp *DemoServiceProvider) Boot(c framework.Container) error  {
	fmt.Println("demo service boot")
	return nil
}


