package net

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	co "aixc/utils/color"
)

type Ucpe []string

func (u Ucpe) Len() int { return len(u) }

// 字段为空
func (u Ucpe) NotNull() Ucpe {
	for i, length := 0, u.Len(); i < length; i++ {
        ui_str :=  u[i]
		if len(ui_str) == 0 {
			break
		}
	}
	return u
}

// 版本不一致
func (u Ucpe) Version(max string) Ucpe {
    max_version_str := u[2]
    if len(max_version_str) != 0 {
	    if max_version_str == max {
		    u[2] = co.Cyan(max_version_str)
	    } else {
		    u[2] = co.Red(max_version_str)
	    }
    }
	return u
}

// 时间不更新
func (u Ucpe) Time() Ucpe {
	var now = time.Now()
    time_str := u[3]
    if len(time_str) != 0 {
	    synctime, _ := time.Parse("2006-01-02 15:04:05", time_str)
	    if synctime.Year() != now.Year() || synctime.Month() != now.Month() || synctime.Day() != now.Day() || synctime.Hour() != now.Hour() {
		    u[3] = fmt.Sprintf("%s✗%s", co.Red(strings.Split(time_str, " ")[0]), co.Red(strings.Split(time_str, " ")[1]))
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
