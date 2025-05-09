package tools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const Long = "2006-01-02 15:04:05"
const Short = "2006-01-02"

// Convert string to int
func GetInt(v string) int {
	if strings.TrimSpace(v) == "" {
		return 0
	}

	i, err := strconv.Atoi(v)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return 0
	}

	return i
}

// Convert string to uint
func GetuInt(v string) uint {
	if strings.TrimSpace(v) == "" {
		return 0
	}

	i, err := strconv.Atoi(v)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return 0
	}

	return uint(i)
}

// Convert string to int
func GetInt64(v string) int64 {
	if strings.TrimSpace(v) == "" {
		return 0
	}

	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return 0
	}

	return i
}

// Convert string to float64
func GetFloat64(v string) float64 {
	if strings.TrimSpace(v) == "" {
		return float64(0)
	}

	i, err := strconv.ParseFloat(v, 64)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return float64(0)
	}

	return i
}

// 将普通时间格式转化为 RFC3339
func GetRFC3339(v string) string {
	if v == "" {
		return ""
	}
	t, err := time.Parse("2006-01-02 15:04:05", v)
	if err != nil {
		fmt.Printf("[GetRFC3339]=>%s\n", err.Error())
		return ""
	}

	return t.Format("2006-01-02T15:04:05Z07:00")
}

func GetDate(v string) (time.Time, error) {
	if v == "" {
		return time.Time{}, fmt.Errorf("empty date string")
	}

	// Try common date formats
	formats := []string{
		"2006-01-02",
		"2006/01/02",
		"02-01-2006",
		"02/01/2006",
		"2006.01.02",
		"02.01.2006",
	}

	for _, format := range formats {
		t, err := time.Parse(format, v)
		if err == nil {
			return t.Local(), nil
		}
	}

	return time.Time{}, fmt.Errorf("invalid date format: %s", v)
}

// 将普通时间格式转化为 time.Time
func GetDatetime(v string) (time.Time, error) {
	if v == "" {
		return time.Time{}, fmt.Errorf("empty date string")
	}

	// Try parsing with time component
	formats := []string{
		"2006-01-02 15:04:05",
		"2006/01/02 15:04:05",
		"02-01-2006 15:04:05",
		"02/01/2006 15:04:05",
	}

	for _, format := range formats {
		t, err := time.Parse(format, v)
		if err == nil {
			return t.Local(), nil
		}
	}

	return time.Time{}, fmt.Errorf("invalid date format: %s", v)
}

// https://juejin.cn/post/6844903648045039624
func IsChinesePhone(phone string) bool {
	reg1 := regexp.MustCompile(`^1(?:3[0-9]|4[5-9]|5[0-9]|6[12456]|7[0-8]|8[0-9]|9[0-9])[0-9]{8}$`)
	if reg1 == nil {
		return false
	}
	//根据规则提取关键信息
	if len(reg1.FindAllStringSubmatch(phone, 1)) > 0 {
		return true
	}

	return false
}

func ToBytes(t interface{}) []byte {
	bytes, _ := json.Marshal(t)
	return bytes
}

// 禁用Unicode转义
func ToBytes2(t interface{}) []byte {
	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(t); err != nil {
		log.Printf("%s\n", err.Error())
	}
	return buf.Bytes()
}

// 手机手机号 136****1389
func HideMiddlePhone(phoneNumber string) string {
	if len(phoneNumber) != 11 {
		return "Invalid phone number"
	}
	// 隐藏中间6位
	return phoneNumber[:3] + "****" + phoneNumber[7:]
}
