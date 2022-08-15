package middleware

import (
	"context"
	"coredemo/framework"
	"fmt"
	"time"
)

func Timeout(d time.Duration) framework.ControllerHandler {
	return func(c *framework.Context) error {
		finish := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)

		//执行业务逻辑前预操作： 初始化context
		duration, cancel := context.WithTimeout(c.BaseContext(), d)
		defer cancel()

		//考虑边界情况，系统的异常退出捕捉
		go func() {
			defer func() {
				if err := recover(); err != nil {
					panicChan <- struct{}{}
				}
			}()
			//执行具体逻辑
			c.Next()

			finish <- struct{}{}
		}()

		//执行业务逻辑后判断
		select {
		case <-finish:
			fmt.Println("finish")
		case p := <-panicChan:
			c.SetStatus(500).Json( "time out")
			fmt.Println(p)
		case <-duration.Done():
			c.SetHasTimeout()
			c.SetStatus(500).Json( "time out")
		}
		return nil
	}
}
