// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCommandWebhookPreSave(t *testing.T) {
	h := CommandWebhook{}
	h.PreSave()

	require.Len(t, h.ID, 26, "Id should be generated")
	require.NotEqual(t, 0, h.CreateAt, "CreateAt should be set")
}

func TestCommandWebhookIsValid(t *testing.T) {
	h := CommandWebhook{}
	h.ID = NewID()
	h.CreateAt = GetMillis()
	h.CommandID = NewID()
	h.UserID = NewID()
	h.ChannelID = NewID()

	for _, test := range []struct {
		Transform     func()
		ExpectedError string
	}{
		{func() {}, ""},
		{func() { h.ID = "asd" }, "model.command_hook.id.app_error"},
		{func() { h.ID = NewID() }, ""},
		{func() { h.CreateAt = 0 }, "model.command_hook.create_at.app_error"},
		{func() { h.CreateAt = GetMillis() }, ""},
		{func() { h.CommandID = "asd" }, "model.command_hook.command_id.app_error"},
		{func() { h.CommandID = NewID() }, ""},
		{func() { h.UserID = "asd" }, "model.command_hook.user_id.app_error"},
		{func() { h.UserID = NewID() }, ""},
		{func() { h.ChannelID = "asd" }, "model.command_hook.channel_id.app_error"},
		{func() { h.ChannelID = NewID() }, ""},
		{func() { h.RootID = "asd" }, "model.command_hook.root_id.app_error"},
		{func() { h.RootID = NewID() }, ""},
		{func() { h.ParentID = "asd" }, "model.command_hook.parent_id.app_error"},
		{func() { h.ParentID = NewID() }, ""},
	} {
		tmp := h
		test.Transform()
		err := h.IsValid()

		if test.ExpectedError == "" {
			assert.Nil(t, err, "hook should be valid")
		} else {
			require.NotNil(t, err)
			assert.Equal(t, test.ExpectedError, err.ID, "expected "+test.ExpectedError+" error")
		}

		h = tmp
	}
}
