package auth

import (
	"github.com/gin-gonic/gin"
	"my-im/src/kernel/server"
	"my-im/src/service/auth"
)

type LoginController struct {
	LoginService *auth.LoginService
}

func NewLoginController() *LoginController {
	return &LoginController{
		LoginService: auth.NewLoginService(),
	}
}

func (c *LoginController) Build(s *server.Server) {
	s.Handle("POST", "/login", c.Login)
	s.Handle("POST", "/register", c.Register)
}

//登录
func (c *LoginController) Login(ctx *gin.Context) {
	var loginForm LoginForm
	if ctx.ShouldBind(&loginForm) != nil {
		ctx.JSON(401, gin.H{
			"code":    401,
			"message": "unauthorized",
		})
	}

	token, err := c.LoginService.Login(loginForm.Mobile, loginForm.Password)
	if err != nil {
		ctx.JSON(401, gin.H{
			"code":    401,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"code": 200,
		"data": map[string]string{
			"token": token,
		},
	})
}

func (c *LoginController) Register(ctx *gin.Context) {
	var registerForm RegisterForm
	if ctx.ShouldBind(&registerForm) != nil {
		ctx.JSON(401, gin.H{
			"code":    401,
			"message": "unauthorized",
		})
	}

	token, err := c.LoginService.Register(registerForm.Mobile, registerForm.Password)
	if err != nil {
		ctx.JSON(401, gin.H{
			"code":    401,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"code": 200,
		"data": map[string]string{
			"token": token,
		},
	})
}
