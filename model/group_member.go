// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import "net/http"

type GroupMember struct {
	GroupID  string `json:"group_id"`
	UserID   string `json:"user_id"`
	CreateAt int64  `json:"create_at"`
	DeleteAt int64  `json:"delete_at"`
}

func (gm *GroupMember) IsValid() *AppError {
	if !IsValidID(gm.GroupID) {
		return NewAppError("GroupMember.IsValid", "model.group_member.group_id.app_error", nil, "", http.StatusBadRequest)
	}
	if !IsValidID(gm.UserID) {
		return NewAppError("GroupMember.IsValid", "model.group_member.user_id.app_error", nil, "", http.StatusBadRequest)
	}
	return nil
}
