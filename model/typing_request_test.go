// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTypingRequestJson(t *testing.T) {
	o := TypingRequest{ChannelID: NewID(), ParentID: NewID()}
	json := o.ToJson()
	ro := TypingRequestFromJson(strings.NewReader(json))

	require.Equal(t, o.ChannelID, ro.ChannelID, "ChannelIds do not match")
	require.Equal(t, o.ParentID, ro.ParentID, "ParentIds do not match")
}
