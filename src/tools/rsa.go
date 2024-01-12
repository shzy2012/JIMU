package tools

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

/*
通过公钥加密

	plaintext: 待加密明文
	pubkey:公钥.

rsa_public.pem
格式如下:

-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDL6629+ogb1agpruLeIy4lvpOe
6n+6uLz/GqgbOXvS7SjEEuBDGyrdwjqZwMvKwl2otALqYt+FhraGL3UlspdFUEoK
UK1EtFvtyuEQahgjuddjBv5KiE75mwPHzzM25nJfRgjIPjBKam20GRZzaaV/yPyG
NE329bI61hUMe6JEfwIDAQAB
-----END PUBLIC KEY-----
*/
func RsaEncode(plaintext, pubkey string) (string, error) {

	// 生成加密后的密文
	block, _ := pem.Decode([]byte(pubkey))
	if block == nil {
		return "", errors.New("公钥错误")
	}

	pubKey, err2 := x509.ParsePKIXPublicKey(block.Bytes)
	if err2 != nil {
		return "", err2
	}

	encryptedData, err3 := rsa.EncryptPKCS1v15(rand.Reader, pubKey.(*rsa.PublicKey), []byte(plaintext))
	if err3 != nil {
		return "", err2
	}

	uEnc := base64.URLEncoding.EncodeToString([]byte(encryptedData))
	// fmt.Println(uEnc)
	return uEnc, nil
}
