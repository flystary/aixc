package cpe

import (
	"strconv"
	"strings"
)

func compareVersion(a, b string) int {
	as := strings.Split(a, ".")
	bs := strings.Split(b, ".")

	n := len(as)
	if len(bs) > n {
		n = len(bs)
	}

	for i := 0; i < n; i++ {
		ai, bi := 0, 0

		if i < len(as) {
			ai, _ = strconv.Atoi(as[i])
		}
		if i < len(bs) {
			bi, _ = strconv.Atoi(bs[i])
		}

		if ai > bi {
			return 1
		} else if ai < bi {
			return -1
		}
	}
	return 0
}
