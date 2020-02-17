package fameex_api_demo

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func PreHashString(_timestamp, _method, _requestPath, _body string) string {
	return _timestamp + strings.ToUpper(_method) + _requestPath + _body
}

func HmacSha256HexSigner(message string, secretKey string) (string, error) {
	mac := hmac.New(sha256.New, []byte(secretKey))
	_, err := mac.Write([]byte(message))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(mac.Sum(nil)), nil
}
