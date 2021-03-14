package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"my-im/src/kernel/orm"
	"my-im/src/model"
)

type LoginService struct {
	DB *gorm.DB
}

func NewLoginService() *LoginService {
	return &LoginService{
		DB: orm.DB,
	}
}

//登录
func (s *LoginService) Login(mobile, password string) (token string, err error) {
	user := &model.User{}
	result := s.DB.First(user, "mobile = ?", mobile)

	if result.Error != nil {
		return "", result.Error
	}
	if result.RowsAffected == 0 {
		return "", errors.New("用户不存在")
	}

	err = user.CheckPassword(password)
	if err != nil {
		return "", err
	}

	token, err = user.GenerateJWT()
	if err != nil {
		return "", err
	}

	return token, err
}

//注册
func (s *LoginService) Register(mobile, password string) (token string, err error) {
	user := model.User{
		Mobile: mobile,
	}
	passwordHas, err := user.GeneratePassword(password)
	if err != nil {
		return "", err
	}
	user.Password = passwordHas

	result := s.DB.Create(&user)
	if result.Error != nil {
		return "", result.Error
	}

	token, err = user.GenerateJWT()
	if err != nil {
		return "", err
	}

	return token, err
}

//检查token
func CheckAuth(token string) (user model.UserClaim, err error) {
	sec := []byte("acb123")

	//u := &model.UserClaim{}
	getToken, err := jwt.ParseWithClaims(token, &user, func(token *jwt.Token) (interface{}, error) {
		return sec, nil
	})

	if err != nil {
		return model.UserClaim{}, err
	}

	if !getToken.Valid {
		return model.UserClaim{}, errors.New("token解析错误")
	}

	return user, nil
}
