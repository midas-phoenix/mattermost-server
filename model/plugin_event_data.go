// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"encoding/json"
	"io"
)

// PluginEventData used to notify peers about plugin changes.
type PluginEventData struct {
	ID string `json:"id"`
}

func (p *PluginEventData) ToJSON() string {
	b, _ := json.Marshal(p)
	return string(b)
}

func PluginEventDataFromJSON(data io.Reader) PluginEventData {
	var m PluginEventData
	json.NewDecoder(data).Decode(&m)
	return m
}
