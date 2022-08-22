package kernel

import (
	"github.com/gohade/hade/framework"
	"github.com/gohade/hade/framework/gin"
)

type HadeKernelProvider struct {
	HttpEngine *gin.Engine
}

func (provider *HadeKernelProvider) Register(c framework.Container) framework.NewInstance {
	return NewHadeKernelService
}

func (provider *HadeKernelProvider) Boot(c framework.Container) error  {
	if provider.HttpEngine == nil {
		provider.HttpEngine = gin.Default()
	}
}
