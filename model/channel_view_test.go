// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChannelViewJSON(t *testing.T) {
	o := ChannelView{ChannelID: NewID(), PrevChannelID: NewID()}
	json := o.ToJSON()
	ro := ChannelViewFromJSON(strings.NewReader(json))

	assert.Equal(t, o.ChannelID, ro.ChannelID, "ChannelIdIds do not match")
	assert.Equal(t, o.PrevChannelID, ro.PrevChannelID, "PrevChannelIds do not match")
}

func TestChannelViewResponseJSON(t *testing.T) {
	id := NewID()
	o := ChannelViewResponse{Status: "OK", LastViewedAtTimes: map[string]int64{id: 12345}}
	json := o.ToJSON()
	ro := ChannelViewResponseFromJSON(strings.NewReader(json))

	assert.Equal(t, o.Status, ro.Status, "ChannelIdIds do not match")
	assert.Equal(t, o.LastViewedAtTimes[id], ro.LastViewedAtTimes[id], "LastViewedAtTimes do not match")
}
