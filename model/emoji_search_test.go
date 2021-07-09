// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEmojiSearchJSON(t *testing.T) {
	emojiSearch := EmojiSearch{Term: NewID()}
	json := emojiSearch.ToJSON()
	remojiSearch := EmojiSearchFromJSON(strings.NewReader(json))

	require.Equal(t, emojiSearch.Term, remojiSearch.Term, "Terms do not match")
}
