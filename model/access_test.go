// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAccessJSON(t *testing.T) {
	a1 := AccessData{}
	a1.ClientID = NewID()
	a1.UserID = NewID()
	a1.Token = NewID()
	a1.RefreshToken = NewID()

	json := a1.ToJSON()
	ra1 := AccessDataFromJSON(strings.NewReader(json))

	require.Equal(t, a1.Token, ra1.Token)
}

func TestAccessIsValid(t *testing.T) {
	ad := AccessData{}

	require.NotNil(t, ad.IsValid())

	ad.ClientID = NewRandomString(28)
	require.NotNil(t, ad.IsValid())

	ad.ClientID = ""
	require.NotNil(t, ad.IsValid())

	ad.ClientID = NewID()
	require.NotNil(t, ad.IsValid())

	ad.UserID = NewRandomString(28)
	require.NotNil(t, ad.IsValid())

	ad.UserID = ""
	require.NotNil(t, ad.IsValid())

	ad.UserID = NewID()
	require.NotNil(t, ad.IsValid())

	ad.Token = NewRandomString(22)
	require.NotNil(t, ad.IsValid())

	ad.Token = NewID()
	require.NotNil(t, ad.IsValid())

	ad.RefreshToken = NewRandomString(28)
	require.NotNil(t, ad.IsValid())

	ad.RefreshToken = NewID()
	require.NotNil(t, ad.IsValid())

	ad.RedirectURI = ""
	require.NotNil(t, ad.IsValid())

	ad.RedirectURI = NewRandomString(28)
	require.NotNil(t, ad.IsValid())

	ad.RedirectURI = "http://example.com"
	require.Nil(t, ad.IsValid())
}
