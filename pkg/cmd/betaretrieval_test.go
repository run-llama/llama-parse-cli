// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/llamacloud-prod-cli/internal/mocktest"
	"github.com/stainless-sdks/llamacloud-prod-cli/internal/requestflag"
)

func TestBetaRetrievalRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"beta:retrieval", "retrieve",
			"--index-id", "idx-abc123",
			"--query", "What are the key findings?",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--custom-filters", "{foo: {operator: eq, value: string}}",
			"--full-text-pipeline-weight", "0",
			"--num-candidates", "0",
			"--rerank", "{enabled: true, top_n: 5}",
			"--score-threshold", "0",
			"--static-filters", "{parsed_directory_file_id: {operator: eq, value: string}}",
			"--top-k", "10",
			"--vector-pipeline-weight", "0",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(betaRetrievalRetrieve)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"beta:retrieval", "retrieve",
			"--index-id", "idx-abc123",
			"--query", "What are the key findings?",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--custom-filters", "{foo: {operator: eq, value: string}}",
			"--full-text-pipeline-weight", "0",
			"--num-candidates", "0",
			"--rerank.enabled=true",
			"--rerank.top-n", "5",
			"--score-threshold", "0",
			"--static-filters.parsed-directory-file-id", "{operator: eq, value: string}",
			"--top-k", "10",
			"--vector-pipeline-weight", "0",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"index_id: idx-abc123\n" +
			"query: What are the key findings?\n" +
			"custom_filters:\n" +
			"  foo:\n" +
			"    operator: eq\n" +
			"    value: string\n" +
			"full_text_pipeline_weight: 0\n" +
			"num_candidates: 0\n" +
			"rerank:\n" +
			"  enabled: true\n" +
			"  top_n: 5\n" +
			"score_threshold: 0\n" +
			"static_filters:\n" +
			"  parsed_directory_file_id:\n" +
			"    operator: eq\n" +
			"    value: string\n" +
			"top_k: 10\n" +
			"vector_pipeline_weight: 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"beta:retrieval", "retrieve",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestBetaRetrievalFind(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"beta:retrieval", "find",
			"--max-items", "10",
			"--index-id", "idx-abc123",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--file-name", "file_name",
			"--file-name-contains", "file_name_contains",
			"--page-size", "0",
			"--page-token", "page_token",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"index_id: idx-abc123\n" +
			"file_name: file_name\n" +
			"file_name_contains: file_name_contains\n" +
			"page_size: 0\n" +
			"page_token: page_token\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"beta:retrieval", "find",
			"--max-items", "10",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestBetaRetrievalGrep(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"beta:retrieval", "grep",
			"--max-items", "10",
			"--file-id", "file_id",
			"--index-id", "idx-abc123",
			"--pattern", "revenue|profit",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--context-chars", "0",
			"--page-size", "0",
			"--page-token", "page_token",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"file_id: file_id\n" +
			"index_id: idx-abc123\n" +
			"pattern: revenue|profit\n" +
			"context_chars: 0\n" +
			"page_size: 0\n" +
			"page_token: page_token\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"beta:retrieval", "grep",
			"--max-items", "10",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestBetaRetrievalRead(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"beta:retrieval", "read",
			"--file-id", "file_id",
			"--index-id", "idx-abc123",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--max-length", "0",
			"--offset", "0",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"file_id: file_id\n" +
			"index_id: idx-abc123\n" +
			"max_length: 0\n" +
			"offset: 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"beta:retrieval", "read",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
