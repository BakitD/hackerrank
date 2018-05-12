package warmup

import (
	"strconv"
	"strings"
)

func TimeConversion(s string) string {
	n := len(s)
	parts := strings.Split(s, ":")
	last := parts[2][:2]
	var result string
	if parts[2][2:] == "AM" {
		if parts[0] == "12" {
			result = "00" + s[2:n-2]
		} else {
			result = s[:n-2]
		}
	} else {
		newTime, _ := strconv.Atoi(parts[0])
		first := strconv.Itoa(newTime + 12)
		if first == "24" {
			first = "12"
		}
		result = strings.Join([]string{first, parts[1], last}, ":")
	}
	return result
}
