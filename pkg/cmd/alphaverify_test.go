// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/run-llama/llama-parse-cli/internal/mocktest"
	"github.com/run-llama/llama-parse-cli/internal/requestflag"
)

func TestAlphaVerifyCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"alpha:verify", "create",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--configuration", "{target_pages: '1,3,5-7', tier: agentic}",
			"--file-id", "dfl-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
			"--file-input", "dfl-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
			"--transaction-id", "tx-unique-idempotency-key",
			"--webhook-configuration-id", "[whc-..., whc-...]",
			"--webhook-configuration", "[{webhook_events: [parse.success, parse.error], webhook_headers: {Authorization: Bearer sk-...}, webhook_output_format: json, webhook_signing_secret: whsec_..., webhook_url: https://example.com/webhooks/llamacloud}]",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(alphaVerifyCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"alpha:verify", "create",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--configuration.target-pages", "1,3,5-7",
			"--configuration.tier", "agentic",
			"--file-id", "dfl-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
			"--file-input", "dfl-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
			"--transaction-id", "tx-unique-idempotency-key",
			"--webhook-configuration-id", "[whc-..., whc-...]",
			"--webhook-configuration.webhook-events", "[parse.success, parse.error]",
			"--webhook-configuration.webhook-headers", "{Authorization: Bearer sk-...}",
			"--webhook-configuration.webhook-output-format", "json",
			"--webhook-configuration.webhook-signing-secret", "whsec_...",
			"--webhook-configuration.webhook-url", "https://example.com/webhooks/llamacloud",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"configuration:\n" +
			"  target_pages: 1,3,5-7\n" +
			"  tier: agentic\n" +
			"file_id: dfl-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee\n" +
			"file_input: dfl-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee\n" +
			"transaction_id: tx-unique-idempotency-key\n" +
			"webhook_configuration_ids:\n" +
			"  - whc-...\n" +
			"  - whc-...\n" +
			"webhook_configurations:\n" +
			"  - webhook_events:\n" +
			"      - parse.success\n" +
			"      - parse.error\n" +
			"    webhook_headers:\n" +
			"      Authorization: Bearer sk-...\n" +
			"    webhook_output_format: json\n" +
			"    webhook_signing_secret: whsec_...\n" +
			"    webhook_url: https://example.com/webhooks/llamacloud\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"alpha:verify", "create",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAlphaVerifyList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"alpha:verify", "list",
			"--max-items", "10",
			"--created-at-on-or-after", "'2019-12-27T18:11:19.117Z'",
			"--created-at-on-or-before", "'2019-12-27T18:11:19.117Z'",
			"--expand", "string",
			"--job-id", "[string, string]",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--page-size", "1",
			"--page-token", "page_token",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--status", "CANCELLED",
		)
	})
}

func TestAlphaVerifyCancel(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"alpha:verify", "cancel",
			"--job-id", "job_id",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAlphaVerifyGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"alpha:verify", "get",
			"--job-id", "job_id",
			"--expand", "string",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestAlphaVerifyGetDetails(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"alpha:verify", "get-details",
			"--job-id", "job_id",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
