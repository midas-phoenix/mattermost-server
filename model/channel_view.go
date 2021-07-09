// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"encoding/json"
	"io"
)

type ChannelView struct {
	ChannelID                 string `json:"channel_id"`
	PrevChannelID             string `json:"prev_channel_id"`
	CollapsedThreadsSupported bool   `json:"collapsed_threads_supported"`
}

func (o *ChannelView) ToJSON() string {
	b, _ := json.Marshal(o)
	return string(b)
}

func ChannelViewFromJSON(data io.Reader) *ChannelView {
	var o *ChannelView
	json.NewDecoder(data).Decode(&o)
	return o
}

type ChannelViewResponse struct {
	Status            string           `json:"status"`
	LastViewedAtTimes map[string]int64 `json:"last_viewed_at_times"`
}

func (o *ChannelViewResponse) ToJSON() string {
	b, _ := json.Marshal(o)
	return string(b)
}

func ChannelViewResponseFromJSON(data io.Reader) *ChannelViewResponse {
	var o *ChannelViewResponse
	json.NewDecoder(data).Decode(&o)
	return o
}
