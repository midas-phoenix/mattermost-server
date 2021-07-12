// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestChannelMemberJson(t *testing.T) {
	o := ChannelMember{ChannelID: NewID(), UserID: NewID()}
	json := o.ToJson()
	ro := ChannelMemberFromJson(strings.NewReader(json))

	require.Equal(t, o.ChannelID, ro.ChannelID, "ids do not match")
}

func TestChannelMemberIsValid(t *testing.T) {
	o := ChannelMember{}

	require.NotNil(t, o.IsValid(), "should be invalid")

	o.ChannelID = NewID()
	require.NotNil(t, o.IsValid(), "should be invalid")

	o.NotifyProps = GetDefaultChannelNotifyProps()
	o.UserID = NewID()

	o.NotifyProps["desktop"] = "junk"
	require.NotNil(t, o.IsValid(), "should be invalid")

	o.NotifyProps["desktop"] = "123456789012345678901"
	require.NotNil(t, o.IsValid(), "should be invalid")

	o.NotifyProps["desktop"] = ChannelNotifyAll
	require.Nil(t, o.IsValid(), "should be valid")

	o.NotifyProps["mark_unread"] = "123456789012345678901"
	require.NotNil(t, o.IsValid(), "should be invalid")

	o.NotifyProps["mark_unread"] = ChannelMarkUnreadAll
	require.Nil(t, o.IsValid(), "should be valid")

	o.Roles = ""
	require.Nil(t, o.IsValid(), "should be invalid")
}

func TestChannelUnreadJson(t *testing.T) {
	o := ChannelUnread{ChannelID: NewID(), TeamID: NewID(), MsgCount: 5, MentionCount: 3}
	json := o.ToJson()
	ro := ChannelUnreadFromJson(strings.NewReader(json))

	require.Equal(t, o.TeamID, ro.TeamID, "team Ids do not match")
	require.Equal(t, o.MentionCount, ro.MentionCount, "mention count do not match")
}
