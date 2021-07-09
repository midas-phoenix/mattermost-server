// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSharedChannelJSON(t *testing.T) {
	o := SharedChannel{ChannelID: NewID(), ShareName: NewID()}
	json := o.ToJSON()
	ro, err := SharedChannelFromJSON(strings.NewReader(json))

	require.NoError(t, err)
	require.Equal(t, o.ChannelID, ro.ChannelID)
	require.Equal(t, o.ShareName, ro.ShareName)
}

func TestSharedChannelIsValid(t *testing.T) {
	id := NewID()
	now := GetMillis()
	data := []struct {
		name  string
		sc    *SharedChannel
		valid bool
	}{
		{name: "Zero value", sc: &SharedChannel{}, valid: false},
		{name: "Missing team_id", sc: &SharedChannel{ChannelID: id}, valid: false},
		{name: "Missing create_at", sc: &SharedChannel{ChannelID: id, TeamID: id}, valid: false},
		{name: "Missing update_at", sc: &SharedChannel{ChannelID: id, TeamID: id, CreateAt: now}, valid: false},
		{name: "Missing share_name", sc: &SharedChannel{ChannelID: id, TeamID: id, CreateAt: now, UpdateAt: now}, valid: false},
		{name: "Invalid share_name", sc: &SharedChannel{ChannelID: id, TeamID: id, CreateAt: now, UpdateAt: now,
			ShareName: "@test@"}, valid: false},
		{name: "Too long share_name", sc: &SharedChannel{ChannelID: id, TeamID: id, CreateAt: now, UpdateAt: now,
			ShareName: strings.Repeat("01234567890", 100)}, valid: false},
		{name: "Missing creator_id", sc: &SharedChannel{ChannelID: id, TeamID: id, CreateAt: now, UpdateAt: now,
			ShareName: "test"}, valid: false},
		{name: "Missing remote_id", sc: &SharedChannel{ChannelID: id, TeamID: id, CreateAt: now, UpdateAt: now,
			ShareName: "test", CreatorID: id}, valid: false},
		{name: "Valid shared channel", sc: &SharedChannel{ChannelID: id, TeamID: id, CreateAt: now, UpdateAt: now,
			ShareName: "test", CreatorID: id, RemoteID: id}, valid: true},
	}

	for _, item := range data {
		err := item.sc.IsValid()
		if item.valid {
			assert.Nil(t, err, item.name)
		} else {
			assert.NotNil(t, err, item.name)
		}
	}
}

func TestSharedChannelPreSave(t *testing.T) {
	now := GetMillis()

	o := SharedChannel{ChannelID: NewID(), ShareName: "test"}
	o.PreSave()

	require.GreaterOrEqual(t, o.CreateAt, now)
	require.GreaterOrEqual(t, o.UpdateAt, now)
}

func TestSharedChannelPreUpdate(t *testing.T) {
	now := GetMillis()

	o := SharedChannel{ChannelID: NewID(), ShareName: "test"}
	o.PreUpdate()

	require.GreaterOrEqual(t, o.UpdateAt, now)
}

func TestSharedChannelRemoteJSON(t *testing.T) {
	o := SharedChannelRemote{ID: NewID(), ChannelID: NewID()}
	json := o.ToJSON()
	ro, err := SharedChannelRemoteFromJSON(strings.NewReader(json))

	require.NoError(t, err)
	require.Equal(t, o.ID, ro.ID)
	require.Equal(t, o.ChannelID, ro.ChannelID)
}
