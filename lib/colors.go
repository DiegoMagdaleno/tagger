package lib

import (
	"strconv"
	"strings"
)

func foregroundColor(code int) string {
	colored := []string{"\x1b[38;5;", strconv.Itoa(code), "m"}
	return strings.Join(colored, "")
}

const Reset = "\x1b[0m"

var (
	ColorOfTag = map[string][2]string{
		"orange": [2]string{foregroundColor(166)},
		"red":    [2]string{foregroundColor(160)},
		"yellow": [2]string{foregroundColor(226)},
		"green":  [2]string{foregroundColor(040)},
		"blue":   [2]string{foregroundColor(033)},
		"purple": [2]string{foregroundColor(129)},
		"grey":   [2]string{foregroundColor(240)},
	}
)
