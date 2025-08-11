package param

import (
	_ "embed"
	"encoding/json"
	"time"

	"github.com/pactus-project/pactus/crypto"
	"github.com/pactus-project/pactus/genesis"
	"github.com/pactus-project/pactus/types/amount"
)

//go:embed foundation_testnet.json
var foundationTestnetBytes []byte

type Params struct {
	BlockVersion              uint8
	BlockIntervalInSecond     int
	MaxTransactionsPerBlock   int
	CommitteeSize             int
	BlockReward               amount.Amount
	TransactionToLiveInterval uint32
	BondInterval              uint32
	UnbondInterval            uint32
	SortitionInterval         uint32
	MinimumStake              amount.Amount
	MaximumStake              amount.Amount
	FoundationReward          amount.Amount
	FoundationAddress         []crypto.Address
	SplitRewardForkHeight     uint32
}

func FromGenesis(genDoc *genesis.Genesis) *Params {
	params := &Params{
		// genesis parameters
		BlockVersion:              genDoc.Params().BlockVersion,
		BlockIntervalInSecond:     genDoc.Params().BlockIntervalInSecond,
		CommitteeSize:             genDoc.Params().CommitteeSize,
		BlockReward:               genDoc.Params().BlockReward,
		TransactionToLiveInterval: genDoc.Params().TransactionToLiveInterval,
		BondInterval:              genDoc.Params().BondInterval,
		UnbondInterval:            genDoc.Params().UnbondInterval,
		SortitionInterval:         genDoc.Params().SortitionInterval,
		MaximumStake:              genDoc.Params().MaximumStake,
		MinimumStake:              genDoc.Params().MinimumStake,

		// chain parameters
		MaxTransactionsPerBlock: 1000,
		FoundationAddress:       []crypto.Address{},
		FoundationReward:        amount.Amount(300_000_000),
		SplitRewardForkHeight:   0,
	}

	foundationAddressList := make([]string, 0)
	switch genDoc.ChainType() {
	case genesis.Mainnet:
		params.SplitRewardForkHeight = 0

	case genesis.Testnet:
		params.SplitRewardForkHeight = 3_680_000
		if err := json.Unmarshal(foundationTestnetBytes, &foundationAddressList); err != nil {
			panic(err)
		}

	case genesis.Localnet:
		params.SplitRewardForkHeight = 5
		if err := json.Unmarshal(foundationTestnetBytes, &foundationAddressList); err != nil {
			panic(err)
		}

	default:
		params.SplitRewardForkHeight = 0
	}

	for _, addrStr := range foundationAddressList {
		addr, err := crypto.AddressFromString(addrStr)
		if err != nil {
			panic(err)
		}
		params.FoundationAddress = append(params.FoundationAddress, addr)
	}

	return params
}

func (p *Params) BlockInterval() time.Duration {
	return time.Duration(p.BlockIntervalInSecond) * time.Second
}
