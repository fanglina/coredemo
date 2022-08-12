package main

import "coredemo/framework"

func SubjectAddController(c *framework.Context) error  {
	c.Json(200, "OK, SubjectAddController")
	return nil
}

func SubjectListController(c *framework.Context) error  {
	c.Json(200, "OK, SubjectListController")
	return nil
}


func SubjectDelController(c *framework.Context) error  {
	c.Json(200, "OK, SubjectDelController")
	return nil
}

func SubjectUpdateController(c *framework.Context) error  {
	c.Json(200, "OK, SubjectDelController")
	return nil
}

func SubjectGetController(c *framework.Context) error  {
	c.Json(200, "OK, SubjectGetController")
	return nil
}

func SubjectNameController(c *framework.Context) error  {
	c.Json(200, "OK, SubjectNameController")
	return nil
}