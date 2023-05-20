package tools

import (
	"fmt"
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

func Unique(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
