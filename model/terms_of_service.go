// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"unicode/utf8"
)

type TermsOfService struct {
	ID       string `json:"id"`
	CreateAt int64  `json:"create_at"`
	UserID   string `json:"user_id"`
	Text     string `json:"text"`
}

func (t *TermsOfService) IsValid() *AppError {
	if !IsValidID(t.ID) {
		return InvalidTermsOfServiceError("id", "")
	}

	if t.CreateAt == 0 {
		return InvalidTermsOfServiceError("create_at", t.ID)
	}

	if !IsValidID(t.UserID) {
		return InvalidTermsOfServiceError("user_id", t.ID)
	}

	if utf8.RuneCountInString(t.Text) > PostMessageMaxRunesV2 {
		return InvalidTermsOfServiceError("text", t.ID)
	}

	return nil
}

func (t *TermsOfService) ToJSON() string {
	b, _ := json.Marshal(t)
	return string(b)
}

func TermsOfServiceFromJSON(data io.Reader) *TermsOfService {
	var termsOfService *TermsOfService
	json.NewDecoder(data).Decode(&termsOfService)
	return termsOfService
}

func InvalidTermsOfServiceError(fieldName string, termsOfServiceID string) *AppError {
	id := fmt.Sprintf("model.terms_of_service.is_valid.%s.app_error", fieldName)
	details := ""
	if termsOfServiceID != "" {
		details = "terms_of_service_id=" + termsOfServiceID
	}
	return NewAppError("TermsOfService.IsValid", id, map[string]interface{}{"MaxLength": PostMessageMaxRunesV2}, details, http.StatusBadRequest)
}

func (t *TermsOfService) PreSave() {
	if t.ID == "" {
		t.ID = NewID()
	}

	t.CreateAt = GetMillis()
}
