package service

import (
	"errors"
	"github.com/CiroLong/Horizon-social/internal/model"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	// 这里可以注入其他依赖，如数据库连接
}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) Register(username, password string) error {
	// 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 这里应该是数据库操作，暂时用内存存储模拟
	model.Users[username] = model.User{
		Username: username,
		Password: string(hashedPassword),
	}

	return nil
}

func (us *UserService) Login(username, password string) (string, error) {
	user, exists := model.Users[username]
	if !exists {
		return "", errors.New("user not found")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid password")
	}

	// 这里应该生成JWT token，暂时返回模拟token
	return "mock-token-for-" + username, nil
}
