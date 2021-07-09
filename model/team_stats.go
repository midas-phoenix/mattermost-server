// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"encoding/json"
	"io"
)

type TeamStats struct {
	TeamID            string `json:"team_id"`
	TotalMemberCount  int64  `json:"total_member_count"`
	ActiveMemberCount int64  `json:"active_member_count"`
}

func (o *TeamStats) ToJSON() string {
	b, _ := json.Marshal(o)
	return string(b)
}

func TeamStatsFromJSON(data io.Reader) *TeamStats {
	var o *TeamStats
	json.NewDecoder(data).Decode(&o)
	return o
}
