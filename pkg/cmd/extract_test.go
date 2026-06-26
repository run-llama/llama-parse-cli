// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/run-llama/llama-parse-cli/internal/mocktest"
	"github.com/run-llama/llama-parse-cli/internal/requestflag"
)

func TestExtractCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"extract", "create",
			"--file-input", "dfl-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--configuration", "{data_schema: {properties: {vendor_name: bar, total_amount: bar}, required: [vendor_name, total_amount], type: object}, cite_sources: true, confidence_scores: true, extraction_target: per_doc, max_pages: 10, parse_config_id: cfg-11111111-2222-3333-4444-555555555555, parse_tier: fast, system_prompt: 'Extract all monetary values in USD. If a currency is not specified, assume USD.', target_pages: '1,3,5-7', tier: cost_effective, version: latest}",
			"--configuration-id", "cfg-11111111-2222-3333-4444-555555555555",
			"--webhook-configuration", "[{webhook_events: [parse.success, parse.error], webhook_headers: {Authorization: Bearer sk-...}, webhook_output_format: json, webhook_signing_secret: whsec_..., webhook_url: https://example.com/webhooks/llamacloud}]",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(extractCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"extract", "create",
			"--file-input", "dfl-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--configuration.data-schema", "{properties: {vendor_name: bar, total_amount: bar}, required: [vendor_name, total_amount], type: object}",
			"--configuration.cite-sources=true",
			"--configuration.confidence-scores=true",
			"--configuration.extraction-target", "per_doc",
			"--configuration.max-pages", "10",
			"--configuration.parse-config-id", "cfg-11111111-2222-3333-4444-555555555555",
			"--configuration.parse-tier", "fast",
			"--configuration.system-prompt", "Extract all monetary values in USD. If a currency is not specified, assume USD.",
			"--configuration.target-pages", "1,3,5-7",
			"--configuration.tier", "cost_effective",
			"--configuration.version", "latest",
			"--configuration-id", "cfg-11111111-2222-3333-4444-555555555555",
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
			"  data_schema:\n" +
			"    properties:\n" +
			"      vendor_name: bar\n" +
			"      total_amount: bar\n" +
			"    required:\n" +
			"      - vendor_name\n" +
			"      - total_amount\n" +
			"    type: object\n" +
			"  cite_sources: true\n" +
			"  confidence_scores: true\n" +
			"  extraction_target: per_doc\n" +
			"  max_pages: 10\n" +
			"  parse_config_id: cfg-11111111-2222-3333-4444-555555555555\n" +
			"  parse_tier: fast\n" +
			"  system_prompt: >-\n" +
			"    Extract all monetary values in USD. If a currency is not specified, assume\n" +
			"    USD.\n" +
			"  target_pages: 1,3,5-7\n" +
			"  tier: cost_effective\n" +
			"  version: latest\n" +
			"configuration_id: cfg-11111111-2222-3333-4444-555555555555\n" +
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
			"extract", "create",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestExtractList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"extract", "list",
			"--max-items", "10",
			"--configuration-id", "cfg-11111111-2222-3333-4444-555555555555",
			"--created-at-on-or-after", "'2019-12-27T18:11:19.117Z'",
			"--created-at-on-or-before", "'2019-12-27T18:11:19.117Z'",
			"--document-input-type", "document_input_type",
			"--document-input-value", "document_input_value",
			"--expand", "string",
			"--file-input", "file_input",
			"--job-id", "[string, string]",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--page-size", "0",
			"--page-token", "page_token",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--status", "PENDING",
		)
	})
}

func TestExtractDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"extract", "delete",
			"--job-id", "job_id",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestExtractGenerateSchema(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"extract", "generate-schema",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--data-schema", "{foo: {foo: bar}}",
			"--file-id", "dfl-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
			"--name", "invoice_extraction",
			"--prompt", "Extract vendor name, invoice number, date, line items with descriptions and amounts, and total amount from invoices.",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"data_schema:\n" +
			"  foo:\n" +
			"    foo: bar\n" +
			"file_id: dfl-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee\n" +
			"name: invoice_extraction\n" +
			"prompt: >-\n" +
			"  Extract vendor name, invoice number, date, line items with descriptions and\n" +
			"  amounts, and total amount from invoices.\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"extract", "generate-schema",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestExtractGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"extract", "get",
			"--job-id", "job_id",
			"--expand", "string",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestExtractValidateSchema(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"extract", "validate-schema",
			"--data-schema", "{properties: {vendor_name: bar, invoice_number: bar, total_amount: bar, line_items: bar}, required: [vendor_name, invoice_number, total_amount], type: object}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"data_schema:\n" +
			"  properties:\n" +
			"    vendor_name: bar\n" +
			"    invoice_number: bar\n" +
			"    total_amount: bar\n" +
			"    line_items: bar\n" +
			"  required:\n" +
			"    - vendor_name\n" +
			"    - invoice_number\n" +
			"    - total_amount\n" +
			"  type: object\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"extract", "validate-schema",
		)
	})
}
