package model

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    User   `json:"data"`
}

type EmailReq struct {
	ToUser   string `json:"toUser"`
	UserName string `json:"userName"`
}

type User struct {
	Id           int    `json:"id"`
	NickName     string `json:"nickName"`
	Password     string `json:"password"`
	RegisterDate int64  `json:"registerDate"`
	Sex          int    `json:"sex"`
	UserName     string `json:"userName"`
}
