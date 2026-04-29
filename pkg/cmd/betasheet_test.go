// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/llamacloud-prod-cli/internal/mocktest"
	"github.com/stainless-sdks/llamacloud-prod-cli/internal/requestflag"
)

func TestBetaSheetsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"beta:sheets", "create",
			"--file-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--config", "{extraction_range: extraction_range, flatten_hierarchical_tables: true, generate_additional_metadata: true, include_hidden_cells: true, sheet_names: [string], specialization: specialization, table_merge_sensitivity: strong, use_experimental_processing: true}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(betaSheetsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"beta:sheets", "create",
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
			"--config.use-experimental-processing=true",
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
			"  use_experimental_processing: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"beta:sheets", "create",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestBetaSheetsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"beta:sheets", "list",
			"--max-items", "10",
			"--created-at-on-or-after", "'2019-12-27T18:11:19.117Z'",
			"--created-at-on-or-before", "'2019-12-27T18:11:19.117Z'",
			"--include-results=true",
			"--job-id", "[string, string]",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--page-size", "0",
			"--page-token", "page_token",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--status", "PENDING",
		)
	})
}

func TestBetaSheetsDeleteJob(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"beta:sheets", "delete-job",
			"--spreadsheet-job-id", "spreadsheet_job_id",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestBetaSheetsGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"beta:sheets", "get",
			"--spreadsheet-job-id", "spreadsheet_job_id",
			"--include-results=true",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestBetaSheetsGetResultTable(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"beta:sheets", "get-result-table",
			"--spreadsheet-job-id", "spreadsheet_job_id",
			"--region-id", "region_id",
			"--region-type", "table",
			"--expires-at-seconds", "0",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
