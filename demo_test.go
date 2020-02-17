package fameex_api_demo

import (
	"fmt"
	"testing"
	"time"
)

const (
	GET = "GET"
	SecretKey = "your_secret_key"
)

func TestSign(t *testing.T) {
	requestPath := "/v1/api/spot/orderlist?buyType=-1&pageNum=1&pageSize=5"

	preHashString := PreHashString(fmt.Sprintf("%d", time.Now().Unix()), GET, requestPath, "")

	fmt.Println("preHash:", preHashString)

	signature, err := HmacSha256HexSigner(preHashString, SecretKey)
	fmt.Println("signature:", signature, "\nerror:", err)
}
