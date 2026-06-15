// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"strings"
	"testing"

	"github.com/run-llama/llama-parse-cli/internal/mocktest"
)

func TestBetaDirectoriesFilesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"beta:directories:files", "update",
			"--directory-id", "directory_id",
			"--directory-file-id", "directory_file_id",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--display-name", "display_name",
			"--metadata", "{foo: string}",
			"--target-directory-id", "target_directory_id",
			"--unique-id", "x",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"display_name: display_name\n" +
			"metadata:\n" +
			"  foo: string\n" +
			"target_directory_id: target_directory_id\n" +
			"unique_id: x\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"beta:directories:files", "update",
			"--directory-id", "directory_id",
			"--directory-file-id", "directory_file_id",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestBetaDirectoriesFilesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"beta:directories:files", "list",
			"--max-items", "10",
			"--directory-id", "directory_id",
			"--display-name", "display_name",
			"--display-name-contains", "display_name_contains",
			"--expand", "[string, string]",
			"--file-id", "file_id",
			"--include-deleted=true",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--page-size", "0",
			"--page-token", "page_token",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--unique-id", "unique_id",
			"--updated-at-on-or-after", "'2019-12-27T18:11:19.117Z'",
			"--updated-at-on-or-before", "'2019-12-27T18:11:19.117Z'",
		)
	})
}

func TestBetaDirectoriesFilesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"beta:directories:files", "delete",
			"--directory-id", "directory_id",
			"--directory-file-id", "directory_file_id",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestBetaDirectoriesFilesAdd(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"beta:directories:files", "add",
			"--directory-id", "directory_id",
			"--file-id", "file_id",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--display-name", "display_name",
			"--metadata", "{foo: string}",
			"--unique-id", "unique_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"file_id: file_id\n" +
			"display_name: display_name\n" +
			"metadata:\n" +
			"  foo: string\n" +
			"unique_id: unique_id\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"beta:directories:files", "add",
			"--directory-id", "directory_id",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestBetaDirectoriesFilesGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"beta:directories:files", "get",
			"--directory-id", "directory_id",
			"--directory-file-id", "directory_file_id",
			"--expand", "[string, string]",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestBetaDirectoriesFilesUpload(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"beta:directories:files", "upload",
			"--directory-id", "directory_id",
			"--upload-file", mocktest.TestFile(t, "Example data"),
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--display-name", "display_name",
			"--external-file-id", "external_file_id",
			"--metadata", `{"source": "web", "priority": 1}`,
			"--unique-id", "unique_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		testFile := mocktest.TestFile(t, "Example data")
		// Test piping YAML data over stdin
		pipeDataStr := "" +
			"upload_file: Example data\n" +
			"display_name: display_name\n" +
			"external_file_id: external_file_id\n" +
			"metadata: '{\"source\": \"web\", \"priority\": 1}'\n" +
			"unique_id: unique_id\n"
		pipeDataStr = strings.ReplaceAll(pipeDataStr, "Example data", testFile)
		pipeData := []byte(pipeDataStr)
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"beta:directories:files", "upload",
			"--directory-id", "directory_id",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
