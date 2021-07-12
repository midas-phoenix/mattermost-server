// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"encoding/json"
	"errors"
)

type OrphanedRecord struct {
	ParentID *string `json:"parent_id"`
	ChildID  *string `json:"child_id"`
}

type RelationalIntegrityCheckData struct {
	ParentName   string           `json:"parent_name"`
	ChildName    string           `json:"child_name"`
	ParentIDAttr string           `json:"parent_id_attr"`
	ChildIDAttr  string           `json:"child_id_attr"`
	Records      []OrphanedRecord `json:"records"`
}

type IntegrityCheckResult struct {
	Data interface{} `json:"data"`
	Err  error       `json:"err"`
}

func (r *IntegrityCheckResult) UnmarshalJSON(b []byte) error {
	var data map[string]interface{}
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}
	if d, ok := data["data"]; ok && d != nil {
		var rdata RelationalIntegrityCheckData
		m := d.(map[string]interface{})
		rdata.ParentName = m["parent_name"].(string)
		rdata.ChildName = m["child_name"].(string)
		rdata.ParentIDAttr = m["parent_id_attr"].(string)
		rdata.ChildIDAttr = m["child_id_attr"].(string)
		for _, recData := range m["records"].([]interface{}) {
			var record OrphanedRecord
			m := recData.(map[string]interface{})
			if val := m["parent_id"]; val != nil {
				record.ParentID = NewString(val.(string))
			}
			if val := m["child_id"]; val != nil {
				record.ChildID = NewString(val.(string))
			}
			rdata.Records = append(rdata.Records, record)
		}
		r.Data = rdata
	}
	if err, ok := data["err"]; ok && err != nil {
		r.Err = errors.New(data["err"].(string))
	}
	return nil
}
