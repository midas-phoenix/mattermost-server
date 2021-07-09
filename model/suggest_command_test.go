// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuggestCommandJSON(t *testing.T) {
	command := &SuggestCommand{Suggestion: NewID()}
	json := command.ToJSON()
	result := SuggestCommandFromJSON(strings.NewReader(json))

	assert.Equal(t, command.Suggestion, result.Suggestion, "Ids do not match")
}
