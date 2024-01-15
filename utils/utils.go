package utils

import "regexp"

func IsInGroupS(str int64, groupUids []int64) bool {
	for _, s := range groupUids {
		if s == str {
			return true
		}
	}
	return false
}

func IsAdmins(uid int64, adminUids []int64) bool {
	for _, s := range adminUids {
		if s == uid {
			return true
		}
	}
	return false
}

func ContainsURL(str string) bool {
	// 定义匹配网址的正则表达式模式
	pattern := `(?i)(https?:\/\/)?([\w-]+\.)?([\w-]+\.[\w-]+)`
	// 编译正则表达式模式
	regex := regexp.MustCompile(pattern)
	// 使用正则表达式进行匹配
	return regex.MatchString(str)
}
