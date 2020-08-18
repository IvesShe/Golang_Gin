package main

import (
	"fmt"
	_ "log"
	"net/http"
	_ "net/http"
	"time"

	_ "path"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

func ShowHandler(c *gin.Context) {
	fmt.Println("ShowHandler")
	c.JSON(http.StatusOK, gin.H{
		"msg": "使用了ShowHandler中間件...",
	})
}

func Middleware_m1(c *gin.Context) {
	fmt.Println("m1 進入...")
	start := time.Now()
	c.Next()
	cost := time.Since(start)
	fmt.Printf("cost: %v\n", cost)
	fmt.Println("m1 離開...")
}

func Middleware_m2(c *gin.Context) {
	fmt.Println("m2 進入...")
	c.Next()
	//c.Abort()
	fmt.Println("m2 離開...")
}

// 中間件 使用的舉例
func authMiddleware(doCheck bool) gin.HandlerFunc {
	// 連接數據庫
	// 或作一些其它準備工作
	return func(c *gin.Context) {
		if doCheck {
			// 存放具體的邏輯
			// 是否登錄的判斷
			// if 是登錄用戶
			//		c.Next()
			// else
			// 		c.Abort()
		} else {
			c.Next()
		}
	}
}

func ShowTest() {
	fmt.Println("===========================")
	fmt.Println("我來自Handler.go")
	fmt.Println("===========================")
}
