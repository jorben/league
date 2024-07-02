package errs

// 定义业务错误码
const (
	Success = 0 // 成功

	ErrAuthLoginUrl        = -10001 // 获取登录URL失败
	ErrAuthUnknownProvider = -10002 // 未知的登录方式
	ErrAuthUserinfo        = -10003 // 获取用户信息失败
	ErrAuthLoginFailed     = -10004 // 登录失败
	ErrAuthNoLogin         = -10005 // 未登录
	ErrAuthUnauthorized    = -10006 // 未授权
)

// 定义错误码对应的错误描述
var errorMsg = map[int]string{
	Success:                "success",
	ErrAuthLoginUrl:        "获取登录url失败，请稍后重试",
	ErrAuthUnknownProvider: "未知的登录方式，请使用正确的登录渠道",
	ErrAuthUserinfo:        "获取用户信息失败",
	ErrAuthLoginFailed:     "登录/注册失败，请稍后重试",
	ErrAuthNoLogin:         "未登录或登录态已过期",
	ErrAuthUnauthorized:    "未授权或权限不足",
}

// GetErrorMsg 获取错误码对应的错误描述
func GetErrorMsg(code int) string {
	if msg, ok := errorMsg[code]; ok {
		return msg
	}
	return ""
}
