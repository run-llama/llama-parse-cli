// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/run-llama/llama-parse-cli/internal/mocktest"
	"github.com/run-llama/llama-parse-cli/internal/requestflag"
)

func TestRetrieversQuerySearch(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"retrievers:query", "search",
			"--retriever-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--query", "x",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--mode", "routing",
			"--rerank-config", "{top_n: 1, type: system_default}",
			"--rerank-top-n", "0",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(retrieversQuerySearch)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"retrievers:query", "search",
			"--retriever-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--query", "x",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--mode", "routing",
			"--rerank-config.top-n", "1",
			"--rerank-config.type", "system_default",
			"--rerank-top-n", "0",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"query: x\n" +
			"mode: routing\n" +
			"rerank_config:\n" +
			"  top_n: 1\n" +
			"  type: system_default\n" +
			"rerank_top_n: 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"retrievers:query", "search",
			"--retriever-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
