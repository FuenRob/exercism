package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	r := regexp.MustCompile(`(^\[TRC\]|^\[DBG\]|^\[INF\]|^[WRN]|^\[ERR\]|^\[FTL\])`)
    return r.MatchString(text)
}

func SplitLogLine(text string) []string {
	r := regexp.MustCompile(`<[~*=-]*>`)
	return r.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	r := regexp.MustCompile(`"(.*)(?i)password(.*)"`)
	var sum int

	for _, val := range lines {
		res := r.FindAllString(val, -1)
		sum += len(res)
	}
	return sum
}

func RemoveEndOfLineText(text string) string {
	return regexp.MustCompile(`end-of-line\d+`).ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	r := regexp.MustCompile(`User\s+(\w+)`)
	for n, line := range lines {
		if match := r.FindStringSubmatch(line); match != nil {
			lines[n] = "[USR] " + match[1] + " " + line
		}
	}
	return lines
}
