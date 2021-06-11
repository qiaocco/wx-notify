package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/xen0n/go-workwx"
	"net/http"
)

type Config struct {
	CorpID  string
	Secret  string
	AgentID int64
	UserIDs []string
}

var conf Config

func init() {
	_, err := toml.DecodeFile("config.toml", &conf)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	r.GET("/", func(c *gin.Context) {
		msg := c.Query("msg")
		SendMsg(msg)
		c.JSON(http.StatusOK, gin.H{
			"message": msg,
		})
	})
	r.Run()
}

func SendMsg(msg string) {
	corpID := conf.CorpID
	corpSecret := conf.Secret
	agentID := conf.AgentID
	userIDs := conf.UserIDs
	fmt.Printf("corpID:%v, corpSecret:%v, agentID:%v,userIDs:%v\n",
		corpID, corpSecret, agentID, userIDs)

	client := workwx.New(corpID)

	app := client.WithApp(corpSecret, agentID)
	// preferably do this at app initialization
	app.SpawnAccessTokenRefresher()

	// send to user(s)
	to1 := workwx.Recipient{
		UserIDs: userIDs,
	}
	err := app.SendTextMessage(&to1, msg, false)
	if err != nil {
		fmt.Printf("SendTextMessage err:%v\n", err)
	}

}
