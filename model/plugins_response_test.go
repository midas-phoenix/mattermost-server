// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPluginsResponseJSON(t *testing.T) {
	manifest := &Manifest{
		ID: "theid",
		Server: &ManifestServer{
			Executable: "theexecutable",
		},
		Webapp: &ManifestWebapp{
			BundlePath: "thebundlepath",
		},
	}

	response := &PluginsResponse{
		Active:   []*PluginInfo{{Manifest: *manifest}},
		Inactive: []*PluginInfo{},
	}

	json := response.ToJSON()
	newResponse := PluginsResponseFromJSON(strings.NewReader(json))
	assert.Equal(t, newResponse, response)
	assert.Equal(t, newResponse.ToJSON(), json)
	assert.Equal(t, PluginsResponseFromJSON(strings.NewReader("junk")), (*PluginsResponse)(nil))
}
