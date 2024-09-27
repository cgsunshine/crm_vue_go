package comm

import (
	"fmt"
	"time"
)

func TimeFormatYearMonth() string {
	// 当前时间
	now := time.Now()

	// 获取年份的后两位
	year := now.Year() % 100

	// 获取月份（1-12）
	month := now.Month()

	// 格式化输出
	return fmt.Sprintf("%02d%02d", year, month)
}
