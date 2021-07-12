// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

type MessageExport struct {
	TeamID          *string
	TeamName        *string
	TeamDisplayName *string

	ChannelID          *string
	ChannelName        *string
	ChannelDisplayName *string
	ChannelType        *string

	UserID    *string
	UserEmail *string
	Username  *string
	IsBot     bool

	PostID         *string
	PostCreateAt   *int64
	PostUpdateAt   *int64
	PostDeleteAt   *int64
	PostMessage    *string
	PostType       *string
	PostRootID     *string
	PostProps      *string
	PostOriginalID *string
	PostFileIDs    StringArray
}

type MessageExportCursor struct {
	LastPostUpdateAt int64
	LastPostID       string
}
