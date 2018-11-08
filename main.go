package main

import (
	"github.com/gin-gonic/gin"
	"os/exec"
	"fmt"
)

func main() {
	r := gin.Default()
	r.GET("/pull", func(c *gin.Context) {
		cmd := exec.Command("./gitpull.sh"," front")
		opBytes, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
		}
		c.JSON(200, gin.H{
			"message": string(opBytes),
		})
	})
	r.Run(":8989") // listen and serve on 0.0.0.0:8080
}