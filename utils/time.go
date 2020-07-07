/****
	高性能时间，精确到0.1秒
****/
package utils

import (
	"fmt"
	"strings"
	"time"
)

var Now time.Time
var Local *time.Location

func init() {
	Now = time.Now().Round(time.Second)
	Local, _ = time.LoadLocation("Local")
	go refresh()
}

func refresh() {
	for {
		Now = time.Now().Round(time.Second)
		time.Sleep(100 * time.Millisecond)
	}
}

// 获取时间界限，如：today  返回stm: 2015-05-01 00:00:00  etm: 2015-05-02: 00:00:00
func TmLime(flag string) (start, end string) {
	start = "1970-01-01 00:00:00"
	end = "2070-01-01 00:00:00"
	if "today" == flag {
		start = Now.Format("2006-01-02") + " 00:00:00"
		end = Now.AddDate(0, 0, 1).Format("2006-01-02") + " 00:00:00"
	} else if "yesterday" == flag {
		start = Now.AddDate(0, 0, -1).Format("2006-01-02") + " 00:00:00"
		end = Now.Format("2006-01-02") + " 00:00:00"
	}
	return
}

func FormatPrevLogin(tm time.Time) (st string) {
	if Now.Before(tm) {
		return "在线"
	}
	du := Now.Sub(tm)
	switch {
	case int(du.Minutes()) == 0:
		return "在线"
	case du.Minutes() < 60:
		return ToString(int(du.Minutes())) + "分钟前"
	case du.Hours() < 24:
		return ToString(int(du.Hours())) + "小时前"
	case du.Hours() < 24*7:
		return ToString(int(du.Hours()/24)) + "天前"
	default:
		return "七天前"
	}

}

func FormatPrevLive(tm time.Time) (st string) {
	if Now.Before(tm) {
		return "1分钟前"
	}
	du := Now.Sub(tm)
	switch {
	case int(du.Minutes()) == 0:
		return "1分钟前"
	case du.Minutes() < 60:
		return ToString(int(du.Minutes())) + "分钟前"
	case du.Hours() < 24:
		return ToString(int(du.Hours())) + "小时前"
	case du.Hours() < 24*7:
		return ToString(int(du.Hours()/24)) + "天前"
	default:
		return "7天前"
	}
	return
}

func FormatTimeForbid(t int64) (st string) {
	tm := time.Unix(t, 0)
	du := tm.Sub(Now)

	switch {
	case int(du.Minutes()) == 0:
		return "1分钟内"
	case du.Minutes() < 60:
		return ToString(int(du.Minutes())) + "分钟"
	case du.Hours() < 24:
		return ToString(int(du.Hours())) + "小时"
	case du.Hours() < 24*7:
		return ToString(int(du.Hours()/24)) + "天"
	default:
		return "七天"
	}
	return
}

//打印超过exceed时长的时间
//Prams:
// 	key: 用于识别的关键字
// 	start: 起始时间
// 	exceed: 持续时间超过多久才打印
func PrintDuration(key string, start time.Time, exceed time.Duration) {
	dur := time.Now().Sub(start)
	if dur >= exceed {
		fmt.Println("Duration", key, ":", dur.Seconds())
	}
}

//从现在到[days]天后的[H:M:S]时刻的时长
func DurationTo(days int, H, M, S int) time.Duration {
	fmt.Println(Now.Minute(), Now.Second())
	seconds := (days*24+H-Now.Hour())*3600 + (M-Now.Minute())*60 + S - Now.Second()
	return time.Duration(seconds) * time.Second
}

func SecondsOfTheDay(tm time.Time) int64 {
	return int64(tm.Hour()*3600 + tm.Minute()*60 + tm.Second())
}

func SecondsOfTheMouth(tm time.Time) int64 {
	return int64((tm.Day()-1)*3600*24 + tm.Hour()*3600 + tm.Minute()*60 + tm.Second())
}

//这一周的秒数 从周一开始
func SecondsOfTheWeek(tm time.Time) int64 {
	w := int(tm.Weekday())
	if w == 0 {
		w = 7
	}
	w--
	return int64(w*3600*24 + tm.Hour()*3600 + tm.Minute()*60 + tm.Second())
}

//0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23
func CheckHourStr(hourStr string) bool {
	s := fmt.Sprintf(",%d,", Now.Hour())
	return strings.Contains(","+hourStr+",", s)
}

func CheckWeekStr(weekStr string) bool {
	s := fmt.Sprintf(",%d,", Now.Weekday())
	return strings.Contains(","+weekStr+",", s)
}

func DayBegin() int64 {
	return Now.Unix() - SecondsOfTheDay(Now)
}

func HourBegin() time.Time {
	t := Now
	tm := t.Unix() - int64(t.Minute()*60-t.Second())
	return time.Unix(tm, 0)
}

func DayMid() int64 {
	t := SecondsOfTheDay(Now)
	if t >= 22*3600 {
		return Now.Unix() - SecondsOfTheDay(Now) + 22*3600
	} else {
		return Now.Unix() - SecondsOfTheDay(Now) - 2*3600
	}
}

func DayBeginThen(t time.Time) int64 {
	return t.Unix() - SecondsOfTheDay(t)
}

func TodayTime(t time.Time) int64 {
	return DayBegin() + (t.Unix() - DayBeginThen(t))
}

func MouthBegin() int64 {
	return Now.Unix() - SecondsOfTheMouth(Now)
}

func MouthBeginThen(t time.Time) int64 {
	return t.Unix() - SecondsOfTheMouth(t)
}

func WeekBegin() int64 {
	return Now.Unix() - SecondsOfTheWeek(Now)
}

func WeekBeginThen(t time.Time) int64 {
	return t.Unix() - SecondsOfTheWeek(t)
}

func DayLeft() int64 {
	return 24*3600 - SecondsOfTheDay(Now)
}

//将UNIX时间戳格式化到时分秒
func FormatUnixTime(t int64) string {
	tm := time.Unix(t, 0)
	return tm.Format("2006-01-02 15:04:05")
}

//将Weekday枚举转成int
func ChineseWeekday(wd time.Weekday) int {
	if wd == time.Monday {
		return 0
	} else if wd == time.Tuesday {
		return 1
	} else if wd == time.Wednesday {
		return 2
	} else if wd == time.Thursday {
		return 3
	} else if wd == time.Friday {
		return 4
	} else if wd == time.Saturday {
		return 5
	} else if wd == time.Sunday {
		return 6
	}

	return 0
}

// AddDays returns the time t+d.
func AddDays(t time.Time, d int) time.Time {
	return t.Add(time.Hour * 24 * time.Duration(d))
}
