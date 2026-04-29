// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/llamacloud-prod-cli/internal/mocktest"
)

func TestBetaBatchCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"beta:batch", "create",
			"--job-config", "{correlation_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, job_name: parse_raw_file_job, parameters: {adaptive_long_table: true, aggressive_table_extraction: true, annotate_links: true, auto_mode: true, auto_mode_configuration_json: auto_mode_configuration_json, auto_mode_trigger_on_image_in_page: true, auto_mode_trigger_on_regexp_in_page: auto_mode_trigger_on_regexp_in_page, auto_mode_trigger_on_table_in_page: true, auto_mode_trigger_on_text_in_page: auto_mode_trigger_on_text_in_page, azure_openai_api_version: azure_openai_api_version, azure_openai_deployment_name: azure_openai_deployment_name, azure_openai_endpoint: azure_openai_endpoint, azure_openai_key: azure_openai_key, bbox_bottom: 0, bbox_left: 0, bbox_right: 0, bbox_top: 0, bounding_box: bounding_box, compact_markdown_table: true, complemental_formatting_instruction: complemental_formatting_instruction, content_guideline_instruction: content_guideline_instruction, continuous_mode: true, custom_metadata: {foo: bar}, disable_image_extraction: true, disable_ocr: true, disable_reconstruction: true, do_not_cache: true, do_not_unroll_columns: true, enable_cost_optimizer: true, extract_charts: true, extract_layout: true, extract_printed_page_number: true, fast_mode: true, formatting_instruction: formatting_instruction, gpt4o_api_key: gpt4o_api_key, gpt4o_mode: true, guess_xlsx_sheet_name: true, hide_footers: true, hide_headers: true, high_res_ocr: true, html_make_all_elements_visible: true, html_remove_fixed_elements: true, html_remove_navigation_elements: true, http_proxy: http_proxy, ignore_document_elements_for_layout_detection: true, images_to_save: [screenshot], inline_images_in_markdown: true, input_s3_path: input_s3_path, input_s3_region: input_s3_region, input_url: input_url, internal_is_screenshot_job: true, invalidate_cache: true, is_formatting_instruction: true, job_timeout_extra_time_per_page_in_seconds: 0, job_timeout_in_seconds: 0, keep_page_separator_when_merging_tables: true, lang: lang, languages: [af], layout_aware: true, line_level_bounding_box: true, markdown_table_multiline_header_separator: markdown_table_multiline_header_separator, max_pages: 0, max_pages_enforced: 0, merge_tables_across_pages_in_markdown: true, model: model, outlined_table_extraction: true, output_pdf_of_document: true, output_s3_path_prefix: output_s3_path_prefix, output_s3_region: output_s3_region, output_tables_as_HTML: true, outputBucket: outputBucket, page_error_tolerance: 0, page_footer_prefix: page_footer_prefix, page_footer_suffix: page_footer_suffix, page_header_prefix: page_header_prefix, page_header_suffix: page_header_suffix, page_prefix: page_prefix, page_separator: page_separator, page_suffix: page_suffix, parse_mode: parse_page_without_llm, parsing_instruction: parsing_instruction, pipeline_id: pipeline_id, precise_bounding_box: true, premium_mode: true, presentation_out_of_bounds_content: true, presentation_skip_embedded_data: true, preserve_layout_alignment_across_pages: true, preserve_very_small_text: true, preset: preset, priority: low, project_id: project_id, remove_hidden_text: true, replace_failed_page_mode: raw_text, replace_failed_page_with_error_message_prefix: replace_failed_page_with_error_message_prefix, replace_failed_page_with_error_message_suffix: replace_failed_page_with_error_message_suffix, resource_info: {foo: bar}, save_images: true, skip_diagonal_text: true, specialized_chart_parsing_agentic: true, specialized_chart_parsing_efficient: true, specialized_chart_parsing_plus: true, specialized_image_parsing: true, spreadsheet_extract_sub_tables: true, spreadsheet_force_formula_computation: true, spreadsheet_include_hidden_sheets: true, strict_mode_buggy_font: true, strict_mode_image_extraction: true, strict_mode_image_ocr: true, strict_mode_reconstruction: true, structured_output: true, structured_output_json_schema: structured_output_json_schema, structured_output_json_schema_name: structured_output_json_schema_name, system_prompt: system_prompt, system_prompt_append: system_prompt_append, take_screenshot: true, target_pages: target_pages, tier: tier, type: parse, use_vendor_multimodal_model: true, user_prompt: user_prompt, vendor_multimodal_api_key: vendor_multimodal_api_key, vendor_multimodal_model_name: vendor_multimodal_model_name, version: version, webhook_configurations: [{webhook_events: [parse.success, parse.error], webhook_headers: {Authorization: Bearer sk-...}, webhook_output_format: json, webhook_url: https://example.com/webhooks/llamacloud}], webhook_url: webhook_url}, parent_job_execution_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, partitions: {foo: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e}, project_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, session_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, user_id: user_id, webhook_url: webhook_url}",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--continue-as-new-threshold", "0",
			"--directory-id", "dir-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
			"--item-id", "[dfl-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee, dfl-11111111-2222-3333-4444-555555555555]",
			"--page-size", "1",
			"--temporal-namespace", "temporal-namespace",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"job_config:\n" +
			"  correlation_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"  job_name: parse_raw_file_job\n" +
			"  parameters:\n" +
			"    adaptive_long_table: true\n" +
			"    aggressive_table_extraction: true\n" +
			"    annotate_links: true\n" +
			"    auto_mode: true\n" +
			"    auto_mode_configuration_json: auto_mode_configuration_json\n" +
			"    auto_mode_trigger_on_image_in_page: true\n" +
			"    auto_mode_trigger_on_regexp_in_page: auto_mode_trigger_on_regexp_in_page\n" +
			"    auto_mode_trigger_on_table_in_page: true\n" +
			"    auto_mode_trigger_on_text_in_page: auto_mode_trigger_on_text_in_page\n" +
			"    azure_openai_api_version: azure_openai_api_version\n" +
			"    azure_openai_deployment_name: azure_openai_deployment_name\n" +
			"    azure_openai_endpoint: azure_openai_endpoint\n" +
			"    azure_openai_key: azure_openai_key\n" +
			"    bbox_bottom: 0\n" +
			"    bbox_left: 0\n" +
			"    bbox_right: 0\n" +
			"    bbox_top: 0\n" +
			"    bounding_box: bounding_box\n" +
			"    compact_markdown_table: true\n" +
			"    complemental_formatting_instruction: complemental_formatting_instruction\n" +
			"    content_guideline_instruction: content_guideline_instruction\n" +
			"    continuous_mode: true\n" +
			"    custom_metadata:\n" +
			"      foo: bar\n" +
			"    disable_image_extraction: true\n" +
			"    disable_ocr: true\n" +
			"    disable_reconstruction: true\n" +
			"    do_not_cache: true\n" +
			"    do_not_unroll_columns: true\n" +
			"    enable_cost_optimizer: true\n" +
			"    extract_charts: true\n" +
			"    extract_layout: true\n" +
			"    extract_printed_page_number: true\n" +
			"    fast_mode: true\n" +
			"    formatting_instruction: formatting_instruction\n" +
			"    gpt4o_api_key: gpt4o_api_key\n" +
			"    gpt4o_mode: true\n" +
			"    guess_xlsx_sheet_name: true\n" +
			"    hide_footers: true\n" +
			"    hide_headers: true\n" +
			"    high_res_ocr: true\n" +
			"    html_make_all_elements_visible: true\n" +
			"    html_remove_fixed_elements: true\n" +
			"    html_remove_navigation_elements: true\n" +
			"    http_proxy: http_proxy\n" +
			"    ignore_document_elements_for_layout_detection: true\n" +
			"    images_to_save:\n" +
			"      - screenshot\n" +
			"    inline_images_in_markdown: true\n" +
			"    input_s3_path: input_s3_path\n" +
			"    input_s3_region: input_s3_region\n" +
			"    input_url: input_url\n" +
			"    internal_is_screenshot_job: true\n" +
			"    invalidate_cache: true\n" +
			"    is_formatting_instruction: true\n" +
			"    job_timeout_extra_time_per_page_in_seconds: 0\n" +
			"    job_timeout_in_seconds: 0\n" +
			"    keep_page_separator_when_merging_tables: true\n" +
			"    lang: lang\n" +
			"    languages:\n" +
			"      - af\n" +
			"    layout_aware: true\n" +
			"    line_level_bounding_box: true\n" +
			"    markdown_table_multiline_header_separator: markdown_table_multiline_header_separator\n" +
			"    max_pages: 0\n" +
			"    max_pages_enforced: 0\n" +
			"    merge_tables_across_pages_in_markdown: true\n" +
			"    model: model\n" +
			"    outlined_table_extraction: true\n" +
			"    output_pdf_of_document: true\n" +
			"    output_s3_path_prefix: output_s3_path_prefix\n" +
			"    output_s3_region: output_s3_region\n" +
			"    output_tables_as_HTML: true\n" +
			"    outputBucket: outputBucket\n" +
			"    page_error_tolerance: 0\n" +
			"    page_footer_prefix: page_footer_prefix\n" +
			"    page_footer_suffix: page_footer_suffix\n" +
			"    page_header_prefix: page_header_prefix\n" +
			"    page_header_suffix: page_header_suffix\n" +
			"    page_prefix: page_prefix\n" +
			"    page_separator: page_separator\n" +
			"    page_suffix: page_suffix\n" +
			"    parse_mode: parse_page_without_llm\n" +
			"    parsing_instruction: parsing_instruction\n" +
			"    pipeline_id: pipeline_id\n" +
			"    precise_bounding_box: true\n" +
			"    premium_mode: true\n" +
			"    presentation_out_of_bounds_content: true\n" +
			"    presentation_skip_embedded_data: true\n" +
			"    preserve_layout_alignment_across_pages: true\n" +
			"    preserve_very_small_text: true\n" +
			"    preset: preset\n" +
			"    priority: low\n" +
			"    project_id: project_id\n" +
			"    remove_hidden_text: true\n" +
			"    replace_failed_page_mode: raw_text\n" +
			"    replace_failed_page_with_error_message_prefix: replace_failed_page_with_error_message_prefix\n" +
			"    replace_failed_page_with_error_message_suffix: replace_failed_page_with_error_message_suffix\n" +
			"    resource_info:\n" +
			"      foo: bar\n" +
			"    save_images: true\n" +
			"    skip_diagonal_text: true\n" +
			"    specialized_chart_parsing_agentic: true\n" +
			"    specialized_chart_parsing_efficient: true\n" +
			"    specialized_chart_parsing_plus: true\n" +
			"    specialized_image_parsing: true\n" +
			"    spreadsheet_extract_sub_tables: true\n" +
			"    spreadsheet_force_formula_computation: true\n" +
			"    spreadsheet_include_hidden_sheets: true\n" +
			"    strict_mode_buggy_font: true\n" +
			"    strict_mode_image_extraction: true\n" +
			"    strict_mode_image_ocr: true\n" +
			"    strict_mode_reconstruction: true\n" +
			"    structured_output: true\n" +
			"    structured_output_json_schema: structured_output_json_schema\n" +
			"    structured_output_json_schema_name: structured_output_json_schema_name\n" +
			"    system_prompt: system_prompt\n" +
			"    system_prompt_append: system_prompt_append\n" +
			"    take_screenshot: true\n" +
			"    target_pages: target_pages\n" +
			"    tier: tier\n" +
			"    type: parse\n" +
			"    use_vendor_multimodal_model: true\n" +
			"    user_prompt: user_prompt\n" +
			"    vendor_multimodal_api_key: vendor_multimodal_api_key\n" +
			"    vendor_multimodal_model_name: vendor_multimodal_model_name\n" +
			"    version: version\n" +
			"    webhook_configurations:\n" +
			"      - webhook_events:\n" +
			"          - parse.success\n" +
			"          - parse.error\n" +
			"        webhook_headers:\n" +
			"          Authorization: Bearer sk-...\n" +
			"        webhook_output_format: json\n" +
			"        webhook_url: https://example.com/webhooks/llamacloud\n" +
			"    webhook_url: webhook_url\n" +
			"  parent_job_execution_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"  partitions:\n" +
			"    foo: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"  project_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"  session_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"  user_id: user_id\n" +
			"  webhook_url: webhook_url\n" +
			"continue_as_new_threshold: 0\n" +
			"directory_id: dir-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee\n" +
			"item_ids:\n" +
			"  - dfl-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee\n" +
			"  - dfl-11111111-2222-3333-4444-555555555555\n" +
			"page_size: 1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"beta:batch", "create",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--temporal-namespace", "temporal-namespace",
		)
	})
}

func TestBetaBatchList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"beta:batch", "list",
			"--max-items", "10",
			"--directory-id", "directory_id",
			"--job-type", "parse",
			"--limit", "1",
			"--offset", "0",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--status", "pending",
		)
	})
}

func TestBetaBatchCancel(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"beta:batch", "cancel",
			"--job-id", "job_id",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--reason", "reason",
			"--temporal-namespace", "temporal-namespace",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("reason: reason")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"beta:batch", "cancel",
			"--job-id", "job_id",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--temporal-namespace", "temporal-namespace",
		)
	})
}

func TestBetaBatchGetStatus(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"beta:batch", "get-status",
			"--job-id", "job_id",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
