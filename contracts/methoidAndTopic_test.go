package contracts

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"

	"github.com/ethereum/go-ethereum/common"

	"github.com/qmarliu/hdwallet"
)

func TestPrintMethodID(t *testing.T) {
	t.Log("flipallowclaim: " +
		SignToMethodID("flipAllowClaim()"))
}

func TestPrintTopic(t *testing.T) {
	t.Log("harvest: " +
		SignToTopic("Harvest(address,uint256,uint256)"))
	t.Log("deposit: " +
		SignToTopic("Deposit(address,uint256)"))
}

var (
	ethHdKey     *hdwallet.Key
	masterWallet hdwallet.Wallet
	masterAddr   common.Address
)

// go test -run TestNewMneomonic
func TestGetPKAndAddress(t *testing.T) {

	var err error
	ethHdKey, err = hdwallet.NewKey(
		hdwallet.Mnemonic(
			"shaft key hidden talent margin game behind belt wolf fall two helmet"),
	)
	if err != nil {
		t.Errorf("NewKey %s\n", err.Error())
		return
	}

	masterWallet, err = ethHdKey.GetWallet(
		hdwallet.CoinType(hdwallet.ETH), hdwallet.AddressIndex(0))
	if err != nil {
		t.Errorf("GetWallet %s\n", err.Error())
		return
	}

	masterAddrStr, err := masterWallet.GetAddress()
	if err != nil {
		t.Errorf("GetAddress %s\n", err.Error())
		return
	}
	masterAddr = common.HexToAddress(masterAddrStr)

	wallet, err := ethHdKey.GetWallet(
		hdwallet.CoinType(hdwallet.ETH),
		hdwallet.AddressIndex(1000))

	// pk, err := crypto.HexToECDSA(wallet.GetKey().PrivateHex())
	// if err != nil {
	// 	t.Errorf("HexToECDSA %s\n", err.Error())
	// 	return
	// }
	t.Logf("pk %v\n", wallet.GetKey().PrivateHex())
	addr, err := wallet.GetAddress()
	if err != nil {
		t.Errorf("GetAddress %s\n", err.Error())
		return
	}
	t.Logf("address %v\n", addr)
}

// go test -run TestSendETH
func TestSendETH(t *testing.T) {
	pk, err := crypto.HexToECDSA("pk")
	if err != nil {
		t.Errorf("HexToECDSA %s\n", err.Error())
		return
	}
	EthClient, err := ethclient.Dial(
		"https://eth-goerli.nodereal.io/v1/cf41970999784178af8cf1d4f175defe")
	if err != nil {
		t.Errorf("eth dial %s\n", err.Error())
		return
	}

	ChainID, err := EthClient.ChainID(context.Background())
	if err != nil {
		panic(err.Error())
	}

	value := decimal.NewFromBigInt(big.NewInt(2), 16).BigInt() // 0.02
	to := common.HexToAddress("0x3D64968faa49729c0b42B444Bb220557d30Ea360")
	tran, err := SendEth(EthClient, pk, to, value, ChainID)
	if err != nil {
		t.Errorf("sendeth %s\n", err.Error())
		return
	}

	t.Logf("hash %v", tran.Hash().String())
}
