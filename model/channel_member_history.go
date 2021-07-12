// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

type ChannelMemberHistory struct {
	ChannelID string
	UserID    string
	JoinTime  int64
	LeaveTime *int64
}
