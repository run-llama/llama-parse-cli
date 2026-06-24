// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/run-llama/llama-parse-cli/internal/mocktest"
	"github.com/run-llama/llama-parse-cli/internal/requestflag"
)

func TestPipelinesFilesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pipelines:files", "create",
			"--pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--body", "{file_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, custom_metadata: {foo: {foo: bar}}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(pipelinesFilesCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pipelines:files", "create",
			"--pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--body.file-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--body.custom-metadata", "{foo: {foo: bar}}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"- file_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"  custom_metadata:\n" +
			"    foo:\n" +
			"      foo: bar\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"pipelines:files", "create",
			"--pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPipelinesFilesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pipelines:files", "update",
			"--pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--file-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--custom-metadata", "{foo: {foo: bar}}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"custom_metadata:\n" +
			"  foo:\n" +
			"    foo: bar\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"pipelines:files", "update",
			"--pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--file-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPipelinesFilesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pipelines:files", "list",
			"--max-items", "10",
			"--pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--data-source-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--file-name-contains", "file_name_contains",
			"--limit", "0",
			"--offset", "0",
			"--only-manually-uploaded=true",
			"--order-by", "order_by",
			"--status", "[NOT_STARTED, IN_PROGRESS]",
		)
	})
}

func TestPipelinesFilesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pipelines:files", "delete",
			"--pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--file-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPipelinesFilesGetStatus(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pipelines:files", "get-status",
			"--pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--file-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPipelinesFilesGetStatusCounts(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pipelines:files", "get-status-counts",
			"--pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--data-source-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--only-manually-uploaded=true",
		)
	})
}
