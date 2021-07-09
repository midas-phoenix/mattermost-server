// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"encoding/json"
	"io"
)

type MfaSecret struct {
	Secret string `json:"secret"`
	QRCode string `json:"qr_code"`
}

func (mfa *MfaSecret) ToJSON() string {
	b, _ := json.Marshal(mfa)
	return string(b)
}

func MfaSecretFromJSON(data io.Reader) *MfaSecret {
	var mfa *MfaSecret
	json.NewDecoder(data).Decode(&mfa)
	return mfa
}
