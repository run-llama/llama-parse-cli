// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/run-llama/llama-parse-cli/internal/mocktest"
	"github.com/run-llama/llama-parse-cli/internal/requestflag"
)

func TestSheetsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sheets", "create",
			"--file-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--config", "{extraction_range: extraction_range, flatten_hierarchical_tables: true, generate_additional_metadata: true, include_hidden_cells: true, sheet_names: [string], specialization: specialization, table_merge_sensitivity: strong, tier: agentic, use_experimental_processing: true}",
			"--configuration", "{extraction_range: extraction_range, flatten_hierarchical_tables: true, generate_additional_metadata: true, include_hidden_cells: true, sheet_names: [string], specialization: specialization, table_merge_sensitivity: strong, tier: agentic, use_experimental_processing: true}",
			"--configuration-id", "cfg-11111111-2222-3333-4444-555555555555",
			"--webhook-configuration", "[{webhook_events: [parse.success, parse.error], webhook_headers: {Authorization: Bearer sk-...}, webhook_output_format: json, webhook_signing_secret: whsec_..., webhook_url: https://example.com/webhooks/llamacloud}]",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(sheetsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sheets", "create",
			"--file-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--config.extraction-range", "extraction_range",
			"--config.flatten-hierarchical-tables=true",
			"--config.generate-additional-metadata=true",
			"--config.include-hidden-cells=true",
			"--config.sheet-names", "[string]",
			"--config.specialization", "specialization",
			"--config.table-merge-sensitivity", "strong",
			"--config.tier", "agentic",
			"--config.use-experimental-processing=true",
			"--configuration.extraction-range", "extraction_range",
			"--configuration.flatten-hierarchical-tables=true",
			"--configuration.generate-additional-metadata=true",
			"--configuration.include-hidden-cells=true",
			"--configuration.sheet-names", "[string]",
			"--configuration.specialization", "specialization",
			"--configuration.table-merge-sensitivity", "strong",
			"--configuration.tier", "agentic",
			"--configuration.use-experimental-processing=true",
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
			"file_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"config:\n" +
			"  extraction_range: extraction_range\n" +
			"  flatten_hierarchical_tables: true\n" +
			"  generate_additional_metadata: true\n" +
			"  include_hidden_cells: true\n" +
			"  sheet_names:\n" +
			"    - string\n" +
			"  specialization: specialization\n" +
			"  table_merge_sensitivity: strong\n" +
			"  tier: agentic\n" +
			"  use_experimental_processing: true\n" +
			"configuration:\n" +
			"  extraction_range: extraction_range\n" +
			"  flatten_hierarchical_tables: true\n" +
			"  generate_additional_metadata: true\n" +
			"  include_hidden_cells: true\n" +
			"  sheet_names:\n" +
			"    - string\n" +
			"  specialization: specialization\n" +
			"  table_merge_sensitivity: strong\n" +
			"  tier: agentic\n" +
			"  use_experimental_processing: true\n" +
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
			"sheets", "create",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestSheetsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sheets", "list",
			"--max-items", "10",
			"--configuration-id", "configuration_id",
			"--created-at-on-or-after", "'2019-12-27T18:11:19.117Z'",
			"--created-at-on-or-before", "'2019-12-27T18:11:19.117Z'",
			"--include-results=true",
			"--job-id", "[string, string]",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--page-size", "0",
			"--page-token", "page_token",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--status", "CANCELLED",
		)
	})
}

func TestSheetsDeleteJob(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sheets", "delete-job",
			"--spreadsheet-job-id", "spreadsheet_job_id",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestSheetsGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sheets", "get",
			"--spreadsheet-job-id", "spreadsheet_job_id",
			"--expand", "string",
			"--include-results=true",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestSheetsGetResultTable(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sheets", "get-result-table",
			"--spreadsheet-job-id", "spreadsheet_job_id",
			"--region-id", "region_id",
			"--region-type", "cell_metadata",
			"--expires-at-seconds", "0",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
