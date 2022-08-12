package middleware

import (
	"coredemo/framework"
	"log"
	"time"
)

// Cost 查看程序花费多少时间
func Cost() framework.ControllerHandler  {
	return func(c *framework.Context) error {
		//记录开始时间
		start := time.Now()

		//使用next执行业务
		c.Next()

		//结束时间
		end := time.Now()
		cost := end.Sub(start)
		log.Printf("api uri: %v, cost: %v", c.GetRequest().RequestURI, cost.Seconds())

		return nil
	}
}
