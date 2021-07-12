// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type UserTermsOfService struct {
	UserID           string `json:"user_id"`
	TermsOfServiceID string `json:"terms_of_service_id"`
	CreateAt         int64  `json:"create_at"`
}

func (ut *UserTermsOfService) IsValid() *AppError {
	if !IsValidID(ut.UserID) {
		return InvalidUserTermsOfServiceError("user_id", ut.UserID)
	}

	if !IsValidID(ut.TermsOfServiceID) {
		return InvalidUserTermsOfServiceError("terms_of_service_id", ut.UserID)
	}

	if ut.CreateAt == 0 {
		return InvalidUserTermsOfServiceError("create_at", ut.UserID)
	}

	return nil
}

func (ut *UserTermsOfService) ToJson() string {
	b, _ := json.Marshal(ut)
	return string(b)
}

func (ut *UserTermsOfService) PreSave() {
	if ut.UserID == "" {
		ut.UserID = NewID()
	}

	ut.CreateAt = GetMillis()
}

func UserTermsOfServiceFromJson(data io.Reader) *UserTermsOfService {
	var userTermsOfService *UserTermsOfService
	json.NewDecoder(data).Decode(&userTermsOfService)
	return userTermsOfService
}

func InvalidUserTermsOfServiceError(fieldName string, userTermsOfServiceID string) *AppError {
	id := fmt.Sprintf("model.user_terms_of_service.is_valid.%s.app_error", fieldName)
	details := ""
	if userTermsOfServiceID != "" {
		details = "user_terms_of_service_user_id=" + userTermsOfServiceID
	}
	return NewAppError("UserTermsOfService.IsValid", id, nil, details, http.StatusBadRequest)
}
