// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClusterDiscovery(t *testing.T) {
	o := ClusterDiscovery{
		Type:        "test_type",
		ClusterName: "cluster_name",
		Hostname:    "test_hostname",
	}

	json := o.ToJSON()
	result1 := ClusterDiscoveryFromJSON(strings.NewReader(json))

	assert.NotNil(t, result1)
	assert.Equal(t, "cluster_name", result1.ClusterName)

	result2 := ClusterDiscoveryFromJSON(strings.NewReader(json))
	result3 := ClusterDiscoveryFromJSON(strings.NewReader(json))

	o.ID = "0"
	result1.ID = "1"
	result2.ID = "2"
	result3.ID = "3"
	result3.Hostname = "something_diff"

	assert.True(t, o.IsEqual(result1))

	list := make([]*ClusterDiscovery, 0)
	list = append(list, &o)
	list = append(list, result1)
	list = append(list, result2)
	list = append(list, result3)

	rlist := FilterClusterDiscovery(list, func(in *ClusterDiscovery) bool {
		return !o.IsEqual(in)
	})

	assert.Len(t, rlist, 1)

	o.AutoFillHostname()
	o.Hostname = ""
	o.AutoFillHostname()

	o.AutoFillIPAddress("", "")
	o.Hostname = ""
	o.AutoFillIPAddress("", "")
}
