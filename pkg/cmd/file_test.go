// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"strings"
	"testing"

	"github.com/stainless-sdks/llamacloud-prod-cli/internal/mocktest"
	"github.com/stainless-sdks/llamacloud-prod-cli/internal/requestflag"
)

func TestFilesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"files", "create",
			"--file", mocktest.TestFile(t, "Example data"),
			"--purpose", "purpose",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--external-file-id", "external_file_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		testFile := mocktest.TestFile(t, "Example data")
		// Test piping YAML data over stdin
		pipeDataStr := "" +
			"file: Example data\n" +
			"purpose: purpose\n" +
			"external_file_id: external_file_id\n"
		pipeDataStr = strings.ReplaceAll(pipeDataStr, "Example data", testFile)
		pipeData := []byte(pipeDataStr)
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"files", "create",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestFilesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"files", "list",
			"--max-items", "10",
			"--external-file-id", "external_file_id",
			"--file-id", "[182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e]",
			"--file-name", "file_name",
			"--order-by", "order_by",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--page-size", "1",
			"--page-token", "page_token",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestFilesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"files", "delete",
			"--file-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestFilesGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"files", "get",
			"--file-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--expires-at-seconds", "0",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestFilesQuery(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"files", "query",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--filter", "{data_source_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, external_file_id: external_file_id, file_ids: [182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e], file_name: file_name, only_manually_uploaded: true, project_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e}",
			"--order-by", "order_by",
			"--page-size", "0",
			"--page-token", "page_token",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(filesQuery)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"files", "query",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--filter.data-source-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--filter.external-file-id", "external_file_id",
			"--filter.file-ids", "[182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e]",
			"--filter.file-name", "file_name",
			"--filter.only-manually-uploaded=true",
			"--filter.project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--order-by", "order_by",
			"--page-size", "0",
			"--page-token", "page_token",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"filter:\n" +
			"  data_source_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"  external_file_id: external_file_id\n" +
			"  file_ids:\n" +
			"    - 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"  file_name: file_name\n" +
			"  only_manually_uploaded: true\n" +
			"  project_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"order_by: order_by\n" +
			"page_size: 0\n" +
			"page_token: page_token\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"files", "query",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
