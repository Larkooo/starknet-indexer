// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/cartridge-gg/starknet-indexer/ent/balance"
	"github.com/cartridge-gg/starknet-indexer/ent/contract"
	"github.com/cartridge-gg/starknet-indexer/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	balanceFields := schema.Balance{}.Fields()
	_ = balanceFields
	// balanceDescBalance is the schema descriptor for balance field.
	balanceDescBalance := balanceFields[1].Descriptor()
	// balance.DefaultBalance holds the default value on creation for the balance field.
	balance.DefaultBalance = balanceDescBalance.Default.(uint64)
	contractFields := schema.Contract{}.Fields()
	_ = contractFields
	// contractDescCreatedAt is the schema descriptor for created_at field.
	contractDescCreatedAt := contractFields[2].Descriptor()
	// contract.DefaultCreatedAt holds the default value on creation for the created_at field.
	contract.DefaultCreatedAt = contractDescCreatedAt.Default.(func() time.Time)
	// contractDescUpdatedAt is the schema descriptor for updated_at field.
	contractDescUpdatedAt := contractFields[3].Descriptor()
	// contract.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	contract.DefaultUpdatedAt = contractDescUpdatedAt.Default.(func() time.Time)
	// contract.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	contract.UpdateDefaultUpdatedAt = contractDescUpdatedAt.UpdateDefault.(func() time.Time)
}
