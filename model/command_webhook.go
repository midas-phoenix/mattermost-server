// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"net/http"
)

type CommandWebhook struct {
	ID        string
	CreateAt  int64
	CommandID string
	UserID    string
	ChannelID string
	RootID    string
	ParentID  string
	UseCount  int
}

const (
	CommandWebhookLifetime = 1000 * 60 * 30
)

func (o *CommandWebhook) PreSave() {
	if o.ID == "" {
		o.ID = NewID()
	}

	if o.CreateAt == 0 {
		o.CreateAt = GetMillis()
	}
}

func (o *CommandWebhook) IsValid() *AppError {
	if !IsValidID(o.ID) {
		return NewAppError("CommandWebhook.IsValid", "model.command_hook.id.app_error", nil, "", http.StatusBadRequest)
	}

	if o.CreateAt == 0 {
		return NewAppError("CommandWebhook.IsValid", "model.command_hook.create_at.app_error", nil, "id="+o.ID, http.StatusBadRequest)
	}

	if !IsValidID(o.CommandID) {
		return NewAppError("CommandWebhook.IsValid", "model.command_hook.command_id.app_error", nil, "", http.StatusBadRequest)
	}

	if !IsValidID(o.UserID) {
		return NewAppError("CommandWebhook.IsValid", "model.command_hook.user_id.app_error", nil, "", http.StatusBadRequest)
	}

	if !IsValidID(o.ChannelID) {
		return NewAppError("CommandWebhook.IsValid", "model.command_hook.channel_id.app_error", nil, "", http.StatusBadRequest)
	}

	if o.RootID != "" && !IsValidID(o.RootID) {
		return NewAppError("CommandWebhook.IsValid", "model.command_hook.root_id.app_error", nil, "", http.StatusBadRequest)
	}

	if o.ParentID != "" && !IsValidID(o.ParentID) {
		return NewAppError("CommandWebhook.IsValid", "model.command_hook.parent_id.app_error", nil, "", http.StatusBadRequest)
	}

	return nil
}
