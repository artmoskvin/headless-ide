/*
Daytona Server API

Testing ContainerRegistryAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package daytona

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/hide-org/hide/pkg/daytona"
)

func Test_daytona_ContainerRegistryAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContainerRegistryAPIService GetContainerRegistry", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var server string

		resp, httpRes, err := apiClient.ContainerRegistryAPI.GetContainerRegistry(context.Background(), server).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContainerRegistryAPIService ListContainerRegistries", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContainerRegistryAPI.ListContainerRegistries(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContainerRegistryAPIService RemoveContainerRegistry", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var server string

		httpRes, err := apiClient.ContainerRegistryAPI.RemoveContainerRegistry(context.Background(), server).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContainerRegistryAPIService SetContainerRegistry", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var server string

		httpRes, err := apiClient.ContainerRegistryAPI.SetContainerRegistry(context.Background(), server).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}