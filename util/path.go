// 文件路径相关
package util

import (
	"os"
	"path/filepath"
	"strings"
)

// 获取当前路径
func GetCurrentDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		ErrRecord("%s", err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
