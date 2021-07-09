// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"crypto/rand"
	"encoding/json"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRemoteClusterJSON(t *testing.T) {
	o := RemoteCluster{RemoteID: NewID(), Name: "test"}

	json, err := o.ToJSON()
	require.NoError(t, err)

	ro, appErr := RemoteClusterFromJSON(strings.NewReader(json))
	require.Nil(t, appErr)

	require.Equal(t, o.RemoteID, ro.RemoteID)
	require.Equal(t, o.Name, ro.Name)
}

func TestRemoteClusterIsValid(t *testing.T) {
	id := NewID()
	creator := NewID()
	now := GetMillis()
	data := []struct {
		name  string
		rc    *RemoteCluster
		valid bool
	}{
		{name: "Zero value", rc: &RemoteCluster{}, valid: false},
		{name: "Missing cluster_name", rc: &RemoteCluster{RemoteID: id}, valid: false},
		{name: "Missing host_name", rc: &RemoteCluster{RemoteID: id, Name: NewID()}, valid: false},
		{name: "Missing create_at", rc: &RemoteCluster{RemoteID: id, Name: NewID(), SiteURL: "example.com"}, valid: false},
		{name: "Missing last_ping_at", rc: &RemoteCluster{RemoteID: id, Name: NewID(), SiteURL: "example.com", CreatorID: creator, CreateAt: now}, valid: true},
		{name: "Missing creator", rc: &RemoteCluster{RemoteID: id, Name: NewID(), SiteURL: "example.com", CreateAt: now, LastPingAt: now}, valid: false},
		{name: "RemoteCluster valid", rc: &RemoteCluster{RemoteID: id, Name: NewID(), SiteURL: "example.com", CreateAt: now, LastPingAt: now, CreatorID: creator}, valid: true},
		{name: "Include protocol", rc: &RemoteCluster{RemoteID: id, Name: NewID(), SiteURL: "http://example.com", CreateAt: now, LastPingAt: now, CreatorID: creator}, valid: true},
		{name: "Include protocol & port", rc: &RemoteCluster{RemoteID: id, Name: NewID(), SiteURL: "http://example.com:8065", CreateAt: now, LastPingAt: now, CreatorID: creator}, valid: true},
	}

	for _, item := range data {
		err := item.rc.IsValid()
		if item.valid {
			assert.Nil(t, err, item.name)
		} else {
			assert.NotNil(t, err, item.name)
		}
	}
}

func TestRemoteClusterPreSave(t *testing.T) {
	now := GetMillis()

	o := RemoteCluster{RemoteID: NewID(), Name: NewID()}
	o.PreSave()

	require.GreaterOrEqual(t, o.CreateAt, now)
}

func TestRemoteClusterMsgJSON(t *testing.T) {
	o := NewRemoteClusterMsg("shared_channel", []byte("{\"hello\":\"world\"}"))

	json, err := json.Marshal(o)
	require.NoError(t, err)

	ro, appErr := RemoteClusterMsgFromJSON(strings.NewReader(string(json)))
	require.Nil(t, appErr)

	require.Equal(t, o.ID, ro.ID)
	require.Equal(t, o.CreateAt, ro.CreateAt)
	require.Equal(t, o.Topic, ro.Topic)
}

func TestRemoteClusterMsgIsValid(t *testing.T) {
	id := NewID()
	now := GetMillis()
	data := []struct {
		name  string
		msg   *RemoteClusterMsg
		valid bool
	}{
		{name: "Zero value", msg: &RemoteClusterMsg{}, valid: false},
		{name: "Missing remote id", msg: &RemoteClusterMsg{ID: id}, valid: false},
		{name: "Missing Topic", msg: &RemoteClusterMsg{ID: id}, valid: false},
		{name: "Missing Payload", msg: &RemoteClusterMsg{ID: id, CreateAt: now, Topic: "shared_channel"}, valid: false},
		{name: "RemoteClusterMsg valid", msg: &RemoteClusterMsg{ID: id, CreateAt: now, Topic: "shared_channel", Payload: []byte("{\"hello\":\"world\"}")}, valid: true},
	}

	for _, item := range data {
		err := item.msg.IsValid()
		if item.valid {
			assert.Nil(t, err, item.name)
		} else {
			assert.NotNil(t, err, item.name)
		}
	}
}

func TestFixTopics(t *testing.T) {
	testData := []struct {
		topics   string
		expected string
	}{
		{topics: "", expected: ""},
		{topics: "   ", expected: ""},
		{topics: "share", expected: " share "},
		{topics: "share incident", expected: " share incident "},
		{topics: " share incident ", expected: " share incident "},
		{topics: "    share     incident    ", expected: " share incident "},
	}

	for _, tt := range testData {
		rc := &RemoteCluster{Topics: tt.topics}
		rc.fixTopics()
		assert.Equal(t, tt.expected, rc.Topics)
	}
}

func TestRemoteClusterInviteEncryption(t *testing.T) {
	testData := []struct {
		name       string
		badDecrypt bool
		password   string
		invite     RemoteClusterInvite
	}{
		{name: "empty password", badDecrypt: false, password: "", invite: makeInvite("https://example.com:8065")},
		{name: "good password", badDecrypt: false, password: "Ultra secret password!", invite: makeInvite("https://example.com:8065")},
		{name: "bad decrypt", badDecrypt: true, password: "correct horse battery staple", invite: makeInvite("https://example.com:8065")},
	}

	for _, tt := range testData {
		encrypted, err := tt.invite.Encrypt(tt.password)
		require.NoError(t, err)

		invite := RemoteClusterInvite{}
		if tt.badDecrypt {
			buf := make([]byte, len(encrypted))
			_, err = io.ReadFull(rand.Reader, buf)
			assert.NoError(t, err)

			err = invite.Decrypt(buf, tt.password)
			require.Error(t, err)
		} else {
			err = invite.Decrypt(encrypted, tt.password)
			require.NoError(t, err)
			assert.Equal(t, tt.invite, invite)
		}
	}
}

func makeInvite(url string) RemoteClusterInvite {
	return RemoteClusterInvite{
		RemoteID:     NewID(),
		RemoteTeamID: NewID(),
		SiteURL:      url,
		Token:        NewID(),
	}
}
