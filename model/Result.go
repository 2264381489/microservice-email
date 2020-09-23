package model

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

type EmailReq struct {
	ToUser   string `json:"toUser"`
	UserName string `json:"userName"`
}
