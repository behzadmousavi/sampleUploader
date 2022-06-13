package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
)

func upload(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	filename := header.Filename
	out, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	oldLocation := filename
	newLocation := "/home/behzad/Desktop/" + filename
	err = os.Rename(oldLocation, newLocation)
	if err != nil {
		log.Fatal(err)
	}
	filepath := newLocation
	c.JSON(http.StatusOK, gin.H{"filepath": filepath})
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("template/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "select_file.html", gin.H{})
	})
	router.POST("/upload", upload)
	router.StaticFS("/file", http.Dir("public"))
	router.Run(":8080")
}