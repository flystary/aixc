package net

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	. "aixc/tools/color"
)

type Ucpe []string

func (u Ucpe) Len() int { return len(u) }

// 字段为空
func (u Ucpe) NotNull() Ucpe {
	k := 0
	for _, v := range u {
		if v != "" {
			u[k] = v
			k++
		}
	}
	return u[:k]
}

// 版本不一致
func (u Ucpe) Version(max string) Ucpe {
	if len(u) <= 2 {
		return u
	}

	v := u[2]
	if v == "" {
		return u
	}

	if v == max {
		u[2] = Cyan(v)
	} else {
		u[2] = Red(v)
	}

	return u
}

// 时间不更新
func (u Ucpe) Time() Ucpe {
	if len(u) <= 3 {
		return u
	}

	timeStr := u[3]
	if timeStr == "" {
		return u
	}
	t, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		return u
	}

	now := time.Now()

	if t.Before(now.Add(-1 * time.Hour)) {
		parts := strings.SplitN(timeStr, " ", 2)
		if len(parts) == 2 {
			u[3] = fmt.Sprintf("%s✗%s",
				Red(parts[0]),
				Red(parts[1]),
			)
		}
	}
	return u
}

type Ucpes [][]string

func (us Ucpes) Len() int { return len(us) }

// 字段为空
func (us Ucpes) NotNull() Ucpes {
	for i, lengths := 0, us.Len(); i < lengths; i++ {
		for j, length := 0, len(us[i]); j < length; j++ {
			if us[i][j] == " " {
				break
			}
		}
	}
	return us
}

// 实现model倒序排序
func (us Ucpes) Less(i, j int) bool {
	if len(us[i]) <= 1 || len(us[j]) <= 1 {
		return false
	}

	leftParts := (strings.Split(us[i][1], "-"))
	rightParts := (strings.Split(us[j][1], "-"))

	if len(leftParts) < 2 || len(rightParts) < 2 {
		return false
	}

	left, _ := strconv.Atoi(leftParts[1])
	right, _ := strconv.Atoi(rightParts[1])

	return left > right
}

func (us Ucpes) Swap(i, j int) { us[i], us[j] = us[j], us[i] }
