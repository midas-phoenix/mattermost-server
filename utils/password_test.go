// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package utils

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mattermost/mattermost-server/v5/model"
)

func TestIsPasswordValidWithSettings(t *testing.T) {
	for name, tc := range map[string]struct {
		Password      string
		Settings      *model.PasswordSettings
		ExpectedError string
	}{
		"Short": {
			Password: strings.Repeat("x", 3),
			Settings: &model.PasswordSettings{
				MinimumLength: model.NewInt(3),
				Lowercase:     model.NewBool(false),
				Uppercase:     model.NewBool(false),
				Number:        model.NewBool(false),
				Symbol:        model.NewBool(false),
			},
		},
		"Long": {
			Password: strings.Repeat("x", model.PASSWORD_MAXIMUM_LENGTH),
			Settings: &model.PasswordSettings{
				Lowercase: model.NewBool(false),
				Uppercase: model.NewBool(false),
				Number:    model.NewBool(false),
				Symbol:    model.NewBool(false),
			},
		},
		"TooShort": {
			Password: strings.Repeat("x", 2),
			Settings: &model.PasswordSettings{
				MinimumLength: model.NewInt(3),
				Lowercase:     model.NewBool(false),
				Uppercase:     model.NewBool(false),
				Number:        model.NewBool(false),
				Symbol:        model.NewBool(false),
			},
			ExpectedError: "model.user.is_valid.pwd.app_error",
		},
		"TooLong": {
			Password: strings.Repeat("x", model.PASSWORD_MAXIMUM_LENGTH+1),
			Settings: &model.PasswordSettings{
				Lowercase: model.NewBool(false),
				Uppercase: model.NewBool(false),
				Number:    model.NewBool(false),
				Symbol:    model.NewBool(false),
			},
			ExpectedError: "model.user.is_valid.pwd.app_error",
		},
		"MissingLower": {
			Password: "AAAAAAAAAAASD123!@#",
			Settings: &model.PasswordSettings{
				Lowercase: model.NewBool(true),
				Uppercase: model.NewBool(false),
				Number:    model.NewBool(false),
				Symbol:    model.NewBool(false),
			},
			ExpectedError: "model.user.is_valid.pwd_lowercase.app_error",
		},
		"MissingUpper": {
			Password: "aaaaaaaaaaaaasd123!@#",
			Settings: &model.PasswordSettings{
				Uppercase: model.NewBool(true),
				Lowercase: model.NewBool(false),
				Number:    model.NewBool(false),
				Symbol:    model.NewBool(false),
			},
			ExpectedError: "model.user.is_valid.pwd_uppercase.app_error",
		},
		"MissingNumber": {
			Password: "asasdasdsadASD!@#",
			Settings: &model.PasswordSettings{
				Number:    model.NewBool(true),
				Lowercase: model.NewBool(false),
				Uppercase: model.NewBool(false),
				Symbol:    model.NewBool(false),
			},
			ExpectedError: "model.user.is_valid.pwd_number.app_error",
		},
		"MissingSymbol": {
			Password: "asdasdasdasdasdASD123",
			Settings: &model.PasswordSettings{
				Symbol:    model.NewBool(true),
				Lowercase: model.NewBool(false),
				Uppercase: model.NewBool(false),
				Number:    model.NewBool(false),
			},
			ExpectedError: "model.user.is_valid.pwd_symbol.app_error",
		},
		"MissingMultiple": {
			Password: "asdasdasdasdasdasd",
			Settings: &model.PasswordSettings{
				Lowercase: model.NewBool(true),
				Uppercase: model.NewBool(true),
				Number:    model.NewBool(true),
				Symbol:    model.NewBool(true),
			},
			ExpectedError: "model.user.is_valid.pwd_lowercase_uppercase_number_symbol.app_error",
		},
		"Everything": {
			Password: "asdASD!@#123",
			Settings: &model.PasswordSettings{
				Lowercase: model.NewBool(true),
				Uppercase: model.NewBool(true),
				Number:    model.NewBool(true),
				Symbol:    model.NewBool(true),
			},
		},
	} {
		tc.Settings.SetDefaults()
		t.Run(name, func(t *testing.T) {
			if err := IsPasswordValidWithSettings(tc.Password, tc.Settings); tc.ExpectedError == "" {
				assert.Nil(t, err)
			} else {
				assert.Equal(t, tc.ExpectedError, err.Id)
			}
		})
	}
}