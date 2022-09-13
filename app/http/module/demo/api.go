package demo

import (
	demoService "github.com/gohade/hade/app/provider/demo"
	"github.com/gohade/hade/framework/gin"
)

type Api struct {
	service *Service
}

func Register(r *gin.Engine) error  {
	api := NewDemoApi()
	r.Bind(&demoService.Provider{})

	r.GET("/demo/demo",api.Demo)
	r.GET("/demo/demo2", api.Demo2)
	r.POST("/demo/demo_post", api.DemoPost)
	return nil
}

func NewDemoApi() *Api {
	Service := NewService()
	return &Api{service:Service}
}

func (api *Api) Demo(c *gin.Context)  {
	users := api.service.GetUsers()
	userDTO := UserModelsUserDTOs(users)
	c.JSON(200, userDTO)
}

func (api *Api) Demo2(c *gin.Context)  {
	demoProvider := c.MustMake(demoService.Key).(demoService.IService)
	students := demoProvider.GetALlStudent()
	UserDTO := StudentToUserDTOs(students)
	c.JSON(200, UserDTO)
}

func (api *Api) DemoPost(c *gin.Context)  {
	type Foo struct {
		Name string
	}
	foo := &Foo{}
	err := c.BindJSON(&foo)
	if err != nil {
		c.AbortWithError(500, err)
	}
	c.JSON(200, nil)
}

