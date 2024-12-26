package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var gevents []string

func Events(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"list": gevents})
}

func GetEventList(c *gin.Context) {
	errormessage := "Index out of range"
	indexstring := c.Param("index")
	if index, err := strconv.Atoi(indexstring); err == nil && index < len(gevents) {
		c.JSON(http.StatusOK, gin.H{"item": gevents[index]})
	} else {
		if err != nil {
			errormessage = "Number expected: " + indexstring
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": errormessage})
	}
}

func AddEventToList(c *gin.Context) {

	item := c.PostForm("item")
	gevents = append(gevents, item)
	c.String(http.StatusCreated, c.FullPath()+"/"+strconv.Itoa(len(gevents)-1))
}

func main() {
	gevents = append(gevents, "Go For a Walk : 8 A.M.")
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./eventsmgr/dist", false)))
	r.GET("/api/events", Events)
	r.GET("/api/events/:index", GetEventList)
	r.POST("/api/events", AddEventToList)
	r.Run()
}
