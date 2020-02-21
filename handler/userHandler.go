package handler

import (
	"gintest/model"
	"net/http"

	//    "encoding/json"

	"github.com/gin-gonic/gin"
)

// GetUser function
func GetUser(ctx *gin.Context) {
	uid := ctx.Param("uid")
	ctx.JSON(http.StatusOK, gin.H{"userId": uid})
	/*
	   type User struct {
	       UserId string `json:"userId"`
	   }

	   user := User{
	       UserId: uid,
	   }

	   expectedBody, _ := json.Marshal(user)

	   ctx.Data(http.StatusOK, "application/json", expectedBody)
	*/
}

// UserLogin function
func UserLogin(ctx *gin.Context) {
	var user model.UserLogin

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)

}
