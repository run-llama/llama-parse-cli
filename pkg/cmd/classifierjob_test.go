// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/llamacloud-prod-cli/internal/mocktest"
	"github.com/stainless-sdks/llamacloud-prod-cli/internal/requestflag"
)

func TestClassifierJobsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"classifier:jobs", "create",
			"--file-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--rule", "{description: 'contains invoice number, line items, and total amount', type: invoice}",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--mode", "FAST",
			"--parsing-configuration", "{lang: af, max_pages: 0, target_pages: [0]}",
			"--webhook-configuration", "{webhook_events: [parse.success, parse.error], webhook_headers: {foo: bar}, webhook_output_format: json, webhook_url: 'https:'}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(classifierJobsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"classifier:jobs", "create",
			"--file-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--rule.description", "contains invoice number, line items, and total amount",
			"--rule.type", "invoice",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--mode", "FAST",
			"--parsing-configuration.lang", "af",
			"--parsing-configuration.max-pages", "0",
			"--parsing-configuration.target-pages", "[0]",
			"--webhook-configuration.webhook-events", "[parse.success, parse.error]",
			"--webhook-configuration.webhook-headers", "{foo: bar}",
			"--webhook-configuration.webhook-output-format", "json",
			"--webhook-configuration.webhook-url", "https:",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"file_ids:\n" +
			"  - 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"rules:\n" +
			"  - description: contains invoice number, line items, and total amount\n" +
			"    type: invoice\n" +
			"mode: FAST\n" +
			"parsing_configuration:\n" +
			"  lang: af\n" +
			"  max_pages: 0\n" +
			"  target_pages:\n" +
			"    - 0\n" +
			"webhook_configurations:\n" +
			"  - webhook_events:\n" +
			"      - parse.success\n" +
			"      - parse.error\n" +
			"    webhook_headers:\n" +
			"      foo: bar\n" +
			"    webhook_output_format: json\n" +
			"    webhook_url: 'https:'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"classifier:jobs", "create",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestClassifierJobsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"classifier:jobs", "list",
			"--max-items", "10",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--page-size", "0",
			"--page-token", "page_token",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestClassifierJobsGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"classifier:jobs", "get",
			"--classify-job-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestClassifierJobsGetResults(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"classifier:jobs", "get-results",
			"--classify-job-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
