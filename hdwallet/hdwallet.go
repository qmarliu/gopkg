package hdwallet

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/qmarliu/hdwallet"
)

var (
	hdKey      *hdwallet.Key
	MasterPk   *ecdsa.PrivateKey
	MasterAddr common.Address
)

func InitHDWallet(mneomonic string) (err error) {
	hdKey, err = hdwallet.NewKey(
		hdwallet.Mnemonic(mneomonic),
	)
	if err != nil {
		return SetMasterPk(0)
	}
	return err
}

func SetMasterPk(hdid uint32) (err error) {
	MasterPk, err = GetHDPvkey(hdid)
	if err != nil {
		return
	}
	addr := ""
	addr, err = GetHDAddr(hdid)
	MasterAddr = common.HexToAddress(addr)
	return
}

func GetHDAddr(hdid uint32) (string, error) {
	wallet, err := hdKey.GetWallet(hdwallet.CoinType(hdwallet.ETH), hdwallet.AddressIndex(hdid))
	if err != nil {
		return "", err
	}
	address, err := wallet.GetAddress()
	return address, err
}

func GetHDPvkey(hdid uint32) (*ecdsa.PrivateKey, error) {
	wallet, err := hdKey.GetWallet(hdwallet.CoinType(hdwallet.ETH), hdwallet.AddressIndex(hdid))
	if err != nil {
		return nil, err
	}
	return wallet.GetKey().PrivateECDSA, nil
}
