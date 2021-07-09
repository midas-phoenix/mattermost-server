// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMfaSecretJSON(t *testing.T) {
	secret := MfaSecret{Secret: NewID(), QRCode: NewID()}
	json := secret.ToJSON()
	result := MfaSecretFromJSON(strings.NewReader(json))

	require.Equal(t, secret.Secret, result.Secret, "Secrets do not match")
}
