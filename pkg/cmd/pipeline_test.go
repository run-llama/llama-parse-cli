// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/run-llama/llama-parse-cli/internal/mocktest"
	"github.com/run-llama/llama-parse-cli/internal/requestflag"
)

func TestPipelinesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pipelines", "create",
			"--name", "x",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--data-sink", "{component: {foo: bar}, name: name, sink_type: PINECONE}",
			"--data-sink-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--embedding-config", "{component: {additional_kwargs: {foo: bar}, api_base: api_base, api_key: api_key, api_version: api_version, azure_deployment: azure_deployment, azure_endpoint: azure_endpoint, class_name: class_name, default_headers: {foo: string}, dimensions: 0, embed_batch_size: 1, max_retries: 0, model_name: model_name, num_workers: 0, reuse_client: true, timeout: 0}, type: AZURE_EMBEDDING}",
			"--embedding-model-config-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--llama-parse-parameters", "{adaptive_long_table: true, aggressive_table_extraction: true, annotate_links: true, auto_mode: true, auto_mode_configuration_json: auto_mode_configuration_json, auto_mode_trigger_on_image_in_page: true, auto_mode_trigger_on_regexp_in_page: auto_mode_trigger_on_regexp_in_page, auto_mode_trigger_on_table_in_page: true, auto_mode_trigger_on_text_in_page: auto_mode_trigger_on_text_in_page, azure_openai_api_version: azure_openai_api_version, azure_openai_deployment_name: azure_openai_deployment_name, azure_openai_endpoint: azure_openai_endpoint, azure_openai_key: azure_openai_key, bbox_bottom: 0, bbox_left: 0, bbox_right: 0, bbox_top: 0, bounding_box: bounding_box, compact_markdown_table: true, complemental_formatting_instruction: complemental_formatting_instruction, content_guideline_instruction: content_guideline_instruction, continuous_mode: true, disable_image_extraction: true, disable_ocr: true, disable_reconstruction: true, do_not_cache: true, do_not_unroll_columns: true, enable_cost_optimizer: true, extract_charts: true, extract_layout: true, extract_printed_page_number: true, fast_mode: true, formatting_instruction: formatting_instruction, gpt4o_api_key: gpt4o_api_key, gpt4o_mode: true, guess_xlsx_sheet_name: true, hide_footers: true, hide_headers: true, high_res_ocr: true, html_make_all_elements_visible: true, html_remove_fixed_elements: true, html_remove_navigation_elements: true, http_proxy: http_proxy, ignore_document_elements_for_layout_detection: true, images_to_save: [screenshot], inline_images_in_markdown: true, input_s3_path: input_s3_path, input_s3_region: input_s3_region, input_url: input_url, internal_is_screenshot_job: true, invalidate_cache: true, is_formatting_instruction: true, job_timeout_extra_time_per_page_in_seconds: 0, job_timeout_in_seconds: 0, keep_page_separator_when_merging_tables: true, languages: [af], layout_aware: true, line_level_bounding_box: true, markdown_table_multiline_header_separator: markdown_table_multiline_header_separator, max_pages: 0, max_pages_enforced: 0, merge_tables_across_pages_in_markdown: true, model: model, outlined_table_extraction: true, output_pdf_of_document: true, output_s3_path_prefix: output_s3_path_prefix, output_s3_region: output_s3_region, output_tables_as_HTML: true, page_error_tolerance: 0, page_footer_prefix: page_footer_prefix, page_footer_suffix: page_footer_suffix, page_header_prefix: page_header_prefix, page_header_suffix: page_header_suffix, page_prefix: page_prefix, page_separator: page_separator, page_suffix: page_suffix, parse_mode: parse_page_without_llm, parsing_instruction: parsing_instruction, precise_bounding_box: true, premium_mode: true, presentation_out_of_bounds_content: true, presentation_skip_embedded_data: true, preserve_layout_alignment_across_pages: true, preserve_very_small_text: true, preset: preset, priority: low, project_id: project_id, remove_hidden_text: true, replace_failed_page_mode: raw_text, replace_failed_page_with_error_message_prefix: replace_failed_page_with_error_message_prefix, replace_failed_page_with_error_message_suffix: replace_failed_page_with_error_message_suffix, save_images: true, skip_diagonal_text: true, specialized_chart_parsing_agentic: true, specialized_chart_parsing_efficient: true, specialized_chart_parsing_plus: true, specialized_image_parsing: true, spreadsheet_extract_sub_tables: true, spreadsheet_force_formula_computation: true, spreadsheet_include_hidden_sheets: true, strict_mode_buggy_font: true, strict_mode_image_extraction: true, strict_mode_image_ocr: true, strict_mode_reconstruction: true, structured_output: true, structured_output_json_schema: structured_output_json_schema, structured_output_json_schema_name: structured_output_json_schema_name, system_prompt: system_prompt, system_prompt_append: system_prompt_append, take_screenshot: true, target_pages: target_pages, tier: tier, use_vendor_multimodal_model: true, user_prompt: user_prompt, vendor_multimodal_api_key: vendor_multimodal_api_key, vendor_multimodal_model_name: vendor_multimodal_model_name, version: version, webhook_configurations: [{webhook_events: [parse.success, parse.error], webhook_headers: {Authorization: Bearer sk-...}, webhook_output_format: json, webhook_url: https://example.com/webhooks/llamacloud}], webhook_url: webhook_url}",
			"--managed-pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--metadata-config", "{excluded_embed_metadata_keys: [string], excluded_llm_metadata_keys: [string]}",
			"--pipeline-type", "PLAYGROUND",
			"--preset-retrieval-parameters", "{alpha: 0, class_name: class_name, dense_similarity_cutoff: 0, dense_similarity_top_k: 1, enable_reranking: true, files_top_k: 1, rerank_top_n: 1, retrieval_mode: chunks, retrieve_image_nodes: true, retrieve_page_figure_nodes: true, retrieve_page_screenshot_nodes: true, search_filters: {filters: [{key: key, value: 0, operator: '=='}], condition: and}, search_filters_inference_schema: {foo: {foo: bar}}, sparse_similarity_top_k: 1}",
			"--sparse-model-config", "{class_name: class_name, model_type: splade}",
			"--status", "status",
			"--transform-config", "{chunk_overlap: 0, chunk_size: 1, mode: auto}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(pipelinesCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pipelines", "create",
			"--name", "x",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--data-sink.component", "{foo: bar}",
			"--data-sink.name", "name",
			"--data-sink.sink-type", "PINECONE",
			"--data-sink-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--embedding-config", "{component: {additional_kwargs: {foo: bar}, api_base: api_base, api_key: api_key, api_version: api_version, azure_deployment: azure_deployment, azure_endpoint: azure_endpoint, class_name: class_name, default_headers: {foo: string}, dimensions: 0, embed_batch_size: 1, max_retries: 0, model_name: model_name, num_workers: 0, reuse_client: true, timeout: 0}, type: AZURE_EMBEDDING}",
			"--embedding-model-config-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--llama-parse-parameters.adaptive-long-table=true",
			"--llama-parse-parameters.aggressive-table-extraction=true",
			"--llama-parse-parameters.annotate-links=true",
			"--llama-parse-parameters.auto-mode=true",
			"--llama-parse-parameters.auto-mode-configuration-json", "auto_mode_configuration_json",
			"--llama-parse-parameters.auto-mode-trigger-on-image-in-page=true",
			"--llama-parse-parameters.auto-mode-trigger-on-regexp-in-page", "auto_mode_trigger_on_regexp_in_page",
			"--llama-parse-parameters.auto-mode-trigger-on-table-in-page=true",
			"--llama-parse-parameters.auto-mode-trigger-on-text-in-page", "auto_mode_trigger_on_text_in_page",
			"--llama-parse-parameters.azure-openai-api-version", "azure_openai_api_version",
			"--llama-parse-parameters.azure-openai-deployment-name", "azure_openai_deployment_name",
			"--llama-parse-parameters.azure-openai-endpoint", "azure_openai_endpoint",
			"--llama-parse-parameters.azure-openai-key", "azure_openai_key",
			"--llama-parse-parameters.bbox-bottom", "0",
			"--llama-parse-parameters.bbox-left", "0",
			"--llama-parse-parameters.bbox-right", "0",
			"--llama-parse-parameters.bbox-top", "0",
			"--llama-parse-parameters.bounding-box", "bounding_box",
			"--llama-parse-parameters.compact-markdown-table=true",
			"--llama-parse-parameters.complemental-formatting-instruction", "complemental_formatting_instruction",
			"--llama-parse-parameters.content-guideline-instruction", "content_guideline_instruction",
			"--llama-parse-parameters.continuous-mode=true",
			"--llama-parse-parameters.disable-image-extraction=true",
			"--llama-parse-parameters.disable-ocr=true",
			"--llama-parse-parameters.disable-reconstruction=true",
			"--llama-parse-parameters.do-not-cache=true",
			"--llama-parse-parameters.do-not-unroll-columns=true",
			"--llama-parse-parameters.enable-cost-optimizer=true",
			"--llama-parse-parameters.extract-charts=true",
			"--llama-parse-parameters.extract-layout=true",
			"--llama-parse-parameters.extract-printed-page-number=true",
			"--llama-parse-parameters.fast-mode=true",
			"--llama-parse-parameters.formatting-instruction", "formatting_instruction",
			"--llama-parse-parameters.gpt4o-api-key", "gpt4o_api_key",
			"--llama-parse-parameters.gpt4o-mode=true",
			"--llama-parse-parameters.guess-xlsx-sheet-name=true",
			"--llama-parse-parameters.hide-footers=true",
			"--llama-parse-parameters.hide-headers=true",
			"--llama-parse-parameters.high-res-ocr=true",
			"--llama-parse-parameters.html-make-all-elements-visible=true",
			"--llama-parse-parameters.html-remove-fixed-elements=true",
			"--llama-parse-parameters.html-remove-navigation-elements=true",
			"--llama-parse-parameters.http-proxy", "http_proxy",
			"--llama-parse-parameters.ignore-document-elements-for-layout-detection=true",
			"--llama-parse-parameters.images-to-save", "[screenshot]",
			"--llama-parse-parameters.inline-images-in-markdown=true",
			"--llama-parse-parameters.input-s3-path", "input_s3_path",
			"--llama-parse-parameters.input-s3-region", "input_s3_region",
			"--llama-parse-parameters.input-url", "input_url",
			"--llama-parse-parameters.internal-is-screenshot-job=true",
			"--llama-parse-parameters.invalidate-cache=true",
			"--llama-parse-parameters.is-formatting-instruction=true",
			"--llama-parse-parameters.job-timeout-extra-time-per-page-in-seconds", "0",
			"--llama-parse-parameters.job-timeout-in-seconds", "0",
			"--llama-parse-parameters.keep-page-separator-when-merging-tables=true",
			"--llama-parse-parameters.languages", "[af]",
			"--llama-parse-parameters.layout-aware=true",
			"--llama-parse-parameters.line-level-bounding-box=true",
			"--llama-parse-parameters.markdown-table-multiline-header-separator", "markdown_table_multiline_header_separator",
			"--llama-parse-parameters.max-pages", "0",
			"--llama-parse-parameters.max-pages-enforced", "0",
			"--llama-parse-parameters.merge-tables-across-pages-in-markdown=true",
			"--llama-parse-parameters.model", "model",
			"--llama-parse-parameters.outlined-table-extraction=true",
			"--llama-parse-parameters.output-pdf-of-document=true",
			"--llama-parse-parameters.output-s3-path-prefix", "output_s3_path_prefix",
			"--llama-parse-parameters.output-s3-region", "output_s3_region",
			"--llama-parse-parameters.output-tables-as-html=true",
			"--llama-parse-parameters.page-error-tolerance", "0",
			"--llama-parse-parameters.page-footer-prefix", "page_footer_prefix",
			"--llama-parse-parameters.page-footer-suffix", "page_footer_suffix",
			"--llama-parse-parameters.page-header-prefix", "page_header_prefix",
			"--llama-parse-parameters.page-header-suffix", "page_header_suffix",
			"--llama-parse-parameters.page-prefix", "page_prefix",
			"--llama-parse-parameters.page-separator", "page_separator",
			"--llama-parse-parameters.page-suffix", "page_suffix",
			"--llama-parse-parameters.parse-mode", "parse_page_without_llm",
			"--llama-parse-parameters.parsing-instruction", "parsing_instruction",
			"--llama-parse-parameters.precise-bounding-box=true",
			"--llama-parse-parameters.premium-mode=true",
			"--llama-parse-parameters.presentation-out-of-bounds-content=true",
			"--llama-parse-parameters.presentation-skip-embedded-data=true",
			"--llama-parse-parameters.preserve-layout-alignment-across-pages=true",
			"--llama-parse-parameters.preserve-very-small-text=true",
			"--llama-parse-parameters.preset", "preset",
			"--llama-parse-parameters.priority", "low",
			"--llama-parse-parameters.project-id", "project_id",
			"--llama-parse-parameters.remove-hidden-text=true",
			"--llama-parse-parameters.replace-failed-page-mode", "raw_text",
			"--llama-parse-parameters.replace-failed-page-with-error-message-prefix", "replace_failed_page_with_error_message_prefix",
			"--llama-parse-parameters.replace-failed-page-with-error-message-suffix", "replace_failed_page_with_error_message_suffix",
			"--llama-parse-parameters.save-images=true",
			"--llama-parse-parameters.skip-diagonal-text=true",
			"--llama-parse-parameters.specialized-chart-parsing-agentic=true",
			"--llama-parse-parameters.specialized-chart-parsing-efficient=true",
			"--llama-parse-parameters.specialized-chart-parsing-plus=true",
			"--llama-parse-parameters.specialized-image-parsing=true",
			"--llama-parse-parameters.spreadsheet-extract-sub-tables=true",
			"--llama-parse-parameters.spreadsheet-force-formula-computation=true",
			"--llama-parse-parameters.spreadsheet-include-hidden-sheets=true",
			"--llama-parse-parameters.strict-mode-buggy-font=true",
			"--llama-parse-parameters.strict-mode-image-extraction=true",
			"--llama-parse-parameters.strict-mode-image-ocr=true",
			"--llama-parse-parameters.strict-mode-reconstruction=true",
			"--llama-parse-parameters.structured-output=true",
			"--llama-parse-parameters.structured-output-json-schema", "structured_output_json_schema",
			"--llama-parse-parameters.structured-output-json-schema-name", "structured_output_json_schema_name",
			"--llama-parse-parameters.system-prompt", "system_prompt",
			"--llama-parse-parameters.system-prompt-append", "system_prompt_append",
			"--llama-parse-parameters.take-screenshot=true",
			"--llama-parse-parameters.target-pages", "target_pages",
			"--llama-parse-parameters.tier", "tier",
			"--llama-parse-parameters.use-vendor-multimodal-model=true",
			"--llama-parse-parameters.user-prompt", "user_prompt",
			"--llama-parse-parameters.vendor-multimodal-api-key", "vendor_multimodal_api_key",
			"--llama-parse-parameters.vendor-multimodal-model-name", "vendor_multimodal_model_name",
			"--llama-parse-parameters.version", "version",
			"--llama-parse-parameters.webhook-configurations", "[{webhook_events: [parse.success, parse.error], webhook_headers: {Authorization: Bearer sk-...}, webhook_output_format: json, webhook_url: https://example.com/webhooks/llamacloud}]",
			"--llama-parse-parameters.webhook-url", "webhook_url",
			"--managed-pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--metadata-config.excluded-embed-metadata-keys", "[string]",
			"--metadata-config.excluded-llm-metadata-keys", "[string]",
			"--pipeline-type", "PLAYGROUND",
			"--preset-retrieval-parameters.alpha", "0",
			"--preset-retrieval-parameters.class-name", "class_name",
			"--preset-retrieval-parameters.dense-similarity-cutoff", "0",
			"--preset-retrieval-parameters.dense-similarity-top-k", "1",
			"--preset-retrieval-parameters.enable-reranking=true",
			"--preset-retrieval-parameters.files-top-k", "1",
			"--preset-retrieval-parameters.rerank-top-n", "1",
			"--preset-retrieval-parameters.retrieval-mode", "chunks",
			"--preset-retrieval-parameters.retrieve-image-nodes=true",
			"--preset-retrieval-parameters.retrieve-page-figure-nodes=true",
			"--preset-retrieval-parameters.retrieve-page-screenshot-nodes=true",
			"--preset-retrieval-parameters.search-filters", "{filters: [{key: key, value: 0, operator: '=='}], condition: and}",
			"--preset-retrieval-parameters.search-filters-inference-schema", "{foo: {foo: bar}}",
			"--preset-retrieval-parameters.sparse-similarity-top-k", "1",
			"--sparse-model-config.class-name", "class_name",
			"--sparse-model-config.model-type", "splade",
			"--status", "status",
			"--transform-config", "{chunk_overlap: 0, chunk_size: 1, mode: auto}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: x\n" +
			"data_sink:\n" +
			"  component:\n" +
			"    foo: bar\n" +
			"  name: name\n" +
			"  sink_type: PINECONE\n" +
			"data_sink_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"embedding_config:\n" +
			"  component:\n" +
			"    additional_kwargs:\n" +
			"      foo: bar\n" +
			"    api_base: api_base\n" +
			"    api_key: api_key\n" +
			"    api_version: api_version\n" +
			"    azure_deployment: azure_deployment\n" +
			"    azure_endpoint: azure_endpoint\n" +
			"    class_name: class_name\n" +
			"    default_headers:\n" +
			"      foo: string\n" +
			"    dimensions: 0\n" +
			"    embed_batch_size: 1\n" +
			"    max_retries: 0\n" +
			"    model_name: model_name\n" +
			"    num_workers: 0\n" +
			"    reuse_client: true\n" +
			"    timeout: 0\n" +
			"  type: AZURE_EMBEDDING\n" +
			"embedding_model_config_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"llama_parse_parameters:\n" +
			"  adaptive_long_table: true\n" +
			"  aggressive_table_extraction: true\n" +
			"  annotate_links: true\n" +
			"  auto_mode: true\n" +
			"  auto_mode_configuration_json: auto_mode_configuration_json\n" +
			"  auto_mode_trigger_on_image_in_page: true\n" +
			"  auto_mode_trigger_on_regexp_in_page: auto_mode_trigger_on_regexp_in_page\n" +
			"  auto_mode_trigger_on_table_in_page: true\n" +
			"  auto_mode_trigger_on_text_in_page: auto_mode_trigger_on_text_in_page\n" +
			"  azure_openai_api_version: azure_openai_api_version\n" +
			"  azure_openai_deployment_name: azure_openai_deployment_name\n" +
			"  azure_openai_endpoint: azure_openai_endpoint\n" +
			"  azure_openai_key: azure_openai_key\n" +
			"  bbox_bottom: 0\n" +
			"  bbox_left: 0\n" +
			"  bbox_right: 0\n" +
			"  bbox_top: 0\n" +
			"  bounding_box: bounding_box\n" +
			"  compact_markdown_table: true\n" +
			"  complemental_formatting_instruction: complemental_formatting_instruction\n" +
			"  content_guideline_instruction: content_guideline_instruction\n" +
			"  continuous_mode: true\n" +
			"  disable_image_extraction: true\n" +
			"  disable_ocr: true\n" +
			"  disable_reconstruction: true\n" +
			"  do_not_cache: true\n" +
			"  do_not_unroll_columns: true\n" +
			"  enable_cost_optimizer: true\n" +
			"  extract_charts: true\n" +
			"  extract_layout: true\n" +
			"  extract_printed_page_number: true\n" +
			"  fast_mode: true\n" +
			"  formatting_instruction: formatting_instruction\n" +
			"  gpt4o_api_key: gpt4o_api_key\n" +
			"  gpt4o_mode: true\n" +
			"  guess_xlsx_sheet_name: true\n" +
			"  hide_footers: true\n" +
			"  hide_headers: true\n" +
			"  high_res_ocr: true\n" +
			"  html_make_all_elements_visible: true\n" +
			"  html_remove_fixed_elements: true\n" +
			"  html_remove_navigation_elements: true\n" +
			"  http_proxy: http_proxy\n" +
			"  ignore_document_elements_for_layout_detection: true\n" +
			"  images_to_save:\n" +
			"    - screenshot\n" +
			"  inline_images_in_markdown: true\n" +
			"  input_s3_path: input_s3_path\n" +
			"  input_s3_region: input_s3_region\n" +
			"  input_url: input_url\n" +
			"  internal_is_screenshot_job: true\n" +
			"  invalidate_cache: true\n" +
			"  is_formatting_instruction: true\n" +
			"  job_timeout_extra_time_per_page_in_seconds: 0\n" +
			"  job_timeout_in_seconds: 0\n" +
			"  keep_page_separator_when_merging_tables: true\n" +
			"  languages:\n" +
			"    - af\n" +
			"  layout_aware: true\n" +
			"  line_level_bounding_box: true\n" +
			"  markdown_table_multiline_header_separator: markdown_table_multiline_header_separator\n" +
			"  max_pages: 0\n" +
			"  max_pages_enforced: 0\n" +
			"  merge_tables_across_pages_in_markdown: true\n" +
			"  model: model\n" +
			"  outlined_table_extraction: true\n" +
			"  output_pdf_of_document: true\n" +
			"  output_s3_path_prefix: output_s3_path_prefix\n" +
			"  output_s3_region: output_s3_region\n" +
			"  output_tables_as_HTML: true\n" +
			"  page_error_tolerance: 0\n" +
			"  page_footer_prefix: page_footer_prefix\n" +
			"  page_footer_suffix: page_footer_suffix\n" +
			"  page_header_prefix: page_header_prefix\n" +
			"  page_header_suffix: page_header_suffix\n" +
			"  page_prefix: page_prefix\n" +
			"  page_separator: page_separator\n" +
			"  page_suffix: page_suffix\n" +
			"  parse_mode: parse_page_without_llm\n" +
			"  parsing_instruction: parsing_instruction\n" +
			"  precise_bounding_box: true\n" +
			"  premium_mode: true\n" +
			"  presentation_out_of_bounds_content: true\n" +
			"  presentation_skip_embedded_data: true\n" +
			"  preserve_layout_alignment_across_pages: true\n" +
			"  preserve_very_small_text: true\n" +
			"  preset: preset\n" +
			"  priority: low\n" +
			"  project_id: project_id\n" +
			"  remove_hidden_text: true\n" +
			"  replace_failed_page_mode: raw_text\n" +
			"  replace_failed_page_with_error_message_prefix: replace_failed_page_with_error_message_prefix\n" +
			"  replace_failed_page_with_error_message_suffix: replace_failed_page_with_error_message_suffix\n" +
			"  save_images: true\n" +
			"  skip_diagonal_text: true\n" +
			"  specialized_chart_parsing_agentic: true\n" +
			"  specialized_chart_parsing_efficient: true\n" +
			"  specialized_chart_parsing_plus: true\n" +
			"  specialized_image_parsing: true\n" +
			"  spreadsheet_extract_sub_tables: true\n" +
			"  spreadsheet_force_formula_computation: true\n" +
			"  spreadsheet_include_hidden_sheets: true\n" +
			"  strict_mode_buggy_font: true\n" +
			"  strict_mode_image_extraction: true\n" +
			"  strict_mode_image_ocr: true\n" +
			"  strict_mode_reconstruction: true\n" +
			"  structured_output: true\n" +
			"  structured_output_json_schema: structured_output_json_schema\n" +
			"  structured_output_json_schema_name: structured_output_json_schema_name\n" +
			"  system_prompt: system_prompt\n" +
			"  system_prompt_append: system_prompt_append\n" +
			"  take_screenshot: true\n" +
			"  target_pages: target_pages\n" +
			"  tier: tier\n" +
			"  use_vendor_multimodal_model: true\n" +
			"  user_prompt: user_prompt\n" +
			"  vendor_multimodal_api_key: vendor_multimodal_api_key\n" +
			"  vendor_multimodal_model_name: vendor_multimodal_model_name\n" +
			"  version: version\n" +
			"  webhook_configurations:\n" +
			"    - webhook_events:\n" +
			"        - parse.success\n" +
			"        - parse.error\n" +
			"      webhook_headers:\n" +
			"        Authorization: Bearer sk-...\n" +
			"      webhook_output_format: json\n" +
			"      webhook_url: https://example.com/webhooks/llamacloud\n" +
			"  webhook_url: webhook_url\n" +
			"managed_pipeline_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"metadata_config:\n" +
			"  excluded_embed_metadata_keys:\n" +
			"    - string\n" +
			"  excluded_llm_metadata_keys:\n" +
			"    - string\n" +
			"pipeline_type: PLAYGROUND\n" +
			"preset_retrieval_parameters:\n" +
			"  alpha: 0\n" +
			"  class_name: class_name\n" +
			"  dense_similarity_cutoff: 0\n" +
			"  dense_similarity_top_k: 1\n" +
			"  enable_reranking: true\n" +
			"  files_top_k: 1\n" +
			"  rerank_top_n: 1\n" +
			"  retrieval_mode: chunks\n" +
			"  retrieve_image_nodes: true\n" +
			"  retrieve_page_figure_nodes: true\n" +
			"  retrieve_page_screenshot_nodes: true\n" +
			"  search_filters:\n" +
			"    filters:\n" +
			"      - key: key\n" +
			"        value: 0\n" +
			"        operator: '=='\n" +
			"    condition: and\n" +
			"  search_filters_inference_schema:\n" +
			"    foo:\n" +
			"      foo: bar\n" +
			"  sparse_similarity_top_k: 1\n" +
			"sparse_model_config:\n" +
			"  class_name: class_name\n" +
			"  model_type: splade\n" +
			"status: status\n" +
			"transform_config:\n" +
			"  chunk_overlap: 0\n" +
			"  chunk_size: 1\n" +
			"  mode: auto\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"pipelines", "create",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPipelinesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pipelines", "update",
			"--pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--data-sink", "{component: {foo: bar}, name: name, sink_type: PINECONE}",
			"--data-sink-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--embedding-config", "{component: {additional_kwargs: {foo: bar}, api_base: api_base, api_key: api_key, api_version: api_version, azure_deployment: azure_deployment, azure_endpoint: azure_endpoint, class_name: class_name, default_headers: {foo: string}, dimensions: 0, embed_batch_size: 1, max_retries: 0, model_name: model_name, num_workers: 0, reuse_client: true, timeout: 0}, type: AZURE_EMBEDDING}",
			"--embedding-model-config-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--llama-parse-parameters", "{adaptive_long_table: true, aggressive_table_extraction: true, annotate_links: true, auto_mode: true, auto_mode_configuration_json: auto_mode_configuration_json, auto_mode_trigger_on_image_in_page: true, auto_mode_trigger_on_regexp_in_page: auto_mode_trigger_on_regexp_in_page, auto_mode_trigger_on_table_in_page: true, auto_mode_trigger_on_text_in_page: auto_mode_trigger_on_text_in_page, azure_openai_api_version: azure_openai_api_version, azure_openai_deployment_name: azure_openai_deployment_name, azure_openai_endpoint: azure_openai_endpoint, azure_openai_key: azure_openai_key, bbox_bottom: 0, bbox_left: 0, bbox_right: 0, bbox_top: 0, bounding_box: bounding_box, compact_markdown_table: true, complemental_formatting_instruction: complemental_formatting_instruction, content_guideline_instruction: content_guideline_instruction, continuous_mode: true, disable_image_extraction: true, disable_ocr: true, disable_reconstruction: true, do_not_cache: true, do_not_unroll_columns: true, enable_cost_optimizer: true, extract_charts: true, extract_layout: true, extract_printed_page_number: true, fast_mode: true, formatting_instruction: formatting_instruction, gpt4o_api_key: gpt4o_api_key, gpt4o_mode: true, guess_xlsx_sheet_name: true, hide_footers: true, hide_headers: true, high_res_ocr: true, html_make_all_elements_visible: true, html_remove_fixed_elements: true, html_remove_navigation_elements: true, http_proxy: http_proxy, ignore_document_elements_for_layout_detection: true, images_to_save: [screenshot], inline_images_in_markdown: true, input_s3_path: input_s3_path, input_s3_region: input_s3_region, input_url: input_url, internal_is_screenshot_job: true, invalidate_cache: true, is_formatting_instruction: true, job_timeout_extra_time_per_page_in_seconds: 0, job_timeout_in_seconds: 0, keep_page_separator_when_merging_tables: true, languages: [af], layout_aware: true, line_level_bounding_box: true, markdown_table_multiline_header_separator: markdown_table_multiline_header_separator, max_pages: 0, max_pages_enforced: 0, merge_tables_across_pages_in_markdown: true, model: model, outlined_table_extraction: true, output_pdf_of_document: true, output_s3_path_prefix: output_s3_path_prefix, output_s3_region: output_s3_region, output_tables_as_HTML: true, page_error_tolerance: 0, page_footer_prefix: page_footer_prefix, page_footer_suffix: page_footer_suffix, page_header_prefix: page_header_prefix, page_header_suffix: page_header_suffix, page_prefix: page_prefix, page_separator: page_separator, page_suffix: page_suffix, parse_mode: parse_page_without_llm, parsing_instruction: parsing_instruction, precise_bounding_box: true, premium_mode: true, presentation_out_of_bounds_content: true, presentation_skip_embedded_data: true, preserve_layout_alignment_across_pages: true, preserve_very_small_text: true, preset: preset, priority: low, project_id: project_id, remove_hidden_text: true, replace_failed_page_mode: raw_text, replace_failed_page_with_error_message_prefix: replace_failed_page_with_error_message_prefix, replace_failed_page_with_error_message_suffix: replace_failed_page_with_error_message_suffix, save_images: true, skip_diagonal_text: true, specialized_chart_parsing_agentic: true, specialized_chart_parsing_efficient: true, specialized_chart_parsing_plus: true, specialized_image_parsing: true, spreadsheet_extract_sub_tables: true, spreadsheet_force_formula_computation: true, spreadsheet_include_hidden_sheets: true, strict_mode_buggy_font: true, strict_mode_image_extraction: true, strict_mode_image_ocr: true, strict_mode_reconstruction: true, structured_output: true, structured_output_json_schema: structured_output_json_schema, structured_output_json_schema_name: structured_output_json_schema_name, system_prompt: system_prompt, system_prompt_append: system_prompt_append, take_screenshot: true, target_pages: target_pages, tier: tier, use_vendor_multimodal_model: true, user_prompt: user_prompt, vendor_multimodal_api_key: vendor_multimodal_api_key, vendor_multimodal_model_name: vendor_multimodal_model_name, version: version, webhook_configurations: [{webhook_events: [parse.success, parse.error], webhook_headers: {Authorization: Bearer sk-...}, webhook_output_format: json, webhook_url: https://example.com/webhooks/llamacloud}], webhook_url: webhook_url}",
			"--managed-pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--metadata-config", "{excluded_embed_metadata_keys: [string], excluded_llm_metadata_keys: [string]}",
			"--name", "name",
			"--preset-retrieval-parameters", "{alpha: 0, class_name: class_name, dense_similarity_cutoff: 0, dense_similarity_top_k: 1, enable_reranking: true, files_top_k: 1, rerank_top_n: 1, retrieval_mode: chunks, retrieve_image_nodes: true, retrieve_page_figure_nodes: true, retrieve_page_screenshot_nodes: true, search_filters: {filters: [{key: key, value: 0, operator: '=='}], condition: and}, search_filters_inference_schema: {foo: {foo: bar}}, sparse_similarity_top_k: 1}",
			"--sparse-model-config", "{class_name: class_name, model_type: splade}",
			"--status", "status",
			"--transform-config", "{chunk_overlap: 0, chunk_size: 1, mode: auto}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(pipelinesUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pipelines", "update",
			"--pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--data-sink.component", "{foo: bar}",
			"--data-sink.name", "name",
			"--data-sink.sink-type", "PINECONE",
			"--data-sink-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--embedding-config", "{component: {additional_kwargs: {foo: bar}, api_base: api_base, api_key: api_key, api_version: api_version, azure_deployment: azure_deployment, azure_endpoint: azure_endpoint, class_name: class_name, default_headers: {foo: string}, dimensions: 0, embed_batch_size: 1, max_retries: 0, model_name: model_name, num_workers: 0, reuse_client: true, timeout: 0}, type: AZURE_EMBEDDING}",
			"--embedding-model-config-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--llama-parse-parameters.adaptive-long-table=true",
			"--llama-parse-parameters.aggressive-table-extraction=true",
			"--llama-parse-parameters.annotate-links=true",
			"--llama-parse-parameters.auto-mode=true",
			"--llama-parse-parameters.auto-mode-configuration-json", "auto_mode_configuration_json",
			"--llama-parse-parameters.auto-mode-trigger-on-image-in-page=true",
			"--llama-parse-parameters.auto-mode-trigger-on-regexp-in-page", "auto_mode_trigger_on_regexp_in_page",
			"--llama-parse-parameters.auto-mode-trigger-on-table-in-page=true",
			"--llama-parse-parameters.auto-mode-trigger-on-text-in-page", "auto_mode_trigger_on_text_in_page",
			"--llama-parse-parameters.azure-openai-api-version", "azure_openai_api_version",
			"--llama-parse-parameters.azure-openai-deployment-name", "azure_openai_deployment_name",
			"--llama-parse-parameters.azure-openai-endpoint", "azure_openai_endpoint",
			"--llama-parse-parameters.azure-openai-key", "azure_openai_key",
			"--llama-parse-parameters.bbox-bottom", "0",
			"--llama-parse-parameters.bbox-left", "0",
			"--llama-parse-parameters.bbox-right", "0",
			"--llama-parse-parameters.bbox-top", "0",
			"--llama-parse-parameters.bounding-box", "bounding_box",
			"--llama-parse-parameters.compact-markdown-table=true",
			"--llama-parse-parameters.complemental-formatting-instruction", "complemental_formatting_instruction",
			"--llama-parse-parameters.content-guideline-instruction", "content_guideline_instruction",
			"--llama-parse-parameters.continuous-mode=true",
			"--llama-parse-parameters.disable-image-extraction=true",
			"--llama-parse-parameters.disable-ocr=true",
			"--llama-parse-parameters.disable-reconstruction=true",
			"--llama-parse-parameters.do-not-cache=true",
			"--llama-parse-parameters.do-not-unroll-columns=true",
			"--llama-parse-parameters.enable-cost-optimizer=true",
			"--llama-parse-parameters.extract-charts=true",
			"--llama-parse-parameters.extract-layout=true",
			"--llama-parse-parameters.extract-printed-page-number=true",
			"--llama-parse-parameters.fast-mode=true",
			"--llama-parse-parameters.formatting-instruction", "formatting_instruction",
			"--llama-parse-parameters.gpt4o-api-key", "gpt4o_api_key",
			"--llama-parse-parameters.gpt4o-mode=true",
			"--llama-parse-parameters.guess-xlsx-sheet-name=true",
			"--llama-parse-parameters.hide-footers=true",
			"--llama-parse-parameters.hide-headers=true",
			"--llama-parse-parameters.high-res-ocr=true",
			"--llama-parse-parameters.html-make-all-elements-visible=true",
			"--llama-parse-parameters.html-remove-fixed-elements=true",
			"--llama-parse-parameters.html-remove-navigation-elements=true",
			"--llama-parse-parameters.http-proxy", "http_proxy",
			"--llama-parse-parameters.ignore-document-elements-for-layout-detection=true",
			"--llama-parse-parameters.images-to-save", "[screenshot]",
			"--llama-parse-parameters.inline-images-in-markdown=true",
			"--llama-parse-parameters.input-s3-path", "input_s3_path",
			"--llama-parse-parameters.input-s3-region", "input_s3_region",
			"--llama-parse-parameters.input-url", "input_url",
			"--llama-parse-parameters.internal-is-screenshot-job=true",
			"--llama-parse-parameters.invalidate-cache=true",
			"--llama-parse-parameters.is-formatting-instruction=true",
			"--llama-parse-parameters.job-timeout-extra-time-per-page-in-seconds", "0",
			"--llama-parse-parameters.job-timeout-in-seconds", "0",
			"--llama-parse-parameters.keep-page-separator-when-merging-tables=true",
			"--llama-parse-parameters.languages", "[af]",
			"--llama-parse-parameters.layout-aware=true",
			"--llama-parse-parameters.line-level-bounding-box=true",
			"--llama-parse-parameters.markdown-table-multiline-header-separator", "markdown_table_multiline_header_separator",
			"--llama-parse-parameters.max-pages", "0",
			"--llama-parse-parameters.max-pages-enforced", "0",
			"--llama-parse-parameters.merge-tables-across-pages-in-markdown=true",
			"--llama-parse-parameters.model", "model",
			"--llama-parse-parameters.outlined-table-extraction=true",
			"--llama-parse-parameters.output-pdf-of-document=true",
			"--llama-parse-parameters.output-s3-path-prefix", "output_s3_path_prefix",
			"--llama-parse-parameters.output-s3-region", "output_s3_region",
			"--llama-parse-parameters.output-tables-as-html=true",
			"--llama-parse-parameters.page-error-tolerance", "0",
			"--llama-parse-parameters.page-footer-prefix", "page_footer_prefix",
			"--llama-parse-parameters.page-footer-suffix", "page_footer_suffix",
			"--llama-parse-parameters.page-header-prefix", "page_header_prefix",
			"--llama-parse-parameters.page-header-suffix", "page_header_suffix",
			"--llama-parse-parameters.page-prefix", "page_prefix",
			"--llama-parse-parameters.page-separator", "page_separator",
			"--llama-parse-parameters.page-suffix", "page_suffix",
			"--llama-parse-parameters.parse-mode", "parse_page_without_llm",
			"--llama-parse-parameters.parsing-instruction", "parsing_instruction",
			"--llama-parse-parameters.precise-bounding-box=true",
			"--llama-parse-parameters.premium-mode=true",
			"--llama-parse-parameters.presentation-out-of-bounds-content=true",
			"--llama-parse-parameters.presentation-skip-embedded-data=true",
			"--llama-parse-parameters.preserve-layout-alignment-across-pages=true",
			"--llama-parse-parameters.preserve-very-small-text=true",
			"--llama-parse-parameters.preset", "preset",
			"--llama-parse-parameters.priority", "low",
			"--llama-parse-parameters.project-id", "project_id",
			"--llama-parse-parameters.remove-hidden-text=true",
			"--llama-parse-parameters.replace-failed-page-mode", "raw_text",
			"--llama-parse-parameters.replace-failed-page-with-error-message-prefix", "replace_failed_page_with_error_message_prefix",
			"--llama-parse-parameters.replace-failed-page-with-error-message-suffix", "replace_failed_page_with_error_message_suffix",
			"--llama-parse-parameters.save-images=true",
			"--llama-parse-parameters.skip-diagonal-text=true",
			"--llama-parse-parameters.specialized-chart-parsing-agentic=true",
			"--llama-parse-parameters.specialized-chart-parsing-efficient=true",
			"--llama-parse-parameters.specialized-chart-parsing-plus=true",
			"--llama-parse-parameters.specialized-image-parsing=true",
			"--llama-parse-parameters.spreadsheet-extract-sub-tables=true",
			"--llama-parse-parameters.spreadsheet-force-formula-computation=true",
			"--llama-parse-parameters.spreadsheet-include-hidden-sheets=true",
			"--llama-parse-parameters.strict-mode-buggy-font=true",
			"--llama-parse-parameters.strict-mode-image-extraction=true",
			"--llama-parse-parameters.strict-mode-image-ocr=true",
			"--llama-parse-parameters.strict-mode-reconstruction=true",
			"--llama-parse-parameters.structured-output=true",
			"--llama-parse-parameters.structured-output-json-schema", "structured_output_json_schema",
			"--llama-parse-parameters.structured-output-json-schema-name", "structured_output_json_schema_name",
			"--llama-parse-parameters.system-prompt", "system_prompt",
			"--llama-parse-parameters.system-prompt-append", "system_prompt_append",
			"--llama-parse-parameters.take-screenshot=true",
			"--llama-parse-parameters.target-pages", "target_pages",
			"--llama-parse-parameters.tier", "tier",
			"--llama-parse-parameters.use-vendor-multimodal-model=true",
			"--llama-parse-parameters.user-prompt", "user_prompt",
			"--llama-parse-parameters.vendor-multimodal-api-key", "vendor_multimodal_api_key",
			"--llama-parse-parameters.vendor-multimodal-model-name", "vendor_multimodal_model_name",
			"--llama-parse-parameters.version", "version",
			"--llama-parse-parameters.webhook-configurations", "[{webhook_events: [parse.success, parse.error], webhook_headers: {Authorization: Bearer sk-...}, webhook_output_format: json, webhook_url: https://example.com/webhooks/llamacloud}]",
			"--llama-parse-parameters.webhook-url", "webhook_url",
			"--managed-pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--metadata-config.excluded-embed-metadata-keys", "[string]",
			"--metadata-config.excluded-llm-metadata-keys", "[string]",
			"--name", "name",
			"--preset-retrieval-parameters.alpha", "0",
			"--preset-retrieval-parameters.class-name", "class_name",
			"--preset-retrieval-parameters.dense-similarity-cutoff", "0",
			"--preset-retrieval-parameters.dense-similarity-top-k", "1",
			"--preset-retrieval-parameters.enable-reranking=true",
			"--preset-retrieval-parameters.files-top-k", "1",
			"--preset-retrieval-parameters.rerank-top-n", "1",
			"--preset-retrieval-parameters.retrieval-mode", "chunks",
			"--preset-retrieval-parameters.retrieve-image-nodes=true",
			"--preset-retrieval-parameters.retrieve-page-figure-nodes=true",
			"--preset-retrieval-parameters.retrieve-page-screenshot-nodes=true",
			"--preset-retrieval-parameters.search-filters", "{filters: [{key: key, value: 0, operator: '=='}], condition: and}",
			"--preset-retrieval-parameters.search-filters-inference-schema", "{foo: {foo: bar}}",
			"--preset-retrieval-parameters.sparse-similarity-top-k", "1",
			"--sparse-model-config.class-name", "class_name",
			"--sparse-model-config.model-type", "splade",
			"--status", "status",
			"--transform-config", "{chunk_overlap: 0, chunk_size: 1, mode: auto}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"data_sink:\n" +
			"  component:\n" +
			"    foo: bar\n" +
			"  name: name\n" +
			"  sink_type: PINECONE\n" +
			"data_sink_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"embedding_config:\n" +
			"  component:\n" +
			"    additional_kwargs:\n" +
			"      foo: bar\n" +
			"    api_base: api_base\n" +
			"    api_key: api_key\n" +
			"    api_version: api_version\n" +
			"    azure_deployment: azure_deployment\n" +
			"    azure_endpoint: azure_endpoint\n" +
			"    class_name: class_name\n" +
			"    default_headers:\n" +
			"      foo: string\n" +
			"    dimensions: 0\n" +
			"    embed_batch_size: 1\n" +
			"    max_retries: 0\n" +
			"    model_name: model_name\n" +
			"    num_workers: 0\n" +
			"    reuse_client: true\n" +
			"    timeout: 0\n" +
			"  type: AZURE_EMBEDDING\n" +
			"embedding_model_config_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"llama_parse_parameters:\n" +
			"  adaptive_long_table: true\n" +
			"  aggressive_table_extraction: true\n" +
			"  annotate_links: true\n" +
			"  auto_mode: true\n" +
			"  auto_mode_configuration_json: auto_mode_configuration_json\n" +
			"  auto_mode_trigger_on_image_in_page: true\n" +
			"  auto_mode_trigger_on_regexp_in_page: auto_mode_trigger_on_regexp_in_page\n" +
			"  auto_mode_trigger_on_table_in_page: true\n" +
			"  auto_mode_trigger_on_text_in_page: auto_mode_trigger_on_text_in_page\n" +
			"  azure_openai_api_version: azure_openai_api_version\n" +
			"  azure_openai_deployment_name: azure_openai_deployment_name\n" +
			"  azure_openai_endpoint: azure_openai_endpoint\n" +
			"  azure_openai_key: azure_openai_key\n" +
			"  bbox_bottom: 0\n" +
			"  bbox_left: 0\n" +
			"  bbox_right: 0\n" +
			"  bbox_top: 0\n" +
			"  bounding_box: bounding_box\n" +
			"  compact_markdown_table: true\n" +
			"  complemental_formatting_instruction: complemental_formatting_instruction\n" +
			"  content_guideline_instruction: content_guideline_instruction\n" +
			"  continuous_mode: true\n" +
			"  disable_image_extraction: true\n" +
			"  disable_ocr: true\n" +
			"  disable_reconstruction: true\n" +
			"  do_not_cache: true\n" +
			"  do_not_unroll_columns: true\n" +
			"  enable_cost_optimizer: true\n" +
			"  extract_charts: true\n" +
			"  extract_layout: true\n" +
			"  extract_printed_page_number: true\n" +
			"  fast_mode: true\n" +
			"  formatting_instruction: formatting_instruction\n" +
			"  gpt4o_api_key: gpt4o_api_key\n" +
			"  gpt4o_mode: true\n" +
			"  guess_xlsx_sheet_name: true\n" +
			"  hide_footers: true\n" +
			"  hide_headers: true\n" +
			"  high_res_ocr: true\n" +
			"  html_make_all_elements_visible: true\n" +
			"  html_remove_fixed_elements: true\n" +
			"  html_remove_navigation_elements: true\n" +
			"  http_proxy: http_proxy\n" +
			"  ignore_document_elements_for_layout_detection: true\n" +
			"  images_to_save:\n" +
			"    - screenshot\n" +
			"  inline_images_in_markdown: true\n" +
			"  input_s3_path: input_s3_path\n" +
			"  input_s3_region: input_s3_region\n" +
			"  input_url: input_url\n" +
			"  internal_is_screenshot_job: true\n" +
			"  invalidate_cache: true\n" +
			"  is_formatting_instruction: true\n" +
			"  job_timeout_extra_time_per_page_in_seconds: 0\n" +
			"  job_timeout_in_seconds: 0\n" +
			"  keep_page_separator_when_merging_tables: true\n" +
			"  languages:\n" +
			"    - af\n" +
			"  layout_aware: true\n" +
			"  line_level_bounding_box: true\n" +
			"  markdown_table_multiline_header_separator: markdown_table_multiline_header_separator\n" +
			"  max_pages: 0\n" +
			"  max_pages_enforced: 0\n" +
			"  merge_tables_across_pages_in_markdown: true\n" +
			"  model: model\n" +
			"  outlined_table_extraction: true\n" +
			"  output_pdf_of_document: true\n" +
			"  output_s3_path_prefix: output_s3_path_prefix\n" +
			"  output_s3_region: output_s3_region\n" +
			"  output_tables_as_HTML: true\n" +
			"  page_error_tolerance: 0\n" +
			"  page_footer_prefix: page_footer_prefix\n" +
			"  page_footer_suffix: page_footer_suffix\n" +
			"  page_header_prefix: page_header_prefix\n" +
			"  page_header_suffix: page_header_suffix\n" +
			"  page_prefix: page_prefix\n" +
			"  page_separator: page_separator\n" +
			"  page_suffix: page_suffix\n" +
			"  parse_mode: parse_page_without_llm\n" +
			"  parsing_instruction: parsing_instruction\n" +
			"  precise_bounding_box: true\n" +
			"  premium_mode: true\n" +
			"  presentation_out_of_bounds_content: true\n" +
			"  presentation_skip_embedded_data: true\n" +
			"  preserve_layout_alignment_across_pages: true\n" +
			"  preserve_very_small_text: true\n" +
			"  preset: preset\n" +
			"  priority: low\n" +
			"  project_id: project_id\n" +
			"  remove_hidden_text: true\n" +
			"  replace_failed_page_mode: raw_text\n" +
			"  replace_failed_page_with_error_message_prefix: replace_failed_page_with_error_message_prefix\n" +
			"  replace_failed_page_with_error_message_suffix: replace_failed_page_with_error_message_suffix\n" +
			"  save_images: true\n" +
			"  skip_diagonal_text: true\n" +
			"  specialized_chart_parsing_agentic: true\n" +
			"  specialized_chart_parsing_efficient: true\n" +
			"  specialized_chart_parsing_plus: true\n" +
			"  specialized_image_parsing: true\n" +
			"  spreadsheet_extract_sub_tables: true\n" +
			"  spreadsheet_force_formula_computation: true\n" +
			"  spreadsheet_include_hidden_sheets: true\n" +
			"  strict_mode_buggy_font: true\n" +
			"  strict_mode_image_extraction: true\n" +
			"  strict_mode_image_ocr: true\n" +
			"  strict_mode_reconstruction: true\n" +
			"  structured_output: true\n" +
			"  structured_output_json_schema: structured_output_json_schema\n" +
			"  structured_output_json_schema_name: structured_output_json_schema_name\n" +
			"  system_prompt: system_prompt\n" +
			"  system_prompt_append: system_prompt_append\n" +
			"  take_screenshot: true\n" +
			"  target_pages: target_pages\n" +
			"  tier: tier\n" +
			"  use_vendor_multimodal_model: true\n" +
			"  user_prompt: user_prompt\n" +
			"  vendor_multimodal_api_key: vendor_multimodal_api_key\n" +
			"  vendor_multimodal_model_name: vendor_multimodal_model_name\n" +
			"  version: version\n" +
			"  webhook_configurations:\n" +
			"    - webhook_events:\n" +
			"        - parse.success\n" +
			"        - parse.error\n" +
			"      webhook_headers:\n" +
			"        Authorization: Bearer sk-...\n" +
			"      webhook_output_format: json\n" +
			"      webhook_url: https://example.com/webhooks/llamacloud\n" +
			"  webhook_url: webhook_url\n" +
			"managed_pipeline_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"metadata_config:\n" +
			"  excluded_embed_metadata_keys:\n" +
			"    - string\n" +
			"  excluded_llm_metadata_keys:\n" +
			"    - string\n" +
			"name: name\n" +
			"preset_retrieval_parameters:\n" +
			"  alpha: 0\n" +
			"  class_name: class_name\n" +
			"  dense_similarity_cutoff: 0\n" +
			"  dense_similarity_top_k: 1\n" +
			"  enable_reranking: true\n" +
			"  files_top_k: 1\n" +
			"  rerank_top_n: 1\n" +
			"  retrieval_mode: chunks\n" +
			"  retrieve_image_nodes: true\n" +
			"  retrieve_page_figure_nodes: true\n" +
			"  retrieve_page_screenshot_nodes: true\n" +
			"  search_filters:\n" +
			"    filters:\n" +
			"      - key: key\n" +
			"        value: 0\n" +
			"        operator: '=='\n" +
			"    condition: and\n" +
			"  search_filters_inference_schema:\n" +
			"    foo:\n" +
			"      foo: bar\n" +
			"  sparse_similarity_top_k: 1\n" +
			"sparse_model_config:\n" +
			"  class_name: class_name\n" +
			"  model_type: splade\n" +
			"status: status\n" +
			"transform_config:\n" +
			"  chunk_overlap: 0\n" +
			"  chunk_size: 1\n" +
			"  mode: auto\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"pipelines", "update",
			"--pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPipelinesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pipelines", "list",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--pipeline-name", "pipeline_name",
			"--pipeline-type", "PLAYGROUND",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-name", "project_name",
		)
	})
}

func TestPipelinesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pipelines", "delete",
			"--pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPipelinesGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pipelines", "get",
			"--pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPipelinesGetStatus(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pipelines", "get-status",
			"--pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--full-details=true",
		)
	})
}

func TestPipelinesRunSearch(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pipelines", "run-search",
			"--pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--query", "x",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--alpha", "0",
			"--class-name", "class_name",
			"--dense-similarity-cutoff", "0",
			"--dense-similarity-top-k", "1",
			"--enable-reranking=true",
			"--files-top-k", "1",
			"--rerank-top-n", "1",
			"--retrieval-mode", "chunks",
			"--retrieve-image-nodes=true",
			"--retrieve-page-figure-nodes=true",
			"--retrieve-page-screenshot-nodes=true",
			"--search-filters", "{filters: [{key: key, value: 0, operator: '=='}], condition: and}",
			"--search-filters-inference-schema", "{foo: {foo: bar}}",
			"--sparse-similarity-top-k", "1",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(pipelinesRunSearch)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pipelines", "run-search",
			"--pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--query", "x",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--alpha", "0",
			"--class-name", "class_name",
			"--dense-similarity-cutoff", "0",
			"--dense-similarity-top-k", "1",
			"--enable-reranking=true",
			"--files-top-k", "1",
			"--rerank-top-n", "1",
			"--retrieval-mode", "chunks",
			"--retrieve-image-nodes=true",
			"--retrieve-page-figure-nodes=true",
			"--retrieve-page-screenshot-nodes=true",
			"--search-filters.filters", "[{key: key, value: 0, operator: '=='}]",
			"--search-filters.condition", "and",
			"--search-filters-inference-schema", "{foo: {foo: bar}}",
			"--sparse-similarity-top-k", "1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"query: x\n" +
			"alpha: 0\n" +
			"class_name: class_name\n" +
			"dense_similarity_cutoff: 0\n" +
			"dense_similarity_top_k: 1\n" +
			"enable_reranking: true\n" +
			"files_top_k: 1\n" +
			"rerank_top_n: 1\n" +
			"retrieval_mode: chunks\n" +
			"retrieve_image_nodes: true\n" +
			"retrieve_page_figure_nodes: true\n" +
			"retrieve_page_screenshot_nodes: true\n" +
			"search_filters:\n" +
			"  filters:\n" +
			"    - key: key\n" +
			"      value: 0\n" +
			"      operator: '=='\n" +
			"  condition: and\n" +
			"search_filters_inference_schema:\n" +
			"  foo:\n" +
			"    foo: bar\n" +
			"sparse_similarity_top_k: 1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"pipelines", "run-search",
			"--pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPipelinesUpsert(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pipelines", "upsert",
			"--name", "x",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--data-sink", "{component: {foo: bar}, name: name, sink_type: PINECONE}",
			"--data-sink-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--embedding-config", "{component: {additional_kwargs: {foo: bar}, api_base: api_base, api_key: api_key, api_version: api_version, azure_deployment: azure_deployment, azure_endpoint: azure_endpoint, class_name: class_name, default_headers: {foo: string}, dimensions: 0, embed_batch_size: 1, max_retries: 0, model_name: model_name, num_workers: 0, reuse_client: true, timeout: 0}, type: AZURE_EMBEDDING}",
			"--embedding-model-config-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--llama-parse-parameters", "{adaptive_long_table: true, aggressive_table_extraction: true, annotate_links: true, auto_mode: true, auto_mode_configuration_json: auto_mode_configuration_json, auto_mode_trigger_on_image_in_page: true, auto_mode_trigger_on_regexp_in_page: auto_mode_trigger_on_regexp_in_page, auto_mode_trigger_on_table_in_page: true, auto_mode_trigger_on_text_in_page: auto_mode_trigger_on_text_in_page, azure_openai_api_version: azure_openai_api_version, azure_openai_deployment_name: azure_openai_deployment_name, azure_openai_endpoint: azure_openai_endpoint, azure_openai_key: azure_openai_key, bbox_bottom: 0, bbox_left: 0, bbox_right: 0, bbox_top: 0, bounding_box: bounding_box, compact_markdown_table: true, complemental_formatting_instruction: complemental_formatting_instruction, content_guideline_instruction: content_guideline_instruction, continuous_mode: true, disable_image_extraction: true, disable_ocr: true, disable_reconstruction: true, do_not_cache: true, do_not_unroll_columns: true, enable_cost_optimizer: true, extract_charts: true, extract_layout: true, extract_printed_page_number: true, fast_mode: true, formatting_instruction: formatting_instruction, gpt4o_api_key: gpt4o_api_key, gpt4o_mode: true, guess_xlsx_sheet_name: true, hide_footers: true, hide_headers: true, high_res_ocr: true, html_make_all_elements_visible: true, html_remove_fixed_elements: true, html_remove_navigation_elements: true, http_proxy: http_proxy, ignore_document_elements_for_layout_detection: true, images_to_save: [screenshot], inline_images_in_markdown: true, input_s3_path: input_s3_path, input_s3_region: input_s3_region, input_url: input_url, internal_is_screenshot_job: true, invalidate_cache: true, is_formatting_instruction: true, job_timeout_extra_time_per_page_in_seconds: 0, job_timeout_in_seconds: 0, keep_page_separator_when_merging_tables: true, languages: [af], layout_aware: true, line_level_bounding_box: true, markdown_table_multiline_header_separator: markdown_table_multiline_header_separator, max_pages: 0, max_pages_enforced: 0, merge_tables_across_pages_in_markdown: true, model: model, outlined_table_extraction: true, output_pdf_of_document: true, output_s3_path_prefix: output_s3_path_prefix, output_s3_region: output_s3_region, output_tables_as_HTML: true, page_error_tolerance: 0, page_footer_prefix: page_footer_prefix, page_footer_suffix: page_footer_suffix, page_header_prefix: page_header_prefix, page_header_suffix: page_header_suffix, page_prefix: page_prefix, page_separator: page_separator, page_suffix: page_suffix, parse_mode: parse_page_without_llm, parsing_instruction: parsing_instruction, precise_bounding_box: true, premium_mode: true, presentation_out_of_bounds_content: true, presentation_skip_embedded_data: true, preserve_layout_alignment_across_pages: true, preserve_very_small_text: true, preset: preset, priority: low, project_id: project_id, remove_hidden_text: true, replace_failed_page_mode: raw_text, replace_failed_page_with_error_message_prefix: replace_failed_page_with_error_message_prefix, replace_failed_page_with_error_message_suffix: replace_failed_page_with_error_message_suffix, save_images: true, skip_diagonal_text: true, specialized_chart_parsing_agentic: true, specialized_chart_parsing_efficient: true, specialized_chart_parsing_plus: true, specialized_image_parsing: true, spreadsheet_extract_sub_tables: true, spreadsheet_force_formula_computation: true, spreadsheet_include_hidden_sheets: true, strict_mode_buggy_font: true, strict_mode_image_extraction: true, strict_mode_image_ocr: true, strict_mode_reconstruction: true, structured_output: true, structured_output_json_schema: structured_output_json_schema, structured_output_json_schema_name: structured_output_json_schema_name, system_prompt: system_prompt, system_prompt_append: system_prompt_append, take_screenshot: true, target_pages: target_pages, tier: tier, use_vendor_multimodal_model: true, user_prompt: user_prompt, vendor_multimodal_api_key: vendor_multimodal_api_key, vendor_multimodal_model_name: vendor_multimodal_model_name, version: version, webhook_configurations: [{webhook_events: [parse.success, parse.error], webhook_headers: {Authorization: Bearer sk-...}, webhook_output_format: json, webhook_url: https://example.com/webhooks/llamacloud}], webhook_url: webhook_url}",
			"--managed-pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--metadata-config", "{excluded_embed_metadata_keys: [string], excluded_llm_metadata_keys: [string]}",
			"--pipeline-type", "PLAYGROUND",
			"--preset-retrieval-parameters", "{alpha: 0, class_name: class_name, dense_similarity_cutoff: 0, dense_similarity_top_k: 1, enable_reranking: true, files_top_k: 1, rerank_top_n: 1, retrieval_mode: chunks, retrieve_image_nodes: true, retrieve_page_figure_nodes: true, retrieve_page_screenshot_nodes: true, search_filters: {filters: [{key: key, value: 0, operator: '=='}], condition: and}, search_filters_inference_schema: {foo: {foo: bar}}, sparse_similarity_top_k: 1}",
			"--sparse-model-config", "{class_name: class_name, model_type: splade}",
			"--status", "status",
			"--transform-config", "{chunk_overlap: 0, chunk_size: 1, mode: auto}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(pipelinesUpsert)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"pipelines", "upsert",
			"--name", "x",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--data-sink.component", "{foo: bar}",
			"--data-sink.name", "name",
			"--data-sink.sink-type", "PINECONE",
			"--data-sink-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--embedding-config", "{component: {additional_kwargs: {foo: bar}, api_base: api_base, api_key: api_key, api_version: api_version, azure_deployment: azure_deployment, azure_endpoint: azure_endpoint, class_name: class_name, default_headers: {foo: string}, dimensions: 0, embed_batch_size: 1, max_retries: 0, model_name: model_name, num_workers: 0, reuse_client: true, timeout: 0}, type: AZURE_EMBEDDING}",
			"--embedding-model-config-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--llama-parse-parameters.adaptive-long-table=true",
			"--llama-parse-parameters.aggressive-table-extraction=true",
			"--llama-parse-parameters.annotate-links=true",
			"--llama-parse-parameters.auto-mode=true",
			"--llama-parse-parameters.auto-mode-configuration-json", "auto_mode_configuration_json",
			"--llama-parse-parameters.auto-mode-trigger-on-image-in-page=true",
			"--llama-parse-parameters.auto-mode-trigger-on-regexp-in-page", "auto_mode_trigger_on_regexp_in_page",
			"--llama-parse-parameters.auto-mode-trigger-on-table-in-page=true",
			"--llama-parse-parameters.auto-mode-trigger-on-text-in-page", "auto_mode_trigger_on_text_in_page",
			"--llama-parse-parameters.azure-openai-api-version", "azure_openai_api_version",
			"--llama-parse-parameters.azure-openai-deployment-name", "azure_openai_deployment_name",
			"--llama-parse-parameters.azure-openai-endpoint", "azure_openai_endpoint",
			"--llama-parse-parameters.azure-openai-key", "azure_openai_key",
			"--llama-parse-parameters.bbox-bottom", "0",
			"--llama-parse-parameters.bbox-left", "0",
			"--llama-parse-parameters.bbox-right", "0",
			"--llama-parse-parameters.bbox-top", "0",
			"--llama-parse-parameters.bounding-box", "bounding_box",
			"--llama-parse-parameters.compact-markdown-table=true",
			"--llama-parse-parameters.complemental-formatting-instruction", "complemental_formatting_instruction",
			"--llama-parse-parameters.content-guideline-instruction", "content_guideline_instruction",
			"--llama-parse-parameters.continuous-mode=true",
			"--llama-parse-parameters.disable-image-extraction=true",
			"--llama-parse-parameters.disable-ocr=true",
			"--llama-parse-parameters.disable-reconstruction=true",
			"--llama-parse-parameters.do-not-cache=true",
			"--llama-parse-parameters.do-not-unroll-columns=true",
			"--llama-parse-parameters.enable-cost-optimizer=true",
			"--llama-parse-parameters.extract-charts=true",
			"--llama-parse-parameters.extract-layout=true",
			"--llama-parse-parameters.extract-printed-page-number=true",
			"--llama-parse-parameters.fast-mode=true",
			"--llama-parse-parameters.formatting-instruction", "formatting_instruction",
			"--llama-parse-parameters.gpt4o-api-key", "gpt4o_api_key",
			"--llama-parse-parameters.gpt4o-mode=true",
			"--llama-parse-parameters.guess-xlsx-sheet-name=true",
			"--llama-parse-parameters.hide-footers=true",
			"--llama-parse-parameters.hide-headers=true",
			"--llama-parse-parameters.high-res-ocr=true",
			"--llama-parse-parameters.html-make-all-elements-visible=true",
			"--llama-parse-parameters.html-remove-fixed-elements=true",
			"--llama-parse-parameters.html-remove-navigation-elements=true",
			"--llama-parse-parameters.http-proxy", "http_proxy",
			"--llama-parse-parameters.ignore-document-elements-for-layout-detection=true",
			"--llama-parse-parameters.images-to-save", "[screenshot]",
			"--llama-parse-parameters.inline-images-in-markdown=true",
			"--llama-parse-parameters.input-s3-path", "input_s3_path",
			"--llama-parse-parameters.input-s3-region", "input_s3_region",
			"--llama-parse-parameters.input-url", "input_url",
			"--llama-parse-parameters.internal-is-screenshot-job=true",
			"--llama-parse-parameters.invalidate-cache=true",
			"--llama-parse-parameters.is-formatting-instruction=true",
			"--llama-parse-parameters.job-timeout-extra-time-per-page-in-seconds", "0",
			"--llama-parse-parameters.job-timeout-in-seconds", "0",
			"--llama-parse-parameters.keep-page-separator-when-merging-tables=true",
			"--llama-parse-parameters.languages", "[af]",
			"--llama-parse-parameters.layout-aware=true",
			"--llama-parse-parameters.line-level-bounding-box=true",
			"--llama-parse-parameters.markdown-table-multiline-header-separator", "markdown_table_multiline_header_separator",
			"--llama-parse-parameters.max-pages", "0",
			"--llama-parse-parameters.max-pages-enforced", "0",
			"--llama-parse-parameters.merge-tables-across-pages-in-markdown=true",
			"--llama-parse-parameters.model", "model",
			"--llama-parse-parameters.outlined-table-extraction=true",
			"--llama-parse-parameters.output-pdf-of-document=true",
			"--llama-parse-parameters.output-s3-path-prefix", "output_s3_path_prefix",
			"--llama-parse-parameters.output-s3-region", "output_s3_region",
			"--llama-parse-parameters.output-tables-as-html=true",
			"--llama-parse-parameters.page-error-tolerance", "0",
			"--llama-parse-parameters.page-footer-prefix", "page_footer_prefix",
			"--llama-parse-parameters.page-footer-suffix", "page_footer_suffix",
			"--llama-parse-parameters.page-header-prefix", "page_header_prefix",
			"--llama-parse-parameters.page-header-suffix", "page_header_suffix",
			"--llama-parse-parameters.page-prefix", "page_prefix",
			"--llama-parse-parameters.page-separator", "page_separator",
			"--llama-parse-parameters.page-suffix", "page_suffix",
			"--llama-parse-parameters.parse-mode", "parse_page_without_llm",
			"--llama-parse-parameters.parsing-instruction", "parsing_instruction",
			"--llama-parse-parameters.precise-bounding-box=true",
			"--llama-parse-parameters.premium-mode=true",
			"--llama-parse-parameters.presentation-out-of-bounds-content=true",
			"--llama-parse-parameters.presentation-skip-embedded-data=true",
			"--llama-parse-parameters.preserve-layout-alignment-across-pages=true",
			"--llama-parse-parameters.preserve-very-small-text=true",
			"--llama-parse-parameters.preset", "preset",
			"--llama-parse-parameters.priority", "low",
			"--llama-parse-parameters.project-id", "project_id",
			"--llama-parse-parameters.remove-hidden-text=true",
			"--llama-parse-parameters.replace-failed-page-mode", "raw_text",
			"--llama-parse-parameters.replace-failed-page-with-error-message-prefix", "replace_failed_page_with_error_message_prefix",
			"--llama-parse-parameters.replace-failed-page-with-error-message-suffix", "replace_failed_page_with_error_message_suffix",
			"--llama-parse-parameters.save-images=true",
			"--llama-parse-parameters.skip-diagonal-text=true",
			"--llama-parse-parameters.specialized-chart-parsing-agentic=true",
			"--llama-parse-parameters.specialized-chart-parsing-efficient=true",
			"--llama-parse-parameters.specialized-chart-parsing-plus=true",
			"--llama-parse-parameters.specialized-image-parsing=true",
			"--llama-parse-parameters.spreadsheet-extract-sub-tables=true",
			"--llama-parse-parameters.spreadsheet-force-formula-computation=true",
			"--llama-parse-parameters.spreadsheet-include-hidden-sheets=true",
			"--llama-parse-parameters.strict-mode-buggy-font=true",
			"--llama-parse-parameters.strict-mode-image-extraction=true",
			"--llama-parse-parameters.strict-mode-image-ocr=true",
			"--llama-parse-parameters.strict-mode-reconstruction=true",
			"--llama-parse-parameters.structured-output=true",
			"--llama-parse-parameters.structured-output-json-schema", "structured_output_json_schema",
			"--llama-parse-parameters.structured-output-json-schema-name", "structured_output_json_schema_name",
			"--llama-parse-parameters.system-prompt", "system_prompt",
			"--llama-parse-parameters.system-prompt-append", "system_prompt_append",
			"--llama-parse-parameters.take-screenshot=true",
			"--llama-parse-parameters.target-pages", "target_pages",
			"--llama-parse-parameters.tier", "tier",
			"--llama-parse-parameters.use-vendor-multimodal-model=true",
			"--llama-parse-parameters.user-prompt", "user_prompt",
			"--llama-parse-parameters.vendor-multimodal-api-key", "vendor_multimodal_api_key",
			"--llama-parse-parameters.vendor-multimodal-model-name", "vendor_multimodal_model_name",
			"--llama-parse-parameters.version", "version",
			"--llama-parse-parameters.webhook-configurations", "[{webhook_events: [parse.success, parse.error], webhook_headers: {Authorization: Bearer sk-...}, webhook_output_format: json, webhook_url: https://example.com/webhooks/llamacloud}]",
			"--llama-parse-parameters.webhook-url", "webhook_url",
			"--managed-pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--metadata-config.excluded-embed-metadata-keys", "[string]",
			"--metadata-config.excluded-llm-metadata-keys", "[string]",
			"--pipeline-type", "PLAYGROUND",
			"--preset-retrieval-parameters.alpha", "0",
			"--preset-retrieval-parameters.class-name", "class_name",
			"--preset-retrieval-parameters.dense-similarity-cutoff", "0",
			"--preset-retrieval-parameters.dense-similarity-top-k", "1",
			"--preset-retrieval-parameters.enable-reranking=true",
			"--preset-retrieval-parameters.files-top-k", "1",
			"--preset-retrieval-parameters.rerank-top-n", "1",
			"--preset-retrieval-parameters.retrieval-mode", "chunks",
			"--preset-retrieval-parameters.retrieve-image-nodes=true",
			"--preset-retrieval-parameters.retrieve-page-figure-nodes=true",
			"--preset-retrieval-parameters.retrieve-page-screenshot-nodes=true",
			"--preset-retrieval-parameters.search-filters", "{filters: [{key: key, value: 0, operator: '=='}], condition: and}",
			"--preset-retrieval-parameters.search-filters-inference-schema", "{foo: {foo: bar}}",
			"--preset-retrieval-parameters.sparse-similarity-top-k", "1",
			"--sparse-model-config.class-name", "class_name",
			"--sparse-model-config.model-type", "splade",
			"--status", "status",
			"--transform-config", "{chunk_overlap: 0, chunk_size: 1, mode: auto}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: x\n" +
			"data_sink:\n" +
			"  component:\n" +
			"    foo: bar\n" +
			"  name: name\n" +
			"  sink_type: PINECONE\n" +
			"data_sink_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"embedding_config:\n" +
			"  component:\n" +
			"    additional_kwargs:\n" +
			"      foo: bar\n" +
			"    api_base: api_base\n" +
			"    api_key: api_key\n" +
			"    api_version: api_version\n" +
			"    azure_deployment: azure_deployment\n" +
			"    azure_endpoint: azure_endpoint\n" +
			"    class_name: class_name\n" +
			"    default_headers:\n" +
			"      foo: string\n" +
			"    dimensions: 0\n" +
			"    embed_batch_size: 1\n" +
			"    max_retries: 0\n" +
			"    model_name: model_name\n" +
			"    num_workers: 0\n" +
			"    reuse_client: true\n" +
			"    timeout: 0\n" +
			"  type: AZURE_EMBEDDING\n" +
			"embedding_model_config_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"llama_parse_parameters:\n" +
			"  adaptive_long_table: true\n" +
			"  aggressive_table_extraction: true\n" +
			"  annotate_links: true\n" +
			"  auto_mode: true\n" +
			"  auto_mode_configuration_json: auto_mode_configuration_json\n" +
			"  auto_mode_trigger_on_image_in_page: true\n" +
			"  auto_mode_trigger_on_regexp_in_page: auto_mode_trigger_on_regexp_in_page\n" +
			"  auto_mode_trigger_on_table_in_page: true\n" +
			"  auto_mode_trigger_on_text_in_page: auto_mode_trigger_on_text_in_page\n" +
			"  azure_openai_api_version: azure_openai_api_version\n" +
			"  azure_openai_deployment_name: azure_openai_deployment_name\n" +
			"  azure_openai_endpoint: azure_openai_endpoint\n" +
			"  azure_openai_key: azure_openai_key\n" +
			"  bbox_bottom: 0\n" +
			"  bbox_left: 0\n" +
			"  bbox_right: 0\n" +
			"  bbox_top: 0\n" +
			"  bounding_box: bounding_box\n" +
			"  compact_markdown_table: true\n" +
			"  complemental_formatting_instruction: complemental_formatting_instruction\n" +
			"  content_guideline_instruction: content_guideline_instruction\n" +
			"  continuous_mode: true\n" +
			"  disable_image_extraction: true\n" +
			"  disable_ocr: true\n" +
			"  disable_reconstruction: true\n" +
			"  do_not_cache: true\n" +
			"  do_not_unroll_columns: true\n" +
			"  enable_cost_optimizer: true\n" +
			"  extract_charts: true\n" +
			"  extract_layout: true\n" +
			"  extract_printed_page_number: true\n" +
			"  fast_mode: true\n" +
			"  formatting_instruction: formatting_instruction\n" +
			"  gpt4o_api_key: gpt4o_api_key\n" +
			"  gpt4o_mode: true\n" +
			"  guess_xlsx_sheet_name: true\n" +
			"  hide_footers: true\n" +
			"  hide_headers: true\n" +
			"  high_res_ocr: true\n" +
			"  html_make_all_elements_visible: true\n" +
			"  html_remove_fixed_elements: true\n" +
			"  html_remove_navigation_elements: true\n" +
			"  http_proxy: http_proxy\n" +
			"  ignore_document_elements_for_layout_detection: true\n" +
			"  images_to_save:\n" +
			"    - screenshot\n" +
			"  inline_images_in_markdown: true\n" +
			"  input_s3_path: input_s3_path\n" +
			"  input_s3_region: input_s3_region\n" +
			"  input_url: input_url\n" +
			"  internal_is_screenshot_job: true\n" +
			"  invalidate_cache: true\n" +
			"  is_formatting_instruction: true\n" +
			"  job_timeout_extra_time_per_page_in_seconds: 0\n" +
			"  job_timeout_in_seconds: 0\n" +
			"  keep_page_separator_when_merging_tables: true\n" +
			"  languages:\n" +
			"    - af\n" +
			"  layout_aware: true\n" +
			"  line_level_bounding_box: true\n" +
			"  markdown_table_multiline_header_separator: markdown_table_multiline_header_separator\n" +
			"  max_pages: 0\n" +
			"  max_pages_enforced: 0\n" +
			"  merge_tables_across_pages_in_markdown: true\n" +
			"  model: model\n" +
			"  outlined_table_extraction: true\n" +
			"  output_pdf_of_document: true\n" +
			"  output_s3_path_prefix: output_s3_path_prefix\n" +
			"  output_s3_region: output_s3_region\n" +
			"  output_tables_as_HTML: true\n" +
			"  page_error_tolerance: 0\n" +
			"  page_footer_prefix: page_footer_prefix\n" +
			"  page_footer_suffix: page_footer_suffix\n" +
			"  page_header_prefix: page_header_prefix\n" +
			"  page_header_suffix: page_header_suffix\n" +
			"  page_prefix: page_prefix\n" +
			"  page_separator: page_separator\n" +
			"  page_suffix: page_suffix\n" +
			"  parse_mode: parse_page_without_llm\n" +
			"  parsing_instruction: parsing_instruction\n" +
			"  precise_bounding_box: true\n" +
			"  premium_mode: true\n" +
			"  presentation_out_of_bounds_content: true\n" +
			"  presentation_skip_embedded_data: true\n" +
			"  preserve_layout_alignment_across_pages: true\n" +
			"  preserve_very_small_text: true\n" +
			"  preset: preset\n" +
			"  priority: low\n" +
			"  project_id: project_id\n" +
			"  remove_hidden_text: true\n" +
			"  replace_failed_page_mode: raw_text\n" +
			"  replace_failed_page_with_error_message_prefix: replace_failed_page_with_error_message_prefix\n" +
			"  replace_failed_page_with_error_message_suffix: replace_failed_page_with_error_message_suffix\n" +
			"  save_images: true\n" +
			"  skip_diagonal_text: true\n" +
			"  specialized_chart_parsing_agentic: true\n" +
			"  specialized_chart_parsing_efficient: true\n" +
			"  specialized_chart_parsing_plus: true\n" +
			"  specialized_image_parsing: true\n" +
			"  spreadsheet_extract_sub_tables: true\n" +
			"  spreadsheet_force_formula_computation: true\n" +
			"  spreadsheet_include_hidden_sheets: true\n" +
			"  strict_mode_buggy_font: true\n" +
			"  strict_mode_image_extraction: true\n" +
			"  strict_mode_image_ocr: true\n" +
			"  strict_mode_reconstruction: true\n" +
			"  structured_output: true\n" +
			"  structured_output_json_schema: structured_output_json_schema\n" +
			"  structured_output_json_schema_name: structured_output_json_schema_name\n" +
			"  system_prompt: system_prompt\n" +
			"  system_prompt_append: system_prompt_append\n" +
			"  take_screenshot: true\n" +
			"  target_pages: target_pages\n" +
			"  tier: tier\n" +
			"  use_vendor_multimodal_model: true\n" +
			"  user_prompt: user_prompt\n" +
			"  vendor_multimodal_api_key: vendor_multimodal_api_key\n" +
			"  vendor_multimodal_model_name: vendor_multimodal_model_name\n" +
			"  version: version\n" +
			"  webhook_configurations:\n" +
			"    - webhook_events:\n" +
			"        - parse.success\n" +
			"        - parse.error\n" +
			"      webhook_headers:\n" +
			"        Authorization: Bearer sk-...\n" +
			"      webhook_output_format: json\n" +
			"      webhook_url: https://example.com/webhooks/llamacloud\n" +
			"  webhook_url: webhook_url\n" +
			"managed_pipeline_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"metadata_config:\n" +
			"  excluded_embed_metadata_keys:\n" +
			"    - string\n" +
			"  excluded_llm_metadata_keys:\n" +
			"    - string\n" +
			"pipeline_type: PLAYGROUND\n" +
			"preset_retrieval_parameters:\n" +
			"  alpha: 0\n" +
			"  class_name: class_name\n" +
			"  dense_similarity_cutoff: 0\n" +
			"  dense_similarity_top_k: 1\n" +
			"  enable_reranking: true\n" +
			"  files_top_k: 1\n" +
			"  rerank_top_n: 1\n" +
			"  retrieval_mode: chunks\n" +
			"  retrieve_image_nodes: true\n" +
			"  retrieve_page_figure_nodes: true\n" +
			"  retrieve_page_screenshot_nodes: true\n" +
			"  search_filters:\n" +
			"    filters:\n" +
			"      - key: key\n" +
			"        value: 0\n" +
			"        operator: '=='\n" +
			"    condition: and\n" +
			"  search_filters_inference_schema:\n" +
			"    foo:\n" +
			"      foo: bar\n" +
			"  sparse_similarity_top_k: 1\n" +
			"sparse_model_config:\n" +
			"  class_name: class_name\n" +
			"  model_type: splade\n" +
			"status: status\n" +
			"transform_config:\n" +
			"  chunk_overlap: 0\n" +
			"  chunk_size: 1\n" +
			"  mode: auto\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"pipelines", "upsert",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
