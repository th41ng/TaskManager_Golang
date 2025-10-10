package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func ErrorLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next() // chạy handler

		//kiểm tra lỗi
		errors := c.Errors
		if len(errors) > 0 {
			for _, e := range errors {
				fmt.Printf("[ERROR] %v | %s | %s\n", time.Since(start), c.Request.Method, e.Error())
			}
			c.JSON(-1, gin.H{
				"status":  "error",
				"message": errors[0].Error(),
			})
			return
		}
	}
}
