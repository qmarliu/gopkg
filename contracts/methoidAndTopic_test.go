package contracts

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"

	"github.com/ethereum/go-ethereum/common"
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

// go test -run TestSendETHEip1599
func TestSendETHEip1599(t *testing.T) {
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
	tran, err := SendEthEip1599(EthClient, pk, to, value, ChainID)
	if err != nil {
		t.Errorf("sendeth %s\n", err.Error())
		return
	}

	t.Logf("hash %v", tran.Hash().String())
}

// go test -run TestSendETHBsc
func TestSendETHBsc(t *testing.T) {
	pk, err := crypto.HexToECDSA("be16eb0651b5f7f6c660ae00c3d0ce4ba72bc32e132a55a8633c7f06ab270871")
	if err != nil {
		t.Errorf("HexToECDSA %s\n", err.Error())
		return
	}
	EthClient, err := ethclient.Dial(
		"https://bsc-testnet.nodereal.io/v1/64b76360b0f4466fbed156191d07dd58")
	if err != nil {
		t.Errorf("eth dial %s\n", err.Error())
		return
	}

	ChainID, err := EthClient.ChainID(context.Background())
	if err != nil {
		panic(err.Error())
	}

	value := decimal.NewFromBigInt(big.NewInt(2), 16).BigInt() // 0.02
	to := common.HexToAddress("0xa46d90AeFDb38809bcf99BCb86ae24b9Beb2A950")
	tran, err := SendEth(EthClient, pk, to, value, ChainID)
	if err != nil {
		t.Errorf("sendeth %s\n", err.Error())
		return
	}

	t.Logf("hash %v", tran.Hash().String())
}
