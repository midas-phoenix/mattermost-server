// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutgoingWebhookJson(t *testing.T) {
	o := OutgoingWebhook{ID: NewID()}
	json := o.ToJson()
	ro := OutgoingWebhookFromJson(strings.NewReader(json))

	assert.Equal(t, o.ID, ro.ID, "Ids do not match")
}

func TestOutgoingWebhookIsValid(t *testing.T) {
	o := OutgoingWebhook{}
	assert.NotNil(t, o.IsValid(), "empty declaration should be invalid")

	o.ID = NewID()
	assert.NotNilf(t, o.IsValid(), "Id = NewId; %s should be invalid", o.ID)

	o.CreateAt = GetMillis()
	assert.NotNilf(t, o.IsValid(), "CreateAt = GetMillis; %d should be invalid", o.CreateAt)

	o.UpdateAt = GetMillis()
	assert.NotNilf(t, o.IsValid(), "UpdateAt = GetMillis; %d should be invalid", o.UpdateAt)

	o.CreatorID = "123"
	assert.NotNilf(t, o.IsValid(), "CreatorId %s should be invalid", o.CreatorID)

	o.CreatorID = NewID()
	assert.NotNilf(t, o.IsValid(), "CreatorId = NewId; %s should be invalid", o.CreatorID)

	o.Token = "123"
	assert.NotNilf(t, o.IsValid(), "Token %s should be invalid", o.Token)

	o.Token = NewID()
	assert.NotNilf(t, o.IsValid(), "Token = NewId; %s should be invalid", o.Token)

	o.ChannelID = "123"
	assert.NotNilf(t, o.IsValid(), "ChannelId %s should be invalid", o.ChannelID)

	o.ChannelID = NewID()
	assert.NotNilf(t, o.IsValid(), "ChannelId = NewId; %s should be invalid", o.ChannelID)

	o.TeamID = "123"
	assert.NotNilf(t, o.IsValid(), "TeamId %s should be invalid", o.TeamID)

	o.TeamID = NewID()
	assert.NotNilf(t, o.IsValid(), "TeamId = NewId; %s should be invalid", o.TeamID)

	o.CallbackURLs = []string{"nowhere.com/"}
	assert.NotNilf(t, o.IsValid(), "%v for CallbackURLs should be invalid", o.CallbackURLs)

	o.CallbackURLs = []string{"http://nowhere.com/"}
	assert.Nilf(t, o.IsValid(), "%v for CallbackURLs should be valid", o.CallbackURLs)

	o.DisplayName = strings.Repeat("1", 65)
	assert.NotNilf(t, o.IsValid(), "DisplayName length %d invalid, max length 64", len(o.DisplayName))

	o.DisplayName = strings.Repeat("1", 64)
	assert.Nilf(t, o.IsValid(), "DisplayName length %d should be valid, max length 64", len(o.DisplayName))

	o.Description = strings.Repeat("1", 501)
	assert.NotNilf(t, o.IsValid(), "Description length %d should be invalid, max length 500", len(o.Description))

	o.Description = strings.Repeat("1", 500)
	assert.Nilf(t, o.IsValid(), "Description length %d should be valid, max length 500", len(o.Description))

	o.ContentType = strings.Repeat("1", 129)
	assert.NotNilf(t, o.IsValid(), "ContentType length %d should be invalid, max length 128", len(o.ContentType))

	o.ContentType = strings.Repeat("1", 128)
	assert.Nilf(t, o.IsValid(), "ContentType length %d should be valid", len(o.ContentType))

	o.Username = strings.Repeat("1", 65)
	assert.NotNilf(t, o.IsValid(), "Username length %d should be invalid, max length 64", len(o.Username))

	o.Username = strings.Repeat("1", 64)
	assert.Nilf(t, o.IsValid(), "Username length %d should be valid", len(o.Username))

	o.IconURL = strings.Repeat("1", 1025)
	assert.NotNilf(t, o.IsValid(), "IconURL length %d should be invalid, max length 1024", len(o.IconURL))

	o.IconURL = strings.Repeat("1", 1024)
	assert.Nilf(t, o.IsValid(), "IconURL length %d should be valid", len(o.IconURL))
}

func TestOutgoingWebhookPayloadToFormValues(t *testing.T) {
	p := &OutgoingWebhookPayload{
		Token:       "Token",
		TeamID:      "TeamId",
		TeamDomain:  "TeamDomain",
		ChannelID:   "ChannelId",
		ChannelName: "ChannelName",
		Timestamp:   123000,
		UserID:      "UserId",
		UserName:    "UserName",
		PostID:      "PostId",
		Text:        "Text",
		TriggerWord: "TriggerWord",
		FileIDs:     "FileIds",
	}
	v := url.Values{}
	v.Set("token", "Token")
	v.Set("team_id", "TeamId")
	v.Set("team_domain", "TeamDomain")
	v.Set("channel_id", "ChannelId")
	v.Set("channel_name", "ChannelName")
	v.Set("timestamp", "123")
	v.Set("user_id", "UserId")
	v.Set("user_name", "UserName")
	v.Set("post_id", "PostId")
	v.Set("text", "Text")
	v.Set("trigger_word", "TriggerWord")
	v.Set("file_ids", "FileIds")
	got := p.ToFormValues()
	want := v.Encode()
	assert.Equalf(t, got, want, "Got %+v, wanted %+v", got, want)
}

func TestOutgoingWebhookPreSave(t *testing.T) {
	o := OutgoingWebhook{}
	o.PreSave()
}

func TestOutgoingWebhookPreUpdate(t *testing.T) {
	o := OutgoingWebhook{}
	o.PreUpdate()
}

func TestOutgoingWebhookTriggerWordStartsWith(t *testing.T) {
	o := OutgoingWebhook{ID: NewID()}
	o.TriggerWords = append(o.TriggerWords, "foo")
	assert.True(t, o.TriggerWordStartsWith("foobar"), "Should return true")
	assert.False(t, o.TriggerWordStartsWith("barfoo"), "Should return false")
}

func TestOutgoingWebhookResponseJson(t *testing.T) {
	o := OutgoingWebhookResponse{}
	o.Text = NewString("some text")

	json := o.ToJson()
	ro, _ := OutgoingWebhookResponseFromJson(strings.NewReader(json))

	assert.Equal(t, *o.Text, *ro.Text, "Text does not match")
}
