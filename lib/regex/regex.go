package regex

import "regexp"

// 获取时间,匹配 ： 1990-05-04 15:04:05
func regexDataTime(text string) string {
	datePattern := `(\d{4}-\d{1,2}-\d{1,2}\s*\d{1,2}:\d{1,2}:\d{1,2})`
	dateRegex := regexp.MustCompile(datePattern)
	dates := dateRegex.FindAllString(text, -1)
	if len(dates) > 0 {
		return dates[0]
	}
	return ""
}
