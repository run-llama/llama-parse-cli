// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/run-llama/llama-parse-cli/internal/mocktest"
)

func TestWebhookConfigsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"webhook-configs", "create",
			"--webhook-url", "https://example.com/webhooks/llamacloud",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--webhook-event", "[parse.success, parse.error]",
			"--webhook-headers", "{Authorization: Bearer sk-...}",
			"--webhook-output-format", "json",
			"--webhook-signing-secret", "whsec_...",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"webhook_url: https://example.com/webhooks/llamacloud\n" +
			"webhook_events:\n" +
			"  - parse.success\n" +
			"  - parse.error\n" +
			"webhook_headers:\n" +
			"  Authorization: Bearer sk-...\n" +
			"webhook_output_format: json\n" +
			"webhook_signing_secret: whsec_...\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"webhook-configs", "create",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestWebhookConfigsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"webhook-configs", "retrieve",
			"--config-id", "config_id",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestWebhookConfigsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"webhook-configs", "update",
			"--config-id", "config_id",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--webhook-event", "[batch.cancelled]",
			"--webhook-headers", "{foo: string}",
			"--webhook-output-format", "json",
			"--webhook-signing-secret", "webhook_signing_secret",
			"--webhook-url", "webhook_url",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"webhook_events:\n" +
			"  - batch.cancelled\n" +
			"webhook_headers:\n" +
			"  foo: string\n" +
			"webhook_output_format: json\n" +
			"webhook_signing_secret: webhook_signing_secret\n" +
			"webhook_url: webhook_url\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"webhook-configs", "update",
			"--config-id", "config_id",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestWebhookConfigsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"webhook-configs", "list",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestWebhookConfigsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"webhook-configs", "delete",
			"--config-id", "config_id",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
