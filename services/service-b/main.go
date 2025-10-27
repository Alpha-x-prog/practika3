package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

type aResp struct {
	Message string `json:"message"`
}

func main() {
	port := getenv("PORT", "5002")
	aURL := getenv("SERVICE_A_URL", "http://service-a:5001")

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		var ar aResp
		resp, err := resty.New().R().SetResult(&ar).Get(aURL + "/")
		if err != nil || resp.IsError() {
			c.JSON(http.StatusBadGateway, gin.H{"message": "Hi from service B. Service A is unavailable"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Hi from service B. Service A said: " + ar.Message,
		})
	})

	r.GET("/healthz", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"ok": true}) })
	_ = r.Run(":" + port)
}

func getenv(k, d string) string { if v := os.Getenv(k); v != "" { return v }; return d }

