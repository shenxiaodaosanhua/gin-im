package model

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Id uint `gorm:"primaryKey" json:"id"`
	Mobile string `gorm:"unique" json:"mobile" form:"mobile"`
	Password string `json:"-"`
	Avatar string `json:"avatar"`
	State int `json:"state"`
	Info UserInfo `gorm:"foreignkey:UserId" json:"info"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserClaim struct {
	Id uint `json:"id"`
	Mobile string `json:"mobile"`
	State int `json:"state"`
	jwt.StandardClaims
}

//设置表名
func (u *User) TableName() string {
	return "user"
}

//生成token
func (u *User) GenerateJWT() (token string, err error) {
	sec := []byte("acb123")
	userClaim := UserClaim{
		Id: u.Id,
		Mobile: u.Mobile,
		State: u.State,
	}
	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaim)
	token, err = tokenObj.SignedString(sec)
	if err != nil {
		return "", err
	}

	return token, err
}

//生成密码
func (u *User) GeneratePassword(password string) (pass string, err error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashed), err
}

//校验密码
func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}