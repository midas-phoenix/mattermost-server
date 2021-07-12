// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"encoding/json"
	"io"
	"net/http"
	"regexp"
)

type Reaction struct {
	UserID    string  `json:"user_id"`
	PostID    string  `json:"post_id"`
	EmojiName string  `json:"emoji_name"`
	CreateAt  int64   `json:"create_at"`
	UpdateAt  int64   `json:"update_at"`
	DeleteAt  int64   `json:"delete_at"`
	RemoteID  *string `json:"remote_id"`
}

func (o *Reaction) ToJson() string {
	b, _ := json.Marshal(o)
	return string(b)
}

func ReactionFromJson(data io.Reader) *Reaction {
	var o Reaction

	if err := json.NewDecoder(data).Decode(&o); err != nil {
		return nil
	}
	return &o
}

func ReactionsToJson(o []*Reaction) string {
	b, _ := json.Marshal(o)
	return string(b)
}

func MapPostIDToReactionsToJson(o map[string][]*Reaction) string {
	b, _ := json.Marshal(o)
	return string(b)
}

func MapPostIDToReactionsFromJson(data io.Reader) map[string][]*Reaction {
	decoder := json.NewDecoder(data)

	var objmap map[string][]*Reaction
	if err := decoder.Decode(&objmap); err != nil {
		return make(map[string][]*Reaction)
	}
	return objmap
}

func ReactionsFromJson(data io.Reader) []*Reaction {
	var o []*Reaction

	if err := json.NewDecoder(data).Decode(&o); err != nil {
		return nil
	}
	return o
}

func (o *Reaction) IsValid() *AppError {
	if !IsValidID(o.UserID) {
		return NewAppError("Reaction.IsValid", "model.reaction.is_valid.user_id.app_error", nil, "user_id="+o.UserID, http.StatusBadRequest)
	}

	if !IsValidID(o.PostID) {
		return NewAppError("Reaction.IsValid", "model.reaction.is_valid.post_id.app_error", nil, "post_id="+o.PostID, http.StatusBadRequest)
	}

	validName := regexp.MustCompile(`^[a-zA-Z0-9\-\+_]+$`)

	if o.EmojiName == "" || len(o.EmojiName) > EmojiNameMaxLength || !validName.MatchString(o.EmojiName) {
		return NewAppError("Reaction.IsValid", "model.reaction.is_valid.emoji_name.app_error", nil, "emoji_name="+o.EmojiName, http.StatusBadRequest)
	}

	if o.CreateAt == 0 {
		return NewAppError("Reaction.IsValid", "model.reaction.is_valid.create_at.app_error", nil, "", http.StatusBadRequest)
	}

	if o.UpdateAt == 0 {
		return NewAppError("Reaction.IsValid", "model.reaction.is_valid.update_at.app_error", nil, "", http.StatusBadRequest)
	}

	return nil
}

func (o *Reaction) PreSave() {
	if o.CreateAt == 0 {
		o.CreateAt = GetMillis()
	}
	o.UpdateAt = GetMillis()
	o.DeleteAt = 0

	if o.RemoteID == nil {
		o.RemoteID = NewString("")
	}
}

func (o *Reaction) PreUpdate() {
	o.UpdateAt = GetMillis()

	if o.RemoteID == nil {
		o.RemoteID = NewString("")
	}
}
