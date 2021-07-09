// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"encoding/json"
	"io"
)

type SuggestCommand struct {
	Suggestion  string `json:"suggestion"`
	Description string `json:"description"`
}

func (o *SuggestCommand) ToJSON() string {
	b, _ := json.Marshal(o)
	return string(b)
}

func SuggestCommandFromJSON(data io.Reader) *SuggestCommand {
	var o *SuggestCommand
	json.NewDecoder(data).Decode(&o)
	return o
}
