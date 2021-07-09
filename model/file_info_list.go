// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"encoding/json"
	"io"
	"sort"
)

type FileInfoList struct {
	Order          []string             `json:"order"`
	FileInfos      map[string]*FileInfo `json:"file_infos"`
	NextFileInfoID string               `json:"next_file_info_id"`
	PrevFileInfoID string               `json:"prev_file_info_id"`
}

func NewFileInfoList() *FileInfoList {
	return &FileInfoList{
		Order:          make([]string, 0),
		FileInfos:      make(map[string]*FileInfo),
		NextFileInfoID: "",
		PrevFileInfoID: "",
	}
}

func (o *FileInfoList) ToSlice() []*FileInfo {
	var fileInfos []*FileInfo
	for _, id := range o.Order {
		fileInfos = append(fileInfos, o.FileInfos[id])
	}
	return fileInfos
}

func (o *FileInfoList) ToJSON() string {
	b, err := json.Marshal(o)
	if err != nil {
		return ""
	}

	return string(b)
}

func (o *FileInfoList) MakeNonNil() {
	if o.Order == nil {
		o.Order = make([]string, 0)
	}

	if o.FileInfos == nil {
		o.FileInfos = make(map[string]*FileInfo)
	}
}

func (o *FileInfoList) AddOrder(id string) {
	if o.Order == nil {
		o.Order = make([]string, 0, 128)
	}

	o.Order = append(o.Order, id)
}

func (o *FileInfoList) AddFileInfo(fileInfo *FileInfo) {
	if o.FileInfos == nil {
		o.FileInfos = make(map[string]*FileInfo)
	}

	o.FileInfos[fileInfo.ID] = fileInfo
}

func (o *FileInfoList) UniqueOrder() {
	keys := make(map[string]bool)
	order := []string{}
	for _, fileInfoID := range o.Order {
		if _, value := keys[fileInfoID]; !value {
			keys[fileInfoID] = true
			order = append(order, fileInfoID)
		}
	}

	o.Order = order
}

func (o *FileInfoList) Extend(other *FileInfoList) {
	for fileInfoID := range other.FileInfos {
		o.AddFileInfo(other.FileInfos[fileInfoID])
	}

	for _, fileInfoID := range other.Order {
		o.AddOrder(fileInfoID)
	}

	o.UniqueOrder()
}

func (o *FileInfoList) SortByCreateAt() {
	sort.Slice(o.Order, func(i, j int) bool {
		return o.FileInfos[o.Order[i]].CreateAt > o.FileInfos[o.Order[j]].CreateAt
	})
}

func (o *FileInfoList) Etag() string {
	id := "0"
	var t int64 = 0

	for _, v := range o.FileInfos {
		if v.UpdateAt > t {
			t = v.UpdateAt
			id = v.ID
		} else if v.UpdateAt == t && v.ID > id {
			t = v.UpdateAt
			id = v.ID
		}
	}

	orderID := ""
	if len(o.Order) > 0 {
		orderID = o.Order[0]
	}

	return Etag(orderID, id, t)
}

func FileInfoListFromJSON(data io.Reader) *FileInfoList {
	var o *FileInfoList
	json.NewDecoder(data).Decode(&o)
	return o
}
