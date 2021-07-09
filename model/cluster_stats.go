// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"encoding/json"
	"io"
)

type ClusterStats struct {
	ID                        string `json:"id"`
	TotalWebsocketConnections int    `json:"total_websocket_connections"`
	TotalReadDbConnections    int    `json:"total_read_db_connections"`
	TotalMasterDbConnections  int    `json:"total_master_db_connections"`
}

func (cs *ClusterStats) ToJSON() string {
	b, _ := json.Marshal(cs)
	return string(b)
}

func ClusterStatsFromJSON(data io.Reader) *ClusterStats {
	var cs *ClusterStats
	json.NewDecoder(data).Decode(&cs)
	return cs
}
