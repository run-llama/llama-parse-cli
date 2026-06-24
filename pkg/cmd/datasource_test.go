// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/run-llama/llama-parse-cli/internal/mocktest"
)

func TestDataSourcesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"data-sources", "create",
			"--component", "{foo: bar}",
			"--name", "name",
			"--source-type", "S3",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--custom-metadata", "{foo: {foo: bar}}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"component:\n" +
			"  foo: bar\n" +
			"name: name\n" +
			"source_type: S3\n" +
			"custom_metadata:\n" +
			"  foo:\n" +
			"    foo: bar\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"data-sources", "create",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestDataSourcesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"data-sources", "update",
			"--data-source-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--source-type", "S3",
			"--component", "{foo: bar}",
			"--custom-metadata", "{foo: {foo: bar}}",
			"--name", "name",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"source_type: S3\n" +
			"component:\n" +
			"  foo: bar\n" +
			"custom_metadata:\n" +
			"  foo:\n" +
			"    foo: bar\n" +
			"name: name\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"data-sources", "update",
			"--data-source-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestDataSourcesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"data-sources", "list",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestDataSourcesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"data-sources", "delete",
			"--data-source-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestDataSourcesGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"data-sources", "get",
			"--data-source-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
