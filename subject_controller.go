package main

import (
	"coredemo/framework"
	"fmt"
)

func SubjectAddController(c *framework.Context) error  {
	c.SetOkStatus().Json( "OK, SubjectAddController")
	return nil
}

func SubjectListController(c *framework.Context) error  {
	c.SetOkStatus().Json( "OK, SubjectListController")
	return nil
}


func SubjectDelController(c *framework.Context) error  {
	c.SetOkStatus().Json( "OK, SubjectDelController")
	return nil
}

func SubjectUpdateController(c *framework.Context) error  {
	c.SetOkStatus().Json( "OK, SubjectDelController")
	return nil
}

func SubjectGetController(c *framework.Context) error  {
	subjectId, _ := c.ParamInt("id", 0)
	c.SetOkStatus().Json( "OK, SubjectGetController" + fmt.Sprint(subjectId) )
	return nil
}

func SubjectNameController(c *framework.Context) error  {
	c.SetOkStatus().Json("OK, SubjectNameController")
	return nil
}