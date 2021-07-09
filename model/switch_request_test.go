// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSwitchRequestJSON(t *testing.T) {
	o := SwitchRequest{Email: NewID(), Password: NewID()}
	json := o.ToJSON()
	ro := SwitchRequestFromJSON(strings.NewReader(json))

	require.Equal(t, o.Email, ro.Email, "Emails do not match")
}
