// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCommandJSON(t *testing.T) {
	o := Command{ID: NewID()}
	json := o.ToJSON()
	ro := CommandFromJSON(strings.NewReader(json))

	require.Equal(t, o.ID, ro.ID, "Ids do not match")
}

func TestCommandIsValid(t *testing.T) {
	o := Command{
		ID:          NewID(),
		Token:       NewID(),
		CreateAt:    GetMillis(),
		UpdateAt:    GetMillis(),
		CreatorID:   NewID(),
		TeamID:      NewID(),
		Trigger:     "trigger",
		URL:         "http://example.com",
		Method:      CommandMethodGet,
		DisplayName: "",
		Description: "",
	}

	require.Nil(t, o.IsValid())

	o.ID = ""
	require.NotNil(t, o.IsValid(), "should be invalid")

	o.ID = NewID()
	require.Nil(t, o.IsValid())

	o.Token = ""
	require.NotNil(t, o.IsValid(), "should be invalid")

	o.Token = NewID()
	require.Nil(t, o.IsValid())

	o.CreateAt = 0
	require.NotNil(t, o.IsValid(), "should be invalid")

	o.CreateAt = GetMillis()
	require.Nil(t, o.IsValid())

	o.UpdateAt = 0
	require.NotNil(t, o.IsValid(), "should be invalid")

	o.UpdateAt = GetMillis()
	require.Nil(t, o.IsValid())

	o.CreatorID = ""
	require.NotNil(t, o.IsValid(), "should be invalid")

	o.CreatorID = NewID()
	require.Nil(t, o.IsValid())

	o.TeamID = ""
	require.NotNil(t, o.IsValid(), "should be invalid")

	o.TeamID = NewID()
	require.Nil(t, o.IsValid())

	o.Trigger = ""
	require.NotNil(t, o.IsValid(), "should be invalid")

	o.Trigger = strings.Repeat("1", 129)
	require.NotNil(t, o.IsValid(), "should be invalid")

	o.Trigger = strings.Repeat("1", 128)
	require.Nil(t, o.IsValid())

	o.URL = ""
	require.NotNil(t, o.IsValid(), "should be invalid")

	o.URL = "1234"
	require.NotNil(t, o.IsValid(), "should be invalid")

	o.URL = "https:////example.com"
	require.NotNil(t, o.IsValid(), "should be invalid")

	o.URL = "https://example.com"
	require.Nil(t, o.IsValid())

	o.Method = "https://example.com"
	require.NotNil(t, o.IsValid(), "should be invalid")

	o.Method = CommandMethodGet
	require.Nil(t, o.IsValid())

	o.Method = CommandMethodPost
	require.Nil(t, o.IsValid())

	o.DisplayName = strings.Repeat("1", 65)
	require.NotNil(t, o.IsValid(), "should be invalid")

	o.DisplayName = strings.Repeat("1", 64)
	require.Nil(t, o.IsValid())

	o.Description = strings.Repeat("1", 129)
	require.NotNil(t, o.IsValid(), "should be invalid")

	o.Description = strings.Repeat("1", 128)
	require.Nil(t, o.IsValid())
}

func TestCommandPreSave(t *testing.T) {
	o := Command{}
	o.PreSave()
}

func TestCommandPreUpdate(t *testing.T) {
	o := Command{}
	o.PreUpdate()
}
