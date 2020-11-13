package main

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/getOpenId", getOpenId)
	r.Run(":9011") // listen and serve on 0.0.0.0:9011
}

//{
//session_key: "WXBFOAU9q/tyYUtC8VVacw==",
//openid: "o9Id85XFrFQuX0y0cM6HQmHsimVo"
//}
type getOpenIdResult struct {
	SessionKey string `json:"session_key"`
	OpenId     string `json:"openid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

//git pull
func getOpenId(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")

	code := c.Query("code")
	var res getOpenIdResult
	getOpenIdUrl := "https://api.weixin.qq.com/sns/jscode2session?appid=wxb6c50476dbddd673&secret=fb7db724cb2d77f0db6c15b12f9af47c&grant_type=authorization_code&js_code=" + code
	fmt.Println("getOpenIdUrl:", getOpenIdUrl)
	if err := getJSON(getOpenIdUrl, &res); err != nil {
		c.JSON(200, res)
	} else {
		c.JSON(200, res)
	}
	return
}

// GetJSON executes HTTP GET against specified url and tried to parse
// the response into out object.
func getJSON(url string, out interface{}) error {

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var reader io.ReadCloser
	switch res.Header.Get("Content-Encoding") {
	case "gzip":
		reader, _ = gzip.NewReader(res.Body)
		defer reader.Close()
	default:
		reader = res.Body
	}

	if res.StatusCode >= 400 {
		body, err := ioutil.ReadAll(reader)
		if err != nil {
			return err
		}
		return errors.New(string(body))
	}
	decoder := json.NewDecoder(reader)
	return decoder.Decode(out)
}
