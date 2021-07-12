// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAuthJson(t *testing.T) {
	a1 := AuthData{}
	a1.ClientID = NewID()
	a1.UserID = NewID()
	a1.Code = NewID()

	json := a1.ToJson()
	ra1 := AuthDataFromJson(strings.NewReader(json))
	require.Equal(t, a1.Code, ra1.Code, "codes didn't match")

	a2 := AuthorizeRequest{}
	a2.ClientID = NewID()
	a2.Scope = NewID()

	json = a2.ToJson()
	ra2 := AuthorizeRequestFromJson(strings.NewReader(json))

	require.Equal(t, a2.ClientID, ra2.ClientID, "client ids didn't match")
}

func TestAuthPreSave(t *testing.T) {
	a1 := AuthData{}
	a1.ClientID = NewID()
	a1.UserID = NewID()
	a1.Code = NewID()
	a1.PreSave()
	a1.IsExpired()
}

func TestAuthIsValid(t *testing.T) {

	ad := AuthData{}

	require.NotNil(t, ad.IsValid())

	ad.ClientID = NewRandomString(28)
	require.NotNil(t, ad.IsValid(), "Should have failed Client Id")

	ad.ClientID = NewID()
	require.NotNil(t, ad.IsValid())

	ad.UserID = NewRandomString(28)
	require.NotNil(t, ad.IsValid(), "Should have failed User Id")

	ad.UserID = NewID()
	require.NotNil(t, ad.IsValid())

	ad.Code = NewRandomString(129)
	require.NotNil(t, ad.IsValid(), "Should have failed Code to long")

	ad.Code = ""
	require.NotNil(t, ad.IsValid(), "Should have failed Code not set")

	ad.Code = NewID()
	require.NotNil(t, ad.IsValid())

	ad.ExpiresIn = 0
	require.NotNil(t, ad.IsValid(), "Should have failed invalid ExpiresIn")

	ad.ExpiresIn = 1
	require.NotNil(t, ad.IsValid())

	ad.CreateAt = 0
	require.NotNil(t, ad.IsValid(), "Should have failed Invalid Create At")

	ad.CreateAt = 1
	require.NotNil(t, ad.IsValid())

	ad.State = NewRandomString(129)
	require.NotNil(t, ad.IsValid(), "Should have failed invalid State")

	ad.State = NewRandomString(128)
	require.NotNil(t, ad.IsValid())

	ad.Scope = NewRandomString(1025)
	require.NotNil(t, ad.IsValid(), "Should have failed invalid Scope")

	ad.Scope = NewRandomString(128)
	require.NotNil(t, ad.IsValid())

	ad.RedirectUri = ""
	require.NotNil(t, ad.IsValid(), "Should have failed Redirect URI not set")

	ad.RedirectUri = NewRandomString(28)
	require.NotNil(t, ad.IsValid(), "Should have failed invalid URL")

	ad.RedirectUri = NewRandomString(257)
	require.NotNil(t, ad.IsValid(), "Should have failed invalid URL")

	ad.RedirectUri = "http://example.com"
	require.Nil(t, ad.IsValid())
}
