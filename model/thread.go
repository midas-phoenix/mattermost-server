// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"encoding/json"
)

type Thread struct {
	PostID       string      `json:"id"`
	ChannelID    string      `json:"channel_id"`
	ReplyCount   int64       `json:"reply_count"`
	LastReplyAt  int64       `json:"last_reply_at"`
	Participants StringArray `json:"participants"`
}

type ThreadResponse struct {
	PostID         string  `json:"id"`
	ReplyCount     int64   `json:"reply_count"`
	LastReplyAt    int64   `json:"last_reply_at"`
	LastViewedAt   int64   `json:"last_viewed_at"`
	Participants   []*User `json:"participants"`
	Post           *Post   `json:"post"`
	UnreadReplies  int64   `json:"unread_replies"`
	UnreadMentions int64   `json:"unread_mentions"`
}

type Threads struct {
	Total               int64             `json:"total"`
	TotalUnreadThreads  int64             `json:"total_unread_threads"`
	TotalUnreadMentions int64             `json:"total_unread_mentions"`
	Threads             []*ThreadResponse `json:"threads"`
}

type GetUserThreadsOpts struct {
	// PageSize specifies the size of the returned chunk of results. Default = 30
	PageSize uint64

	// Extended will enrich the response with participant details. Default = false
	Extended bool

	// Deleted will specify that even deleted threads should be returned (For mobile sync). Default = false
	Deleted bool

	// Since filters the threads based on their LastUpdateAt timestamp.
	Since uint64

	// Before specifies thread id as a cursor for pagination and will return `PageSize` threads before the cursor
	Before string

	// After specifies thread id as a cursor for pagination and will return `PageSize` threads after the cursor
	After string

	// Unread will make sure that only threads with unread replies are returned
	Unread bool
}

func (o *ThreadResponse) ToJSON() string {
	b, _ := json.Marshal(o)
	return string(b)
}

func ThreadResponseFromJSON(s string) (*ThreadResponse, error) {
	var t ThreadResponse
	err := json.Unmarshal([]byte(s), &t)
	return &t, err
}

func (o *Threads) ToJSON() string {
	b, _ := json.Marshal(o)
	return string(b)
}

func (o *Thread) ToJSON() string {
	b, _ := json.Marshal(o)
	return string(b)
}

func ThreadFromJSON(s string) (*Thread, error) {
	var t Thread
	err := json.Unmarshal([]byte(s), &t)
	return &t, err
}

func (o *Thread) Etag() string {
	return Etag(o.PostID, o.LastReplyAt)
}

type ThreadMembership struct {
	PostID         string `json:"post_id"`
	UserID         string `json:"user_id"`
	Following      bool   `json:"following"`
	LastViewed     int64  `json:"last_view_at"`
	LastUpdated    int64  `json:"last_update_at"`
	UnreadMentions int64  `json:"unread_mentions"`
}

func (o *ThreadMembership) ToJSON() string {
	b, _ := json.Marshal(o)
	return string(b)
}
