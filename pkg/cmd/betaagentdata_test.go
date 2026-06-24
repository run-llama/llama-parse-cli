// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/run-llama/llama-parse-cli/internal/mocktest"
)

func TestBetaAgentDataCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"beta:agent-data", "create",
			"--data", "{foo: bar}",
			"--deployment-name", "deployment_name",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--collection", "collection",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"data:\n" +
			"  foo: bar\n" +
			"deployment_name: deployment_name\n" +
			"collection: collection\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"beta:agent-data", "create",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestBetaAgentDataUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"beta:agent-data", "update",
			"--item-id", "item_id",
			"--data", "{foo: bar}",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"data:\n" +
			"  foo: bar\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"beta:agent-data", "update",
			"--item-id", "item_id",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestBetaAgentDataDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"beta:agent-data", "delete",
			"--item-id", "item_id",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestBetaAgentDataAggregate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"beta:agent-data", "aggregate",
			"--max-items", "10",
			"--deployment-name", "deployment_name",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--collection", "collection",
			"--count=true",
			"--filter", "{foo: {eq: 0, excludes: [0], gt: 0, gte: 0, includes: [0], lt: 0, lte: 0, ne: 0}}",
			"--first=true",
			"--group-by", "[string]",
			"--offset", "0",
			"--order-by", "order_by",
			"--page-size", "0",
			"--page-token", "page_token",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"deployment_name: deployment_name\n" +
			"collection: collection\n" +
			"count: true\n" +
			"filter:\n" +
			"  foo:\n" +
			"    eq: 0\n" +
			"    excludes:\n" +
			"      - 0\n" +
			"    gt: 0\n" +
			"    gte: 0\n" +
			"    includes:\n" +
			"      - 0\n" +
			"    lt: 0\n" +
			"    lte: 0\n" +
			"    ne: 0\n" +
			"first: true\n" +
			"group_by:\n" +
			"  - string\n" +
			"offset: 0\n" +
			"order_by: order_by\n" +
			"page_size: 0\n" +
			"page_token: page_token\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"beta:agent-data", "aggregate",
			"--max-items", "10",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestBetaAgentDataDeleteByQuery(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"beta:agent-data", "delete-by-query",
			"--deployment-name", "deployment_name",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--collection", "collection",
			"--filter", "{foo: {eq: 0, excludes: [0], gt: 0, gte: 0, includes: [0], lt: 0, lte: 0, ne: 0}}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"deployment_name: deployment_name\n" +
			"collection: collection\n" +
			"filter:\n" +
			"  foo:\n" +
			"    eq: 0\n" +
			"    excludes:\n" +
			"      - 0\n" +
			"    gt: 0\n" +
			"    gte: 0\n" +
			"    includes:\n" +
			"      - 0\n" +
			"    lt: 0\n" +
			"    lte: 0\n" +
			"    ne: 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"beta:agent-data", "delete-by-query",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestBetaAgentDataGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"beta:agent-data", "get",
			"--item-id", "item_id",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestBetaAgentDataSearch(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"beta:agent-data", "search",
			"--max-items", "10",
			"--deployment-name", "deployment_name",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--collection", "collection",
			"--filter", "{foo: {eq: 0, excludes: [0], gt: 0, gte: 0, includes: [0], lt: 0, lte: 0, ne: 0}}",
			"--include-total=true",
			"--offset", "0",
			"--order-by", "order_by",
			"--page-size", "0",
			"--page-token", "page_token",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"deployment_name: deployment_name\n" +
			"collection: collection\n" +
			"filter:\n" +
			"  foo:\n" +
			"    eq: 0\n" +
			"    excludes:\n" +
			"      - 0\n" +
			"    gt: 0\n" +
			"    gte: 0\n" +
			"    includes:\n" +
			"      - 0\n" +
			"    lt: 0\n" +
			"    lte: 0\n" +
			"    ne: 0\n" +
			"include_total: true\n" +
			"offset: 0\n" +
			"order_by: order_by\n" +
			"page_size: 0\n" +
			"page_token: page_token\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"beta:agent-data", "search",
			"--max-items", "10",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
