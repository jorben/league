package errs

import "errors"

// 定义特定error
var (
	ErrorHasChildren    = errors.New("has children")
	ErrorRecordNotFound = errors.New("record not found")
)

// 定义业务错误码
const (
	Success = 0 // 成功

	ErrAuthLoginUrl        = -10001 // 获取登录URL失败
	ErrAuthUnknownProvider = -10002 // 未知的登录方式
	ErrAuthUserinfo        = -10003 // 获取用户信息失败
	ErrAuthLoginFailed     = -10004 // 登录失败
	ErrAuthNoLogin         = -10005 // 未登录
	ErrAuthUnauthorized    = -10006 // 未授权
	ErrAuthUnexpired       = -10007 // 未过期
	ErrAuthGroup           = -10008 //权限组操作失败
	ErrMenu                = -10030 // 获取菜单失败

	ErrNoRecord = -10900 // 没有匹配到预期的记录，无数据
	ErrDbSelect = -10901 // 查询异常
	ErrParam    = -10902 // 参数错误，缺少必要参数
	ErrDbUpdate = -10903 // 数据库更新异常
	ErrDbDelete = -10904 // 数据库删除异常
	ErrLogic    = -10905 // 逻辑错误
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
	ErrAuthUnexpired:       "刷新登录态失败，当前登录态还有足够长的有效期",
	ErrAuthGroup:           "用户组操作失败",
	ErrMenu:                "获取菜单失败，请稍后重试",
	ErrNoRecord:            "未查询到相关数据",
	ErrDbSelect:            "查询失败，请稍后重试",
	ErrParam:               "缺少必要参数或参数错误",
	ErrDbUpdate:            "更新失败，请稍后重试",
	ErrDbDelete:            "删除失败，请稍后重试",
	ErrLogic:               "逻辑错误",
}

// GetErrorMsg 获取错误码对应的错误描述
func GetErrorMsg(code int) string {
	if msg, ok := errorMsg[code]; ok {
		return msg
	}
	return ""
}
