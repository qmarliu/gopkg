package encrypt

import (
	"encoding/base64"
	"fmt"
	"testing"
)

// go test -run TestDes
func TestDes(t *testing.T) {
	//定义明文
	plaintext := []byte("abcdefg")
	fmt.Println("明文", string(plaintext))
	//填充分组
	plaintext = GroupFill(plaintext)
	//加密
	ciptext := DesEnc(plaintext)
	fmt.Println("密文", base64.StdEncoding.EncodeToString(ciptext))
	//解密
	plaintext = DesDec(ciptext)
	fmt.Println("解密", string(plaintext))
}

// go test -run TestDecode
func TestDecode(t *testing.T) {
	result, err := base64.StdEncoding.DecodeString("6sC+GXGssE3NeZcWL5l/jQCsuayiaCKN")
	//result, err := base64.StdEncoding.DecodeString("GJIGXaRcJZSivjbTH2mRz/9Ca/IVc6eSzrtzReHdJEMKi4tY//KRDgMd7xZucgYYwhBvr047TPyo3kxCLSqo0t40AbQbd0cq")
	if err != nil {
		t.Error(err)
	}
	resultPlaintext := DesDec(result)
	fmt.Println("解密", string(resultPlaintext))
}
