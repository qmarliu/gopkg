package hdwallet

import (
	"encoding/base64"
	"testing"

	encrypt "github.com/qmarliu/gopkg/encrypt/aes"

	"github.com/qmarliu/hdwallet"
)

// go test -run TestDes
func TestNewMneomonic(t *testing.T) {
	mneomonic, err := hdwallet.NewMnemonic(12, hdwallet.English)
	if err != nil {
		t.Errorf("NewMnemonic: %s\n", err.Error())
		return
	}
	t.Logf("请妥善保管好您的助记词\n----------\n%s\n----------\n", mneomonic)

	// 助记词加密
	aesKey := "1234567890abcdef"
	if len(aesKey) != 16 {
		t.Errorf("aesKey 参数必须有且仅有16位长度才能加密\n")
		return
	}
	origData := []byte(mneomonic)
	key := []byte(aesKey)
	en := encrypt.AESEncrypt(origData, key)
	t.Logf("aesCrypt: %v", base64.StdEncoding.EncodeToString(en))
	de := encrypt.AESDecrypt(en, key)
	t.Logf("dec: %v", string(de))
	return
}

// go test -run TestDes
func TestMneomonicHdid(t *testing.T) {
	//定义明文
	mneomonic := "random cruise delay salmon butter track toss absorb science breeze scan file"
	hdKey, err := hdwallet.NewKey(
		hdwallet.Mnemonic(mneomonic),
	)
	if err != nil {
		t.Errorf("new key err: %v", err)
		return
	}
	var hdid uint32 = 0
	wallet, err := hdKey.GetWallet(hdwallet.CoinType(hdwallet.ETH), hdwallet.AddressIndex(hdid))
	if err != nil {
		t.Errorf("get wallet: %v", err)
		return
	}
	address, err := wallet.GetAddress()
	if err != nil {
		t.Errorf("get address: %v", address)
		return
	}
	pk := wallet.GetKey().PrivateHex()
	if err != nil {
		t.Errorf("get pk: %v", address)
		return
	}
	t.Log("hdid", hdid)
	t.Log("address", address)
	t.Log("pk", pk)
	return
}
