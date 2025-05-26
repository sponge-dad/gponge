package main

import (
	"gee"
	"log"
	"net/http"
)

func main() {
	r := gee.Default()
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
	v1 := r.Group("////v1")
	v1_1 := v1.Group("v1_1")
	v1_1_1 := v1_1.Group("v1_1_1")
	v1_1_1.GET("/user/:id", func(c *gee.Context) {
		id := c.Param("id")
		c.JSON(200, id)
	})
	r.GET("/panic", func(c *gee.Context) {
		a := "qwrer"
		c.String(200, "%s", a[100])
	})
	log.Fatal(r.Run(":7799"))
}
