// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAuditsJSON(t *testing.T) {
	audit := Audit{ID: NewID(), UserID: NewID(), CreateAt: GetMillis()}
	json := audit.ToJSON()
	result := AuditFromJSON(strings.NewReader(json))

	require.Equal(t, audit.ID, result.ID, "Ids do not match")

	var audits Audits = make([]Audit, 1)
	audits[0] = audit

	ljson := audits.ToJSON()
	results := AuditsFromJSON(strings.NewReader(ljson))

	require.Equal(t, audits[0].ID, results[0].ID, "Ids do not match")
}
