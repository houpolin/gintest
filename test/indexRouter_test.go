package test

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "gintest/router"
)

var engine *gin.Engine

func init() {
    gin.SetMode(gin.TestMode)
    engine = router.SetupRouter()
}

func TestIndexGetRouter(t *testing.T) {
    w := httptest.NewRecorder()
    req, _ := http.NewRequest(http.MethodGet, "/", nil)

    engine.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
}
