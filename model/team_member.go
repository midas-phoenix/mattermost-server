// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	USERNAME = "Username"
)

//msgp:tuple TeamMember
// This struct's serializer methods are auto-generated. If a new field is added/removed,
// please run make gen-serialized.
type TeamMember struct {
	TeamID        string `json:"team_id"`
	UserID        string `json:"user_id"`
	Roles         string `json:"roles"`
	DeleteAt      int64  `json:"delete_at"`
	SchemeGuest   bool   `json:"scheme_guest"`
	SchemeUser    bool   `json:"scheme_user"`
	SchemeAdmin   bool   `json:"scheme_admin"`
	ExplicitRoles string `json:"explicit_roles"`
}

//msgp:ignore TeamUnread
type TeamUnread struct {
	TeamID           string `json:"team_id"`
	MsgCount         int64  `json:"msg_count"`
	MentionCount     int64  `json:"mention_count"`
	MentionCountRoot int64  `json:"mention_count_root"`
	MsgCountRoot     int64  `json:"msg_count_root"`
}

//msgp:ignore TeamMemberForExport
type TeamMemberForExport struct {
	TeamMember
	TeamName string
}

//msgp:ignore TeamMemberWithError
type TeamMemberWithError struct {
	UserID string      `json:"user_id"`
	Member *TeamMember `json:"member"`
	Error  *AppError   `json:"error"`
}

//msgp:ignore EmailInviteWithError
type EmailInviteWithError struct {
	Email string    `json:"email"`
	Error *AppError `json:"error"`
}

//msgp:ignore TeamMembersGetOptions
type TeamMembersGetOptions struct {
	// Sort the team members. Accepts "Username", but defaults to "Id".
	Sort string

	// If true, exclude team members whose corresponding user is deleted.
	ExcludeDeletedUsers bool

	// Restrict to search in a list of teams and channels
	ViewRestrictions *ViewUsersRestrictions
}

func (o *TeamMember) ToJSON() string {
	b, _ := json.Marshal(o)
	return string(b)
}

func (o *TeamUnread) ToJSON() string {
	b, _ := json.Marshal(o)
	return string(b)
}

func TeamMemberFromJSON(data io.Reader) *TeamMember {
	var o *TeamMember
	json.NewDecoder(data).Decode(&o)
	return o
}

func TeamUnreadFromJSON(data io.Reader) *TeamUnread {
	var o *TeamUnread
	json.NewDecoder(data).Decode(&o)
	return o
}

func EmailInviteWithErrorFromJSON(data io.Reader) []*EmailInviteWithError {
	var o []*EmailInviteWithError
	json.NewDecoder(data).Decode(&o)
	return o
}

func EmailInviteWithErrorToEmails(o []*EmailInviteWithError) []string {
	var ret []string
	for _, o := range o {
		if o.Error == nil {
			ret = append(ret, o.Email)
		}
	}
	return ret
}

func EmailInviteWithErrorToJSON(o []*EmailInviteWithError) string {
	b, err := json.Marshal(o)
	if err != nil {
		return "[]"
	}
	return string(b)
}

func EmailInviteWithErrorToString(o *EmailInviteWithError) string {
	return fmt.Sprintf("%s:%s", o.Email, o.Error.Error())
}

func TeamMembersWithErrorToTeamMembers(o []*TeamMemberWithError) []*TeamMember {
	var ret []*TeamMember
	for _, o := range o {
		if o.Error == nil {
			ret = append(ret, o.Member)
		}
	}
	return ret
}

func TeamMembersWithErrorToJSON(o []*TeamMemberWithError) string {
	b, err := json.Marshal(o)
	if err != nil {
		return "[]"
	}
	return string(b)
}

func TeamMemberWithErrorToString(o *TeamMemberWithError) string {
	return fmt.Sprintf("%s:%s", o.UserID, o.Error.Error())
}

func TeamMembersWithErrorFromJSON(data io.Reader) []*TeamMemberWithError {
	var o []*TeamMemberWithError
	json.NewDecoder(data).Decode(&o)
	return o
}

func TeamMembersToJSON(o []*TeamMember) string {
	b, err := json.Marshal(o)
	if err != nil {
		return "[]"
	}
	return string(b)
}

func TeamMembersFromJSON(data io.Reader) []*TeamMember {
	var o []*TeamMember
	json.NewDecoder(data).Decode(&o)
	return o
}

func TeamsUnreadToJSON(o []*TeamUnread) string {
	b, err := json.Marshal(o)
	if err != nil {
		return "[]"
	}
	return string(b)
}

func TeamsUnreadFromJSON(data io.Reader) []*TeamUnread {
	var o []*TeamUnread
	json.NewDecoder(data).Decode(&o)
	return o
}

func (o *TeamMember) IsValid() *AppError {

	if !IsValidID(o.TeamID) {
		return NewAppError("TeamMember.IsValid", "model.team_member.is_valid.team_id.app_error", nil, "", http.StatusBadRequest)
	}

	if !IsValidID(o.UserID) {
		return NewAppError("TeamMember.IsValid", "model.team_member.is_valid.user_id.app_error", nil, "", http.StatusBadRequest)
	}

	return nil
}

func (o *TeamMember) PreUpdate() {
}

func (o *TeamMember) GetRoles() []string {
	return strings.Fields(o.Roles)
}
