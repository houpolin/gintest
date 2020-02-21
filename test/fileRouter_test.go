package test

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"

	"testing"

	"gintest/router"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUploadSingleRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)
	engine := router.SetupRouter()

	file, err := os.Open("./img/test.png")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	body := &bytes.Buffer{}

	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", filepath.Base("test.png"))
	if err != nil {
		t.Error(err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		t.Error(err)
	}
	_ = writer.Close()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodPost, "/file/uploadSingle", body)
	req.Header.Add("Content-type", writer.FormDataContentType())
	engine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	// r := gofight.New()
	// r.POST("file/uploadSingle").SetDebug(true).SetFileFromPath([]gofight.UploadFile{
	// 	{
	// 		Path: "./img/test.png",
	// 		Name: "file",
	// 	},
	// }).Run(engine, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
	// 	assert.Equal(t, http.StatusOK, r.Code)
	// })
}
