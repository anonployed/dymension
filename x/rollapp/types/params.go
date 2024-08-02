package types

import (
	"errors"
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	// KeyDisputePeriodInBlocks is store's key for DisputePeriodInBlocks Params
	KeyDisputePeriodInBlocks            = []byte("DisputePeriodInBlocks")
	DefaultDisputePeriodInBlocks uint64 = 3
	// MinDisputePeriodInBlocks is the minimum number of blocks for dispute period
	MinDisputePeriodInBlocks uint64 = 1
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(disputePeriodInBlocks uint64) Params {
	return Params{
		DisputePeriodInBlocks: disputePeriodInBlocks,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(DefaultDisputePeriodInBlocks)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyDisputePeriodInBlocks, &p.DisputePeriodInBlocks, validateDisputePeriodInBlocks),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return validateDisputePeriodInBlocks(p.DisputePeriodInBlocks)
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateDisputePeriodInBlocks validates the DisputePeriodInBlocks param
func validateDisputePeriodInBlocks(v interface{}) error {
	disputePeriodInBlocks, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	if disputePeriodInBlocks < MinDisputePeriodInBlocks {
		return errors.New("dispute period cannot be lower than 1 block")
	}

	return nil
}
