// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"encoding/json"
	"io"
	"net/http"
)

type UserAccessToken struct {
	ID          string `json:"id"`
	Token       string `json:"token,omitempty"`
	UserID      string `json:"user_id"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}

func (t *UserAccessToken) IsValid() *AppError {
	if !IsValidID(t.ID) {
		return NewAppError("UserAccessToken.IsValid", "model.user_access_token.is_valid.id.app_error", nil, "", http.StatusBadRequest)
	}

	if len(t.Token) != 26 {
		return NewAppError("UserAccessToken.IsValid", "model.user_access_token.is_valid.token.app_error", nil, "", http.StatusBadRequest)
	}

	if !IsValidID(t.UserID) {
		return NewAppError("UserAccessToken.IsValid", "model.user_access_token.is_valid.user_id.app_error", nil, "", http.StatusBadRequest)
	}

	if len(t.Description) > 255 {
		return NewAppError("UserAccessToken.IsValid", "model.user_access_token.is_valid.description.app_error", nil, "", http.StatusBadRequest)
	}

	return nil
}

func (t *UserAccessToken) PreSave() {
	t.ID = NewID()
	t.IsActive = true
}

func (t *UserAccessToken) ToJson() string {
	b, _ := json.Marshal(t)
	return string(b)
}

func UserAccessTokenFromJson(data io.Reader) *UserAccessToken {
	var t *UserAccessToken
	json.NewDecoder(data).Decode(&t)
	return t
}

func UserAccessTokenListToJson(t []*UserAccessToken) string {
	b, _ := json.Marshal(t)
	return string(b)
}

func UserAccessTokenListFromJson(data io.Reader) []*UserAccessToken {
	var t []*UserAccessToken
	json.NewDecoder(data).Decode(&t)
	return t
}
