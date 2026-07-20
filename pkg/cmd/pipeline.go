// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/run-llama/llama-parse-cli/internal/apiquery"
	"github.com/run-llama/llama-parse-cli/internal/requestflag"
	"github.com/run-llama/llama-parse-go"
	"github.com/run-llama/llama-parse-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var pipelinesCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a new managed ingestion pipeline.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[*string]{
			Name:      "organization-id",
			QueryPath: "organization_id",
		},
		&requestflag.Flag[*string]{
			Name:      "project-id",
			QueryPath: "project_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "data-sink",
			Usage:    "Schema for creating a data sink.",
			BodyPath: "data_sink",
		},
		&requestflag.Flag[*string]{
			Name:     "data-sink-id",
			Usage:    "Data sink ID. When provided instead of data_sink, the data sink will be looked up by ID.",
			BodyPath: "data_sink_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "embedding-config",
			BodyPath: "embedding_config",
		},
		&requestflag.Flag[*string]{
			Name:     "embedding-model-config-id",
			Usage:    "Embedding model config ID. When provided instead of embedding_config, the embedding model config will be looked up by ID.",
			BodyPath: "embedding_model_config_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "llama-parse-parameters",
			BodyPath: "llama_parse_parameters",
		},
		&requestflag.Flag[*string]{
			Name:     "managed-pipeline-id",
			Usage:    "The ID of the ManagedPipeline this playground pipeline is linked to.",
			BodyPath: "managed_pipeline_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata-config",
			BodyPath: "metadata_config",
		},
		&requestflag.Flag[string]{
			Name:     "pipeline-type",
			Usage:    "Enum for representing the type of a pipeline",
			BodyPath: "pipeline_type",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "preset-retrieval-parameters",
			Usage:    "Schema for the search params for an retrieval execution that can be preset for a pipeline.",
			BodyPath: "preset_retrieval_parameters",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "sparse-model-config",
			Usage:    "Configuration for sparse embedding models used in hybrid search.\n\nThis allows users to choose between Splade and BM25 models for\nsparse retrieval in managed data sinks.",
			BodyPath: "sparse_model_config",
		},
		&requestflag.Flag[*string]{
			Name:     "status",
			Usage:    "Status of the pipeline deployment.",
			BodyPath: "status",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "transform-config",
			Usage:    "Configuration for the transformation.",
			BodyPath: "transform_config",
		},
	},
	Action:          handlePipelinesCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"data-sink": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "data-sink.component",
			Usage:      "Component that implements the data sink",
			InnerField: "component",
		},
		&requestflag.InnerFlag[string]{
			Name:       "data-sink.name",
			Usage:      "The name of the data sink.",
			InnerField: "name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "data-sink.sink-type",
			Usage:      `Allowed values: "PINECONE", "POSTGRES", "QDRANT", "AZUREAI_SEARCH", "MONGODB_ATLAS", "MILVUS", "ASTRA_DB".`,
			InnerField: "sink_type",
		},
	},
	"llama-parse-parameters": {
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.adaptive-long-table",
			InnerField: "adaptive_long_table",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.aggressive-table-extraction",
			InnerField: "aggressive_table_extraction",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.annotate-links",
			InnerField: "annotate_links",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.auto-mode",
			InnerField: "auto_mode",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.auto-mode-configuration-json",
			InnerField: "auto_mode_configuration_json",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.auto-mode-trigger-on-image-in-page",
			InnerField: "auto_mode_trigger_on_image_in_page",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.auto-mode-trigger-on-regexp-in-page",
			InnerField: "auto_mode_trigger_on_regexp_in_page",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.auto-mode-trigger-on-table-in-page",
			InnerField: "auto_mode_trigger_on_table_in_page",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.auto-mode-trigger-on-text-in-page",
			InnerField: "auto_mode_trigger_on_text_in_page",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.azure-openai-api-version",
			InnerField: "azure_openai_api_version",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.azure-openai-deployment-name",
			InnerField: "azure_openai_deployment_name",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.azure-openai-endpoint",
			InnerField: "azure_openai_endpoint",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.azure-openai-key",
			InnerField: "azure_openai_key",
		},
		&requestflag.InnerFlag[*float64]{
			Name:       "llama-parse-parameters.bbox-bottom",
			InnerField: "bbox_bottom",
		},
		&requestflag.InnerFlag[*float64]{
			Name:       "llama-parse-parameters.bbox-left",
			InnerField: "bbox_left",
		},
		&requestflag.InnerFlag[*float64]{
			Name:       "llama-parse-parameters.bbox-right",
			InnerField: "bbox_right",
		},
		&requestflag.InnerFlag[*float64]{
			Name:       "llama-parse-parameters.bbox-top",
			InnerField: "bbox_top",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.bounding-box",
			InnerField: "bounding_box",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.compact-markdown-table",
			InnerField: "compact_markdown_table",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.complemental-formatting-instruction",
			InnerField: "complemental_formatting_instruction",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.confidence-score-effort",
			InnerField: "confidence_score_effort",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.content-guideline-instruction",
			InnerField: "content_guideline_instruction",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.continuous-mode",
			InnerField: "continuous_mode",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.disable-image-extraction",
			InnerField: "disable_image_extraction",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.disable-ocr",
			InnerField: "disable_ocr",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.disable-reconstruction",
			InnerField: "disable_reconstruction",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.do-not-cache",
			InnerField: "do_not_cache",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.do-not-unroll-columns",
			InnerField: "do_not_unroll_columns",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.enable-cost-optimizer",
			InnerField: "enable_cost_optimizer",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.extract-charts",
			InnerField: "extract_charts",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.extract-layout",
			InnerField: "extract_layout",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.extract-printed-page-number",
			InnerField: "extract_printed_page_number",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.fast-mode",
			InnerField: "fast_mode",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.formatting-instruction",
			InnerField: "formatting_instruction",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.gpt4o-api-key",
			InnerField: "gpt4o_api_key",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.gpt4o-mode",
			InnerField: "gpt4o_mode",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.guess-xlsx-sheet-name",
			InnerField: "guess_xlsx_sheet_name",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.hide-footers",
			InnerField: "hide_footers",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.hide-headers",
			InnerField: "hide_headers",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.high-res-ocr",
			InnerField: "high_res_ocr",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.html-make-all-elements-visible",
			InnerField: "html_make_all_elements_visible",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.html-remove-fixed-elements",
			InnerField: "html_remove_fixed_elements",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.html-remove-navigation-elements",
			InnerField: "html_remove_navigation_elements",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.http-proxy",
			InnerField: "http_proxy",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.ignore-document-elements-for-layout-detection",
			InnerField: "ignore_document_elements_for_layout_detection",
		},
		&requestflag.InnerFlag[any]{
			Name:       "llama-parse-parameters.images-to-save",
			InnerField: "images_to_save",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.inline-images-in-markdown",
			InnerField: "inline_images_in_markdown",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.input-s3-path",
			InnerField: "input_s3_path",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.input-s3-region",
			InnerField: "input_s3_region",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.input-url",
			InnerField: "input_url",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.internal-is-screenshot-job",
			InnerField: "internal_is_screenshot_job",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.invalidate-cache",
			InnerField: "invalidate_cache",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.is-formatting-instruction",
			InnerField: "is_formatting_instruction",
		},
		&requestflag.InnerFlag[*float64]{
			Name:       "llama-parse-parameters.job-timeout-extra-time-per-page-in-seconds",
			InnerField: "job_timeout_extra_time_per_page_in_seconds",
		},
		&requestflag.InnerFlag[*float64]{
			Name:       "llama-parse-parameters.job-timeout-in-seconds",
			InnerField: "job_timeout_in_seconds",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.keep-page-separator-when-merging-tables",
			InnerField: "keep_page_separator_when_merging_tables",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "llama-parse-parameters.languages",
			InnerField: "languages",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.layout-aware",
			InnerField: "layout_aware",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.line-level-bounding-box",
			InnerField: "line_level_bounding_box",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.markdown-table-multiline-header-separator",
			InnerField: "markdown_table_multiline_header_separator",
		},
		&requestflag.InnerFlag[*int64]{
			Name:       "llama-parse-parameters.max-pages",
			InnerField: "max_pages",
		},
		&requestflag.InnerFlag[*int64]{
			Name:       "llama-parse-parameters.max-pages-enforced",
			InnerField: "max_pages_enforced",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.merge-tables-across-pages-in-markdown",
			InnerField: "merge_tables_across_pages_in_markdown",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.model",
			InnerField: "model",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.outlined-table-extraction",
			InnerField: "outlined_table_extraction",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.output-pdf-of-document",
			InnerField: "output_pdf_of_document",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.output-s3-path-prefix",
			InnerField: "output_s3_path_prefix",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.output-s3-region",
			InnerField: "output_s3_region",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.output-tables-as-html",
			InnerField: "output_tables_as_HTML",
		},
		&requestflag.InnerFlag[*float64]{
			Name:       "llama-parse-parameters.page-error-tolerance",
			InnerField: "page_error_tolerance",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.page-footer-prefix",
			InnerField: "page_footer_prefix",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.page-footer-suffix",
			InnerField: "page_footer_suffix",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.page-header-prefix",
			InnerField: "page_header_prefix",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.page-header-suffix",
			InnerField: "page_header_suffix",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.page-prefix",
			InnerField: "page_prefix",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.page-separator",
			InnerField: "page_separator",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.page-suffix",
			InnerField: "page_suffix",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.parse-mode",
			Usage:      "Enum for representing the mode of parsing to be used.",
			InnerField: "parse_mode",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.parsing-instruction",
			InnerField: "parsing_instruction",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.precise-bounding-box",
			InnerField: "precise_bounding_box",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.premium-mode",
			InnerField: "premium_mode",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.presentation-out-of-bounds-content",
			InnerField: "presentation_out_of_bounds_content",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.presentation-skip-embedded-data",
			InnerField: "presentation_skip_embedded_data",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.preserve-layout-alignment-across-pages",
			InnerField: "preserve_layout_alignment_across_pages",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.preserve-very-small-text",
			InnerField: "preserve_very_small_text",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.preset",
			InnerField: "preset",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.priority",
			Usage:      "The priority for the request. This field may be ignored or overwritten depending on the organization tier.",
			InnerField: "priority",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.project-id",
			InnerField: "project_id",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.remove-hidden-text",
			InnerField: "remove_hidden_text",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.replace-failed-page-mode",
			Usage:      "Enum for representing the different available page error handling modes.",
			InnerField: "replace_failed_page_mode",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.replace-failed-page-with-error-message-prefix",
			InnerField: "replace_failed_page_with_error_message_prefix",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.replace-failed-page-with-error-message-suffix",
			InnerField: "replace_failed_page_with_error_message_suffix",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.save-images",
			InnerField: "save_images",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.skip-diagonal-text",
			InnerField: "skip_diagonal_text",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.specialized-chart-parsing-agentic",
			InnerField: "specialized_chart_parsing_agentic",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.specialized-chart-parsing-efficient",
			InnerField: "specialized_chart_parsing_efficient",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.specialized-chart-parsing-plus",
			InnerField: "specialized_chart_parsing_plus",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.specialized-image-parsing",
			InnerField: "specialized_image_parsing",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.spreadsheet-extract-sub-tables",
			InnerField: "spreadsheet_extract_sub_tables",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.spreadsheet-force-formula-computation",
			InnerField: "spreadsheet_force_formula_computation",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.spreadsheet-include-hidden-sheets",
			InnerField: "spreadsheet_include_hidden_sheets",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.strict-mode-buggy-font",
			InnerField: "strict_mode_buggy_font",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.strict-mode-image-extraction",
			InnerField: "strict_mode_image_extraction",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.strict-mode-image-ocr",
			InnerField: "strict_mode_image_ocr",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.strict-mode-reconstruction",
			InnerField: "strict_mode_reconstruction",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.structured-output",
			InnerField: "structured_output",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.structured-output-json-schema",
			InnerField: "structured_output_json_schema",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.structured-output-json-schema-name",
			InnerField: "structured_output_json_schema_name",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.system-prompt",
			InnerField: "system_prompt",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.system-prompt-append",
			InnerField: "system_prompt_append",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.take-screenshot",
			InnerField: "take_screenshot",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.target-pages",
			InnerField: "target_pages",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.tier",
			InnerField: "tier",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.use-vendor-multimodal-model",
			InnerField: "use_vendor_multimodal_model",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.user-prompt",
			InnerField: "user_prompt",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.vendor-multimodal-api-key",
			InnerField: "vendor_multimodal_api_key",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.vendor-multimodal-model-name",
			InnerField: "vendor_multimodal_model_name",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.version",
			InnerField: "version",
		},
		&requestflag.InnerFlag[any]{
			Name:       "llama-parse-parameters.webhook-configurations",
			Usage:      "Outbound webhook endpoints to notify on job status changes",
			InnerField: "webhook_configurations",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.webhook-url",
			InnerField: "webhook_url",
		},
	},
	"metadata-config": {
		&requestflag.InnerFlag[[]string]{
			Name:       "metadata-config.excluded-embed-metadata-keys",
			Usage:      "List of metadata keys to exclude from embeddings",
			InnerField: "excluded_embed_metadata_keys",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "metadata-config.excluded-llm-metadata-keys",
			Usage:      "List of metadata keys to exclude from LLM during retrieval",
			InnerField: "excluded_llm_metadata_keys",
		},
	},
	"preset-retrieval-parameters": {
		&requestflag.InnerFlag[*float64]{
			Name:       "preset-retrieval-parameters.alpha",
			Usage:      "Alpha value for hybrid retrieval to determine the weights between dense and sparse retrieval. 0 is sparse retrieval and 1 is dense retrieval.",
			InnerField: "alpha",
		},
		&requestflag.InnerFlag[string]{
			Name:       "preset-retrieval-parameters.class-name",
			InnerField: "class_name",
		},
		&requestflag.InnerFlag[*float64]{
			Name:       "preset-retrieval-parameters.dense-similarity-cutoff",
			Usage:      "Minimum similarity score wrt query for retrieval",
			InnerField: "dense_similarity_cutoff",
		},
		&requestflag.InnerFlag[*int64]{
			Name:       "preset-retrieval-parameters.dense-similarity-top-k",
			Usage:      "Number of nodes for dense retrieval.",
			InnerField: "dense_similarity_top_k",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "preset-retrieval-parameters.enable-reranking",
			Usage:      "Enable reranking for retrieval",
			InnerField: "enable_reranking",
		},
		&requestflag.InnerFlag[*int64]{
			Name:       "preset-retrieval-parameters.files-top-k",
			Usage:      "Number of files to retrieve (only for retrieval mode files_via_metadata and files_via_content).",
			InnerField: "files_top_k",
		},
		&requestflag.InnerFlag[*int64]{
			Name:       "preset-retrieval-parameters.rerank-top-n",
			Usage:      "Number of reranked nodes for returning.",
			InnerField: "rerank_top_n",
		},
		&requestflag.InnerFlag[string]{
			Name:       "preset-retrieval-parameters.retrieval-mode",
			Usage:      `Allowed values: "chunks", "files_via_metadata", "files_via_content", "auto_routed".`,
			InnerField: "retrieval_mode",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "preset-retrieval-parameters.retrieve-image-nodes",
			Usage:      "Whether to retrieve image nodes.",
			InnerField: "retrieve_image_nodes",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "preset-retrieval-parameters.retrieve-page-figure-nodes",
			Usage:      "Whether to retrieve page figure nodes.",
			InnerField: "retrieve_page_figure_nodes",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "preset-retrieval-parameters.retrieve-page-screenshot-nodes",
			Usage:      "Whether to retrieve page screenshot nodes.",
			InnerField: "retrieve_page_screenshot_nodes",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "preset-retrieval-parameters.search-filters",
			Usage:      "Metadata filters for vector stores.",
			InnerField: "search_filters",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "preset-retrieval-parameters.search-filters-inference-schema",
			Usage:      "JSON Schema that will be used to infer search_filters. Omit or leave as null to skip inference.",
			InnerField: "search_filters_inference_schema",
		},
		&requestflag.InnerFlag[*int64]{
			Name:       "preset-retrieval-parameters.sparse-similarity-top-k",
			Usage:      "Number of nodes for sparse retrieval.",
			InnerField: "sparse_similarity_top_k",
		},
	},
	"sparse-model-config": {
		&requestflag.InnerFlag[string]{
			Name:       "sparse-model-config.class-name",
			InnerField: "class_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "sparse-model-config.model-type",
			Usage:      "The sparse model type to use. 'bm25' uses Qdrant's FastEmbed BM25 model (default for new pipelines), 'splade' uses HuggingFace Splade model, 'auto' selects based on deployment mode (BYOC uses term frequency, Cloud uses Splade).",
			InnerField: "model_type",
		},
	},
})

var pipelinesUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update an existing pipeline's configuration.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "pipeline-id",
			Required:  true,
			PathParam: "pipeline_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "data-sink",
			Usage:    "Schema for creating a data sink.",
			BodyPath: "data_sink",
		},
		&requestflag.Flag[*string]{
			Name:     "data-sink-id",
			Usage:    "Data sink ID. When provided instead of data_sink, the data sink will be looked up by ID.",
			BodyPath: "data_sink_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "embedding-config",
			BodyPath: "embedding_config",
		},
		&requestflag.Flag[*string]{
			Name:     "embedding-model-config-id",
			Usage:    "Embedding model config ID. When provided instead of embedding_config, the embedding model config will be looked up by ID.",
			BodyPath: "embedding_model_config_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "llama-parse-parameters",
			BodyPath: "llama_parse_parameters",
		},
		&requestflag.Flag[*string]{
			Name:     "managed-pipeline-id",
			Usage:    "The ID of the ManagedPipeline this playground pipeline is linked to.",
			BodyPath: "managed_pipeline_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata-config",
			BodyPath: "metadata_config",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "preset-retrieval-parameters",
			Usage:    "Schema for the search params for an retrieval execution that can be preset for a pipeline.",
			BodyPath: "preset_retrieval_parameters",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "sparse-model-config",
			Usage:    "Configuration for sparse embedding models used in hybrid search.\n\nThis allows users to choose between Splade and BM25 models for\nsparse retrieval in managed data sinks.",
			BodyPath: "sparse_model_config",
		},
		&requestflag.Flag[*string]{
			Name:     "status",
			Usage:    "Status of the pipeline deployment.",
			BodyPath: "status",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "transform-config",
			Usage:    "Configuration for the transformation.",
			BodyPath: "transform_config",
		},
	},
	Action:          handlePipelinesUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"data-sink": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "data-sink.component",
			Usage:      "Component that implements the data sink",
			InnerField: "component",
		},
		&requestflag.InnerFlag[string]{
			Name:       "data-sink.name",
			Usage:      "The name of the data sink.",
			InnerField: "name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "data-sink.sink-type",
			Usage:      `Allowed values: "PINECONE", "POSTGRES", "QDRANT", "AZUREAI_SEARCH", "MONGODB_ATLAS", "MILVUS", "ASTRA_DB".`,
			InnerField: "sink_type",
		},
	},
	"llama-parse-parameters": {
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.adaptive-long-table",
			InnerField: "adaptive_long_table",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.aggressive-table-extraction",
			InnerField: "aggressive_table_extraction",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.annotate-links",
			InnerField: "annotate_links",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.auto-mode",
			InnerField: "auto_mode",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.auto-mode-configuration-json",
			InnerField: "auto_mode_configuration_json",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.auto-mode-trigger-on-image-in-page",
			InnerField: "auto_mode_trigger_on_image_in_page",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.auto-mode-trigger-on-regexp-in-page",
			InnerField: "auto_mode_trigger_on_regexp_in_page",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.auto-mode-trigger-on-table-in-page",
			InnerField: "auto_mode_trigger_on_table_in_page",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.auto-mode-trigger-on-text-in-page",
			InnerField: "auto_mode_trigger_on_text_in_page",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.azure-openai-api-version",
			InnerField: "azure_openai_api_version",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.azure-openai-deployment-name",
			InnerField: "azure_openai_deployment_name",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.azure-openai-endpoint",
			InnerField: "azure_openai_endpoint",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.azure-openai-key",
			InnerField: "azure_openai_key",
		},
		&requestflag.InnerFlag[*float64]{
			Name:       "llama-parse-parameters.bbox-bottom",
			InnerField: "bbox_bottom",
		},
		&requestflag.InnerFlag[*float64]{
			Name:       "llama-parse-parameters.bbox-left",
			InnerField: "bbox_left",
		},
		&requestflag.InnerFlag[*float64]{
			Name:       "llama-parse-parameters.bbox-right",
			InnerField: "bbox_right",
		},
		&requestflag.InnerFlag[*float64]{
			Name:       "llama-parse-parameters.bbox-top",
			InnerField: "bbox_top",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.bounding-box",
			InnerField: "bounding_box",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.compact-markdown-table",
			InnerField: "compact_markdown_table",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.complemental-formatting-instruction",
			InnerField: "complemental_formatting_instruction",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.confidence-score-effort",
			InnerField: "confidence_score_effort",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.content-guideline-instruction",
			InnerField: "content_guideline_instruction",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.continuous-mode",
			InnerField: "continuous_mode",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.disable-image-extraction",
			InnerField: "disable_image_extraction",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.disable-ocr",
			InnerField: "disable_ocr",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.disable-reconstruction",
			InnerField: "disable_reconstruction",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.do-not-cache",
			InnerField: "do_not_cache",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.do-not-unroll-columns",
			InnerField: "do_not_unroll_columns",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.enable-cost-optimizer",
			InnerField: "enable_cost_optimizer",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.extract-charts",
			InnerField: "extract_charts",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.extract-layout",
			InnerField: "extract_layout",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.extract-printed-page-number",
			InnerField: "extract_printed_page_number",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.fast-mode",
			InnerField: "fast_mode",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.formatting-instruction",
			InnerField: "formatting_instruction",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.gpt4o-api-key",
			InnerField: "gpt4o_api_key",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.gpt4o-mode",
			InnerField: "gpt4o_mode",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.guess-xlsx-sheet-name",
			InnerField: "guess_xlsx_sheet_name",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.hide-footers",
			InnerField: "hide_footers",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.hide-headers",
			InnerField: "hide_headers",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.high-res-ocr",
			InnerField: "high_res_ocr",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.html-make-all-elements-visible",
			InnerField: "html_make_all_elements_visible",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.html-remove-fixed-elements",
			InnerField: "html_remove_fixed_elements",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.html-remove-navigation-elements",
			InnerField: "html_remove_navigation_elements",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.http-proxy",
			InnerField: "http_proxy",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.ignore-document-elements-for-layout-detection",
			InnerField: "ignore_document_elements_for_layout_detection",
		},
		&requestflag.InnerFlag[any]{
			Name:       "llama-parse-parameters.images-to-save",
			InnerField: "images_to_save",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.inline-images-in-markdown",
			InnerField: "inline_images_in_markdown",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.input-s3-path",
			InnerField: "input_s3_path",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.input-s3-region",
			InnerField: "input_s3_region",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.input-url",
			InnerField: "input_url",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.internal-is-screenshot-job",
			InnerField: "internal_is_screenshot_job",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.invalidate-cache",
			InnerField: "invalidate_cache",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.is-formatting-instruction",
			InnerField: "is_formatting_instruction",
		},
		&requestflag.InnerFlag[*float64]{
			Name:       "llama-parse-parameters.job-timeout-extra-time-per-page-in-seconds",
			InnerField: "job_timeout_extra_time_per_page_in_seconds",
		},
		&requestflag.InnerFlag[*float64]{
			Name:       "llama-parse-parameters.job-timeout-in-seconds",
			InnerField: "job_timeout_in_seconds",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.keep-page-separator-when-merging-tables",
			InnerField: "keep_page_separator_when_merging_tables",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "llama-parse-parameters.languages",
			InnerField: "languages",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.layout-aware",
			InnerField: "layout_aware",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.line-level-bounding-box",
			InnerField: "line_level_bounding_box",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.markdown-table-multiline-header-separator",
			InnerField: "markdown_table_multiline_header_separator",
		},
		&requestflag.InnerFlag[*int64]{
			Name:       "llama-parse-parameters.max-pages",
			InnerField: "max_pages",
		},
		&requestflag.InnerFlag[*int64]{
			Name:       "llama-parse-parameters.max-pages-enforced",
			InnerField: "max_pages_enforced",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.merge-tables-across-pages-in-markdown",
			InnerField: "merge_tables_across_pages_in_markdown",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.model",
			InnerField: "model",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.outlined-table-extraction",
			InnerField: "outlined_table_extraction",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.output-pdf-of-document",
			InnerField: "output_pdf_of_document",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.output-s3-path-prefix",
			InnerField: "output_s3_path_prefix",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.output-s3-region",
			InnerField: "output_s3_region",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.output-tables-as-html",
			InnerField: "output_tables_as_HTML",
		},
		&requestflag.InnerFlag[*float64]{
			Name:       "llama-parse-parameters.page-error-tolerance",
			InnerField: "page_error_tolerance",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.page-footer-prefix",
			InnerField: "page_footer_prefix",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.page-footer-suffix",
			InnerField: "page_footer_suffix",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.page-header-prefix",
			InnerField: "page_header_prefix",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.page-header-suffix",
			InnerField: "page_header_suffix",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.page-prefix",
			InnerField: "page_prefix",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.page-separator",
			InnerField: "page_separator",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.page-suffix",
			InnerField: "page_suffix",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.parse-mode",
			Usage:      "Enum for representing the mode of parsing to be used.",
			InnerField: "parse_mode",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.parsing-instruction",
			InnerField: "parsing_instruction",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.precise-bounding-box",
			InnerField: "precise_bounding_box",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.premium-mode",
			InnerField: "premium_mode",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.presentation-out-of-bounds-content",
			InnerField: "presentation_out_of_bounds_content",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.presentation-skip-embedded-data",
			InnerField: "presentation_skip_embedded_data",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.preserve-layout-alignment-across-pages",
			InnerField: "preserve_layout_alignment_across_pages",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.preserve-very-small-text",
			InnerField: "preserve_very_small_text",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.preset",
			InnerField: "preset",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.priority",
			Usage:      "The priority for the request. This field may be ignored or overwritten depending on the organization tier.",
			InnerField: "priority",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.project-id",
			InnerField: "project_id",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.remove-hidden-text",
			InnerField: "remove_hidden_text",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.replace-failed-page-mode",
			Usage:      "Enum for representing the different available page error handling modes.",
			InnerField: "replace_failed_page_mode",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.replace-failed-page-with-error-message-prefix",
			InnerField: "replace_failed_page_with_error_message_prefix",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.replace-failed-page-with-error-message-suffix",
			InnerField: "replace_failed_page_with_error_message_suffix",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.save-images",
			InnerField: "save_images",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.skip-diagonal-text",
			InnerField: "skip_diagonal_text",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.specialized-chart-parsing-agentic",
			InnerField: "specialized_chart_parsing_agentic",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.specialized-chart-parsing-efficient",
			InnerField: "specialized_chart_parsing_efficient",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.specialized-chart-parsing-plus",
			InnerField: "specialized_chart_parsing_plus",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.specialized-image-parsing",
			InnerField: "specialized_image_parsing",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.spreadsheet-extract-sub-tables",
			InnerField: "spreadsheet_extract_sub_tables",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.spreadsheet-force-formula-computation",
			InnerField: "spreadsheet_force_formula_computation",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.spreadsheet-include-hidden-sheets",
			InnerField: "spreadsheet_include_hidden_sheets",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.strict-mode-buggy-font",
			InnerField: "strict_mode_buggy_font",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.strict-mode-image-extraction",
			InnerField: "strict_mode_image_extraction",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.strict-mode-image-ocr",
			InnerField: "strict_mode_image_ocr",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.strict-mode-reconstruction",
			InnerField: "strict_mode_reconstruction",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.structured-output",
			InnerField: "structured_output",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.structured-output-json-schema",
			InnerField: "structured_output_json_schema",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.structured-output-json-schema-name",
			InnerField: "structured_output_json_schema_name",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.system-prompt",
			InnerField: "system_prompt",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.system-prompt-append",
			InnerField: "system_prompt_append",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.take-screenshot",
			InnerField: "take_screenshot",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.target-pages",
			InnerField: "target_pages",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.tier",
			InnerField: "tier",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.use-vendor-multimodal-model",
			InnerField: "use_vendor_multimodal_model",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.user-prompt",
			InnerField: "user_prompt",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.vendor-multimodal-api-key",
			InnerField: "vendor_multimodal_api_key",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.vendor-multimodal-model-name",
			InnerField: "vendor_multimodal_model_name",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.version",
			InnerField: "version",
		},
		&requestflag.InnerFlag[any]{
			Name:       "llama-parse-parameters.webhook-configurations",
			Usage:      "Outbound webhook endpoints to notify on job status changes",
			InnerField: "webhook_configurations",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.webhook-url",
			InnerField: "webhook_url",
		},
	},
	"metadata-config": {
		&requestflag.InnerFlag[[]string]{
			Name:       "metadata-config.excluded-embed-metadata-keys",
			Usage:      "List of metadata keys to exclude from embeddings",
			InnerField: "excluded_embed_metadata_keys",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "metadata-config.excluded-llm-metadata-keys",
			Usage:      "List of metadata keys to exclude from LLM during retrieval",
			InnerField: "excluded_llm_metadata_keys",
		},
	},
	"preset-retrieval-parameters": {
		&requestflag.InnerFlag[*float64]{
			Name:       "preset-retrieval-parameters.alpha",
			Usage:      "Alpha value for hybrid retrieval to determine the weights between dense and sparse retrieval. 0 is sparse retrieval and 1 is dense retrieval.",
			InnerField: "alpha",
		},
		&requestflag.InnerFlag[string]{
			Name:       "preset-retrieval-parameters.class-name",
			InnerField: "class_name",
		},
		&requestflag.InnerFlag[*float64]{
			Name:       "preset-retrieval-parameters.dense-similarity-cutoff",
			Usage:      "Minimum similarity score wrt query for retrieval",
			InnerField: "dense_similarity_cutoff",
		},
		&requestflag.InnerFlag[*int64]{
			Name:       "preset-retrieval-parameters.dense-similarity-top-k",
			Usage:      "Number of nodes for dense retrieval.",
			InnerField: "dense_similarity_top_k",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "preset-retrieval-parameters.enable-reranking",
			Usage:      "Enable reranking for retrieval",
			InnerField: "enable_reranking",
		},
		&requestflag.InnerFlag[*int64]{
			Name:       "preset-retrieval-parameters.files-top-k",
			Usage:      "Number of files to retrieve (only for retrieval mode files_via_metadata and files_via_content).",
			InnerField: "files_top_k",
		},
		&requestflag.InnerFlag[*int64]{
			Name:       "preset-retrieval-parameters.rerank-top-n",
			Usage:      "Number of reranked nodes for returning.",
			InnerField: "rerank_top_n",
		},
		&requestflag.InnerFlag[string]{
			Name:       "preset-retrieval-parameters.retrieval-mode",
			Usage:      `Allowed values: "chunks", "files_via_metadata", "files_via_content", "auto_routed".`,
			InnerField: "retrieval_mode",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "preset-retrieval-parameters.retrieve-image-nodes",
			Usage:      "Whether to retrieve image nodes.",
			InnerField: "retrieve_image_nodes",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "preset-retrieval-parameters.retrieve-page-figure-nodes",
			Usage:      "Whether to retrieve page figure nodes.",
			InnerField: "retrieve_page_figure_nodes",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "preset-retrieval-parameters.retrieve-page-screenshot-nodes",
			Usage:      "Whether to retrieve page screenshot nodes.",
			InnerField: "retrieve_page_screenshot_nodes",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "preset-retrieval-parameters.search-filters",
			Usage:      "Metadata filters for vector stores.",
			InnerField: "search_filters",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "preset-retrieval-parameters.search-filters-inference-schema",
			Usage:      "JSON Schema that will be used to infer search_filters. Omit or leave as null to skip inference.",
			InnerField: "search_filters_inference_schema",
		},
		&requestflag.InnerFlag[*int64]{
			Name:       "preset-retrieval-parameters.sparse-similarity-top-k",
			Usage:      "Number of nodes for sparse retrieval.",
			InnerField: "sparse_similarity_top_k",
		},
	},
	"sparse-model-config": {
		&requestflag.InnerFlag[string]{
			Name:       "sparse-model-config.class-name",
			InnerField: "class_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "sparse-model-config.model-type",
			Usage:      "The sparse model type to use. 'bm25' uses Qdrant's FastEmbed BM25 model (default for new pipelines), 'splade' uses HuggingFace Splade model, 'auto' selects based on deployment mode (BYOC uses term frequency, Cloud uses Splade).",
			InnerField: "model_type",
		},
	},
})

var pipelinesList = cli.Command{
	Name:    "list",
	Usage:   "Search for pipelines by name, type, or project.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[*string]{
			Name:      "organization-id",
			QueryPath: "organization_id",
		},
		&requestflag.Flag[*string]{
			Name:      "pipeline-name",
			QueryPath: "pipeline_name",
		},
		&requestflag.Flag[*string]{
			Name:      "pipeline-type",
			Usage:     "Enum for representing the type of a pipeline",
			QueryPath: "pipeline_type",
		},
		&requestflag.Flag[*string]{
			Name:      "project-id",
			QueryPath: "project_id",
		},
		&requestflag.Flag[*string]{
			Name:      "project-name",
			QueryPath: "project_name",
		},
	},
	Action:          handlePipelinesList,
	HideHelpCommand: true,
}

var pipelinesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a pipeline and all associated resources.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "pipeline-id",
			Required:  true,
			PathParam: "pipeline_id",
		},
	},
	Action:          handlePipelinesDelete,
	HideHelpCommand: true,
}

var pipelinesGet = cli.Command{
	Name:    "get",
	Usage:   "Get a pipeline by ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "pipeline-id",
			Required:  true,
			PathParam: "pipeline_id",
		},
	},
	Action:          handlePipelinesGet,
	HideHelpCommand: true,
}

var pipelinesGetStatus = cli.Command{
	Name:    "get-status",
	Usage:   "Get the ingestion status of a managed pipeline.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "pipeline-id",
			Required:  true,
			PathParam: "pipeline_id",
		},
		&requestflag.Flag[*bool]{
			Name:      "full-details",
			QueryPath: "full_details",
		},
	},
	Action:          handlePipelinesGetStatus,
	HideHelpCommand: true,
}

var pipelinesRunSearch = requestflag.WithInnerFlags(cli.Command{
	Name:    "run-search",
	Usage:   "Run a retrieval query against a managed pipeline.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "pipeline-id",
			Required:  true,
			PathParam: "pipeline_id",
		},
		&requestflag.Flag[string]{
			Name:     "query",
			Usage:    "The query to retrieve against.",
			Required: true,
			BodyPath: "query",
		},
		&requestflag.Flag[*string]{
			Name:      "organization-id",
			QueryPath: "organization_id",
		},
		&requestflag.Flag[*string]{
			Name:      "project-id",
			QueryPath: "project_id",
		},
		&requestflag.Flag[*float64]{
			Name:     "alpha",
			Usage:    "Alpha value for hybrid retrieval to determine the weights between dense and sparse retrieval. 0 is sparse retrieval and 1 is dense retrieval.",
			BodyPath: "alpha",
		},
		&requestflag.Flag[string]{
			Name:     "class-name",
			Default:  "base_component",
			BodyPath: "class_name",
		},
		&requestflag.Flag[*float64]{
			Name:     "dense-similarity-cutoff",
			Usage:    "Minimum similarity score wrt query for retrieval",
			Default:  requestflag.Ptr[float64](0),
			BodyPath: "dense_similarity_cutoff",
		},
		&requestflag.Flag[*int64]{
			Name:     "dense-similarity-top-k",
			Usage:    "Number of nodes for dense retrieval.",
			Default:  requestflag.Ptr[int64](30),
			BodyPath: "dense_similarity_top_k",
		},
		&requestflag.Flag[*bool]{
			Name:     "enable-reranking",
			Usage:    "Enable reranking for retrieval",
			BodyPath: "enable_reranking",
		},
		&requestflag.Flag[*int64]{
			Name:     "files-top-k",
			Usage:    "Number of files to retrieve (only for retrieval mode files_via_metadata and files_via_content).",
			Default:  requestflag.Ptr[int64](1),
			BodyPath: "files_top_k",
		},
		&requestflag.Flag[*int64]{
			Name:     "rerank-top-n",
			Usage:    "Number of reranked nodes for returning.",
			Default:  requestflag.Ptr[int64](6),
			BodyPath: "rerank_top_n",
		},
		&requestflag.Flag[string]{
			Name:     "retrieval-mode",
			Usage:    `Allowed values: "chunks", "files_via_metadata", "files_via_content", "auto_routed".`,
			BodyPath: "retrieval_mode",
		},
		&requestflag.Flag[bool]{
			Name:     "retrieve-image-nodes",
			Usage:    "Whether to retrieve image nodes.",
			Default:  false,
			BodyPath: "retrieve_image_nodes",
		},
		&requestflag.Flag[bool]{
			Name:     "retrieve-page-figure-nodes",
			Usage:    "Whether to retrieve page figure nodes.",
			Default:  false,
			BodyPath: "retrieve_page_figure_nodes",
		},
		&requestflag.Flag[bool]{
			Name:     "retrieve-page-screenshot-nodes",
			Usage:    "Whether to retrieve page screenshot nodes.",
			Default:  false,
			BodyPath: "retrieve_page_screenshot_nodes",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "search-filters",
			Usage:    "Metadata filters for vector stores.",
			BodyPath: "search_filters",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "search-filters-inference-schema",
			Usage:    "JSON Schema that will be used to infer search_filters. Omit or leave as null to skip inference.",
			BodyPath: "search_filters_inference_schema",
		},
		&requestflag.Flag[*int64]{
			Name:     "sparse-similarity-top-k",
			Usage:    "Number of nodes for sparse retrieval.",
			Default:  requestflag.Ptr[int64](30),
			BodyPath: "sparse_similarity_top_k",
		},
	},
	Action:          handlePipelinesRunSearch,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"search-filters": {
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "search-filters.filters",
			InnerField: "filters",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "search-filters.condition",
			Usage:      "Vector store filter conditions to combine different filters.",
			InnerField: "condition",
		},
	},
})

var pipelinesUpsert = requestflag.WithInnerFlags(cli.Command{
	Name:    "upsert",
	Usage:   "Upsert a pipeline.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[*string]{
			Name:      "organization-id",
			QueryPath: "organization_id",
		},
		&requestflag.Flag[*string]{
			Name:      "project-id",
			QueryPath: "project_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "data-sink",
			Usage:    "Schema for creating a data sink.",
			BodyPath: "data_sink",
		},
		&requestflag.Flag[*string]{
			Name:     "data-sink-id",
			Usage:    "Data sink ID. When provided instead of data_sink, the data sink will be looked up by ID.",
			BodyPath: "data_sink_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "embedding-config",
			BodyPath: "embedding_config",
		},
		&requestflag.Flag[*string]{
			Name:     "embedding-model-config-id",
			Usage:    "Embedding model config ID. When provided instead of embedding_config, the embedding model config will be looked up by ID.",
			BodyPath: "embedding_model_config_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "llama-parse-parameters",
			BodyPath: "llama_parse_parameters",
		},
		&requestflag.Flag[*string]{
			Name:     "managed-pipeline-id",
			Usage:    "The ID of the ManagedPipeline this playground pipeline is linked to.",
			BodyPath: "managed_pipeline_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata-config",
			BodyPath: "metadata_config",
		},
		&requestflag.Flag[string]{
			Name:     "pipeline-type",
			Usage:    "Enum for representing the type of a pipeline",
			BodyPath: "pipeline_type",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "preset-retrieval-parameters",
			Usage:    "Schema for the search params for an retrieval execution that can be preset for a pipeline.",
			BodyPath: "preset_retrieval_parameters",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "sparse-model-config",
			Usage:    "Configuration for sparse embedding models used in hybrid search.\n\nThis allows users to choose between Splade and BM25 models for\nsparse retrieval in managed data sinks.",
			BodyPath: "sparse_model_config",
		},
		&requestflag.Flag[*string]{
			Name:     "status",
			Usage:    "Status of the pipeline deployment.",
			BodyPath: "status",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "transform-config",
			Usage:    "Configuration for the transformation.",
			BodyPath: "transform_config",
		},
	},
	Action:          handlePipelinesUpsert,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"data-sink": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "data-sink.component",
			Usage:      "Component that implements the data sink",
			InnerField: "component",
		},
		&requestflag.InnerFlag[string]{
			Name:       "data-sink.name",
			Usage:      "The name of the data sink.",
			InnerField: "name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "data-sink.sink-type",
			Usage:      `Allowed values: "PINECONE", "POSTGRES", "QDRANT", "AZUREAI_SEARCH", "MONGODB_ATLAS", "MILVUS", "ASTRA_DB".`,
			InnerField: "sink_type",
		},
	},
	"llama-parse-parameters": {
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.adaptive-long-table",
			InnerField: "adaptive_long_table",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.aggressive-table-extraction",
			InnerField: "aggressive_table_extraction",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.annotate-links",
			InnerField: "annotate_links",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.auto-mode",
			InnerField: "auto_mode",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.auto-mode-configuration-json",
			InnerField: "auto_mode_configuration_json",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.auto-mode-trigger-on-image-in-page",
			InnerField: "auto_mode_trigger_on_image_in_page",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.auto-mode-trigger-on-regexp-in-page",
			InnerField: "auto_mode_trigger_on_regexp_in_page",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.auto-mode-trigger-on-table-in-page",
			InnerField: "auto_mode_trigger_on_table_in_page",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.auto-mode-trigger-on-text-in-page",
			InnerField: "auto_mode_trigger_on_text_in_page",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.azure-openai-api-version",
			InnerField: "azure_openai_api_version",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.azure-openai-deployment-name",
			InnerField: "azure_openai_deployment_name",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.azure-openai-endpoint",
			InnerField: "azure_openai_endpoint",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.azure-openai-key",
			InnerField: "azure_openai_key",
		},
		&requestflag.InnerFlag[*float64]{
			Name:       "llama-parse-parameters.bbox-bottom",
			InnerField: "bbox_bottom",
		},
		&requestflag.InnerFlag[*float64]{
			Name:       "llama-parse-parameters.bbox-left",
			InnerField: "bbox_left",
		},
		&requestflag.InnerFlag[*float64]{
			Name:       "llama-parse-parameters.bbox-right",
			InnerField: "bbox_right",
		},
		&requestflag.InnerFlag[*float64]{
			Name:       "llama-parse-parameters.bbox-top",
			InnerField: "bbox_top",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.bounding-box",
			InnerField: "bounding_box",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.compact-markdown-table",
			InnerField: "compact_markdown_table",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.complemental-formatting-instruction",
			InnerField: "complemental_formatting_instruction",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.confidence-score-effort",
			InnerField: "confidence_score_effort",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.content-guideline-instruction",
			InnerField: "content_guideline_instruction",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.continuous-mode",
			InnerField: "continuous_mode",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.disable-image-extraction",
			InnerField: "disable_image_extraction",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.disable-ocr",
			InnerField: "disable_ocr",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.disable-reconstruction",
			InnerField: "disable_reconstruction",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.do-not-cache",
			InnerField: "do_not_cache",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.do-not-unroll-columns",
			InnerField: "do_not_unroll_columns",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.enable-cost-optimizer",
			InnerField: "enable_cost_optimizer",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.extract-charts",
			InnerField: "extract_charts",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.extract-layout",
			InnerField: "extract_layout",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.extract-printed-page-number",
			InnerField: "extract_printed_page_number",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.fast-mode",
			InnerField: "fast_mode",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.formatting-instruction",
			InnerField: "formatting_instruction",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.gpt4o-api-key",
			InnerField: "gpt4o_api_key",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.gpt4o-mode",
			InnerField: "gpt4o_mode",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.guess-xlsx-sheet-name",
			InnerField: "guess_xlsx_sheet_name",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.hide-footers",
			InnerField: "hide_footers",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.hide-headers",
			InnerField: "hide_headers",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.high-res-ocr",
			InnerField: "high_res_ocr",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.html-make-all-elements-visible",
			InnerField: "html_make_all_elements_visible",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.html-remove-fixed-elements",
			InnerField: "html_remove_fixed_elements",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.html-remove-navigation-elements",
			InnerField: "html_remove_navigation_elements",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.http-proxy",
			InnerField: "http_proxy",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.ignore-document-elements-for-layout-detection",
			InnerField: "ignore_document_elements_for_layout_detection",
		},
		&requestflag.InnerFlag[any]{
			Name:       "llama-parse-parameters.images-to-save",
			InnerField: "images_to_save",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.inline-images-in-markdown",
			InnerField: "inline_images_in_markdown",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.input-s3-path",
			InnerField: "input_s3_path",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.input-s3-region",
			InnerField: "input_s3_region",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.input-url",
			InnerField: "input_url",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.internal-is-screenshot-job",
			InnerField: "internal_is_screenshot_job",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.invalidate-cache",
			InnerField: "invalidate_cache",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.is-formatting-instruction",
			InnerField: "is_formatting_instruction",
		},
		&requestflag.InnerFlag[*float64]{
			Name:       "llama-parse-parameters.job-timeout-extra-time-per-page-in-seconds",
			InnerField: "job_timeout_extra_time_per_page_in_seconds",
		},
		&requestflag.InnerFlag[*float64]{
			Name:       "llama-parse-parameters.job-timeout-in-seconds",
			InnerField: "job_timeout_in_seconds",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.keep-page-separator-when-merging-tables",
			InnerField: "keep_page_separator_when_merging_tables",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "llama-parse-parameters.languages",
			InnerField: "languages",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.layout-aware",
			InnerField: "layout_aware",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.line-level-bounding-box",
			InnerField: "line_level_bounding_box",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.markdown-table-multiline-header-separator",
			InnerField: "markdown_table_multiline_header_separator",
		},
		&requestflag.InnerFlag[*int64]{
			Name:       "llama-parse-parameters.max-pages",
			InnerField: "max_pages",
		},
		&requestflag.InnerFlag[*int64]{
			Name:       "llama-parse-parameters.max-pages-enforced",
			InnerField: "max_pages_enforced",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.merge-tables-across-pages-in-markdown",
			InnerField: "merge_tables_across_pages_in_markdown",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.model",
			InnerField: "model",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.outlined-table-extraction",
			InnerField: "outlined_table_extraction",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.output-pdf-of-document",
			InnerField: "output_pdf_of_document",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.output-s3-path-prefix",
			InnerField: "output_s3_path_prefix",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.output-s3-region",
			InnerField: "output_s3_region",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.output-tables-as-html",
			InnerField: "output_tables_as_HTML",
		},
		&requestflag.InnerFlag[*float64]{
			Name:       "llama-parse-parameters.page-error-tolerance",
			InnerField: "page_error_tolerance",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.page-footer-prefix",
			InnerField: "page_footer_prefix",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.page-footer-suffix",
			InnerField: "page_footer_suffix",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.page-header-prefix",
			InnerField: "page_header_prefix",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.page-header-suffix",
			InnerField: "page_header_suffix",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.page-prefix",
			InnerField: "page_prefix",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.page-separator",
			InnerField: "page_separator",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.page-suffix",
			InnerField: "page_suffix",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.parse-mode",
			Usage:      "Enum for representing the mode of parsing to be used.",
			InnerField: "parse_mode",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.parsing-instruction",
			InnerField: "parsing_instruction",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.precise-bounding-box",
			InnerField: "precise_bounding_box",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.premium-mode",
			InnerField: "premium_mode",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.presentation-out-of-bounds-content",
			InnerField: "presentation_out_of_bounds_content",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.presentation-skip-embedded-data",
			InnerField: "presentation_skip_embedded_data",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.preserve-layout-alignment-across-pages",
			InnerField: "preserve_layout_alignment_across_pages",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.preserve-very-small-text",
			InnerField: "preserve_very_small_text",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.preset",
			InnerField: "preset",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.priority",
			Usage:      "The priority for the request. This field may be ignored or overwritten depending on the organization tier.",
			InnerField: "priority",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.project-id",
			InnerField: "project_id",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.remove-hidden-text",
			InnerField: "remove_hidden_text",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.replace-failed-page-mode",
			Usage:      "Enum for representing the different available page error handling modes.",
			InnerField: "replace_failed_page_mode",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.replace-failed-page-with-error-message-prefix",
			InnerField: "replace_failed_page_with_error_message_prefix",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.replace-failed-page-with-error-message-suffix",
			InnerField: "replace_failed_page_with_error_message_suffix",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.save-images",
			InnerField: "save_images",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.skip-diagonal-text",
			InnerField: "skip_diagonal_text",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.specialized-chart-parsing-agentic",
			InnerField: "specialized_chart_parsing_agentic",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.specialized-chart-parsing-efficient",
			InnerField: "specialized_chart_parsing_efficient",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.specialized-chart-parsing-plus",
			InnerField: "specialized_chart_parsing_plus",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.specialized-image-parsing",
			InnerField: "specialized_image_parsing",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.spreadsheet-extract-sub-tables",
			InnerField: "spreadsheet_extract_sub_tables",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.spreadsheet-force-formula-computation",
			InnerField: "spreadsheet_force_formula_computation",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.spreadsheet-include-hidden-sheets",
			InnerField: "spreadsheet_include_hidden_sheets",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.strict-mode-buggy-font",
			InnerField: "strict_mode_buggy_font",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.strict-mode-image-extraction",
			InnerField: "strict_mode_image_extraction",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.strict-mode-image-ocr",
			InnerField: "strict_mode_image_ocr",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.strict-mode-reconstruction",
			InnerField: "strict_mode_reconstruction",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.structured-output",
			InnerField: "structured_output",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.structured-output-json-schema",
			InnerField: "structured_output_json_schema",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.structured-output-json-schema-name",
			InnerField: "structured_output_json_schema_name",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.system-prompt",
			InnerField: "system_prompt",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.system-prompt-append",
			InnerField: "system_prompt_append",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.take-screenshot",
			InnerField: "take_screenshot",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.target-pages",
			InnerField: "target_pages",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.tier",
			InnerField: "tier",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "llama-parse-parameters.use-vendor-multimodal-model",
			InnerField: "use_vendor_multimodal_model",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.user-prompt",
			InnerField: "user_prompt",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.vendor-multimodal-api-key",
			InnerField: "vendor_multimodal_api_key",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.vendor-multimodal-model-name",
			InnerField: "vendor_multimodal_model_name",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.version",
			InnerField: "version",
		},
		&requestflag.InnerFlag[any]{
			Name:       "llama-parse-parameters.webhook-configurations",
			Usage:      "Outbound webhook endpoints to notify on job status changes",
			InnerField: "webhook_configurations",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "llama-parse-parameters.webhook-url",
			InnerField: "webhook_url",
		},
	},
	"metadata-config": {
		&requestflag.InnerFlag[[]string]{
			Name:       "metadata-config.excluded-embed-metadata-keys",
			Usage:      "List of metadata keys to exclude from embeddings",
			InnerField: "excluded_embed_metadata_keys",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "metadata-config.excluded-llm-metadata-keys",
			Usage:      "List of metadata keys to exclude from LLM during retrieval",
			InnerField: "excluded_llm_metadata_keys",
		},
	},
	"preset-retrieval-parameters": {
		&requestflag.InnerFlag[*float64]{
			Name:       "preset-retrieval-parameters.alpha",
			Usage:      "Alpha value for hybrid retrieval to determine the weights between dense and sparse retrieval. 0 is sparse retrieval and 1 is dense retrieval.",
			InnerField: "alpha",
		},
		&requestflag.InnerFlag[string]{
			Name:       "preset-retrieval-parameters.class-name",
			InnerField: "class_name",
		},
		&requestflag.InnerFlag[*float64]{
			Name:       "preset-retrieval-parameters.dense-similarity-cutoff",
			Usage:      "Minimum similarity score wrt query for retrieval",
			InnerField: "dense_similarity_cutoff",
		},
		&requestflag.InnerFlag[*int64]{
			Name:       "preset-retrieval-parameters.dense-similarity-top-k",
			Usage:      "Number of nodes for dense retrieval.",
			InnerField: "dense_similarity_top_k",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "preset-retrieval-parameters.enable-reranking",
			Usage:      "Enable reranking for retrieval",
			InnerField: "enable_reranking",
		},
		&requestflag.InnerFlag[*int64]{
			Name:       "preset-retrieval-parameters.files-top-k",
			Usage:      "Number of files to retrieve (only for retrieval mode files_via_metadata and files_via_content).",
			InnerField: "files_top_k",
		},
		&requestflag.InnerFlag[*int64]{
			Name:       "preset-retrieval-parameters.rerank-top-n",
			Usage:      "Number of reranked nodes for returning.",
			InnerField: "rerank_top_n",
		},
		&requestflag.InnerFlag[string]{
			Name:       "preset-retrieval-parameters.retrieval-mode",
			Usage:      `Allowed values: "chunks", "files_via_metadata", "files_via_content", "auto_routed".`,
			InnerField: "retrieval_mode",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "preset-retrieval-parameters.retrieve-image-nodes",
			Usage:      "Whether to retrieve image nodes.",
			InnerField: "retrieve_image_nodes",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "preset-retrieval-parameters.retrieve-page-figure-nodes",
			Usage:      "Whether to retrieve page figure nodes.",
			InnerField: "retrieve_page_figure_nodes",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "preset-retrieval-parameters.retrieve-page-screenshot-nodes",
			Usage:      "Whether to retrieve page screenshot nodes.",
			InnerField: "retrieve_page_screenshot_nodes",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "preset-retrieval-parameters.search-filters",
			Usage:      "Metadata filters for vector stores.",
			InnerField: "search_filters",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "preset-retrieval-parameters.search-filters-inference-schema",
			Usage:      "JSON Schema that will be used to infer search_filters. Omit or leave as null to skip inference.",
			InnerField: "search_filters_inference_schema",
		},
		&requestflag.InnerFlag[*int64]{
			Name:       "preset-retrieval-parameters.sparse-similarity-top-k",
			Usage:      "Number of nodes for sparse retrieval.",
			InnerField: "sparse_similarity_top_k",
		},
	},
	"sparse-model-config": {
		&requestflag.InnerFlag[string]{
			Name:       "sparse-model-config.class-name",
			InnerField: "class_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "sparse-model-config.model-type",
			Usage:      "The sparse model type to use. 'bm25' uses Qdrant's FastEmbed BM25 model (default for new pipelines), 'splade' uses HuggingFace Splade model, 'auto' selects based on deployment mode (BYOC uses term frequency, Cloud uses Splade).",
			InnerField: "model_type",
		},
	},
})

func handlePipelinesCreate(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatRepeat,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := llamacloud.PipelineNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Pipelines.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "pipelines create",
		Transform:      transform,
	})
}

func handlePipelinesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("pipeline-id") && len(unusedArgs) > 0 {
		cmd.Set("pipeline-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatRepeat,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := llamacloud.PipelineUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Pipelines.Update(
		ctx,
		cmd.Value("pipeline-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "pipelines update",
		Transform:      transform,
	})
}

func handlePipelinesList(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatRepeat,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := llamacloud.PipelineListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Pipelines.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "pipelines list",
		Transform:      transform,
	})
}

func handlePipelinesDelete(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("pipeline-id") && len(unusedArgs) > 0 {
		cmd.Set("pipeline-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatRepeat,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	return client.Pipelines.Delete(ctx, cmd.Value("pipeline-id").(string), options...)
}

func handlePipelinesGet(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("pipeline-id") && len(unusedArgs) > 0 {
		cmd.Set("pipeline-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatRepeat,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Pipelines.Get(ctx, cmd.Value("pipeline-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "pipelines get",
		Transform:      transform,
	})
}

func handlePipelinesGetStatus(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("pipeline-id") && len(unusedArgs) > 0 {
		cmd.Set("pipeline-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatRepeat,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := llamacloud.PipelineGetStatusParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Pipelines.GetStatus(
		ctx,
		cmd.Value("pipeline-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "pipelines get-status",
		Transform:      transform,
	})
}

func handlePipelinesRunSearch(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("pipeline-id") && len(unusedArgs) > 0 {
		cmd.Set("pipeline-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatRepeat,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := llamacloud.PipelineRunSearchParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Pipelines.RunSearch(
		ctx,
		cmd.Value("pipeline-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "pipelines run-search",
		Transform:      transform,
	})
}

func handlePipelinesUpsert(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatRepeat,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := llamacloud.PipelineUpsertParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Pipelines.Upsert(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "pipelines upsert",
		Transform:      transform,
	})
}
