package utils

import (
	"runtime/debug"
	"strings"
)

func IsInTest() bool {
	// 堆栈信息按行分开
	stacks := strings.Split(string(debug.Stack()), "\n")
	for _, line := range stacks {
		// 堆栈中\t开头为文件路径
		if strings.HasPrefix(line, "\t") {
			// 去除堆栈中的行号
			path := strings.Split(strings.TrimSpace(line), ":")[0]
			// 是否为_test.go结尾
			if strings.HasSuffix(path, "_test.go") {
				return true
			}
		}
	}
	return false
}
