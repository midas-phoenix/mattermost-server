// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestChannelJson(t *testing.T) {
	o := Channel{ID: NewID(), Name: NewID()}
	json := o.ToJson()
	ro := ChannelFromJson(strings.NewReader(json))

	require.Equal(t, o.ID, ro.ID)

	p := ChannelPatch{Name: new(string)}
	*p.Name = NewID()
	json = p.ToJson()
	rp := ChannelPatchFromJson(strings.NewReader(json))

	require.Equal(t, *p.Name, *rp.Name)
}

func TestChannelCopy(t *testing.T) {
	o := Channel{ID: NewID(), Name: NewID()}
	ro := o.DeepCopy()

	require.Equal(t, o.ID, ro.ID, "Ids do not match")
}

func TestChannelPatch(t *testing.T) {
	p := &ChannelPatch{Name: new(string), DisplayName: new(string), Header: new(string), Purpose: new(string), GroupConstrained: new(bool)}
	*p.Name = NewID()
	*p.DisplayName = NewID()
	*p.Header = NewID()
	*p.Purpose = NewID()
	*p.GroupConstrained = true

	o := Channel{ID: NewID(), Name: NewID()}
	o.Patch(p)

	require.Equal(t, *p.Name, o.Name)
	require.Equal(t, *p.DisplayName, o.DisplayName)
	require.Equal(t, *p.Header, o.Header)
	require.Equal(t, *p.Purpose, o.Purpose)
	require.Equal(t, *p.GroupConstrained, *o.GroupConstrained)
}

func TestChannelIsValid(t *testing.T) {
	o := Channel{}

	require.NotNil(t, o.IsValid())

	o.ID = NewID()
	require.NotNil(t, o.IsValid())

	o.CreateAt = GetMillis()
	require.NotNil(t, o.IsValid())

	o.UpdateAt = GetMillis()
	require.NotNil(t, o.IsValid())

	o.DisplayName = strings.Repeat("01234567890", 20)
	require.NotNil(t, o.IsValid())

	o.DisplayName = "1234"
	o.Name = "ZZZZZZZ"
	require.NotNil(t, o.IsValid())

	o.Name = "zzzzz"
	require.NotNil(t, o.IsValid())

	o.Type = "U"
	require.NotNil(t, o.IsValid())

	o.Type = "P"
	require.Nil(t, o.IsValid())

	o.Header = strings.Repeat("01234567890", 100)
	require.NotNil(t, o.IsValid())

	o.Header = "1234"
	require.Nil(t, o.IsValid())

	o.Purpose = strings.Repeat("01234567890", 30)
	require.NotNil(t, o.IsValid())

	o.Purpose = "1234"
	require.Nil(t, o.IsValid())

	o.Purpose = strings.Repeat("0123456789", 25)
	require.Nil(t, o.IsValid())
}

func TestChannelPreSave(t *testing.T) {
	o := Channel{Name: "test"}
	o.PreSave()
	o.Etag()
}

func TestChannelPreUpdate(t *testing.T) {
	o := Channel{Name: "test"}
	o.PreUpdate()
}

func TestGetGroupDisplayNameFromUsers(t *testing.T) {
	users := make([]*User, 4)
	users[0] = &User{Username: NewID()}
	users[1] = &User{Username: NewID()}
	users[2] = &User{Username: NewID()}
	users[3] = &User{Username: NewID()}

	name := GetGroupDisplayNameFromUsers(users, true)
	require.LessOrEqual(t, len(name), ChannelNameMaxLength)
}

func TestGetGroupNameFromUserIDs(t *testing.T) {
	name := GetGroupNameFromUserIDs([]string{NewID(), NewID(), NewID(), NewID(), NewID()})

	require.LessOrEqual(t, len(name), ChannelNameMaxLength)
}
