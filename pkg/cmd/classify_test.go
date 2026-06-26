// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/run-llama/llama-parse-cli/internal/mocktest"
	"github.com/run-llama/llama-parse-cli/internal/requestflag"
)

func TestClassifyCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"classify", "create",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--configuration", "{rules: [{description: 'contains invoice number, line items, and total amount', type: invoice}], mode: FAST, parsing_configuration: {lang: en, max_pages: 10, target_pages: '1,3,5-7'}}",
			"--configuration-id", "cfg-11111111-2222-3333-4444-555555555555",
			"--file-id", "dfl-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
			"--file-input", "dfl-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
			"--parse-job-id", "pjb-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
			"--transaction-id", "tx-unique-idempotency-key",
			"--webhook-configuration", "[{webhook_events: [parse.success, parse.error], webhook_headers: {Authorization: Bearer sk-...}, webhook_output_format: json, webhook_url: https://example.com/webhooks/llamacloud}]",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(classifyCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"classify", "create",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--configuration.rules", "[{description: 'contains invoice number, line items, and total amount', type: invoice}]",
			"--configuration.mode", "FAST",
			"--configuration.parsing-configuration", "{lang: en, max_pages: 10, target_pages: '1,3,5-7'}",
			"--configuration-id", "cfg-11111111-2222-3333-4444-555555555555",
			"--file-id", "dfl-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
			"--file-input", "dfl-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
			"--parse-job-id", "pjb-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
			"--transaction-id", "tx-unique-idempotency-key",
			"--webhook-configuration.webhook-events", "[parse.success, parse.error]",
			"--webhook-configuration.webhook-headers", "{Authorization: Bearer sk-...}",
			"--webhook-configuration.webhook-output-format", "json",
			"--webhook-configuration.webhook-url", "https://example.com/webhooks/llamacloud",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"configuration:\n" +
			"  rules:\n" +
			"    - description: contains invoice number, line items, and total amount\n" +
			"      type: invoice\n" +
			"  mode: FAST\n" +
			"  parsing_configuration:\n" +
			"    lang: en\n" +
			"    max_pages: 10\n" +
			"    target_pages: 1,3,5-7\n" +
			"configuration_id: cfg-11111111-2222-3333-4444-555555555555\n" +
			"file_id: dfl-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee\n" +
			"file_input: dfl-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee\n" +
			"parse_job_id: pjb-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee\n" +
			"transaction_id: tx-unique-idempotency-key\n" +
			"webhook_configurations:\n" +
			"  - webhook_events:\n" +
			"      - parse.success\n" +
			"      - parse.error\n" +
			"    webhook_headers:\n" +
			"      Authorization: Bearer sk-...\n" +
			"    webhook_output_format: json\n" +
			"    webhook_url: https://example.com/webhooks/llamacloud\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"classify", "create",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestClassifyList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"classify", "list",
			"--max-items", "10",
			"--configuration-id", "cfg-11111111-2222-3333-4444-555555555555",
			"--created-at-on-or-after", "'2019-12-27T18:11:19.117Z'",
			"--created-at-on-or-before", "'2019-12-27T18:11:19.117Z'",
			"--job-id", "[string, string]",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--page-size", "1",
			"--page-token", "page_token",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--status", "PENDING",
		)
	})
}

func TestClassifyGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"classify", "get",
			"--job-id", "job_id",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
