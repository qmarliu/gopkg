package contracts

import (
	"testing"
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
