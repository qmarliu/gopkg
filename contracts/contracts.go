package contracts

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/qmarliu/gopkg/ethutils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
)

func GetAuth(ethCli *ethclient.Client, privateKey *ecdsa.PrivateKey, chainID *big.Int) (auth *bind.TransactOpts, err error) {
	fromAddress, err := ethutils.GetAddressFromPrivateKey(privateKey)
	if err != nil {
		return
	}
	nonce, err := ethCli.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return
	}
	gasPrice, err := ethCli.SuggestGasPrice(context.Background())
	if err != nil {
		return
	}
	auth, err = bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return auth, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = uint64(0)  // in units
	auth.GasPrice = gasPrice
	return
}

func GetAuth2(ethCli *ethclient.Client, pk string, chainID *big.Int) (auth *bind.TransactOpts, err error) {
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return
	}
	return GetAuth(ethCli, privateKey, chainID)
}

func GetAuthEip1599(ethCli *ethclient.Client, privateKey *ecdsa.PrivateKey, chainID *big.Int) (auth *bind.TransactOpts, err error) {
	fromAddress, err := ethutils.GetAddressFromPrivateKey(privateKey)
	if err != nil {
		return auth, err
	}
	nonce, err := ethCli.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return auth, err
	}
	head, err := ethCli.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return auth, err
	}
	gasTipCap, err := ethCli.SuggestGasTipCap(context.Background())
	if err != nil {
		return auth, err
	}
	auth, err = bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return auth, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = uint64(0)  // in units
	auth.GasTipCap = gasTipCap
	auth.GasFeeCap = new(big.Int).Add(
		gasTipCap,
		new(big.Int).Mul(head.BaseFee, big.NewInt(2)),
	)
	return auth, err
}

func GetAuth2Eip1599(ethCli *ethclient.Client, pk string, chainID *big.Int) (auth *bind.TransactOpts, err error) {
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return
	}
	return GetAuthEip1599(ethCli, privateKey, chainID)
}

func SignToMethodID(funcSign string) (methodID string) {
	var swapFnSignature = []byte(funcSign)
	hash := sha3.NewLegacyKeccak256()
	hash.Write(swapFnSignature)
	methodID = hexutil.Encode(hash.Sum(nil)[:4])
	return
}

func SignToTopic(funcSign string) (topicSign string) {
	eventSignature := []byte(funcSign)
	hash := crypto.Keccak256Hash(eventSignature)
	topicSign = hash.Hex()
	return
}

func GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int,
	swapNumerator *big.Int, swapDenominator *big.Int) (amountOut *big.Int) {
	amountInWithFee := big.NewInt(0)
	numerator := big.NewInt(0)
	denominator := big.NewInt(0)
	amountOut = big.NewInt(0)

	amountInWithFee.Mul(amountIn, swapNumerator)
	numerator.Mul(amountInWithFee, reserveOut)
	denominator.Mul(reserveIn, swapDenominator)
	denominator.Add(denominator, amountInWithFee)
	amountOut.Div(numerator, denominator)
	return
}

func GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int,
	swapNumerator *big.Int, swapDenominator *big.Int) (amountIn *big.Int) {
	numerator := big.NewInt(0)
	denominator := big.NewInt(0)
	amountIn = big.NewInt(0)

	numerator.Mul(reserveIn, amountOut)
	numerator.Mul(numerator, swapDenominator)

	denominator.Sub(reserveOut, amountOut)
	denominator.Mul(denominator, swapNumerator)
	amountIn.Div(numerator, denominator)
	amountIn.Add(amountIn, big.NewInt(1))
	return
}

func SendEth(ethCli *ethclient.Client, pk *ecdsa.PrivateKey, to common.Address, value *big.Int, chainID *big.Int) (*types.Transaction, error) {

	auth, err := GetAuth(ethCli, pk, chainID)
	if err != nil {
		return nil, err
	}
	auth.GasLimit = uint64(21000)
	return SendEthWithAuth(ethCli, pk, to, value, chainID, auth)
}

func SendEthWithAuth(ethCli *ethclient.Client, pk *ecdsa.PrivateKey, to common.Address, value *big.Int, chainID *big.Int, auth *bind.TransactOpts) (*types.Transaction, error) {

	var data []byte
	tx := types.NewTransaction(auth.Nonce.Uint64(), to, value, auth.GasLimit, auth.GasPrice, data)

	chainID, err := ethCli.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), pk)
	if err != nil {
		return nil, err
	}

	err = ethCli.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, err
	}
	return signedTx, nil
}

func SendEthSubFee(ethCli *ethclient.Client, pk *ecdsa.PrivateKey, to common.Address, value *big.Int, chainID *big.Int) (*types.Transaction, error) {

	auth, err := GetAuth(ethCli, pk, chainID)
	if err != nil {
		return nil, err
	}
	auth.GasLimit = 21000
	feeAmount := big.NewInt(int64(auth.GasLimit))
	feeAmount.Mul(feeAmount, auth.GasPrice)
	if value.Cmp(feeAmount) <= 0 {
		return nil, fmt.Errorf("手续费不够 value %v feeAmount %v", value, feeAmount)
	}
	var data []byte
	tx := types.NewTransaction(auth.Nonce.Uint64(), to, value.Sub(value, feeAmount), auth.GasLimit, auth.GasPrice, data)

	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), pk)
	if err != nil {
		return nil, err
	}

	err = ethCli.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, err
	}
	return signedTx, nil
}

func SendEthEip1599(ethCli *ethclient.Client, pk *ecdsa.PrivateKey, to common.Address, value *big.Int, chainID *big.Int) (*types.Transaction, error) {

	auth, err := GetAuthEip1599(ethCli, pk, chainID)
	if err != nil {
		return nil, err
	}
	auth.GasLimit = uint64(21000)
	return SendEthWithAuthEip1599(ethCli, pk, to, value, chainID, auth)
}

func SendEthWithAuthEip1599(ethCli *ethclient.Client, pk *ecdsa.PrivateKey, to common.Address, value *big.Int, chainID *big.Int, auth *bind.TransactOpts) (*types.Transaction, error) {

	var data []byte
	baseTx := &types.DynamicFeeTx{
		To:        &to,
		Nonce:     auth.Nonce.Uint64(),
		GasFeeCap: auth.GasFeeCap,
		GasTipCap: auth.GasTipCap,
		Gas:       auth.GasLimit,
		Value:     value,
		Data:      data,
	}
	tx := types.NewTx(baseTx)

	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), pk)
	if err != nil {
		return nil, err
	}

	err = ethCli.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, err
	}
	return signedTx, nil
}
