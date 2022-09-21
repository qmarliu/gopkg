package ethutils

import (
	"context"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"time"

	"crypto/ecdsa"

	"errors"
	"math/big"

	"reflect"

	"unsafe"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/shopspring/decimal"
	"golang.org/x/crypto/sha3"
)

func ZeroAddr() common.Address {
	return common.HexToAddress("0x0000000000000000000000000000000000000000")
}

// IsValidAddress validate hex address
func IsValidAddress(iaddress interface{}) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	switch v := iaddress.(type) {
	case string:
		return re.MatchString(v)
	case common.Address:
		return re.MatchString(v.Hex())
	default:
		return false
	}
}

// IsZeroAddress validate if it's a 0 address
func IsZeroAddress(iaddress interface{}) bool {
	var address common.Address
	switch v := iaddress.(type) {
	case string:
		address = common.HexToAddress(v)
	case common.Address:
		address = v
	default:
		return false
	}

	zeroAddressBytes := common.FromHex("0x0000000000000000000000000000000000000000")
	addressBytes := address.Bytes()
	return reflect.DeepEqual(addressBytes, zeroAddressBytes)
}

// ToDecimal wei to decimals
func ToDecimal(ivalue interface{}, decimals int) decimal.Decimal {
	value := new(big.Int)
	switch v := ivalue.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(value.String())
	result := num.Div(mul)

	return result
}

// ToWei decimals to wei
func ToWei(iamount interface{}, decimals int) *big.Int {
	amount := decimal.NewFromFloat(0)
	switch v := iamount.(type) {
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)

	wei := new(big.Int)
	wei.SetString(result.String(), 10)

	return wei
}

// CalcGasCost calculate gas cost given gas limit (units) and gas price (wei)
func CalcGasCost(gasLimit uint64, gasPrice *big.Int) *big.Int {
	gasLimitBig := big.NewInt(int64(gasLimit))
	return gasLimitBig.Mul(gasLimitBig, gasPrice)
}

// SigRSV signatures R S V returned as arrays
func SigRSV(isig interface{}) ([32]byte, [32]byte, uint8) {
	var sig []byte
	switch v := isig.(type) {
	case []byte:
		sig = v
	case string:
		sig, _ = hexutil.Decode(v)
	}

	sigstr := common.Bytes2Hex(sig)
	rS := sigstr[0:64]
	sS := sigstr[64:128]
	R := [32]byte{}
	S := [32]byte{}
	copy(R[:], common.FromHex(rS))
	copy(S[:], common.FromHex(sS))
	vStr := sigstr[128:130]
	vI, _ := strconv.Atoi(vStr)
	V := uint8(vI + 27)

	return R, S, V
}

func GetHexAddressFromPrivateKey(privateKey *ecdsa.PrivateKey) (address string, err error) {
	commonAddress, err := GetAddressFromPrivateKey(privateKey)
	if err != nil {
		return
	}
	address = commonAddress.Hex()
	return
}

func GetAddressFromPrivateKey(privateKey *ecdsa.PrivateKey) (address common.Address, err error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		err = errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return
	}

	address = crypto.PubkeyToAddress(*publicKeyECDSA)
	return
}

func GetAddressFromPrivateKey2(pk string) (address common.Address, err error) {
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return
	}

	return GetAddressFromPrivateKey(privateKey)
}

func GetBigFloatBalanceOf(ethCli *ethclient.Client, address common.Address, precision int) (balance *big.Float, err error) {
	balanceBigInt, err := ethCli.BalanceAt(context.Background(), address, nil)
	if err != nil {
		return
	}
	fbalance := new(big.Float)
	fbalance.SetString(balanceBigInt.String())
	balance = new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(precision)))
	return
}

func GigIntToBigFloatEth(number *big.Int, precision int) (balance *big.Float) {
	fbalance := new(big.Float)
	fbalance.SetString(number.String())
	balance = new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(precision)))
	return
}

func WaitForPendingState(ethCli *ethclient.Client, txHash common.Hash, waitTimeSec int) (txReceipt *types.Receipt, err error) {
	isPending := true
	for isPending {
		_, isPending, _ = ethCli.TransactionByHash(context.Background(), txHash)
		time.Sleep(1000 * time.Millisecond)
	}
	timeCounter := 0
	sleepTime := 100 //100毫秒
	if waitTimeSec == 0 {
		waitTimeSec = 36 //最长等待36秒
	}
	totalTimeCounter := waitTimeSec * 1000 / sleepTime
	for {
		txReceipt, err = ethCli.TransactionReceipt(context.Background(), txHash)
		if err == nil {
			return
		} else if timeCounter < totalTimeCounter {
			timeCounter += 1
			time.Sleep(time.Duration(sleepTime) * time.Millisecond)
		} else {
			return
		}
	}
}

func InitEthClient(url string) (ethClient *ethclient.Client, rpcClient *rpc.Client, err error) {
	//ethClient, err = ethclient.Dial("wss://bsc-ws-node.nariox.org:443")//mainnet
	ethClient, err = ethclient.Dial(url)
	if err != nil {
		return
	}
	v := reflect.ValueOf(ethClient).Elem()
	f := v.FieldByName("c")
	rf := reflect.NewAt(f.Type(), unsafe.Pointer(f.UnsafeAddr())).Elem()
	rpcClient, _ = rf.Interface().(*rpc.Client)
	return
}

func GetAuth(ethCli *ethclient.Client, privateKey *ecdsa.PrivateKey, chainID *big.Int) (auth *bind.TransactOpts, err error) {
	fromAddress, err := GetAddressFromPrivateKey(privateKey)
	if err != nil {
		return auth, err
	}
	nonce, err := ethCli.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return auth, err
	}
	gasPrice, err := ethCli.SuggestGasPrice(context.Background())
	if err != nil {
		return auth, err
	}
	if chainID == nil {
		auth = bind.NewKeyedTransactor(privateKey)
	} else {
		auth, err = bind.NewKeyedTransactorWithChainID(privateKey, chainID)
		if err != nil {
			return auth, err
		}
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = uint64(0)  // in units
	auth.GasPrice = gasPrice
	return auth, err
}

func GetAuth2(ethCli *ethclient.Client, pk string, chainID *big.Int) (auth *bind.TransactOpts, err error) {
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return auth, err
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

func SendEth(ethCli *ethclient.Client, pk *ecdsa.PrivateKey, to common.Address, value *big.Int, chainID *big.Int) (*types.Transaction, error) {

	auth, err := GetAuth(ethCli, pk, chainID)
	if err != nil {
		return nil, err
	}
	auth.GasLimit = uint64(21000) // in units
	var data []byte
	tx := types.NewTransaction(auth.Nonce.Uint64(), to, value, auth.GasLimit, auth.GasPrice, data)

	if chainID == nil {
		chainID, err = ethCli.ChainID(context.Background())
		if err != nil {
			return nil, err
		}
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), pk)
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

	if chainID == nil {
		chainID, err = ethCli.ChainID(context.Background())
		if err != nil {
			return nil, err
		}
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), pk)
	if err != nil {
		return nil, err
	}

	err = ethCli.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, err
	}
	return signedTx, nil
}
