// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/llamacloud-prod-cli/internal/mocktest"
	"github.com/stainless-sdks/llamacloud-prod-cli/internal/requestflag"
)

func TestPipelinesDataSourcesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pipelines:data-sources", "update",
			"--pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--data-source-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--sync-interval", "0",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("sync_interval: 0")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"pipelines:data-sources", "update",
			"--pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--data-source-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPipelinesDataSourcesGetDataSources(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pipelines:data-sources", "get-data-sources",
			"--pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPipelinesDataSourcesGetStatus(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pipelines:data-sources", "get-status",
			"--pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--data-source-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPipelinesDataSourcesSync(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pipelines:data-sources", "sync",
			"--pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--data-source-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--pipeline-file-id", "[182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"pipeline_file_ids:\n" +
			"  - 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"pipelines:data-sources", "sync",
			"--pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--data-source-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPipelinesDataSourcesUpdateDataSources(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pipelines:data-sources", "update-data-sources",
			"--pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--body", "{data_source_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, sync_interval: 0}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(pipelinesDataSourcesUpdateDataSources)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pipelines:data-sources", "update-data-sources",
			"--pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--body.data-source-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--body.sync-interval", "0",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"- data_source_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"  sync_interval: 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"pipelines:data-sources", "update-data-sources",
			"--pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
