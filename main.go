package main

import (
        "github.com/gin-gonic/gin"
        "os/exec"
        "fmt"
        "strings"
)

func main() {
        r := gin.Default()
        r.GET("/pull", gitPull)
        r.GET("/deploy", deploy)
        r.Run(":8989") // listen and serve on 0.0.0.0:8080
}

//git pull
func gitPull(c *gin.Context) {
        service := c.Query("service")
        branch := c.Query("branch")
        cmd := exec.Command("./gitpull.sh",service,branch)
        opBytes, err := cmd.Output()
        if err != nil {
                fmt.Println(err)
        }
        str := strings.Replace(string(opBytes),"\n","<br/>",-1)
        c.Header("Content-Type", "text/html; charset=utf-8")
        c.String(200, str)
}

//deploy
func deploy(c *gin.Context) {
        service := c.Query("service")
        branch := c.Query("branch")
        cmd := exec.Command("./deploy.sh",service,branch)
        opBytes, err := cmd.Output()
        if err != nil {
                fmt.Println(err)
        }
        str := strings.Replace(string(opBytes),"\n","<br/>",-1)
        c.Header("Content-Type", "text/html; charset=utf-8")
        c.String(200, str)
}

