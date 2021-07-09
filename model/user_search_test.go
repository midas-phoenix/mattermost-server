// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserSearchJSON(t *testing.T) {
	userSearch := UserSearch{Term: NewID(), TeamID: NewID()}
	json := userSearch.ToJSON()
	ruserSearch := UserSearchFromJSON(bytes.NewReader(json))

	assert.Equal(t, userSearch.Term, ruserSearch.Term, "Terms do not match")
}
