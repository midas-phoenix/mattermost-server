// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"encoding/json"
	"io"
)

const (
	MaxImageSize = int64(6048 * 4032) // 24 megapixels, roughly 36MB as a raw image
)

type FileUploadResponse struct {
	FileInfos []*FileInfo `json:"file_infos"`
	ClientIDs []string    `json:"client_ids"`
}

func FileUploadResponseFromJSON(data io.Reader) *FileUploadResponse {
	var o *FileUploadResponse
	json.NewDecoder(data).Decode(&o)
	return o
}

func (o *FileUploadResponse) ToJSON() string {
	b, _ := json.Marshal(o)
	return string(b)
}
