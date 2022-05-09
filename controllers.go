package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func index(c *gin.Context) {
	msg := c.Query("msg")
	SendMsg(msg)
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})

}

func ping(c *gin.Context) {
	ip := c.ClientIP()
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("get hostname failed, err=%+v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Request processed by %v. Client IP: %v", hostname, ip),
	})
}

func secret(c *gin.Context) {
	corpid := os.Getenv("corpid")
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("corpid:%v", corpid),
	})
}
