// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTeamSearchJSON(t *testing.T) {
	teamSearch := TeamSearch{Term: NewID()}
	json := teamSearch.ToJSON()
	rteamSearch := ChannelSearchFromJSON(strings.NewReader(json))

	assert.Equal(t, teamSearch.Term, rteamSearch.Term, "Terms do not match")
}
