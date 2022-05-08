package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/xen0n/go-workwx"
	"log"
	"net/http"
	"os"
)

type Config struct {
	CorpID  string
	Secret  string
	AgentID int64
	UserIDs []string

	SentryDSN string
}

var conf Config

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}

	return value
}

func init() {
	configFile := getenv("WX_CONFIG_FILE", "config.toml")
	_, err := toml.DecodeFile(configFile, &conf)
	if err != nil {
		log.Fatalf("DecodeFile failed, err=%+v", err)
	}
}

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: conf.SentryDSN,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	app := gin.Default()

	app.GET("/ping", func(c *gin.Context) {
		ip := c.ClientIP()
		hostname, err := os.Hostname()
		if err != nil {
			log.Fatalf("get hostname failed, err=%+v", err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Request processed by %v. Client IP: %v", hostname, ip),
		})
	})

	app.GET("/", func(c *gin.Context) {
		msg := c.Query("msg")
		SendMsg(msg)
		c.JSON(http.StatusOK, gin.H{
			"message": msg,
		})
	})
	app.Run()
}

func SendMsg(msg string) {
	client := workwx.New(conf.CorpID)

	app := client.WithApp(conf.Secret, conf.AgentID)
	// preferably do this at app initialization
	app.SpawnAccessTokenRefresher()

	// send to user(s)
	to := workwx.Recipient{
		UserIDs: conf.UserIDs,
	}
	err := app.SendTextMessage(&to, msg, false)
	if err != nil {
		log.Printf("SendTextMessage err:%v\n", err)
	}

}
