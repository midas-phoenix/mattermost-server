// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAuditJson(t *testing.T) {
	audit := Audit{ID: NewID(), UserID: NewID(), CreateAt: GetMillis()}
	json := audit.ToJson()
	result := AuditFromJson(strings.NewReader(json))
	require.Equal(t, audit.ID, result.ID, "Ids do not match")
}
