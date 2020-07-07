package utils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

/*
字符串截取函数

参数：
	str:带截取字符串
	begin:开始截取位置
	length:截取长度
*/
func SubString(str string, begin, length int) (substr string) {
	// 将字符串的转换成[]rune
	rs := []rune(str)
	lth := len(rs)
	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}
	// 返回子串
	return string(rs[begin:end])
}

// 生成一个随机int64
func RandInt(i int64) int64 {
	if i == 0 {
		return 0
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int63n(i)
}

func MinInt64(a, b int64) (r int64) {
	if a > b {
		r = b
	} else {
		r = a
	}
	return
}

func MaxInt64(a, b int64) (r int64) {
	if a > b {
		r = a
	} else {
		r = b
	}
	return
}

func RandInt2(i int) int {
	if i == 0 {
		return 0
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(i)
}

const (
	snum = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
)

func FilterEmoji(content string) string {
	added := false
	newContent := ""
	for _, value := range content {
		_, size := utf8.DecodeRuneInString(string(value))
		if size <= 3 {
			newContent += string(value)
		} else {
			if !added {
				newContent += "{emoji表情}"
				added = true
			}
		}
	}
	return newContent
}

func TokenGetUid(token string) (uid uint64) {
	var ks []string
	ks = strings.Split(token, "_")
	if len(ks) == 2 {
		uid, _ = ToUint64(ks[0])
	}
	return
}

func SubAtNameString(str string) (nameList []string) {
	if ok := strings.Contains(str, "@all "); ok {
		nameList = append(nameList, "all")
		return
	}
	atIdx := strings.Index(str, "@")
	atStrs := strings.Split(str[atIdx+1:], "@")
	for _, v := range atStrs {
		placeIdx := strings.Index(v, " ")
		if placeIdx == -1 {
			continue
		}
		nameList = append(nameList, v[:placeIdx])
	}
	return
}

func IntFenToYuanStr(fen int64) (yuan string) {
	yuan = strconv.FormatFloat(float64(fen)/100, 'f', -1, 64)
	return
}

func Int01FenToYuanStr(fen int64) (yuan string) {
	yuan = strconv.FormatFloat(float64(fen)/10000, 'f', -1, 64)
	return
}
