package gin

import "github.com/gohade/hade/framework"

func (engine *Engine) SetContainer(container framework.Container)  {
	engine.container = container
}

// engine 实现container的绑定
func (engine *Engine) Bind(provider framework.ServiceProvider) error {
	return engine.container.Bind(provider)
}

func (engine *Engine) IsBind(key string) bool  {
	return engine.container.IsBind(key)
}
