package parsinglogfiles

import (
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
		if len(re.FindStringSubmatch(line)) > 0 {
			line = "[USR] " + re.FindStringSubmatch(line)[1] + " " + line
		}
		linesTaggd = append(linesTaggd, line)
	}

	return linesTaggd

}
