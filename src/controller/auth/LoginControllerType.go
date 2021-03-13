package auth

type LoginForm struct {
	Mobile string `form:"mobile" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type RegisterForm struct {
	Mobile string `form:"mobile" binding:"required"`
	Password string `form:"password" binding:"required"`
}