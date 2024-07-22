package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/url"
	"os"
)

const (
	// TODO replace below sandbox app creds with your own obtained from recipient hub
	ClientId = "<your client id goes here>"
	Secret   = "<your secret goes here>"
)

func main() {
	r := gin.New()
	r.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/"),
		gin.Recovery(),
	)
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Exchanging auth code for a token",
		})
	})

	r.POST("/api/authcode", func(c *gin.Context) {
		tokenUrl := "https://sandbox-idp.ddp.akoya.com/token"
		authCode := c.PostForm("authcode")
		fmt.Println("Auth code: ", authCode)

		bodyData := url.Values{}
		bodyData.Set("grant_type", "authorization_code")
		bodyData.Set("code", authCode)
		bodyData.Set("redirect_uri", "https://recipient.ddp.akoya.com/flow/callback")
		req, err := http.NewRequest("POST", tokenUrl, bytes.NewBufferString(bodyData.Encode()))
		if err != nil {
			panic(err)
		}
		req.Header.Set("accept", "application/json")
		req.Header.Set("content-type", "application/x-www-form-urlencoded")
		req.SetBasicAuth(ClientId, Secret)

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Printf("client: error making http request: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("client: got response!\n")
		fmt.Printf("client: status code: %d\n", res.StatusCode)

		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("client: could not read response body: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("client: response body: %s\n", resBody)

		c.Redirect(302, "/api/received-token?token="+string(resBody))
	})

	r.GET("/api/received-token", func(c *gin.Context) {
		token := c.Query("token")
		c.HTML(http.StatusOK, "token.tmpl", gin.H{
			"title": "Received the access token",
			"token": token,
		})
	})
	r.Run(":8123") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
