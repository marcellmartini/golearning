package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	return regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`).MatchString(text)
}

func SplitLogLine(text string) []string {
	return regexp.MustCompile(`<[~*=-]*>`).Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)(".*password.*")`)
	count := 0

	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	return regexp.MustCompile(`end-of-line\d*`).ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	var linesTaggd []string
	re := regexp.MustCompile(`User *(?P<user>[A-Z][a-z]*[0-9]*)`)

	for _, line := range lines {
		uf := re.FindStringSubmatch(line)
		if uf != nil {
			line = fmt.Sprintf("[USR] %s %s", uf[1], line)
		}
		linesTaggd = append(linesTaggd, line)
	}

	return linesTaggd

}
