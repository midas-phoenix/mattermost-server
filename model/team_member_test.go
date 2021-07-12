// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTeamMemberJson(t *testing.T) {
	o := TeamMember{TeamID: NewID(), UserID: NewID()}
	json := o.ToJson()
	ro := TeamMemberFromJson(strings.NewReader(json))

	require.Equal(t, o.TeamID, ro.TeamID, "Ids do not match")
}

func TestTeamMemberIsValid(t *testing.T) {
	o := TeamMember{}

	require.NotNil(t, o.IsValid(), "should be invalid")

	o.TeamID = NewID()

	require.NotNil(t, o.IsValid(), "should be invalid")
}

func TestUnreadMemberJson(t *testing.T) {
	o := TeamUnread{TeamID: NewID(), MsgCount: 5, MentionCount: 3}
	json := o.ToJson()

	r := TeamUnreadFromJson(strings.NewReader(json))

	require.Equal(t, o.TeamID, r.TeamID, "Ids do not match")

	require.Equal(t, o.MsgCount, r.MsgCount, "MsgCount do not match")
}
