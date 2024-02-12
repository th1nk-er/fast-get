package main

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/th1nk-er/fast-get/server/routers"
)

func main() {
	r := routers.InitRouters()
	r.Use(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/x") {
			c.Request.URL.Path = c.Request.URL.Path[2:]
			r.ServeHTTP(c.Writer, c.Request)
		} else {
			c.File("static/" + c.Request.URL.Path)
		}
	})
	_ = r.Run(":8080")
}
