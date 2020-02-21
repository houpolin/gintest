package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"gintest/router"

	"github.com/stretchr/testify/assert"
)

func TestIUserGetRouter(t *testing.T) {
	type User struct {
		UserID string `json:"userId"`
	}
	user := User{
		UserID: "123",
	}
	expectedBody, _ := json.Marshal(user)
	router := router.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user/"+user.UserID, nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(expectedBody)+"\n", w.Body.String())
}

func TestIUserLoginRouter(t *testing.T) {
	value := url.Values{}
	value.Add("email", "test@test.com")
	value.Add("password", "test")
	value.Add("password-again", "test")

	router := router.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/user/login", bytes.NewBufferString(value.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
