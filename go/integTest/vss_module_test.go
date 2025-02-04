// Copyright Valkey GLIDE Project Contributors - SPDX Identifier: Apache-2.0

package integTest

import (
	"strings"

	"github.com/jamesx-improving/valkey-glide/go/api"
	"github.com/stretchr/testify/assert"
)

func (suite *GlideTestSuite) TestModuleVerifyVssLoaded() {
	client := suite.defaultClusterClient()
	result, err := client.InfoWithOptions(
		api.ClusterInfoOptions{InfoOptions: &api.InfoOptions{Sections: []api.Section{api.Server}}, Route: nil},
	)

	assert.Nil(suite.T(), err)
	for _, value := range result.MultiValue() {
		assert.True(suite.T(), strings.Contains(value, "# search_index_stats"))
	}
}
