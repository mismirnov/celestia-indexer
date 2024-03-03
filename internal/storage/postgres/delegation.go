// SPDX-FileCopyrightText: 2024 PK Lab AG <contact@pklab.io>
// SPDX-License-Identifier: MIT

package postgres

import (
	"context"

	"github.com/celenium-io/celestia-indexer/internal/storage"
	"github.com/dipdup-net/go-lib/database"
	"github.com/dipdup-net/indexer-sdk/pkg/storage/postgres"
)

// Delegation -
type Delegation struct {
	*postgres.Table[*storage.Delegation]
}

// NewDelegation -
func NewDelegation(db *database.Bun) *Delegation {
	return &Delegation{
		Table: postgres.NewTable[*storage.Delegation](db),
	}
}

func (d *Delegation) ByAddress(ctx context.Context, addressId uint64, limit, offset int) (delegations []storage.Delegation, err error) {
	subQuery := d.DB().NewSelect().Model((*storage.Delegation)(nil)).
		Where("address_id = ?", addressId).
		Order("amount desc")

	subQuery = limitScope(subQuery, limit)
	if offset > 0 {
		subQuery = subQuery.Offset(offset)
	}

	err = d.DB().NewSelect().
		TableExpr("(?) as delegation", subQuery).
		ColumnExpr("delegation.*").
		ColumnExpr("validator.id as validator__id, validator.moniker as validator__moniker, validator.cons_address as validator__cons_address").
		Join("left join validator on validator.id = validator_id").
		Scan(ctx, &delegations)

	return
}

func (d *Delegation) ByValidator(ctx context.Context, validatorId uint64, limit, offset int) (delegations []storage.Delegation, err error) {
	subQuery := d.DB().NewSelect().Model((*storage.Delegation)(nil)).
		Where("validator_id = ?", validatorId).
		Order("amount desc")

	subQuery = limitScope(subQuery, limit)
	if offset > 0 {
		subQuery = subQuery.Offset(offset)
	}

	err = d.DB().NewSelect().
		TableExpr("(?) as delegation", subQuery).
		ColumnExpr("delegation.*").
		ColumnExpr("address.id as address__id, address.address as address__address").
		Join("left join address on address.id = address_id").
		Scan(ctx, &delegations)

	return
}
