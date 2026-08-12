// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/run-llama/llama-parse-cli/internal/mocktest"
	"github.com/run-llama/llama-parse-cli/internal/requestflag"
)

func TestSplitCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"split", "create",
			"--file-input", "dfl-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--configuration", "{categories: [{name: x, description: x}], splitting_strategy: {allow_uncategorized: forbid}}",
			"--configuration-id", "cfg-11111111-2222-3333-4444-555555555555",
			"--transaction-id", "tx-unique-idempotency-key",
			"--webhook-configuration-id", "[whc-..., whc-...]",
			"--webhook-configuration", "[{webhook_events: [parse.success, parse.error], webhook_headers: {Authorization: Bearer sk-...}, webhook_output_format: json, webhook_signing_secret: whsec_..., webhook_url: https://example.com/webhooks/llamacloud}]",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(splitCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"split", "create",
			"--file-input", "dfl-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--configuration.categories", "[{name: x, description: x}]",
			"--configuration.splitting-strategy", "{allow_uncategorized: forbid}",
			"--configuration-id", "cfg-11111111-2222-3333-4444-555555555555",
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
			"file_input: dfl-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee\n" +
			"configuration:\n" +
			"  categories:\n" +
			"    - name: x\n" +
			"      description: x\n" +
			"  splitting_strategy:\n" +
			"    allow_uncategorized: forbid\n" +
			"configuration_id: cfg-11111111-2222-3333-4444-555555555555\n" +
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
			"split", "create",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestSplitList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"split", "list",
			"--max-items", "10",
			"--created-at-on-or-after", "'2019-12-27T18:11:19.117Z'",
			"--created-at-on-or-before", "'2019-12-27T18:11:19.117Z'",
			"--job-id", "[string, string]",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--page-size", "0",
			"--page-token", "page_token",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--status", "cancelled",
		)
	})
}

func TestSplitDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"split", "delete",
			"--split-job-id", "split_job_id",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestSplitCancel(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"split", "cancel",
			"--split-job-id", "split_job_id",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestSplitGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"split", "get",
			"--split-job-id", "split_job_id",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
