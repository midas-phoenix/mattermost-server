// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSessionDeepCopy(t *testing.T) {
	sessionID := NewID()
	userID := NewID()
	mapKey := "key"
	mapValue := "val"

	session := &Session{ID: sessionID, Props: map[string]string{mapKey: mapValue}, TeamMembers: []*TeamMember{{UserID: userID, TeamID: "someteamId"}}}

	copySession := session.DeepCopy()
	copySession.ID = "changed"
	copySession.Props[mapKey] = "changed"
	copySession.TeamMembers[0].UserID = "changed"

	assert.Equal(t, sessionID, session.ID)
	assert.Equal(t, mapValue, session.Props[mapKey])
	assert.Equal(t, userID, session.TeamMembers[0].UserID)

	session = &Session{ID: sessionID}
	copySession = session.DeepCopy()

	assert.Equal(t, sessionID, copySession.ID)

	session = &Session{TeamMembers: []*TeamMember{}}
	copySession = session.DeepCopy()

	assert.Equal(t, 0, len(copySession.TeamMembers))
}

func TestSessionJSON(t *testing.T) {
	session := Session{}
	session.PreSave()
	json := session.ToJSON()
	rsession := SessionFromJSON(strings.NewReader(json))

	require.Equal(t, rsession.ID, session.ID, "Ids do not match")

	session.Sanitize()

	require.False(t, session.IsExpired(), "Shouldn't expire")

	session.ExpiresAt = GetMillis()
	time.Sleep(10 * time.Millisecond)

	require.True(t, session.IsExpired(), "Should expire")

	session.SetExpireInDays(10)
}

func TestSessionCSRF(t *testing.T) {
	s := Session{}
	token := s.GetCSRF()
	assert.Empty(t, token)

	token = s.GenerateCSRF()
	assert.NotEmpty(t, token)

	token2 := s.GetCSRF()
	assert.NotEmpty(t, token2)
	assert.Equal(t, token, token2)
}

func TestSessionIsOAuthUser(t *testing.T) {
	testCases := []struct {
		Description string
		Session     Session
		isOAuthUser bool
	}{
		{"False on empty props", Session{}, false},
		{"True when key is set to true", Session{Props: StringMap{UserAuthServiceIsOAuth: strconv.FormatBool(true)}}, true},
		{"False when key is set to false", Session{Props: StringMap{UserAuthServiceIsOAuth: strconv.FormatBool(false)}}, false},
		{"Not affected by Session.IsOAuth being true", Session{IsOAuth: true}, false},
		{"Not affected by Session.IsOAuth being false", Session{IsOAuth: false, Props: StringMap{UserAuthServiceIsOAuth: strconv.FormatBool(true)}}, true},
	}

	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {
			require.Equal(t, tc.isOAuthUser, tc.Session.IsOAuthUser())
		})
	}
}
