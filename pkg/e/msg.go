package e

// MsgFlags 状态码map
var MsgFlags = map[int]string{
	SUCCESS:               "ok",
	UpdatePasswordSuccess: "修改密码成功",
	NotExistIdentifier:    "该第三方账号未绑定",
	InvalidParams:         "请求参数错误",
	ERROR:                 "fail",
}

// GetMsg 获取状态码对应的信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
