// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestSystemJSON(t *testing.T) {
	system := System{Name: "test", Value: NewID()}
	json := system.ToJSON()
	result := SystemFromJSON(strings.NewReader(json))

	require.Equal(t, "test", result.Name, "ids do not match")
}

func TestServerBusyJSON(t *testing.T) {
	now := time.Now()
	sbs := ServerBusyState{Busy: true, Expires: now.Unix()}
	json := sbs.ToJSON()
	result := ServerBusyStateFromJSON(strings.NewReader(json))

	require.Equal(t, sbs.Busy, result.Busy, "busy state does not match")
	require.Equal(t, sbs.Expires, result.Expires, "expiry does not match")
}
