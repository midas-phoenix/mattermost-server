// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"encoding/json"
	"io"
)

type SecurityBulletin struct {
	ID               string `json:"id"`
	AppliesToVersion string `json:"applies_to_version"`
}

type SecurityBulletins []SecurityBulletin

func (sb *SecurityBulletin) ToJSON() string {
	b, _ := json.Marshal(sb)
	return string(b)
}

func SecurityBulletinFromJSON(data io.Reader) *SecurityBulletin {
	var o *SecurityBulletin
	json.NewDecoder(data).Decode(&o)
	return o
}

func (sb SecurityBulletins) ToJSON() string {
	b, err := json.Marshal(sb)
	if err != nil {
		return "[]"
	}
	return string(b)
}

func SecurityBulletinsFromJSON(data io.Reader) SecurityBulletins {
	var o SecurityBulletins
	json.NewDecoder(data).Decode(&o)
	return o
}
