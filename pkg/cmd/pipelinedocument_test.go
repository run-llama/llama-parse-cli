// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/run-llama/llama-parse-cli/internal/mocktest"
	"github.com/run-llama/llama-parse-cli/internal/requestflag"
)

func TestPipelinesDocumentsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pipelines:documents", "create",
			"--pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--body", "{metadata: {foo: bar}, text: text, id: id, excluded_embed_metadata_keys: [string], excluded_llm_metadata_keys: [string], page_positions: [0]}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(pipelinesDocumentsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pipelines:documents", "create",
			"--pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--body.metadata", "{foo: bar}",
			"--body.text", "text",
			"--body.id", "id",
			"--body.excluded-embed-metadata-keys", "[string]",
			"--body.excluded-llm-metadata-keys", "[string]",
			"--body.page-positions", "[0]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"- metadata:\n" +
			"    foo: bar\n" +
			"  text: text\n" +
			"  id: id\n" +
			"  excluded_embed_metadata_keys:\n" +
			"    - string\n" +
			"  excluded_llm_metadata_keys:\n" +
			"    - string\n" +
			"  page_positions:\n" +
			"    - 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"pipelines:documents", "create",
			"--pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPipelinesDocumentsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pipelines:documents", "list",
			"--max-items", "10",
			"--pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--file-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--limit", "0",
			"--only-api-data-source-documents=true",
			"--only-direct-upload=true",
			"--skip", "0",
			"--status-refresh-policy", "cached",
		)
	})
}

func TestPipelinesDocumentsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pipelines:documents", "delete",
			"--pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--document-id", "document_id",
		)
	})
}

func TestPipelinesDocumentsGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pipelines:documents", "get",
			"--pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--document-id", "document_id",
		)
	})
}

func TestPipelinesDocumentsGetChunks(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pipelines:documents", "get-chunks",
			"--pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--document-id", "document_id",
		)
	})
}

func TestPipelinesDocumentsGetStatus(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pipelines:documents", "get-status",
			"--pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--document-id", "document_id",
		)
	})
}

func TestPipelinesDocumentsSync(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pipelines:documents", "sync",
			"--pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--document-id", "document_id",
		)
	})
}

func TestPipelinesDocumentsUpsert(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pipelines:documents", "upsert",
			"--pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--body", "{metadata: {foo: bar}, text: text, id: id, excluded_embed_metadata_keys: [string], excluded_llm_metadata_keys: [string], page_positions: [0]}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(pipelinesDocumentsUpsert)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pipelines:documents", "upsert",
			"--pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--body.metadata", "{foo: bar}",
			"--body.text", "text",
			"--body.id", "id",
			"--body.excluded-embed-metadata-keys", "[string]",
			"--body.excluded-llm-metadata-keys", "[string]",
			"--body.page-positions", "[0]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"- metadata:\n" +
			"    foo: bar\n" +
			"  text: text\n" +
			"  id: id\n" +
			"  excluded_embed_metadata_keys:\n" +
			"    - string\n" +
			"  excluded_llm_metadata_keys:\n" +
			"    - string\n" +
			"  page_positions:\n" +
			"    - 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"pipelines:documents", "upsert",
			"--pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
