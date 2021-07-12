// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"regexp"
	"time"
)

type CompliancePost struct {

	// From Team
	TeamName        string
	TeamDisplayName string

	// From Channel
	ChannelName        string
	ChannelDisplayName string
	ChannelType        string

	// From User
	UserUsername string
	UserEmail    string
	UserNickname string

	// From Post
	PostID         string
	PostCreateAt   int64
	PostUpdateAt   int64
	PostDeleteAt   int64
	PostRootID     string
	PostParentID   string
	PostOriginalID string
	PostMessage    string
	PostType       string
	PostProps      string
	PostHashtags   string
	PostFileIDs    string

	IsBot bool
}

func CompliancePostHeader() []string {
	return []string{
		"TeamName",
		"TeamDisplayName",

		"ChannelName",
		"ChannelDisplayName",
		"ChannelType",

		"UserUsername",
		"UserEmail",
		"UserNickname",
		"UserType",

		"PostId",
		"PostCreateAt",
		"PostUpdateAt",
		"PostDeleteAt",
		"PostRootId",
		"PostParentId",
		"PostOriginalId",
		"PostMessage",
		"PostType",
		"PostProps",
		"PostHashtags",
		"PostFileIds",
	}
}

func cleanComplianceStrings(in string) string {
	if matched, _ := regexp.MatchString("^\\s*(=|\\+|\\-)", in); matched {
		return "'" + in
	}
	return in
}

func (cp *CompliancePost) Row() []string {

	postDeleteAt := ""
	if cp.PostDeleteAt > 0 {
		postDeleteAt = time.Unix(0, cp.PostDeleteAt*int64(1000*1000)).Format(time.RFC3339)
	}

	postUpdateAt := ""
	if cp.PostUpdateAt != cp.PostCreateAt {
		postUpdateAt = time.Unix(0, cp.PostUpdateAt*int64(1000*1000)).Format(time.RFC3339)
	}

	userType := "user"
	if cp.IsBot {
		userType = "bot"
	}

	return []string{
		cleanComplianceStrings(cp.TeamName),
		cleanComplianceStrings(cp.TeamDisplayName),

		cleanComplianceStrings(cp.ChannelName),
		cleanComplianceStrings(cp.ChannelDisplayName),
		cleanComplianceStrings(cp.ChannelType),

		cleanComplianceStrings(cp.UserUsername),
		cleanComplianceStrings(cp.UserEmail),
		cleanComplianceStrings(cp.UserNickname),
		userType,

		cp.PostID,
		time.Unix(0, cp.PostCreateAt*int64(1000*1000)).Format(time.RFC3339),
		postUpdateAt,
		postDeleteAt,

		cp.PostRootID,
		cp.PostParentID,
		cp.PostOriginalID,
		cleanComplianceStrings(cp.PostMessage),
		cp.PostType,
		cp.PostProps,
		cp.PostHashtags,
		cp.PostFileIDs,
	}
}
