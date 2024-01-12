package tools

import (
	"fmt"
	"testing"
)

func TestRsaEncode(t *testing.T) {

	rsa_public := `
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDL6629+ogb1agpruLeIy4lvpOe
6n+6uLz/GqgbOXvS7SjEEuBDGyrdwjqZwMvKwl2otALqYt+FhraGL3UlspdFUEoK
UK1EtFvtyuEQahgjuddjBv5KiE75mwPHzzM25nJfRgjIPjBKam20GRZzaaV/yPyG
NE329bI61hUMe6JEfwIDAQAB
-----END PUBLIC KEY-----
`

	plaintext := "测试文本@123"
	result, err := RsaEncode(plaintext, rsa_public)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}

	fmt.Println(result)
}
