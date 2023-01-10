package net

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Ucpe []string

func (u Ucpe) Len() int { return len(u) }

// 字段为空
func (u Ucpe) Null() Ucpe {
	for i, length := 0, u.Len(); i < length; i++ {
		if u[i] == "" {
			break
		}
	}
	return u
}

// 版本不一致
func (u Ucpe) Version(max string) Ucpe {
	if u[2] == max {
		u[2] = Cyan(u[2])
	} else {
		u[2] = Red(u[2])
	}
	return u
}

// 时间不更新
func (u Ucpe) Time() Ucpe {
	var now = time.Now()
	synctime, _ := time.Parse("2006-01-02 15:04:05", u[3])
	if synctime.Year() != now.Year() || synctime.Month() != now.Month() || synctime.Day() != now.Day() || synctime.Hour() != now.Hour() {
		u[3] = fmt.Sprintf("%s✗%s", Red(strings.Split(u[3], " ")[0]), Red(strings.Split(u[3], " ")[1]))
	}
	return u
}

type Ucpes [][]string

func (us Ucpes) Len() int { return len(us) }

// 字段为空
func (us Ucpes) Null() Ucpes {
	for i, length := 0, us.Len(); i < length; i++ {
		for j, length := 0, len(us[i]); j < length; j++ {
			if us[i][j] == "" {
				break
			}
		}
	}
	return us
}

// 实现model倒序排序
func (us Ucpes) Less(i, j int) bool {
	var one, two int
	left := (strings.Split(us[i][1], "-")[1])
	right := (strings.Split(us[j][1], "-")[1])

	if left != "" && right != "" {
		one, _ = strconv.Atoi(left)
		two, _ = strconv.Atoi(right)
	}
	return one > two
}

func (us Ucpes) Swap(i, j int) { us[i], us[j] = us[j], us[i] }
