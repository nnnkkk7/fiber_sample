package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	now := time.Now()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, world!")
	})
	fmt.Printf("経過: %vms\n", time.Since(now).Nanoseconds())
	r.Run(":3030")
}
