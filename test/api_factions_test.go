/*
SpaceTraders API

Testing FactionsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package spacetraders

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/jabbrwcky/spacetraders-go"
)

func Test_spacetraders_FactionsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FactionsAPIService GetFaction", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var factionSymbol string

		resp, httpRes, err := apiClient.FactionsAPI.GetFaction(context.Background(), factionSymbol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FactionsAPIService GetFactions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FactionsAPI.GetFactions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
