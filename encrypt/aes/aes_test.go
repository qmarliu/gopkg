package encrypt

import (
	"encoding/base64"
	"fmt"
	"testing"
)

//go test -run TestAes
func TestAes(t *testing.T) {
	//定义明文
	// origData := []byte("amcv-IAG_DfnYd9hsMWzhTdgWEftiuvDnkj1cxZc")
	//origData := []byte("890ca2274dae5c8d91c60534a48f435d50f89aacebd27d4bd8128213c53ac8d7")
	origData := []byte("6271102e8b770eea39b241054fb793fb94a22c24ab949af01c9c661d53445ccc")
	//origData := []byte("114102d35fed80f958347d36a09ace5a3d99cc3ea5fbcdf45e0b62faa8fc92ec")
	//origData := []byte("3cefc096562a0f69e8239ea9aae6d824350799912799dfeea8aadea4d06f7e1e")

	key := []byte("k2k38drr4g8cck38")
	fmt.Println("明文", string(origData))
	en := AESEncrypt(origData, key)
	fmt.Println("密文", base64.StdEncoding.EncodeToString(en))
	//解密
	de := AESDecrypt(en, key)
	fmt.Println("解密", string(de))
}

//// go test -run TestDecode
//func TestDecode(t *testing.T) {
//	result, err := base64.StdEncoding.DecodeString("6sC+GXGssE3NeZcWL5l/jQCsuayiaCKN")
//	if err != nil {
//		t.Error(err)
//	}
//	resultPlaintext := AESDecrypt("yWiDWnFpwMjWfmfLeAcM6q4uNi8yWCEIP3Ti6D9kgv8Ig922mn0MDhhWRVUdusrRVtka2CbGqMytjvAq/3KWeB+7mCRwSPD7Xud/0t5kGE/Kt416/Wx6x43v76wTIEWl")
//	fmt.Println("解密", string(resultPlaintext))
//}
