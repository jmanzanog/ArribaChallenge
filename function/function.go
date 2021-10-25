package function

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SaludarFunc() func(c *gin.Context) {
	return func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	}
}

func Saludar1() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Jose Manzano!",
		})
	}
}

func ActionTest() func(c *gin.Context) {
	return func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	}
}

func ActionPost() func(c *gin.Context) {
	return func(c *gin.Context) {
		b := c.FullPath() == "/user/:name/*action" // true
		c.String(http.StatusOK, "%t", b)
	}
}

func CustomRecovery() func(c *gin.Context, recovered interface{}) {
	return func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}
