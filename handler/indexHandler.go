package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetIndex function
func GetIndex(ctx *gin.Context) {
	//    ctx.HTML(http.StatusOK, "index.html", nil)
	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "IT Home again",
	})
}
