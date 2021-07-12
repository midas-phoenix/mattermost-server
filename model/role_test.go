// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChannelModeratedPermissionsChangedByPatch(t *testing.T) {
	testCases := []struct {
		Name             string
		Permissions      []string
		PatchPermissions []string
		Expected         []string
	}{
		{
			"Empty patch returns empty slice",
			[]string{},
			[]string{},
			[]string{},
		},
		{
			"Adds permissions to empty initial permissions list",
			[]string{},
			[]string{PermissionCreatePost.ID, PermissionAddReaction.ID},
			[]string{ChannelModeratedPermissions[0], ChannelModeratedPermissions[1]},
		},
		{
			"Ignores non moderated permissions in initial permissions list",
			[]string{PermissionAssignBot.ID},
			[]string{PermissionCreatePost.ID, PermissionRemoveReaction.ID},
			[]string{ChannelModeratedPermissions[0], ChannelModeratedPermissions[1]},
		},
		{
			"Adds removed moderated permissions from initial permissions list",
			[]string{PermissionCreatePost.ID},
			[]string{},
			[]string{PermissionCreatePost.ID},
		},
		{
			"No changes returns empty slice",
			[]string{PermissionCreatePost.ID, PermissionAssignBot.ID},
			[]string{PermissionCreatePost.ID},
			[]string{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			baseRole := &Role{Permissions: tc.Permissions}
			rolePatch := &RolePatch{Permissions: &tc.PatchPermissions}
			result := ChannelModeratedPermissionsChangedByPatch(baseRole, rolePatch)
			assert.ElementsMatch(t, tc.Expected, result)
		})
	}
}

func TestRolePatchFromChannelModerationsPatch(t *testing.T) {
	createPosts := ChannelModeratedPermissions[0]
	createReactions := ChannelModeratedPermissions[1]
	manageMembers := ChannelModeratedPermissions[2]
	channelMentions := ChannelModeratedPermissions[3]

	basePermissions := []string{
		PermissionAddReaction.ID,
		PermissionRemoveReaction.ID,
		PermissionCreatePost.ID,
		PermissionUseChannelMentions.ID,
		PermissionManagePublicChannelMembers.ID,
		PermissionUploadFile.ID,
		PermissionGetPublicLink.ID,
		PermissionUseSlashCommands.ID,
	}

	baseModeratedPermissions := []string{
		PermissionAddReaction.ID,
		PermissionRemoveReaction.ID,
		PermissionCreatePost.ID,
		PermissionManagePublicChannelMembers.ID,
		PermissionUseChannelMentions.ID,
	}

	testCases := []struct {
		Name                     string
		Permissions              []string
		ChannelModerationsPatch  []*ChannelModerationPatch
		RoleName                 string
		ExpectedPatchPermissions []string
	}{
		{
			"Patch to member role adding a permission that already exists",
			basePermissions,
			[]*ChannelModerationPatch{
				{
					Name:  &createReactions,
					Roles: &ChannelModeratedRolesPatch{Members: NewBool(true)},
				},
			},
			"members",
			baseModeratedPermissions,
		},
		{
			"Patch to member role with moderation patch for guest role",
			basePermissions,
			[]*ChannelModerationPatch{
				{
					Name:  &createReactions,
					Roles: &ChannelModeratedRolesPatch{Guests: NewBool(true)},
				},
			},
			"members",
			baseModeratedPermissions,
		},
		{
			"Patch to guest role with moderation patch for member role",
			basePermissions,
			[]*ChannelModerationPatch{
				{
					Name:  &createReactions,
					Roles: &ChannelModeratedRolesPatch{Members: NewBool(true)},
				},
			},
			"guests",
			baseModeratedPermissions,
		},
		{
			"Patch to member role removing multiple channel moderated permissions",
			basePermissions,
			[]*ChannelModerationPatch{
				{
					Name:  &createReactions,
					Roles: &ChannelModeratedRolesPatch{Members: NewBool(false)},
				},
				{
					Name:  &manageMembers,
					Roles: &ChannelModeratedRolesPatch{Members: NewBool(false)},
				},
				{
					Name:  &channelMentions,
					Roles: &ChannelModeratedRolesPatch{Members: NewBool(false)},
				},
			},
			"members",
			[]string{PermissionCreatePost.ID},
		},
		{
			"Patch to guest role removing multiple channel moderated permissions",
			basePermissions,
			[]*ChannelModerationPatch{
				{
					Name:  &createReactions,
					Roles: &ChannelModeratedRolesPatch{Guests: NewBool(false)},
				},
				{
					Name:  &manageMembers,
					Roles: &ChannelModeratedRolesPatch{Guests: NewBool(false)},
				},
				{
					Name:  &channelMentions,
					Roles: &ChannelModeratedRolesPatch{Guests: NewBool(false)},
				},
			},
			"guests",
			[]string{PermissionCreatePost.ID},
		},
		{
			"Patch enabling and removing multiple channel moderated permissions ",
			[]string{PermissionAddReaction.ID, PermissionManagePublicChannelMembers.ID},
			[]*ChannelModerationPatch{
				{
					Name:  &createReactions,
					Roles: &ChannelModeratedRolesPatch{Members: NewBool(false)},
				},
				{
					Name:  &manageMembers,
					Roles: &ChannelModeratedRolesPatch{Members: NewBool(false)},
				},
				{
					Name:  &channelMentions,
					Roles: &ChannelModeratedRolesPatch{Members: NewBool(true)},
				},
				{
					Name:  &createPosts,
					Roles: &ChannelModeratedRolesPatch{Members: NewBool(true)},
				},
			},
			"members",
			[]string{PermissionCreatePost.ID, PermissionUseChannelMentions.ID},
		},
		{
			"Patch enabling a partially enabled permission",
			[]string{PermissionAddReaction.ID},
			[]*ChannelModerationPatch{
				{
					Name:  &createReactions,
					Roles: &ChannelModeratedRolesPatch{Members: NewBool(true)},
				},
			},
			"members",
			[]string{PermissionAddReaction.ID, PermissionRemoveReaction.ID},
		},
		{
			"Patch disabling a partially disabled permission",
			[]string{PermissionAddReaction.ID},
			[]*ChannelModerationPatch{
				{
					Name:  &createReactions,
					Roles: &ChannelModeratedRolesPatch{Members: NewBool(false)},
				},
				{
					Name:  &createPosts,
					Roles: &ChannelModeratedRolesPatch{Members: NewBool(true)},
				},
			},
			"members",
			[]string{PermissionCreatePost.ID},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			baseRole := &Role{Permissions: tc.Permissions}
			rolePatch := baseRole.RolePatchFromChannelModerationsPatch(tc.ChannelModerationsPatch, tc.RoleName)
			assert.ElementsMatch(t, tc.ExpectedPatchPermissions, *rolePatch.Permissions)
		})
	}
}

func TestGetChannelModeratedPermissions(t *testing.T) {
	tests := []struct {
		Name        string
		Permissions []string
		ChannelType string
		Expected    map[string]bool
	}{
		{
			"Filters non moderated permissions",
			[]string{PermissionCreateBot.ID},
			ChannelTypeOpen,
			map[string]bool{},
		},
		{
			"Returns a map of moderated permissions",
			[]string{PermissionCreatePost.ID, PermissionAddReaction.ID, PermissionRemoveReaction.ID, PermissionManagePublicChannelMembers.ID, PermissionManagePrivateChannelMembers.ID, PermissionUseChannelMentions.ID},
			ChannelTypeOpen,
			map[string]bool{
				ChannelModeratedPermissions[0]: true,
				ChannelModeratedPermissions[1]: true,
				ChannelModeratedPermissions[2]: true,
				ChannelModeratedPermissions[3]: true,
			},
		},
		{
			"Returns a map of moderated permissions when non moderated present",
			[]string{PermissionCreatePost.ID, PermissionCreateDirectChannel.ID},
			ChannelTypeOpen,
			map[string]bool{
				ChannelModeratedPermissions[0]: true,
			},
		},
		{
			"Returns a nothing when no permissions present",
			[]string{},
			ChannelTypeOpen,
			map[string]bool{},
		},
	}
	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			role := &Role{Permissions: tc.Permissions}
			moderatedPermissions := role.GetChannelModeratedPermissions(tc.ChannelType)
			for permission := range moderatedPermissions {
				assert.Equal(t, moderatedPermissions[permission], tc.Expected[permission])
			}
		})
	}
}
