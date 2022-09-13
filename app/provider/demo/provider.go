package demo

import "github.com/gohade/hade/framework"

type Provider struct {
	framework.ServiceProvider
	c framework.Container
}

func (sp *Provider) Name() string  {
	return Key
}

func (sp *Provider) Register(c framework.Container) framework.NewInstance  {
	return NewService
}

func (sp *Provider) IsDefer() bool  {
	return false
}

func (sp *Provider) Boot(c framework.Container) error  {
	sp.c = c
	return nil
}

func (sp *Provider) Params(c framework.Container) []interface{} {
	return []interface{}{sp.c}
}
