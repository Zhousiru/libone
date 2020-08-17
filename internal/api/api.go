package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Run runs the API service.
func Run(addr string) error {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/storage/*path", GetFile)
	r.GET("/api/list_file", ListFile)

	return r.Run(addr)
}

func resp(c *gin.Context, respStatus int, msg string, payload interface{}) {
	c.JSON(statusMap[respStatus].HTTPCode, gin.H{"status": statusMap[respStatus].StatusMsg, "msg": msg, "payload": payload})
}
