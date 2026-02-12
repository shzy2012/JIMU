package tools

// 定义消息
const (
	MsgDefault       = "操作成功"
	MsgFailed        = "操作失败"
	MsgLoginSuccess  = "登陆成功"
	MsgLoginFailed   = "登录失败"
	MsgPasswdError   = "密码不正确"
	MsgSysFailed     = "系统错误"
	MsgSysParam      = "系统参数,不能操作"
	MsgNoRecord      = "没有找到数据"
	MsgParaMissed    = "参数错误"
	MsgParaParse     = "参数解析错误"
	MsgDataExists    = "数据已存在"
	MsgDataNotExists = "数据不存在"
	MsgUnknown       = "未知错误"
)

// 分页
type Paging struct {
	Page       int64  `json:"page" form:"page" url:"page"`
	Size       int64  `json:"size" form:"size" url:"size"`
	SortBy     string `json:"sortBy" form:"sortBy" url:"sortBy"`
	Descending bool   `json:"descending" form:"descending" url:"descending"`
}

// 数据集
type PageDS struct {
	Total int64       `json:"total"`
	Data  interface{} `json:"data"`
}

type Select struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Select2 struct {
	Label string `json:"label"`
	Value string `json:"value"`
}
