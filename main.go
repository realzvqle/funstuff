package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/static", "./resources")
	router.LoadHTMLGlob("templates/*")
	i := 1

	router.GET("/", func(c *gin.Context) {
		name, err := c.Cookie("youfkced")
		if err != nil {
			c.AbortWithError(404, err)
		}
		data := gin.H{
			"Title":   "Trust Game",
			"Counter": i,
		}
		if name == "0" {
			c.HTML(http.StatusOK, "nope.html", data)
		} else {
			c.HTML(http.StatusOK, "index.html", data)
			i++
		}

	})
	router.POST("/nooo", func(c *gin.Context) {
		i = 0
		c.SetCookie("youfkced", "0", -1, "/", "", true, false)
		c.Redirect(http.StatusFound, "/")
	})
	router.Run(":8081")
}
