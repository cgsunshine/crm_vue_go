package comm

import "fmt"

func GetBusinessNumber(prefix string, total int64) string {
	return fmt.Sprintf("%s%s%04d", prefix, TimeFormatYearMonth(), total)
}
