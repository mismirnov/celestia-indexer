// SPDX-FileCopyrightText: 2023 PK Lab AG <contact@pklab.io>
// SPDX-License-Identifier: MIT

package websocket

import (
	"github.com/celenium-io/celestia-indexer/cmd/api/handler/responses"
)

type Filterable[M INotification] interface {
	Filter(c client, msg Notification[M]) bool
}

type BlockFilter struct{}

func (f BlockFilter) Filter(c client, msg Notification[*responses.Block]) bool {
	if msg.Body == nil {
		return false
	}
	fltrs := c.Filters()
	if fltrs == nil {
		return false
	}
	return fltrs.blocks
}

type HeadFilter struct{}

func (f HeadFilter) Filter(c client, msg Notification[*responses.State]) bool {
	if msg.Body == nil {
		return false
	}
	fltrs := c.Filters()
	if fltrs == nil {
		return false
	}
	return fltrs.head
}

type Filters struct {
	head   bool
	blocks bool
}
