package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UploadSingleIndex function
func UploadSingleIndex(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	err = ctx.SaveUploadedFile(file, "./file/"+file.Filename)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"fileName": file.Filename,
		"size":     file.Size,
		"mimeType": file.Header,
	})
}
