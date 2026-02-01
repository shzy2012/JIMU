package user

import "tolo/src/tools"

type ReqDto struct {
	Phone  string `json:"phone"`
	Passwd string `json:"passwd"`
	Code   string `json:"code"` //收集验证码
}

type RespDto struct {
	ID    string `json:"id"`
	Phone string `json:"phone"`
	Name  string `json:"name"`
	Role  string `json:"role"`
	Token string `json:"token"`
}

type Params struct {
	tools.Paging
}
