package core

import (
	"fmt"
	"regexp"
)

const (
	WORLD = "world"
)

type Account struct {
	Address  string   `json:"address" example:"users:001"`
	Metadata Metadata `json:"metadata" swaggertype:"object"`
}

func (a Account) copy() Account {
	a.Metadata = a.Metadata.copy()
	return a
}

func NewAccount(address string) Account {
	return Account{
		Address:  address,
		Metadata: Metadata{},
	}
}

type AccountWithVolumes struct {
	Account
	Volumes AssetsVolumes `json:"volumes"`
	// TODO(gfyrag): Remove from domain layer, it can be computed dynamically when needed
	Balances AssetsBalances `json:"balances" example:"COIN:100"`
}

func (v AccountWithVolumes) Copy() AccountWithVolumes {
	v.Account = v.Account.copy()
	v.Volumes = v.Volumes.copy()
	v.Balances = v.Balances.copy()
	return v
}

const accountPattern = "^[a-zA-Z_]+[a-zA-Z0-9_:]*$"

var accountRegexp = regexp.MustCompile(accountPattern)

type AccountAddress string

func (AccountAddress) GetType() Type { return TypeAccount }
func (a AccountAddress) String() string {
	return fmt.Sprintf("@%v", string(a))
}

func ParseAccountAddress(acc AccountAddress) error {
	// TODO: handle properly in ledger v1.10
	if acc == "" {
		return nil
	}
	if !accountRegexp.MatchString(string(acc)) {
		return fmt.Errorf("accounts should respect pattern %s", accountPattern)
	}
	return nil
}
