package main

import (
	"gee"
	"log"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.JSON(200, gee.H{
			"msg": "nb666",
		})
	})
	r.GET("/hello", func(c *gee.Context) {
		c.JSON(200, c.Req.Header)

	})
	r.GET("/user/:id", func(c *gee.Context) {
		id := c.Param("id")
		c.JSON(200, id)
	})
	r.GET("/assets/*filepath", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
	})
	log.Fatal(r.Run(":7799"))
}
