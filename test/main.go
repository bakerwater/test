package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	f := gin.Default()
	f.LoadHTMLGlob("tempates/*")
	K := make(map[string]string)
	f.GET("/register", func(t *gin.Context) {
		t.HTML(200, "r.html", nil)
	})
	f.POST("/register", func(z *gin.Context) {
		used := z.PostForm("used")
		pwd := z.PostForm("pwd")
		if K[used] == "" {
			K[used] = pwd
			z.HTML(200, "s.html", nil)
		} else {
			z.HTML(200, "u.html", nil)
		}

	})

	f.GET("/login", func(c *gin.Context) {
		c.HTML(200, "h.html", nil)

	})
	f.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		if password == K[username] && K[username] != "" {
			c.HTML(200, "t.html", gin.H{
				"username": username,
				"password": password,
			})
		} else {
			c.HTML(200, "l.html", nil)
		}

	})
	f.Run(":9090")

}
