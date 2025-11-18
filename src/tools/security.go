package tools

import (
	"regexp"
	"strings"
)

// -----------------------------
// 黑名单模式（SQL/JNDI）
// -----------------------------

var blacklistPatterns = []*regexp.Regexp{
	// -----------------------------
	// SQL 注入高频关键词
	// -----------------------------
	regexp.MustCompile(`(?i)s\s*e\s*l\s*e\s*c\s*t`),
	regexp.MustCompile(`(?i)i\s*n\s*s\s*e\s*r\s*t`),
	regexp.MustCompile(`(?i)u\s*p\s*d\s*a\s*t\s*e`),
	regexp.MustCompile(`(?i)d\s*e\s*l\s*e\s*t\s*e`),
	regexp.MustCompile(`(?i)d\s*r\s*o\s*p`),
	regexp.MustCompile(`(?i)t\s*r\s*u\s*n\s*c\s*a\s*t\s*e`),
	regexp.MustCompile(`(?i)u\s*n\s*i\s*o\s*n\s+.*s\s*e\s*l\s*e\s*c\s*t`),
	regexp.MustCompile(`(?i)or\s+1\s*=\s*1`),
	regexp.MustCompile(`(?i)and\s+1\s*=\s*1`),

	// -----------------------------
	// SQL 特征符号攻击
	// -----------------------------
	regexp.MustCompile(`--`),
	regexp.MustCompile(`;`),
	regexp.MustCompile(`/\*.*?\*/`), // /**/ 绕过
	regexp.MustCompile(`%27`),       // URL 编码单引号
	regexp.MustCompile(`%3B`),       // URL 编码分号

	// -----------------------------
	// JNDI / Log4j 远程注入
	// -----------------------------
	regexp.MustCompile(`(?i)\$\s*\{.*j\s*n\s*d\s*i.*\}`),
	regexp.MustCompile(`(?i)\$\s*\{.*ldap.*\}`),
	regexp.MustCompile(`(?i)\$\s*\{.*rmi.*\}`),
	regexp.MustCompile(`(?i)\$\s*\{.*dns.*\}`),

	// 常见绕过格式
	regexp.MustCompile(`(?i)\$\{.*\$\{.*\}.*\}`), // 多层嵌套 payload

	// -----------------------------
	// 模板注入（EL/OGNL）
	// -----------------------------
	regexp.MustCompile(`(?i)#{.*}`),
	regexp.MustCompile(`(?i)\$\{.*\}`),

	// -----------------------------
	// 命令执行攻击
	// -----------------------------
	regexp.MustCompile(`(?i)\bcat\b`),
	regexp.MustCompile(`(?i)\bping\b`),
	regexp.MustCompile(`(?i)\bwget\b`),
	regexp.MustCompile(`(?i)\bcurl\b`),
	regexp.MustCompile(`(?i)\bnc\b`),
	regexp.MustCompile(`(?i)\bsh\b`),
	regexp.MustCompile(`(?i)\bbash\b`),

	// -----------------------------
	// MongoDB 注入 / NoSQL 注入
	// -----------------------------
	regexp.MustCompile(`(?i)\$where`),
	regexp.MustCompile(`(?i)\$regex`),
	regexp.MustCompile(`(?i)\$ne`),

	// -----------------------------
	// Base64 包裹攻击（常见）
	// -----------------------------
	regexp.MustCompile(`(?i)dG1w`),     // "tmp" base64
	regexp.MustCompile(`(?i)Y3VybA==`), // curl
}

/*
加入清洗步骤：统一还原编码（防绕过）

攻击者常用：

%20、%2F%2A%2A%2F（编码绕过）

\u0061\u0062（Unicode）

# Base64

混合大小写

所以在检测前，我们先统一 clean 一次：
*/
func normalizeInput(s string) string {
	s = strings.ToLower(s)

	// 去除空白 tab 换行
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.ReplaceAll(s, "\r", "")
	s = strings.ReplaceAll(s, "\t", "")
	s = strings.Join(strings.Fields(s), "")

	// 去掉 URL %
	s = strings.ReplaceAll(s, "%", "")

	// 去掉多余符号
	s = strings.ReplaceAll(s, "`", "")
	s = strings.ReplaceAll(s, "\\", "")
	s = strings.ReplaceAll(s, "\"", "")

	return s
}

// 黑名单判断函数
func IsMalicious(s string) bool {
	s = normalizeInput(s)

	for _, r := range blacklistPatterns {
		if r.MatchString(s) {
			return true
		}
	}
	return false
}
