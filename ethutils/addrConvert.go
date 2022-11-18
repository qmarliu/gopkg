package ethutils

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

// 调用这个方法前，确保在其它地址执行了如下代码
// const (
// 	// Bech32Prefix defines the Bech32 prefix used for EthAccounts
// 	Bech32Prefix = "gx" //其它值如: cosmos ethm等

//	// Bech32PrefixAccAddr defines the Bech32 prefix of an account's address
//	Bech32PrefixAccAddr = Bech32Prefix
//	// Bech32PrefixAccPub defines the Bech32 prefix of an account's public key
//	Bech32PrefixAccPub = Bech32Prefix + sdk.PrefixPublic
//	// Bech32PrefixValAddr defines the Bech32 prefix of a validator's operator address
//	Bech32PrefixValAddr = Bech32Prefix + sdk.PrefixValidator + sdk.PrefixOperator
//	// Bech32PrefixValPub defines the Bech32 prefix of a validator's operator public key
//	Bech32PrefixValPub = Bech32Prefix + sdk.PrefixValidator + sdk.PrefixOperator + sdk.PrefixPublic
//	// Bech32PrefixConsAddr defines the Bech32 prefix of a consensus node address
//	Bech32PrefixConsAddr = Bech32Prefix + sdk.PrefixValidator + sdk.PrefixConsensus
//	// Bech32PrefixConsPub defines the Bech32 prefix of a consensus node public key
//	Bech32PrefixConsPub = Bech32Prefix + sdk.PrefixValidator + sdk.PrefixConsensus + sdk.PrefixPublic
//
// )
// cfg := sdk.GetConfig()
// cfg.SetBech32PrefixForAccount(Bech32PrefixAccAddr, Bech32PrefixAccPub)
// cfg.SetBech32PrefixForValidator(Bech32PrefixValAddr, Bech32PrefixValPub)
// cfg.SetBech32PrefixForConsensusNode(Bech32PrefixConsAddr, Bech32PrefixConsPub)
func EIP55Bech32Convert(address string) (common.Address, sdk.AccAddress, sdk.ValAddress, error) {
	cfg := sdk.GetConfig()
	var addr []byte
	var err error
	switch {
	case common.IsHexAddress(address):
		addr = common.HexToAddress(address).Bytes()
	case strings.HasPrefix(address, cfg.GetBech32ValidatorAddrPrefix()):
		addr, err = sdk.ValAddressFromBech32(address)
		if err != nil {
			return common.Address{}, nil, nil, err
		}
	case strings.HasPrefix(address, cfg.GetBech32AccountAddrPrefix()):
		addr, err = sdk.AccAddressFromBech32(address)
		if err != nil {
			return common.Address{}, nil, nil, err
		}
	default:
		return common.Address{}, nil, nil,
			fmt.Errorf("expected a valid hex or bech32 address (acc prefix %s), got '%s'\n", cfg.GetBech32AccountAddrPrefix(), address)
	}
	// fmt.Println("Address bytes: ", addr)
	// fmt.Printf("Address (EIP-55): %s\n", common.BytesToAddress(addr))
	// fmt.Printf("Bech32 Acc: %s\n", sdk.AccAddress(addr))
	// fmt.Printf("Bech32 Val: %s\n", sdk.ValAddress(addr))
	return common.BytesToAddress(addr), sdk.AccAddress(addr), sdk.ValAddress(addr), nil
}
