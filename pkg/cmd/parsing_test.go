// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/llamacloud-prod-cli/internal/mocktest"
	"github.com/stainless-sdks/llamacloud-prod-cli/internal/requestflag"
)

func TestParsingCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"parsing", "create",
			"--tier", "fast",
			"--version", "latest",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--agentic-options", "{custom_prompt: custom_prompt}",
			"--client-name", "client_name",
			"--crop-box", "{bottom: 0, left: 0, right: 0, top: 0}",
			"--disable-cache=true",
			"--fast-options", "{}",
			"--file-id", "file_id",
			"--http-proxy", "https:",
			"--input-options", "{html: {make_all_elements_visible: true, remove_fixed_elements: true, remove_navigation_elements: true}, pdf: {}, presentation: {out_of_bounds_content: true, skip_embedded_data: true}, spreadsheet: {detect_sub_tables_in_sheets: true, force_formula_computation_in_sheets: true, include_hidden_sheets: true}}",
			"--output-options", "{additional_outputs: [stripped_md, concatenated_stripped_txt, word_bbox], extract_printed_page_number: true, images_to_save: [screenshot], markdown: {annotate_links: true, inline_images: true, tables: {compact_markdown_tables: true, markdown_table_multiline_separator: markdown_table_multiline_separator, merge_continued_tables: true, output_tables_as_markdown: true}}, spatial_text: {do_not_unroll_columns: true, preserve_layout_alignment_across_pages: true, preserve_very_small_text: true}, tables_as_spreadsheet: {enable: true, guess_sheet_name: true}}",
			"--page-ranges", "{max_pages: 1, target_pages: target_pages}",
			"--processing-control", "{job_failure_conditions: {allowed_page_failure_ratio: 1, fail_on_buggy_font: true, fail_on_image_extraction_error: true, fail_on_image_ocr_error: true, fail_on_markdown_reconstruction_error: true}, timeouts: {base_in_seconds: 1, extra_time_per_page_in_seconds: 1}}",
			"--processing-options", "{aggressive_table_extraction: true, auto_mode_configuration: [{parsing_conf: {adaptive_long_table: true, aggressive_table_extraction: true, crop_box: {bottom: 0, left: 0, right: 0, top: 0}, custom_prompt: custom_prompt, extract_layout: true, high_res_ocr: true, ignore: {ignore_diagonal_text: true, ignore_hidden_text: true}, language: language, outlined_table_extraction: true, presentation: {out_of_bounds_content: true, skip_embedded_data: true}, spatial_text: {do_not_unroll_columns: true, preserve_layout_alignment_across_pages: true, preserve_very_small_text: true}, specialized_chart_parsing: agentic_plus, tier: agentic, version: latest}, filename_match_glob: '*.txt', filename_match_glob_list: [string], filename_regexp: filename_regexp, filename_regexp_mode: filename_regexp_mode, full_page_image_in_page: true, full_page_image_in_page_threshold: 0, image_in_page: true, layout_element_in_page: layout_element_in_page, layout_element_in_page_confidence_threshold: 0, page_contains_at_least_n_charts: 0, page_contains_at_least_n_images: 0, page_contains_at_least_n_layout_elements: 0, page_contains_at_least_n_lines: 0, page_contains_at_least_n_links: 0, page_contains_at_least_n_numbers: 0, page_contains_at_least_n_percent_numbers: 0, page_contains_at_least_n_tables: 0, page_contains_at_least_n_words: 0, page_contains_at_most_n_charts: 0, page_contains_at_most_n_images: 0, page_contains_at_most_n_layout_elements: 0, page_contains_at_most_n_lines: 0, page_contains_at_most_n_links: 0, page_contains_at_most_n_numbers: 0, page_contains_at_most_n_percent_numbers: 0, page_contains_at_most_n_tables: 0, page_contains_at_most_n_words: 0, page_longer_than_n_chars: 0, page_md_error: true, page_shorter_than_n_chars: 0, regexp_in_page: regexp_in_page, regexp_in_page_mode: regexp_in_page_mode, table_in_page: true, text_in_page: text_in_page, trigger_mode: trigger_mode}], cost_optimizer: {enable: true}, disable_heuristics: true, ignore: {ignore_diagonal_text: true, ignore_hidden_text: true, ignore_text_in_image: true}, ocr_parameters: {languages: [af]}, specialized_chart_parsing: agentic_plus}",
			"--source-url", "https:",
			"--webhook-configuration", "{webhook_events: [string], webhook_headers: {foo: bar}, webhook_output_format: json, webhook_url: 'https:'}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(parsingCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"parsing", "create",
			"--tier", "fast",
			"--version", "latest",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--agentic-options.custom-prompt", "custom_prompt",
			"--client-name", "client_name",
			"--crop-box.bottom", "0",
			"--crop-box.left", "0",
			"--crop-box.right", "0",
			"--crop-box.top", "0",
			"--disable-cache=true",
			"--fast-options", "{}",
			"--file-id", "file_id",
			"--http-proxy", "https:",
			"--input-options.html", "{make_all_elements_visible: true, remove_fixed_elements: true, remove_navigation_elements: true}",
			"--input-options.pdf", "{}",
			"--input-options.presentation", "{out_of_bounds_content: true, skip_embedded_data: true}",
			"--input-options.spreadsheet", "{detect_sub_tables_in_sheets: true, force_formula_computation_in_sheets: true, include_hidden_sheets: true}",
			"--output-options.additional-outputs", "[stripped_md, concatenated_stripped_txt, word_bbox]",
			"--output-options.extract-printed-page-number=true",
			"--output-options.images-to-save", "[screenshot]",
			"--output-options.markdown", "{annotate_links: true, inline_images: true, tables: {compact_markdown_tables: true, markdown_table_multiline_separator: markdown_table_multiline_separator, merge_continued_tables: true, output_tables_as_markdown: true}}",
			"--output-options.spatial-text", "{do_not_unroll_columns: true, preserve_layout_alignment_across_pages: true, preserve_very_small_text: true}",
			"--output-options.tables-as-spreadsheet", "{enable: true, guess_sheet_name: true}",
			"--page-ranges.max-pages", "1",
			"--page-ranges.target-pages", "target_pages",
			"--processing-control.job-failure-conditions", "{allowed_page_failure_ratio: 1, fail_on_buggy_font: true, fail_on_image_extraction_error: true, fail_on_image_ocr_error: true, fail_on_markdown_reconstruction_error: true}",
			"--processing-control.timeouts", "{base_in_seconds: 1, extra_time_per_page_in_seconds: 1}",
			"--processing-options.aggressive-table-extraction=true",
			"--processing-options.auto-mode-configuration", "[{parsing_conf: {adaptive_long_table: true, aggressive_table_extraction: true, crop_box: {bottom: 0, left: 0, right: 0, top: 0}, custom_prompt: custom_prompt, extract_layout: true, high_res_ocr: true, ignore: {ignore_diagonal_text: true, ignore_hidden_text: true}, language: language, outlined_table_extraction: true, presentation: {out_of_bounds_content: true, skip_embedded_data: true}, spatial_text: {do_not_unroll_columns: true, preserve_layout_alignment_across_pages: true, preserve_very_small_text: true}, specialized_chart_parsing: agentic_plus, tier: agentic, version: latest}, filename_match_glob: '*.txt', filename_match_glob_list: [string], filename_regexp: filename_regexp, filename_regexp_mode: filename_regexp_mode, full_page_image_in_page: true, full_page_image_in_page_threshold: 0, image_in_page: true, layout_element_in_page: layout_element_in_page, layout_element_in_page_confidence_threshold: 0, page_contains_at_least_n_charts: 0, page_contains_at_least_n_images: 0, page_contains_at_least_n_layout_elements: 0, page_contains_at_least_n_lines: 0, page_contains_at_least_n_links: 0, page_contains_at_least_n_numbers: 0, page_contains_at_least_n_percent_numbers: 0, page_contains_at_least_n_tables: 0, page_contains_at_least_n_words: 0, page_contains_at_most_n_charts: 0, page_contains_at_most_n_images: 0, page_contains_at_most_n_layout_elements: 0, page_contains_at_most_n_lines: 0, page_contains_at_most_n_links: 0, page_contains_at_most_n_numbers: 0, page_contains_at_most_n_percent_numbers: 0, page_contains_at_most_n_tables: 0, page_contains_at_most_n_words: 0, page_longer_than_n_chars: 0, page_md_error: true, page_shorter_than_n_chars: 0, regexp_in_page: regexp_in_page, regexp_in_page_mode: regexp_in_page_mode, table_in_page: true, text_in_page: text_in_page, trigger_mode: trigger_mode}]",
			"--processing-options.cost-optimizer", "{enable: true}",
			"--processing-options.disable-heuristics=true",
			"--processing-options.ignore", "{ignore_diagonal_text: true, ignore_hidden_text: true, ignore_text_in_image: true}",
			"--processing-options.ocr-parameters", "{languages: [af]}",
			"--processing-options.specialized-chart-parsing", "agentic_plus",
			"--source-url", "https:",
			"--webhook-configuration.webhook-events", "[string]",
			"--webhook-configuration.webhook-headers", "{foo: bar}",
			"--webhook-configuration.webhook-output-format", "json",
			"--webhook-configuration.webhook-url", "https:",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"tier: fast\n" +
			"version: latest\n" +
			"agentic_options:\n" +
			"  custom_prompt: custom_prompt\n" +
			"client_name: client_name\n" +
			"crop_box:\n" +
			"  bottom: 0\n" +
			"  left: 0\n" +
			"  right: 0\n" +
			"  top: 0\n" +
			"disable_cache: true\n" +
			"fast_options: {}\n" +
			"file_id: file_id\n" +
			"http_proxy: 'https:'\n" +
			"input_options:\n" +
			"  html:\n" +
			"    make_all_elements_visible: true\n" +
			"    remove_fixed_elements: true\n" +
			"    remove_navigation_elements: true\n" +
			"  pdf: {}\n" +
			"  presentation:\n" +
			"    out_of_bounds_content: true\n" +
			"    skip_embedded_data: true\n" +
			"  spreadsheet:\n" +
			"    detect_sub_tables_in_sheets: true\n" +
			"    force_formula_computation_in_sheets: true\n" +
			"    include_hidden_sheets: true\n" +
			"output_options:\n" +
			"  additional_outputs:\n" +
			"    - stripped_md\n" +
			"    - concatenated_stripped_txt\n" +
			"    - word_bbox\n" +
			"  extract_printed_page_number: true\n" +
			"  images_to_save:\n" +
			"    - screenshot\n" +
			"  markdown:\n" +
			"    annotate_links: true\n" +
			"    inline_images: true\n" +
			"    tables:\n" +
			"      compact_markdown_tables: true\n" +
			"      markdown_table_multiline_separator: markdown_table_multiline_separator\n" +
			"      merge_continued_tables: true\n" +
			"      output_tables_as_markdown: true\n" +
			"  spatial_text:\n" +
			"    do_not_unroll_columns: true\n" +
			"    preserve_layout_alignment_across_pages: true\n" +
			"    preserve_very_small_text: true\n" +
			"  tables_as_spreadsheet:\n" +
			"    enable: true\n" +
			"    guess_sheet_name: true\n" +
			"page_ranges:\n" +
			"  max_pages: 1\n" +
			"  target_pages: target_pages\n" +
			"processing_control:\n" +
			"  job_failure_conditions:\n" +
			"    allowed_page_failure_ratio: 1\n" +
			"    fail_on_buggy_font: true\n" +
			"    fail_on_image_extraction_error: true\n" +
			"    fail_on_image_ocr_error: true\n" +
			"    fail_on_markdown_reconstruction_error: true\n" +
			"  timeouts:\n" +
			"    base_in_seconds: 1\n" +
			"    extra_time_per_page_in_seconds: 1\n" +
			"processing_options:\n" +
			"  aggressive_table_extraction: true\n" +
			"  auto_mode_configuration:\n" +
			"    - parsing_conf:\n" +
			"        adaptive_long_table: true\n" +
			"        aggressive_table_extraction: true\n" +
			"        crop_box:\n" +
			"          bottom: 0\n" +
			"          left: 0\n" +
			"          right: 0\n" +
			"          top: 0\n" +
			"        custom_prompt: custom_prompt\n" +
			"        extract_layout: true\n" +
			"        high_res_ocr: true\n" +
			"        ignore:\n" +
			"          ignore_diagonal_text: true\n" +
			"          ignore_hidden_text: true\n" +
			"        language: language\n" +
			"        outlined_table_extraction: true\n" +
			"        presentation:\n" +
			"          out_of_bounds_content: true\n" +
			"          skip_embedded_data: true\n" +
			"        spatial_text:\n" +
			"          do_not_unroll_columns: true\n" +
			"          preserve_layout_alignment_across_pages: true\n" +
			"          preserve_very_small_text: true\n" +
			"        specialized_chart_parsing: agentic_plus\n" +
			"        tier: agentic\n" +
			"        version: latest\n" +
			"      filename_match_glob: '*.txt'\n" +
			"      filename_match_glob_list:\n" +
			"        - string\n" +
			"      filename_regexp: filename_regexp\n" +
			"      filename_regexp_mode: filename_regexp_mode\n" +
			"      full_page_image_in_page: true\n" +
			"      full_page_image_in_page_threshold: 0\n" +
			"      image_in_page: true\n" +
			"      layout_element_in_page: layout_element_in_page\n" +
			"      layout_element_in_page_confidence_threshold: 0\n" +
			"      page_contains_at_least_n_charts: 0\n" +
			"      page_contains_at_least_n_images: 0\n" +
			"      page_contains_at_least_n_layout_elements: 0\n" +
			"      page_contains_at_least_n_lines: 0\n" +
			"      page_contains_at_least_n_links: 0\n" +
			"      page_contains_at_least_n_numbers: 0\n" +
			"      page_contains_at_least_n_percent_numbers: 0\n" +
			"      page_contains_at_least_n_tables: 0\n" +
			"      page_contains_at_least_n_words: 0\n" +
			"      page_contains_at_most_n_charts: 0\n" +
			"      page_contains_at_most_n_images: 0\n" +
			"      page_contains_at_most_n_layout_elements: 0\n" +
			"      page_contains_at_most_n_lines: 0\n" +
			"      page_contains_at_most_n_links: 0\n" +
			"      page_contains_at_most_n_numbers: 0\n" +
			"      page_contains_at_most_n_percent_numbers: 0\n" +
			"      page_contains_at_most_n_tables: 0\n" +
			"      page_contains_at_most_n_words: 0\n" +
			"      page_longer_than_n_chars: 0\n" +
			"      page_md_error: true\n" +
			"      page_shorter_than_n_chars: 0\n" +
			"      regexp_in_page: regexp_in_page\n" +
			"      regexp_in_page_mode: regexp_in_page_mode\n" +
			"      table_in_page: true\n" +
			"      text_in_page: text_in_page\n" +
			"      trigger_mode: trigger_mode\n" +
			"  cost_optimizer:\n" +
			"    enable: true\n" +
			"  disable_heuristics: true\n" +
			"  ignore:\n" +
			"    ignore_diagonal_text: true\n" +
			"    ignore_hidden_text: true\n" +
			"    ignore_text_in_image: true\n" +
			"  ocr_parameters:\n" +
			"    languages:\n" +
			"      - af\n" +
			"  specialized_chart_parsing: agentic_plus\n" +
			"source_url: 'https:'\n" +
			"webhook_configurations:\n" +
			"  - webhook_events:\n" +
			"      - string\n" +
			"    webhook_headers:\n" +
			"      foo: bar\n" +
			"    webhook_output_format: json\n" +
			"    webhook_url: 'https:'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"parsing", "create",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestParsingList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"parsing", "list",
			"--max-items", "10",
			"--created-at-on-or-after", "'2019-12-27T18:11:19.117Z'",
			"--created-at-on-or-before", "'2019-12-27T18:11:19.117Z'",
			"--job-id", "[string, string]",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--page-size", "0",
			"--page-token", "page_token",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--status", "PENDING",
		)
	})
}

func TestParsingGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"parsing", "get",
			"--job-id", "job_id",
			"--expand", "string",
			"--image-filenames", "image_filenames",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
