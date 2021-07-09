// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChannelSearchJSON(t *testing.T) {
	channelSearch := ChannelSearch{Term: NewID()}
	json := channelSearch.ToJSON()
	rchannelSearch := ChannelSearchFromJSON(strings.NewReader(json))

	assert.Equal(t, channelSearch.Term, rchannelSearch.Term)
}
