package tools

import (
	"fmt"
	"math/rand"
	"net/url"
	"path"
	"strings"
	"time"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// StringBuilderBody 对string进行链式处理
type StringBuilderBody struct {
	origin string
}

// StringBuilder new stringbuilder
func StringBuilder(origin string) *StringBuilderBody {
	return &StringBuilderBody{
		origin: origin,
	}
}

// Replace 替换函数
func (s *StringBuilderBody) Replace(old, new string) *StringBuilderBody {
	s.origin = strings.Replace(s.origin, old, new, -1)
	return s
}

// Build 返回字符串
func (s *StringBuilderBody) Build() string {
	return s.origin
}

// 获取随机字符串
func GetRandomString(length uint64) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	for i := uint64(0); i < length; i++ {
		result = append(result, bytes[r.Intn(int(len(bytes)))])
	}
	return string(result)
}

// 返回包含引号("")的字符串
func Join(a []string, sep string) string {
	switch len(a) {
	case 0:
		return ""
	case 1:
		return "\"" + a[0] + "\""
	}
	n := len(sep) * (len(a) - 1)
	for i := 0; i < len(a); i++ {
		n += len(a[i]) + 2
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString("\"" + a[0] + "\"")
	for _, s := range a[1:] {
		b.WriteString(sep)
		b.WriteString("\"" + s + "\"")
	}

	return b.String()
}

// 返回包含引号("")的字符串 => [1,2,3]
func JoinInt(n []int64, sep string) string {

	if len(n) <= 0 {
		return "[]"
	}

	var b strings.Builder
	b.Write([]byte(fmt.Sprintf("[%v", n[0])))
	for _, s := range n[1:] {
		b.WriteString(sep)
		b.Write([]byte(fmt.Sprintf("%v", s)))
	}
	b.Write([]byte("]"))
	return b.String()
}

// StringKeywordsOrderCheck
func StringKeywordsOrderCheck(s *string, keywords []string) bool {

	c := *s
	if len(c) <= 0 {
		println("[ERROR]=>s is empty")
		return false
	}

	// 预处理文本
	c = Pretext(c)

	n := 0 //记录位置
	for _, word := range keywords {
		if n >= len(c) {
			return false
		}
		i := strings.Index(c[n:], word)
		if i < 0 {
			return false
		}
		n = n + i + len(word)
	}

	return true
}

// Pretext 文本预处理
func Pretext(s string) string {
	m := map[rune]rune{}
	sample := "~!@#$%^&*()_+-=[]{}\\|'\";:/?.>,< `	『』，。，、；‘【】「」（）▼“”👇：！↓……《》——-"
	for _, runeValue := range sample {
		m[runeValue] = runeValue
	}
	m[rune('\n')] = rune('\n')
	m[rune('\t')] = rune('\t')
	m[rune('\r')] = rune('\r')

	news := strings.Builder{}
	for _, a := range s {
		if _, ok := m[a]; ok {
			continue
		}

		news.WriteRune(a)
	}

	return news.String()
}

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func Padding(s string, c int) string {
	for i := 0; i < c; i++ {
		s = s + "#"
	}
	return s
}

func IsEmpty(str string) bool {
	return strings.TrimSpace(str) == ""
}

func JoinURL(base string, paths ...string) string {
	u, _ := url.Parse(base)
	u.Path = path.Join(u.Path, path.Join(paths...))
	return u.String()
}

// Trim  trim string
func Trim(str string) string {
	return strings.TrimSpace(str)
}
