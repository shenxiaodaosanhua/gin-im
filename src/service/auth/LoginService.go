package auth

import (
	"errors"
	"gorm.io/gorm"
	"my-im/src/model"
)

type LoginService struct {
	DB *gorm.DB
}

func NewLoginService() *LoginService {
	return &LoginService{
		DB: model.DB,
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