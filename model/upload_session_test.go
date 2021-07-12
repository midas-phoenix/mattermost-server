// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUploadSessionIsValid(t *testing.T) {
	var session UploadSession

	t.Run("empty session should fail", func(t *testing.T) {
		err := session.IsValid()
		require.NotNil(t, err)
	})

	t.Run("valid session should succeed", func(t *testing.T) {
		session = UploadSession{
			ID:         NewID(),
			Type:       UploadTypeAttachment,
			CreateAt:   GetMillis(),
			UserID:     NewID(),
			ChannelID:  NewID(),
			Filename:   "test",
			Path:       "/tmp/test",
			FileSize:   1024,
			FileOffset: 0,
		}
		err := session.IsValid()
		require.Nil(t, err)
	})

	t.Run("invalid Id should fail", func(t *testing.T) {
		us := session
		us.ID = "invalid"
		err := us.IsValid()
		require.NotNil(t, err)
		require.Equal(t, "model.upload_session.is_valid.id.app_error", err.ID)
	})

	t.Run("invalid type should fail", func(t *testing.T) {
		us := session
		us.Type = "invalid"
		err := us.IsValid()
		require.NotNil(t, err)
		require.Equal(t, "model.upload_session.is_valid.type.app_error", err.ID)
	})

	t.Run("invalid CreateAt should fail", func(t *testing.T) {
		us := session
		us.CreateAt = 0
		err := us.IsValid()
		require.NotNil(t, err)
		require.Equal(t, "model.upload_session.is_valid.create_at.app_error", err.ID)
	})

	t.Run("invalid UserId should fail", func(t *testing.T) {
		us := session
		us.UserID = "invalid"
		err := us.IsValid()
		require.NotNil(t, err)
		require.Equal(t, "model.upload_session.is_valid.user_id.app_error", err.ID)
	})

	t.Run("invalid ChannelId should fail", func(t *testing.T) {
		us := session
		us.ChannelID = "invalid"
		err := us.IsValid()
		require.NotNil(t, err)
		require.Equal(t, "model.upload_session.is_valid.channel_id.app_error", err.ID)
	})

	t.Run("ChannelId is not validated if type is not attachment", func(t *testing.T) {
		us := session
		us.ChannelID = ""
		us.Type = UploadTypeImport
		err := us.IsValid()
		require.Nil(t, err)
	})

	t.Run("invalid Filename should fail", func(t *testing.T) {
		us := session
		us.Filename = ""
		err := us.IsValid()
		require.NotNil(t, err)
		require.Equal(t, "model.upload_session.is_valid.filename.app_error", err.ID)
	})

	t.Run("invalid Path should fail", func(t *testing.T) {
		us := session
		us.Path = ""
		err := us.IsValid()
		require.NotNil(t, err)
		require.Equal(t, "model.upload_session.is_valid.path.app_error", err.ID)
	})

	t.Run("invalid FileSize should fail", func(t *testing.T) {
		us := session
		us.FileSize = 0
		err := us.IsValid()
		require.NotNil(t, err)
		require.Equal(t, "model.upload_session.is_valid.file_size.app_error", err.ID)

		us.FileSize = -1
		err = us.IsValid()
		require.NotNil(t, err)
		require.Equal(t, "model.upload_session.is_valid.file_size.app_error", err.ID)
	})

	t.Run("invalid FileOffset should fail", func(t *testing.T) {
		us := session
		us.FileOffset = us.FileSize + 1
		err := us.IsValid()
		require.NotNil(t, err)
		require.Equal(t, "model.upload_session.is_valid.file_offset.app_error", err.ID)

		us.FileOffset = -1
		err = us.IsValid()
		require.NotNil(t, err)
		require.Equal(t, "model.upload_session.is_valid.file_offset.app_error", err.ID)
	})
}
