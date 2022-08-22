package main

import (
	"fmt"
	"github.com/gohade/hade/app/provider/demo"
	"github.com/gohade/hade/framework/gin"
)

func SubjectAddController(c *gin.Context) {
	c.ISetOkStatus().IJson("OK, SubjectAddController")
}

func SubjectListController(c *gin.Context) {
    //获取实例
	demoservice := c.MustMake(demo.Key).(demo.Service)

	// 调用实例
	foo := demoservice.GetFoo()

	c.ISetOkStatus().IJson(foo)
}

func SubjectDelController(c *gin.Context) {
	c.ISetOkStatus().IJson("OK, SubjectDelController")
}

func SubjectUpdateController(c *gin.Context) {
	c.ISetOkStatus().IJson("OK, SubjectDelController")
}

func SubjectGetController(c *gin.Context) {
	subjectId, _ := c.DefaultParamInt("id", 0)
	c.ISetOkStatus().IJson("OK, SubjectGetController" + fmt.Sprint(subjectId))
}

func SubjectNameController(c *gin.Context) {
	c.ISetOkStatus().IJson("OK, SubjectNameController")
}
