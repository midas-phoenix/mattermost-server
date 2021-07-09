// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWebSocketRequest(t *testing.T) {
	m := WebSocketRequest{Seq: 1, Action: "test"}
	json := m.ToJSON()
	result := WebSocketRequestFromJSON(strings.NewReader(json))

	require.NotNil(t, result)

	badresult := WebSocketRequestFromJSON(strings.NewReader("junk"))

	require.Nil(t, badresult)
}
