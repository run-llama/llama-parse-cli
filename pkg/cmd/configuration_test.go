// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/run-llama/llama-parse-cli/internal/mocktest"
)

func TestConfigurationsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"configurations", "create",
			"--name", "x",
			"--parameters", "{product_type: classify_v2, rules: [{description: 'contains invoice number, line items, and total amount', type: invoice}], mode: FAST, parsing_configuration: {lang: en, max_pages: 10, target_pages: '1,3,5-7'}}",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: x\n" +
			"parameters:\n" +
			"  product_type: classify_v2\n" +
			"  rules:\n" +
			"    - description: contains invoice number, line items, and total amount\n" +
			"      type: invoice\n" +
			"  mode: FAST\n" +
			"  parsing_configuration:\n" +
			"    lang: en\n" +
			"    max_pages: 10\n" +
			"    target_pages: 1,3,5-7\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"configurations", "create",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestConfigurationsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"configurations", "retrieve",
			"--config-id", "config_id",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestConfigurationsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"configurations", "update",
			"--config-id", "config_id",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--name", "x",
			"--parameters", "{product_type: classify_v2, rules: [{description: 'contains invoice number, line items, and total amount', type: invoice}], mode: FAST, parsing_configuration: {lang: en, max_pages: 10, target_pages: '1,3,5-7'}}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: x\n" +
			"parameters:\n" +
			"  product_type: classify_v2\n" +
			"  rules:\n" +
			"    - description: contains invoice number, line items, and total amount\n" +
			"      type: invoice\n" +
			"  mode: FAST\n" +
			"  parsing_configuration:\n" +
			"    lang: en\n" +
			"    max_pages: 10\n" +
			"    target_pages: 1,3,5-7\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"configurations", "update",
			"--config-id", "config_id",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestConfigurationsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"configurations", "list",
			"--max-items", "10",
			"--latest-only=true",
			"--name", "name",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--page-size", "0",
			"--page-token", "page_token",
			"--product-type", "[classify_v2, extract_v2]",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestConfigurationsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"configurations", "delete",
			"--config-id", "config_id",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
