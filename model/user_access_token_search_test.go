// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserAccessTokenSearchJSON(t *testing.T) {
	userAccessTokenSearch := UserAccessTokenSearch{Term: NewID()}
	json := userAccessTokenSearch.ToJSON()
	ruserAccessTokenSearch := UserAccessTokenSearchFromJSON(strings.NewReader(json))
	require.Equal(t, userAccessTokenSearch.Term, ruserAccessTokenSearch.Term, "Terms do not match")
}
