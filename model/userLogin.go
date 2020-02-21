package model

// UserLogin struct
type UserLogin struct {
	Email         string `form:"email" binding:"email"`
	Password      string `form:"password" binding:"required"`
	PasswordAgain string `form:"password-again" binding:"required,eqfield=Password"`
}
