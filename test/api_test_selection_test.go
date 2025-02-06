/*
Test Selection Services

Testing TestSelectionAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/evergreen-ci/test-selection-client"
)

func Test_openapi_TestSelectionAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TestSelectionAPIService ExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNameTestNamesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string
		var requester string
		var buildVariantName string
		var taskId string
		var taskName string
		var testNames string

		resp, httpRes, err := apiClient.TestSelectionAPI.ExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNameTestNamesGet(context.Background(), projectId, requester, buildVariantName, taskId, taskName, testNames).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TestSelectionAPIService SelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNameTestNamesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string
		var requester string
		var buildVariantName string
		var taskId string
		var taskName string
		var testNames string

		resp, httpRes, err := apiClient.TestSelectionAPI.SelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNameTestNamesGet(context.Background(), projectId, requester, buildVariantName, taskId, taskName, testNames).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
