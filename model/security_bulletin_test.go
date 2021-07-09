// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSecurityBulletinToFromJSON(t *testing.T) {
	b := SecurityBulletin{
		ID:               NewID(),
		AppliesToVersion: NewID(),
	}

	j := b.ToJSON()
	b1 := SecurityBulletinFromJSON(strings.NewReader(j))

	require.Equal(t, b, *b1)

	// Malformed JSON
	s2 := `{"wat"`
	b2 := SecurityBulletinFromJSON(strings.NewReader(s2))
	require.Nil(t, b2)
}

func TestSecurityBulletinsToFromJSON(t *testing.T) {
	b := SecurityBulletins{
		{
			ID:               NewID(),
			AppliesToVersion: NewID(),
		},
		{
			ID:               NewID(),
			AppliesToVersion: NewID(),
		},
	}

	j := b.ToJSON()

	b1 := SecurityBulletinsFromJSON(strings.NewReader(j))

	require.Len(t, b1, 2)

	// Malformed JSON
	s2 := `{"wat"`
	b2 := SecurityBulletinsFromJSON(strings.NewReader(s2))

	require.Empty(t, b2)
}
