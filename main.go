package main

import (
	"fmt"
	"log"
	"net/http"

	"path"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World")
	r := gin.Default()

	ShowTest()

	// 全局使用中間件
	r.Use(Middleware_m1, Middleware_m2)
	//r.Use(Middleware_m1, Middleware_m2,authMiddleware(doCheck:false or true))

	r.GET("/", func(c *gin.Context) {
		data := gin.H{"message": "Hi"}
		c.JSON(http.StatusOK, data)
	})

	r.Static("/static", "./static")

	// ==========================================

	// 要載入html，需先load
	r.LoadHTMLFiles("./login.html", "./index.html", "./upload.html")

	// RESTful的風格，同一api名稱，使用Get、Post、Delete、Put來作不同功能的請求
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.POST("/login", func(c *gin.Context) {
		// 方法1
		//username := c.PostForm("username")
		//password := c.PostForm("password")

		// 方法2，帶預設值，但實際試用好像沒出現預設值的效果
		//username := c.DefaultPostForm("username", "somebody")
		//password := c.DefaultPostForm("password", "***")

		// 方法3，帶預設值，但實際試用好像沒出現預設值的效果
		// 有用中斷點查過，即使沒輸入值，ok仍然會回傳true，真是奇怪，之後有空再研究
		username, ok := c.GetPostForm("username")
		if !ok {
			username = "訪客"
		}
		password, ok := c.GetPostForm("password")
		if !ok {
			password = "123456"
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Name":     username,
			"Password": password,
		})
	})

	// ==========================================

	// 返還json格式的資料
	// 方法1，使用gin.H
	r.GET("/json", func(c *gin.Context) {
		data := gin.H{"name": "IvesShe", "message": "Hello Golang", "age": 20}
		c.JSON(http.StatusOK, data)
	})

	// 方法2：使用結構體, 可以使用tag來對結構體字段作定制化操作
	type msg struct {
		Name    string `json:"name"`
		Message string `json:"message"`
		Age     int    `json:"age"`
	}

	r.GET("/more_json", func(c *gin.Context) {
		data := msg{ // 吃上方的msg結構體
			"ChiChi",
			"Hello World",
			18,
		}
		c.JSON(http.StatusOK, data)
	})

	// ==========================================

	// 取Get請求時，取參數的方法
	// http://localhost:9090/web?name=""
	r.GET("/web", func(c *gin.Context) {
		// 方法1
		name := c.Query("name")

		// 方法2 帶預設值
		message := c.DefaultQuery("message", "Hi")
		//message := c.Query("message")

		// 方法3, 取不到name就返回false
		age, ok := c.GetQuery("age")
		if !ok {
			name = "18"
		}

		// 方法2、方法3，只會判斷有沒有name，卻不會判斷name有沒有值

		c.JSON(http.StatusOK, gin.H{
			"name":    name,
			"message": message,
			"age":     age,
		})
	})

	// ==========================================

	// 捉取路徑的參數，要注意路徑不能與其它的路由路徑衝突
	r.GET("/blog/:name/:age", func(c *gin.Context) {

		name := c.Param("name")
		age := c.Param("age")

		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	// ==========================================

	// 可以使用結構體的方式接收資料
	type UserInfo struct {
		Username string `form:"username" json:"username"`
		Password string `form:"password" json:"password"`
	}
	r.GET("/user", func(c *gin.Context) {

		var u UserInfo
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
			c.JSON(http.StatusOK, u)
		}
	})

	r.POST("/form", func(c *gin.Context) {

		var u UserInfo
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)

			c.JSON(http.StatusOK, gin.H{
				"status": "/form POST ok",
			})
			c.JSON(http.StatusOK, u)
		}
	})

	r.POST("/json", func(c *gin.Context) {

		var u UserInfo
		// ShouldBind會根據請求的Content-Type自行選擇綁定器
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)

			c.JSON(http.StatusOK, gin.H{
				"status": "/json POST ok",
			})
			c.JSON(http.StatusOK, u)
		}
	})

	// ==========================================

	// 文件上傳
	r.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", nil)
	})

	// 單個文件上傳
	r.POST("/upload", func(c *gin.Context) {

		f, err := c.FormFile("filename")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			dst := path.Join("./", f.Filename)
			c.SaveUploadedFile(f, dst)
			c.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})
		}
	})

	// 多個文件上傳
	r.POST("/upload_more", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["filename"]

		for index, file := range files {
			log.Println(file.Filename)
			dst := fmt.Sprintf("./%d_%s", index, file.Filename)
			//dst := path.Join("./", file.Filename)
			// 上传文件到指定的目录
			c.SaveUploadedFile(file, dst)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d files uploaded!", len(files)),
		})
	})

	// ==========================================

	// 請求重定向
	r.GET("/turn", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://tw.yahoo.com/")
	})

	r.GET("/turn_this", func(c *gin.Context) {
		// 把請求的uri修改
		c.Request.URL.Path = "/turn_that"
		// 繼續後續的處理
		r.HandleContext(c)
	})

	r.GET("/turn_that", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "turn_that",
		})
	})

	// ==========================================

	// 方法一，分開寫
	r.GET("/shop", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})

	r.POST("/shop", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})

	r.PUT("/shop", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})

	r.DELETE("/shop", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})

	// 方法二，寫一起，並使用定義的名稱
	r.Any("/book", func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet:
			c.JSON(http.StatusOK, gin.H{"method": "GET"})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{"method": "POST"})
		case http.MethodPut:
			c.JSON(http.StatusOK, gin.H{"method": "PUT"})
		case http.MethodDelete:
			c.JSON(http.StatusOK, gin.H{"method": "DELETE"})

		}
	})

	// ==========================================

	// 路由組
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/shop", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"method": "/video/shop",
			})
		})
		videoGroup.GET("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"method": "/video/login",
			})
		})
		videoGroup.GET("/user", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"method": "/video/user",
			})
		})
	}

	// ==========================================

	// 中間件
	r.GET("/member", ShowHandler, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "member",
		})
	})

	r.Run(":9090")

}
