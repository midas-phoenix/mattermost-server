// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"encoding/json"
	"io"
)

const (
	StatusOutOfOffice    = "ooo"
	StatusOffline        = "offline"
	StatusAway           = "away"
	StatusDnd            = "dnd"
	StatusOnline         = "online"
	StatusCacheSize      = SessionCacheSize
	StatusChannelTimeout = 20000  // 20 seconds
	StatusMinUpdateTime  = 120000 // 2 minutes
)

type Status struct {
	UserID         string `json:"user_id"`
	Status         string `json:"status"`
	Manual         bool   `json:"manual"`
	LastActivityAt int64  `json:"last_activity_at"`
	ActiveChannel  string `json:"active_channel,omitempty" db:"-"`
	DNDEndTime     int64  `json:"dnd_end_time"`
	PrevStatus     string `json:"-"`
}

func (o *Status) ToJSON() string {
	oCopy := *o
	oCopy.ActiveChannel = ""
	b, _ := json.Marshal(oCopy)
	return string(b)
}

func (o *Status) ToClusterJSON() string {
	oCopy := *o
	b, _ := json.Marshal(oCopy)
	return string(b)
}

func StatusFromJSON(data io.Reader) *Status {
	var o *Status
	json.NewDecoder(data).Decode(&o)
	return o
}

func StatusListToJSON(u []*Status) string {
	uCopy := make([]Status, len(u))
	for i, s := range u {
		sCopy := *s
		sCopy.ActiveChannel = ""
		uCopy[i] = sCopy
	}

	b, _ := json.Marshal(uCopy)
	return string(b)
}

func StatusListFromJSON(data io.Reader) []*Status {
	var statuses []*Status
	json.NewDecoder(data).Decode(&statuses)
	return statuses
}

func StatusMapToInterfaceMap(statusMap map[string]*Status) map[string]interface{} {
	interfaceMap := map[string]interface{}{}
	for _, s := range statusMap {
		// Omitted statues mean offline
		if s.Status != StatusOffline {
			interfaceMap[s.UserID] = s.Status
		}
	}
	return interfaceMap
}
