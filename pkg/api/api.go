package api

import (
	"encoding/json"
	"net/http"
	"os"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/mmcdole/gofeed"
)

// Start starts the  API server
func Start() {
	router := gin.Default()
	router.GET("/", hello)
	router.GET("/feed", getFeed)
	router.GET("/parse", parse)
	router.Run(":3000")
}

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "heya",
	})
}

func getFeed(c *gin.Context) {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("http://www.producthunt.com/feed")
	c.JSON(200, feed)
}

func parse(c *gin.Context) {
	url := "https://mercury.postlight.com/parser?url=" + c.Query("url")

	mercuryClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("x-api-key", os.Getenv("mercury"))
	resp, err := mercuryClient.Do(req)
	if err != nil {
		c.String(500, "oops! Error in fetching parsed content.")
	}
	defer resp.Body.Close()
	println("say what", resp.Body)
	var content struct {
		Title     string `json:"title"`
		Content   string `json:"content"`
		LeadImage string `json:"lead_image_url"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&content); err != nil {
		c.String(500, "oops! Error in decoding parsed content.")
	}
	c.JSON(200, content)
}
