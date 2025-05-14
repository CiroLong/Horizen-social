package model

type User struct {
	Username string `json:"username"`
	Password string `json:"-"`
}

// Users 内存存储模拟数据库
var Users = make(map[string]User)
