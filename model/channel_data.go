// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"encoding/json"
	"io"
)

type ChannelData struct {
	Channel *Channel       `json:"channel"`
	Member  *ChannelMember `json:"member"`
}

func (o *ChannelData) Etag() string {
	var mt int64 = 0
	if o.Member != nil {
		mt = o.Member.LastUpdateAt
	}

	return Etag(o.Channel.ID, o.Channel.UpdateAt, o.Channel.LastPostAt, mt)
}

func (o *ChannelData) ToJSON() string {
	b, _ := json.Marshal(o)
	return string(b)
}

func ChannelDataFromJSON(data io.Reader) *ChannelData {
	var o *ChannelData
	json.NewDecoder(data).Decode(&o)
	return o
}
