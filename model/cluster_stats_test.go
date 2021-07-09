// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClusterStatsJSON(t *testing.T) {
	cluster := ClusterStats{ID: NewID(), TotalWebsocketConnections: 1, TotalReadDbConnections: 1}
	json := cluster.ToJSON()
	result := ClusterStatsFromJSON(strings.NewReader(json))

	require.Equal(t, cluster.ID, result.ID, "Ids do not match")
}
