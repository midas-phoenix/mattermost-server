// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReactionIsValid(t *testing.T) {
	tests := []struct {
		// reaction
		reaction Reaction
		// error message to print
		errMsg string
		// should there be an error
		shouldErr bool
	}{
		{
			reaction: Reaction{
				UserID:    NewID(),
				PostID:    NewID(),
				EmojiName: "emoji",
				CreateAt:  GetMillis(),
				UpdateAt:  GetMillis(),
			},
			errMsg:    "",
			shouldErr: false,
		},
		{
			reaction: Reaction{
				UserID:    "",
				PostID:    NewID(),
				EmojiName: "emoji",
				CreateAt:  GetMillis(),
				UpdateAt:  GetMillis(),
			},
			errMsg:    "user id should be invalid",
			shouldErr: true,
		},
		{
			reaction: Reaction{
				UserID:    "1234garbage",
				PostID:    NewID(),
				EmojiName: "emoji",
				CreateAt:  GetMillis(),
				UpdateAt:  GetMillis(),
			},
			errMsg:    "user id should be invalid",
			shouldErr: true,
		},
		{
			reaction: Reaction{
				UserID:    NewID(),
				PostID:    "",
				EmojiName: "emoji",
				CreateAt:  GetMillis(),
				UpdateAt:  GetMillis(),
			},
			errMsg:    "post id should be invalid",
			shouldErr: true,
		},
		{
			reaction: Reaction{
				UserID:    NewID(),
				PostID:    "1234garbage",
				EmojiName: "emoji",
				CreateAt:  GetMillis(),
				UpdateAt:  GetMillis(),
			},
			errMsg:    "post id should be invalid",
			shouldErr: true,
		},
		{
			reaction: Reaction{
				UserID:    NewID(),
				PostID:    NewID(),
				EmojiName: strings.Repeat("a", 64),
				CreateAt:  GetMillis(),
				UpdateAt:  GetMillis(),
			},
			errMsg:    "",
			shouldErr: false,
		},
		{
			reaction: Reaction{
				UserID:    NewID(),
				PostID:    NewID(),
				EmojiName: "emoji-",
				CreateAt:  GetMillis(),
				UpdateAt:  GetMillis(),
			},
			errMsg:    "",
			shouldErr: false,
		},
		{
			reaction: Reaction{
				UserID:    NewID(),
				PostID:    NewID(),
				EmojiName: "emoji_",
				CreateAt:  GetMillis(),
				UpdateAt:  GetMillis(),
			},
			errMsg:    "",
			shouldErr: false,
		},
		{
			reaction: Reaction{
				UserID:    NewID(),
				PostID:    NewID(),
				EmojiName: "+1",
				CreateAt:  GetMillis(),
				UpdateAt:  GetMillis(),
			},
			errMsg:    "",
			shouldErr: false,
		},
		{
			reaction: Reaction{
				UserID:    NewID(),
				PostID:    NewID(),
				EmojiName: "emoji:",
				CreateAt:  GetMillis(),
				UpdateAt:  GetMillis(),
			},
			errMsg:    "",
			shouldErr: true,
		},
		{
			reaction: Reaction{
				UserID:    NewID(),
				PostID:    NewID(),
				EmojiName: "",
				CreateAt:  GetMillis(),
				UpdateAt:  GetMillis(),
			},
			errMsg:    "emoji name should be invalid",
			shouldErr: true,
		},
		{
			reaction: Reaction{
				UserID:    NewID(),
				PostID:    NewID(),
				EmojiName: strings.Repeat("a", 65),
				CreateAt:  GetMillis(),
				UpdateAt:  GetMillis(),
			},
			errMsg:    "emoji name should be invalid",
			shouldErr: true,
		},
		{
			reaction: Reaction{
				UserID:    NewID(),
				PostID:    NewID(),
				EmojiName: "emoji",
				CreateAt:  0,
				UpdateAt:  GetMillis(),
			},
			errMsg:    "create at should be invalid",
			shouldErr: true,
		},
		{
			reaction: Reaction{
				UserID:    NewID(),
				PostID:    NewID(),
				EmojiName: "emoji",
				CreateAt:  GetMillis(),
				UpdateAt:  0,
			},
			errMsg:    "update at should be invalid",
			shouldErr: true,
		},
	}

	for _, test := range tests {
		err := test.reaction.IsValid()
		if test.shouldErr {
			// there should be an error here
			require.NotNil(t, err, test.errMsg)
		} else {
			// err should be nil here
			require.Nil(t, err, test.errMsg)
		}
	}
}
