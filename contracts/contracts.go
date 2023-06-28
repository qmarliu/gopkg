package contracts

import (
	"context"
	"crypto/ecdsa"
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

func GetAuth2(ethCli *ethclient.Client, pk string, chainID *big.Int) (auth *bind.TransactOpts, err error) {
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return
	}
	return GetAuth(ethCli, privateKey, chainID)
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
	return SendEthWithAuth(ethCli, pk, to, value, chainID, auth)
}

func SendEthWithAuth(ethCli *ethclient.Client, pk *ecdsa.PrivateKey, to common.Address, value *big.Int, chainID *big.Int, auth *bind.TransactOpts) (*types.Transaction, error) {

	auth.GasLimit = uint64(21000)
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
