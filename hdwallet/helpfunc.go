package hdwallet

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/qmarliu/gopkg/contracts"
	"github.com/qmarliu/gopkg/contracts/ERC20"
	"github.com/qmarliu/gopkg/contracts/ERC721"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func CollectionETH(ethClient *ethclient.Client, hdid uint32, userAddr, collectionAddr common.Address) (*types.Transaction, *big.Int, error) {
	pk, err := GetHDPvkey(hdid)
	if err != nil {
		return nil, nil, fmt.Errorf(err.Error())
	}

	balance, err := ethClient.BalanceAt(context.Background(), userAddr, nil)
	if err != nil {
		return nil, balance, fmt.Errorf(err.Error())
	}

	tx, err := contracts.SendEthSubFee(ethClient, pk, collectionAddr, balance)
	if err != nil {
		return nil, balance, fmt.Errorf(err.Error())
	}
	return tx, balance, nil
}

func CollectionERC20(ethClient *ethclient.Client, contractAddr common.Address,
	hdid uint32, userAddr, collectionAddr common.Address, gasLimit *big.Int) (*types.Transaction, *big.Int, bool, error) {
	transferDone := false
	pk, err := GetHDPvkey(hdid)
	if err != nil {
		return nil, nil, transferDone, fmt.Errorf(err.Error())
	}

	balance, err := ethClient.BalanceAt(context.Background(), userAddr, nil)
	if err != nil {
		return nil, nil, transferDone, fmt.Errorf(err.Error())
	}

	tokenInstance, err := ERC20.NewERC20(contractAddr, ethClient)
	if err != nil {
		return nil, nil, transferDone, fmt.Errorf(err.Error())
	}
	coinBalance, err := tokenInstance.BalanceOf(nil, userAddr)
	if err != nil {
		return nil, nil, transferDone, fmt.Errorf(err.Error())
	}
	if coinBalance.Cmp(big.NewInt(0)) <= 0 {
		return nil, coinBalance, transferDone, fmt.Errorf("addr:" + userAddr.String() +
			"'s'" + contractAddr.String() + "(contract addr) balance is 0")
	}

	gasPrice, err := ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, coinBalance, transferDone, err
	}
	gasValue := new(big.Int)
	gasValue.Mul(gasPrice, gasLimit)
	if balance.Cmp(gasValue) < 0 {
		ethSendTx, err := contracts.SendEth(ethClient, MasterPk, userAddr, gasValue.Sub(gasValue, balance))
		if err != nil {
			return nil, coinBalance, transferDone, fmt.Errorf(err.Error())
		}
		// log.Infof("发送手续费: %v 数量: %v", ethSendTx.Hash().String(), gasValue)
		return ethSendTx, coinBalance, transferDone, nil //等待gas费到账
	}
	auth, err := contracts.GetAuth(ethClient, pk)
	if err != nil {
		return nil, coinBalance, transferDone, fmt.Errorf("GetAuth failed:" + err.Error())
	}
	tx, err := tokenInstance.Transfer(auth, collectionAddr, coinBalance)
	if err != nil {
		if !strings.Contains(err.Error(), "insufficient funds") {
			return nil, coinBalance, transferDone, fmt.Errorf(err.Error())
		}
		gasValue.Mul(gasPrice, gasLimit)
		ethSendTx, err := contracts.SendEth(ethClient, MasterPk, userAddr, gasValue)
		if err != nil {
			return nil, coinBalance, transferDone, fmt.Errorf(err.Error())
		}
		// log.Infof("发送手续费: %v 数量: %v", ethSendTx.Hash().String(), gasValue)
		return ethSendTx, gasValue, transferDone, nil //等待gas费到账
	}
	transferDone = true
	return tx, coinBalance, transferDone, nil
}

func CollectionERC721(ethClient *ethclient.Client, contractAddr common.Address, hdid uint32, userAddr, collectionAddr common.Address, gasLimit *big.Int, tokenID *big.Int) (*types.Transaction, bool, error) {
	transferDone := false
	pk, err := GetHDPvkey(hdid)
	if err != nil {
		return nil, transferDone, fmt.Errorf(err.Error())
	}

	balance, err := ethClient.BalanceAt(context.Background(), userAddr, nil)
	if err != nil {
		return nil, transferDone, fmt.Errorf(err.Error())
	}

	tokenInstance, err := ERC721.NewERC721(contractAddr, ethClient)
	if err != nil {
		return nil, transferDone, fmt.Errorf(err.Error())
	}
	ownerAddr, err := tokenInstance.OwnerOf(nil, tokenID)
	if err != nil {
		return nil, transferDone, fmt.Errorf(err.Error())
	}
	if ownerAddr.String() != userAddr.String() {
		approveAddr, err := tokenInstance.GetApproved(nil, tokenID)
		if err != nil {
			return nil, transferDone, fmt.Errorf(err.Error())
		}
		if approveAddr.String() != userAddr.String() {
			isOperator, err := tokenInstance.IsApprovedForAll(nil, ownerAddr, userAddr)
			if err != nil {
				return nil, transferDone, fmt.Errorf("IsApprovedForAll query failed:" + err.Error())
			}
			if !isOperator {
				return nil, transferDone, fmt.Errorf("addr:" + userAddr.String() +
					";" + contractAddr.String() + " ID " + tokenID.String() + " owner is " +
					ownerAddr.String() + ";approve addr:" + approveAddr.String())
			}
		}
	}
	gasPrice, err := ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, transferDone, err
	}
	gasValue := new(big.Int)
	gasValue.Mul(gasPrice, gasLimit)
	if balance.Cmp(gasValue) < 0 {
		//转transFee
		ethSendTx, err := contracts.SendEth(ethClient, MasterPk, userAddr, gasValue.Sub(gasValue, balance))
		if err != nil {
			return nil, transferDone, fmt.Errorf("transfer:" + err.Error())
		}
		return ethSendTx, transferDone, nil //等待gas费到账
	}
	auth, err := contracts.GetAuth(ethClient, pk)
	if err != nil {
		return nil, transferDone, fmt.Errorf("GetAuth failed:" + err.Error())
	}
	tx, err := tokenInstance.TransferFrom(auth, ownerAddr, collectionAddr, tokenID)
	if err != nil {
		return nil, transferDone, fmt.Errorf("transfer:" + err.Error())
	}
	transferDone = true
	return tx, transferDone, nil
}
