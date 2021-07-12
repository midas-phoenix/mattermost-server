// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"encoding/json"
	"io"
)

type Audit struct {
	ID        string `json:"id"`
	CreateAt  int64  `json:"create_at"`
	UserID    string `json:"user_id"`
	Action    string `json:"action"`
	ExtraInfo string `json:"extra_info"`
	IpAddress string `json:"ip_address"`
	SessionID string `json:"session_id"`
}

func (o *Audit) ToJson() string {
	b, _ := json.Marshal(o)
	return string(b)
}

func AuditFromJson(data io.Reader) *Audit {
	var o *Audit
	json.NewDecoder(data).Decode(&o)
	return o
}
