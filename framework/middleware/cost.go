package middleware

import (
	"github.com/gohade/hade/framework/gin"
	"log"
	"time"
)

// Cost 查看程序花费多少时间
func Cost() gin.HandlerFunc  {
	return func(c *gin.Context)  {
		//记录开始时间
		start := time.Now()
		log.Printf("api uri start: %v", c.Request.RequestURI)
		//使用next执行业务
		c.Next()

		//结束时间
		end := time.Now()
		cost := end.Sub(start)
		log.Printf("api uri: %v, cost: %v", c.Request.RequestURI, cost.Seconds())


	}
}
