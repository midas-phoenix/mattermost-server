// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/base64"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTriggerIDDecodeAndVerification(t *testing.T) {

	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err)

	t.Run("should succeed decoding and validation", func(t *testing.T) {
		userID := NewID()
		clientTriggerID, triggerID, err := GenerateTriggerID(userID, key)
		require.Nil(t, err)
		decodedClientTriggerID, decodedUserID, err := DecodeAndVerifyTriggerID(triggerID, key)
		assert.Nil(t, err)
		assert.Equal(t, clientTriggerID, decodedClientTriggerID)
		assert.Equal(t, userID, decodedUserID)
	})

	t.Run("should succeed decoding and validation through request structs", func(t *testing.T) {
		actionReq := &PostActionIntegrationRequest{
			UserID: NewID(),
		}
		clientTriggerID, triggerID, err := actionReq.GenerateTriggerID(key)
		require.Nil(t, err)
		dialogReq := &OpenDialogRequest{TriggerID: triggerID}
		decodedClientTriggerID, decodedUserID, err := dialogReq.DecodeAndVerifyTriggerID(key)
		assert.Nil(t, err)
		assert.Equal(t, clientTriggerID, decodedClientTriggerID)
		assert.Equal(t, actionReq.UserID, decodedUserID)
	})

	t.Run("should fail on base64 decode", func(t *testing.T) {
		_, _, err := DecodeAndVerifyTriggerID("junk!", key)
		require.NotNil(t, err)
		assert.Equal(t, "interactive_message.decode_trigger_id.base64_decode_failed", err.ID)
	})

	t.Run("should fail on trigger parsing", func(t *testing.T) {
		_, _, err := DecodeAndVerifyTriggerID(base64.StdEncoding.EncodeToString([]byte("junk!")), key)
		require.NotNil(t, err)
		assert.Equal(t, "interactive_message.decode_trigger_id.missing_data", err.ID)
	})

	t.Run("should fail on expired timestamp", func(t *testing.T) {
		_, _, err := DecodeAndVerifyTriggerID(base64.StdEncoding.EncodeToString([]byte("some-trigger-id:some-user-id:1234567890:junksignature")), key)
		require.NotNil(t, err)
		assert.Equal(t, "interactive_message.decode_trigger_id.expired", err.ID)
	})

	t.Run("should fail on base64 decoding signature", func(t *testing.T) {
		_, _, err := DecodeAndVerifyTriggerID(base64.StdEncoding.EncodeToString([]byte("some-trigger-id:some-user-id:12345678900000:junk!")), key)
		require.NotNil(t, err)
		assert.Equal(t, "interactive_message.decode_trigger_id.base64_decode_failed_signature", err.ID)
	})

	t.Run("should fail on bad signature", func(t *testing.T) {
		_, _, err := DecodeAndVerifyTriggerID(base64.StdEncoding.EncodeToString([]byte("some-trigger-id:some-user-id:12345678900000:junk")), key)
		require.NotNil(t, err)
		assert.Equal(t, "interactive_message.decode_trigger_id.signature_decode_failed", err.ID)
	})

	t.Run("should fail on bad key", func(t *testing.T) {
		_, triggerID, err := GenerateTriggerID(NewID(), key)
		require.Nil(t, err)
		newKey, keyErr := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		require.NoError(t, keyErr)
		_, _, err = DecodeAndVerifyTriggerID(triggerID, newKey)
		require.NotNil(t, err)
		assert.Equal(t, "interactive_message.decode_trigger_id.verify_signature_failed", err.ID)
	})
}

func TestPostActionIntegrationRequestToJSON(t *testing.T) {
	o := PostActionIntegrationRequest{UserID: NewID(), Context: StringInterface{"a": "abc"}}
	j := o.ToJSON()
	ro := PostActionIntegrationRequestFromJSON(bytes.NewReader(j))

	assert.NotNil(t, ro)
	assert.Equal(t, o, *ro)
}

func TestPostActionIntegrationRequestFromJSONError(t *testing.T) {
	ro := PostActionIntegrationRequestFromJSON(strings.NewReader(""))
	assert.Nil(t, ro)
}

func TestPostActionIntegrationResponseToJSON(t *testing.T) {
	o := PostActionIntegrationResponse{Update: &Post{ID: NewID(), Message: NewID()}, EphemeralText: NewID()}
	j := o.ToJSON()
	ro := PostActionIntegrationResponseFromJSON(bytes.NewReader(j))

	assert.NotNil(t, ro)
	assert.Equal(t, o, *ro)
}

func TestPostActionIntegrationResponseFromJSONError(t *testing.T) {
	ro := PostActionIntegrationResponseFromJSON(strings.NewReader(""))
	assert.Nil(t, ro)
}

func TestSubmitDialogRequestToJSON(t *testing.T) {
	t.Run("all fine", func(t *testing.T) {
		request := SubmitDialogRequest{
			URL:        "http://example.org",
			CallbackID: NewID(),
			State:      "some state",
			UserID:     NewID(),
			ChannelID:  NewID(),
			TeamID:     NewID(),
			Submission: map[string]interface{}{
				"text":  "some text",
				"float": 1.2,
				"bool":  true,
			},
			Cancelled: true,
		}
		jsonRequest := request.ToJSON()
		r := SubmitDialogRequestFromJSON(bytes.NewReader(jsonRequest))

		require.NotNil(t, r)
		assert.Equal(t, request, *r)
	})
	t.Run("error", func(t *testing.T) {
		r := SubmitDialogRequestFromJSON(strings.NewReader(""))
		assert.Nil(t, r)
	})
}

func TestSubmitDialogResponseToJSON(t *testing.T) {
	t.Run("all fine", func(t *testing.T) {
		request := SubmitDialogResponse{
			Error: "some generic error",
			Errors: map[string]string{
				"text":  "some text",
				"float": "1.2",
				"bool":  "true",
			},
		}
		jsonRequest := request.ToJSON()
		r := SubmitDialogResponseFromJSON(bytes.NewReader(jsonRequest))

		require.NotNil(t, r)
		assert.Equal(t, request, *r)
	})
	t.Run("error", func(t *testing.T) {
		r := SubmitDialogResponseFromJSON(strings.NewReader(""))
		assert.Nil(t, r)
	})
}

func TestPostActionIntegrationEquals(t *testing.T) {
	t.Run("equal uncomparable types", func(t *testing.T) {
		pa1 := &PostAction{
			Integration: &PostActionIntegration{
				Context: map[string]interface{}{
					"a": map[string]interface{}{
						"a": 0,
					},
				},
			},
		}
		pa2 := &PostAction{
			Integration: &PostActionIntegration{
				Context: map[string]interface{}{
					"a": map[string]interface{}{
						"a": 0,
					},
				},
			},
		}
		require.True(t, pa1.Equals(pa2))
	})

	t.Run("equal comparable types", func(t *testing.T) {
		pa1 := &PostAction{
			Integration: &PostActionIntegration{
				Context: map[string]interface{}{
					"a": "test",
				},
			},
		}
		pa2 := &PostAction{
			Integration: &PostActionIntegration{
				Context: map[string]interface{}{
					"a": "test",
				},
			},
		}
		require.True(t, pa1.Equals(pa2))
	})

	t.Run("non-equal types", func(t *testing.T) {
		pa1 := &PostAction{
			Integration: &PostActionIntegration{
				Context: map[string]interface{}{
					"a": map[string]interface{}{
						"a": 0,
					},
				},
			},
		}
		pa2 := &PostAction{
			Integration: &PostActionIntegration{
				Context: map[string]interface{}{
					"a": "test",
				},
			},
		}
		require.False(t, pa1.Equals(pa2))
	})
}
