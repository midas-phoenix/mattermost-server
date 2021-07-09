// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInitialLoadJSON(t *testing.T) {
	u := &User{ID: NewID()}
	o := InitialLoad{User: u}
	json := o.ToJSON()
	ro := InitialLoadFromJSON(strings.NewReader(json))

	require.Equal(t, o.User.ID, ro.User.ID, "Ids do not match")
}
