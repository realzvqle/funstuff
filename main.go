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
		data := gin.H{
			"Title":   "Trust Game",
			"Counter": i,
		}
		sname, err := c.Cookie("allowed")
		if err != nil {
			sname = "none"
		}

		name, err := c.Cookie("youfkced")
		if err != nil {
			c.AbortWithError(404, err)
		}
		if sname == "none" {
			c.HTML(http.StatusOK, "cookie.html", data)
		} else {
			if name == "0" {
				c.HTML(http.StatusOK, "nope.html", data)
			} else {
				c.HTML(http.StatusOK, "index.html", data)
				i++
			}
		}

	})
	router.POST("/nooo", func(c *gin.Context) {
		i = 1
		max := 10 * 365 * 24 * 60 * 60
		c.SetCookie("youfkced", "0", max, "/", "", true, false)
		c.Redirect(http.StatusFound, "/")
	})
	router.POST("/acceptcookies", func(c *gin.Context) {
		max := 10 * 365 * 24 * 60 * 60
		c.SetCookie("allowed", "allowed", max, "/", "", true, false)
		c.Redirect(http.StatusFound, "/")
	})
	router.Run(":8081")
}
