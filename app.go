package main

import "fmt"

import "github.com/gin-gonic/gin"

func main() {
	fmt.Println("mygoweb starting")
	engine := gin.Default()
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "hh"})
	})
	engine.Run()
	fmt.Println("mygoweb running")
}
