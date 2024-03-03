// SPDX-FileCopyrightText: 2024 PK Lab AG <contact@pklab.io>
// SPDX-License-Identifier: MIT

package responses

import (
	"time"

	"github.com/celenium-io/celestia-indexer/internal/storage"
	pkgTypes "github.com/celenium-io/celestia-indexer/pkg/types"
)

type Delegation struct {
	Delegator string `example:"celestia1jc92qdnty48pafummfr8ava2tjtuhfdw774w60" json:"delegator,omitempty" swaggertype:"string"`
	Amount    string `example:"0.1"                                             json:"amount"              swaggertype:"string"`

	Validator *ShortValidator `json:"validator,omitempty"`
}

func NewDelegation(d storage.Delegation) Delegation {
	delegation := Delegation{
		Amount: d.Amount.String(),
	}

	if d.Address != nil {
		delegation.Delegator = d.Address.Address
	}

	if d.Validator != nil {
		delegation.Validator = NewShortValidator(*d.Validator)
	}

	return delegation
}

type Undelegation struct {
	Height         pkgTypes.Level `example:"100"                                             json:"height"              swaggertype:"integer"`
	Time           time.Time      `example:"2023-07-04T03:10:57+00:00"                       json:"time"                swaggertype:"string"`
	CompletionTime time.Time      `example:"2023-07-04T03:10:57+00:00"                       json:"completion_time"     swaggertype:"string"`
	Delegator      string         `example:"celestia1jc92qdnty48pafummfr8ava2tjtuhfdw774w60" json:"delegator,omitempty" swaggertype:"string"`
	Amount         string         `example:"0.1"                                             json:"amount"              swaggertype:"string"`

	Validator *ShortValidator `json:"validator,omitempty"`
}

func NewUndelegation(d storage.Undelegation) Undelegation {
	undelegation := Undelegation{
		Amount:         d.Amount.String(),
		Time:           d.Time,
		CompletionTime: d.CompletionTime,
		Height:         d.Height,
	}

	if d.Address != nil {
		undelegation.Delegator = d.Address.Address
	}

	if d.Validator != nil {
		undelegation.Validator = NewShortValidator(*d.Validator)
	}

	return undelegation
}

type Redelegation struct {
	Height         pkgTypes.Level `example:"100"                                             json:"height"              swaggertype:"integer"`
	Time           time.Time      `example:"2023-07-04T03:10:57+00:00"                       json:"time"                swaggertype:"string"`
	CompletionTime time.Time      `example:"2023-07-04T03:10:57+00:00"                       json:"completion_time"     swaggertype:"string"`
	Delegator      string         `example:"celestia1jc92qdnty48pafummfr8ava2tjtuhfdw774w60" json:"delegator,omitempty" swaggertype:"string"`
	Amount         string         `example:"0.1"                                             json:"amount"              swaggertype:"string"`

	Source      *ShortValidator `json:"source,omitempty"`
	Destination *ShortValidator `json:"destination,omitempty"`
}

func NewRedelegation(d storage.Redelegation) Redelegation {
	redelegation := Redelegation{
		Amount:         d.Amount.String(),
		Time:           d.Time,
		CompletionTime: d.CompletionTime,
		Height:         d.Height,
	}

	if d.Address != nil {
		redelegation.Delegator = d.Address.Address
	}

	if d.Source != nil {
		redelegation.Source = NewShortValidator(*d.Source)
	}

	if d.Destination != nil {
		redelegation.Destination = NewShortValidator(*d.Destination)
	}

	return redelegation
}
