// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBotTrace(t *testing.T) {
	bot := &Bot{
		UserID:         NewID(),
		Username:       "username",
		DisplayName:    "display name",
		Description:    "description",
		OwnerID:        NewID(),
		LastIconUpdate: 1,
		CreateAt:       2,
		UpdateAt:       3,
		DeleteAt:       4,
	}

	require.Equal(t, map[string]interface{}{"user_id": bot.UserID}, bot.Trace())
}

func TestBotClone(t *testing.T) {
	bot := &Bot{
		UserID:         NewID(),
		Username:       "username",
		DisplayName:    "display name",
		Description:    "description",
		OwnerID:        NewID(),
		LastIconUpdate: 1,
		CreateAt:       2,
		UpdateAt:       3,
		DeleteAt:       4,
	}

	clone := bot.Clone()

	require.Equal(t, bot, bot.Clone())
	require.False(t, bot == clone)
}

func TestBotIsValid(t *testing.T) {
	testCases := []struct {
		Description     string
		Bot             *Bot
		ExpectedIsValid bool
	}{
		{
			"nil bot",
			&Bot{},
			false,
		},
		{
			"bot with missing user id",
			&Bot{
				UserID:         "",
				Username:       "username",
				DisplayName:    "display name",
				Description:    "description",
				OwnerID:        NewID(),
				LastIconUpdate: 1,
				CreateAt:       2,
				UpdateAt:       3,
				DeleteAt:       4,
			},
			false,
		},
		{
			"bot with invalid user id",
			&Bot{
				UserID:         "invalid",
				Username:       "username",
				DisplayName:    "display name",
				Description:    "description",
				OwnerID:        NewID(),
				LastIconUpdate: 1,
				CreateAt:       2,
				UpdateAt:       3,
				DeleteAt:       4,
			},
			false,
		},
		{
			"bot with missing username",
			&Bot{
				UserID:         NewID(),
				Username:       "",
				DisplayName:    "display name",
				Description:    "description",
				OwnerID:        NewID(),
				LastIconUpdate: 1,
				CreateAt:       2,
				UpdateAt:       3,
				DeleteAt:       4,
			},
			false,
		},
		{
			"bot with invalid username",
			&Bot{
				UserID:         NewID(),
				Username:       "a@",
				DisplayName:    "display name",
				Description:    "description",
				OwnerID:        NewID(),
				LastIconUpdate: 1,
				CreateAt:       2,
				UpdateAt:       3,
				DeleteAt:       4,
			},
			false,
		},
		{
			"bot with long description",
			&Bot{
				UserID:         "",
				Username:       "username",
				DisplayName:    "display name",
				Description:    strings.Repeat("x", 1025),
				OwnerID:        NewID(),
				LastIconUpdate: 1,
				CreateAt:       2,
				UpdateAt:       3,
				DeleteAt:       4,
			},
			false,
		},
		{
			"bot with missing creator id",
			&Bot{
				UserID:         NewID(),
				Username:       "username",
				DisplayName:    "display name",
				Description:    "description",
				OwnerID:        "",
				LastIconUpdate: 1,
				CreateAt:       2,
				UpdateAt:       3,
				DeleteAt:       4,
			},
			false,
		},
		{
			"bot without create at timestamp",
			&Bot{
				UserID:         NewID(),
				Username:       "username",
				DisplayName:    "display name",
				Description:    "description",
				OwnerID:        NewID(),
				LastIconUpdate: 1,
				CreateAt:       0,
				UpdateAt:       3,
				DeleteAt:       4,
			},
			false,
		},
		{
			"bot without update at timestamp",
			&Bot{
				UserID:         NewID(),
				Username:       "username",
				DisplayName:    "display name",
				Description:    "description",
				OwnerID:        NewID(),
				LastIconUpdate: 1,
				CreateAt:       2,
				UpdateAt:       0,
				DeleteAt:       4,
			},
			false,
		},
		{
			"bot",
			&Bot{
				UserID:         NewID(),
				Username:       "username",
				DisplayName:    "display name",
				Description:    "description",
				OwnerID:        NewID(),
				LastIconUpdate: 1,
				CreateAt:       2,
				UpdateAt:       3,
				DeleteAt:       0,
			},
			true,
		},
		{
			"bot without description",
			&Bot{
				UserID:         NewID(),
				Username:       "username",
				DisplayName:    "display name",
				Description:    "",
				OwnerID:        NewID(),
				LastIconUpdate: 1,
				CreateAt:       2,
				UpdateAt:       3,
				DeleteAt:       0,
			},
			true,
		},
		{
			"deleted bot",
			&Bot{
				UserID:         NewID(),
				Username:       "username",
				DisplayName:    "display name",
				Description:    "a description",
				OwnerID:        NewID(),
				LastIconUpdate: 1,
				CreateAt:       2,
				UpdateAt:       3,
				DeleteAt:       4,
			},
			true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			if testCase.ExpectedIsValid {
				require.Nil(t, testCase.Bot.IsValid())
			} else {
				require.NotNil(t, testCase.Bot.IsValid())
			}
		})
	}
}

func TestBotPreSave(t *testing.T) {
	bot := &Bot{
		UserID:         NewID(),
		Username:       "username",
		DisplayName:    "display name",
		Description:    "description",
		OwnerID:        NewID(),
		LastIconUpdate: 0,
		DeleteAt:       0,
	}

	originalBot := &*bot

	bot.PreSave()
	assert.NotEqual(t, 0, bot.CreateAt)
	assert.NotEqual(t, 0, bot.UpdateAt)

	originalBot.CreateAt = bot.CreateAt
	originalBot.UpdateAt = bot.UpdateAt
	assert.Equal(t, originalBot, bot)
}

func TestBotPreUpdate(t *testing.T) {
	bot := &Bot{
		UserID:         NewID(),
		Username:       "username",
		DisplayName:    "display name",
		Description:    "description",
		OwnerID:        NewID(),
		LastIconUpdate: 1,
		CreateAt:       2,
		DeleteAt:       0,
	}

	originalBot := &*bot

	bot.PreSave()
	assert.NotEqual(t, 0, bot.UpdateAt)

	originalBot.UpdateAt = bot.UpdateAt
	assert.Equal(t, originalBot, bot)
}

func TestBotEtag(t *testing.T) {
	t.Run("same etags", func(t *testing.T) {
		bot1 := &Bot{
			UserID:         NewID(),
			Username:       "username",
			DisplayName:    "display name",
			Description:    "description",
			OwnerID:        NewID(),
			LastIconUpdate: 1,
			CreateAt:       2,
			UpdateAt:       3,
			DeleteAt:       4,
		}
		bot2 := bot1

		assert.Equal(t, bot1.Etag(), bot2.Etag())
	})
	t.Run("different etags", func(t *testing.T) {
		t.Run("different user id", func(t *testing.T) {
			bot1 := &Bot{
				UserID:         NewID(),
				Username:       "username",
				DisplayName:    "display name",
				Description:    "description",
				OwnerID:        NewID(),
				LastIconUpdate: 1,
				CreateAt:       2,
				UpdateAt:       3,
				DeleteAt:       4,
			}
			bot2 := &Bot{
				UserID:         NewID(),
				Username:       "username",
				DisplayName:    "display name",
				Description:    "description",
				OwnerID:        bot1.OwnerID,
				LastIconUpdate: 1,
				CreateAt:       2,
				UpdateAt:       3,
				DeleteAt:       4,
			}

			assert.NotEqual(t, bot1.Etag(), bot2.Etag())
		})
		t.Run("different update at", func(t *testing.T) {
			bot1 := &Bot{
				UserID:         NewID(),
				Username:       "username",
				DisplayName:    "display name",
				Description:    "description",
				OwnerID:        NewID(),
				LastIconUpdate: 1,
				CreateAt:       2,
				UpdateAt:       3,
				DeleteAt:       4,
			}
			bot2 := &Bot{
				UserID:         bot1.UserID,
				Username:       "username",
				DisplayName:    "display name",
				Description:    "description",
				OwnerID:        bot1.OwnerID,
				LastIconUpdate: 1,
				CreateAt:       2,
				UpdateAt:       10,
				DeleteAt:       4,
			}

			assert.NotEqual(t, bot1.Etag(), bot2.Etag())
		})
	})
}

func TestBotToAndFromJSON(t *testing.T) {
	bot1 := &Bot{
		UserID:         NewID(),
		Username:       "username",
		DisplayName:    "display name",
		Description:    "description",
		OwnerID:        NewID(),
		LastIconUpdate: 1,
		CreateAt:       2,
		UpdateAt:       3,
		DeleteAt:       4,
	}

	bot2 := &Bot{
		UserID:         NewID(),
		Username:       "username",
		DisplayName:    "display name",
		Description:    "description 2",
		OwnerID:        NewID(),
		LastIconUpdate: 5,
		CreateAt:       6,
		UpdateAt:       7,
		DeleteAt:       8,
	}

	assert.Equal(t, bot1, BotFromJSON(bytes.NewReader(bot1.ToJSON())))
	assert.Equal(t, bot2, BotFromJSON(bytes.NewReader(bot2.ToJSON())))
}

func sToP(s string) *string {
	return &s
}

func TestBotPatch(t *testing.T) {
	userID1 := NewID()
	creatorID1 := NewID()

	testCases := []struct {
		Description string
		Bot         *Bot
		BotPatch    *BotPatch
		ExpectedBot *Bot
	}{
		{
			"no update",
			&Bot{
				UserID:         userID1,
				Username:       "username",
				DisplayName:    "display name",
				Description:    "description",
				OwnerID:        creatorID1,
				LastIconUpdate: 1,
				CreateAt:       2,
				UpdateAt:       3,
				DeleteAt:       4,
			},
			&BotPatch{},
			&Bot{
				UserID:         userID1,
				Username:       "username",
				DisplayName:    "display name",
				Description:    "description",
				OwnerID:        creatorID1,
				LastIconUpdate: 1,
				CreateAt:       2,
				UpdateAt:       3,
				DeleteAt:       4,
			},
		},
		{
			"partial update",
			&Bot{
				UserID:         userID1,
				Username:       "username",
				DisplayName:    "display name",
				Description:    "description",
				OwnerID:        creatorID1,
				LastIconUpdate: 1,
				CreateAt:       2,
				UpdateAt:       3,
				DeleteAt:       4,
			},
			&BotPatch{
				Username:    sToP("new_username"),
				DisplayName: nil,
				Description: sToP("new description"),
			},
			&Bot{
				UserID:         userID1,
				Username:       "new_username",
				DisplayName:    "display name",
				Description:    "new description",
				OwnerID:        creatorID1,
				LastIconUpdate: 1,
				CreateAt:       2,
				UpdateAt:       3,
				DeleteAt:       4,
			},
		},
		{
			"full update",
			&Bot{
				UserID:         userID1,
				Username:       "username",
				DisplayName:    "display name",
				Description:    "description",
				OwnerID:        creatorID1,
				LastIconUpdate: 1,
				CreateAt:       2,
				UpdateAt:       3,
				DeleteAt:       4,
			},
			&BotPatch{
				Username:    sToP("new_username"),
				DisplayName: sToP("new display name"),
				Description: sToP("new description"),
			},
			&Bot{
				UserID:         userID1,
				Username:       "new_username",
				DisplayName:    "new display name",
				Description:    "new description",
				OwnerID:        creatorID1,
				LastIconUpdate: 1,
				CreateAt:       2,
				UpdateAt:       3,
				DeleteAt:       4,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			testCase.Bot.Patch(testCase.BotPatch)
			assert.Equal(t, testCase.ExpectedBot, testCase.Bot)
		})
	}
}

func TestBotWouldPatch(t *testing.T) {
	b := &Bot{
		UserID: NewID(),
	}

	t.Run("nil patch", func(t *testing.T) {
		ok := b.WouldPatch(nil)
		require.False(t, ok)
	})

	t.Run("nil patch fields", func(t *testing.T) {
		patch := &BotPatch{}
		ok := b.WouldPatch(patch)
		require.False(t, ok)
	})

	t.Run("patch", func(t *testing.T) {
		patch := &BotPatch{
			DisplayName: NewString("BotName"),
		}
		ok := b.WouldPatch(patch)
		require.True(t, ok)
	})

	t.Run("no patch", func(t *testing.T) {
		patch := &BotPatch{
			DisplayName: NewString("BotName"),
		}
		b.Patch(patch)
		ok := b.WouldPatch(patch)
		require.False(t, ok)
	})
}

func TestBotPatchToAndFromJSON(t *testing.T) {
	botPatch1 := &BotPatch{
		Username:    sToP("username"),
		DisplayName: sToP("display name"),
		Description: sToP("description"),
	}

	botPatch2 := &BotPatch{
		Username:    sToP("username"),
		DisplayName: sToP("display name"),
		Description: sToP("description 2"),
	}

	assert.Equal(t, botPatch1, BotPatchFromJSON(bytes.NewReader(botPatch1.ToJSON())))
	assert.Equal(t, botPatch2, BotPatchFromJSON(bytes.NewReader(botPatch2.ToJSON())))
}

func TestUserFromBot(t *testing.T) {
	bot1 := &Bot{
		UserID:         NewID(),
		Username:       "username",
		DisplayName:    "display name",
		Description:    "description",
		OwnerID:        NewID(),
		LastIconUpdate: 1,
		CreateAt:       2,
		UpdateAt:       3,
		DeleteAt:       4,
	}

	bot2 := &Bot{
		UserID:         NewID(),
		Username:       "username2",
		DisplayName:    "display name 2",
		Description:    "description 2",
		OwnerID:        NewID(),
		LastIconUpdate: 5,
		CreateAt:       6,
		UpdateAt:       7,
		DeleteAt:       8,
	}

	assert.Equal(t, &User{
		ID:        bot1.UserID,
		Username:  "username",
		Email:     "username@localhost",
		FirstName: "display name",
		Roles:     "system_user",
	}, UserFromBot(bot1))
	assert.Equal(t, &User{
		ID:        bot2.UserID,
		Username:  "username2",
		Email:     "username2@localhost",
		FirstName: "display name 2",
		Roles:     "system_user",
	}, UserFromBot(bot2))
}

func TestBotFromUser(t *testing.T) {
	user := &User{
		ID:       NewID(),
		Username: "username",
		CreateAt: 1,
		UpdateAt: 2,
		DeleteAt: 3,
	}

	assert.Equal(t, &Bot{
		OwnerID:     user.ID,
		UserID:      user.ID,
		Username:    "username",
		DisplayName: "username",
	}, BotFromUser(user))
}

func TestBotListToAndFromJSON(t *testing.T) {
	testCases := []struct {
		Description string
		BotList     BotList
	}{
		{
			"empty list",
			BotList{},
		},
		{
			"single item",
			BotList{
				&Bot{
					UserID:         NewID(),
					Username:       "username",
					DisplayName:    "display name",
					Description:    "description",
					OwnerID:        NewID(),
					LastIconUpdate: 1,
					CreateAt:       2,
					UpdateAt:       3,
					DeleteAt:       4,
				},
			},
		},
		{
			"multiple items",
			BotList{
				&Bot{
					UserID:         NewID(),
					Username:       "username",
					DisplayName:    "display name",
					Description:    "description",
					OwnerID:        NewID(),
					LastIconUpdate: 1,
					CreateAt:       2,
					UpdateAt:       3,
					DeleteAt:       4,
				},

				&Bot{
					UserID:         NewID(),
					Username:       "username",
					DisplayName:    "display name",
					Description:    "description 2",
					OwnerID:        NewID(),
					LastIconUpdate: 5,
					CreateAt:       6,
					UpdateAt:       7,
					DeleteAt:       8,
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			assert.Equal(t, testCase.BotList, BotListFromJSON(bytes.NewReader(testCase.BotList.ToJSON())))
		})
	}
}

func TestBotListEtag(t *testing.T) {
	bot1 := &Bot{
		UserID:         NewID(),
		Username:       "username",
		DisplayName:    "display name",
		Description:    "description",
		OwnerID:        NewID(),
		LastIconUpdate: 1,
		CreateAt:       2,
		UpdateAt:       3,
		DeleteAt:       4,
	}

	bot1Updated := &Bot{
		UserID:         NewID(),
		Username:       "username",
		DisplayName:    "display name",
		Description:    "description",
		OwnerID:        NewID(),
		LastIconUpdate: 1,
		CreateAt:       2,
		UpdateAt:       10,
		DeleteAt:       4,
	}

	bot2 := &Bot{
		UserID:         NewID(),
		Username:       "username",
		DisplayName:    "display name",
		Description:    "description",
		OwnerID:        NewID(),
		LastIconUpdate: 5,
		CreateAt:       6,
		UpdateAt:       7,
		DeleteAt:       8,
	}

	testCases := []struct {
		Description   string
		BotListA      BotList
		BotListB      BotList
		ExpectedEqual bool
	}{
		{
			"empty lists",
			BotList{},
			BotList{},
			true,
		},
		{
			"single item, same list",
			BotList{bot1},
			BotList{bot1},
			true,
		},
		{
			"single item, different update at",
			BotList{bot1},
			BotList{bot1Updated},
			false,
		},
		{
			"single item vs. multiple items",
			BotList{bot1},
			BotList{bot1, bot2},
			false,
		},
		{
			"multiple items, different update at",
			BotList{bot1, bot2},
			BotList{bot1Updated, bot2},
			false,
		},
		{
			"multiple items, same list",
			BotList{bot1, bot2},
			BotList{bot1, bot2},
			true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			if testCase.ExpectedEqual {
				assert.Equal(t, testCase.BotListA.Etag(), testCase.BotListB.Etag())
			} else {
				assert.NotEqual(t, testCase.BotListA.Etag(), testCase.BotListB.Etag())
			}
		})
	}
}

func TestIsBotChannel(t *testing.T) {
	for _, test := range []struct {
		Name     string
		Channel  *Channel
		Expected bool
	}{
		{
			Name:     "not a direct channel",
			Channel:  &Channel{Type: ChannelTypeOpen},
			Expected: false,
		},
		{
			Name: "a direct channel with another user",
			Channel: &Channel{
				Name: "user1__user2",
				Type: ChannelTypeDirect,
			},
			Expected: false,
		},
		{
			Name: "a direct channel with the name containing the bot's ID first",
			Channel: &Channel{
				Name: "botUserID__user2",
				Type: ChannelTypeDirect,
			},
			Expected: true,
		},
		{
			Name: "a direct channel with the name containing the bot's ID second",
			Channel: &Channel{
				Name: "user1__botUserID",
				Type: ChannelTypeDirect,
			},
			Expected: true,
		},
	} {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.Expected, IsBotDMChannel(test.Channel, "botUserID"))
		})
	}
}
