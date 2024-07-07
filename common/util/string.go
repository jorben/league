package util

import "regexp"

// IsValidEmail 验证字符串是否合法的邮件地址
func IsValidEmail(email string) bool {
	// 使用正则表达式匹配邮箱地址
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
