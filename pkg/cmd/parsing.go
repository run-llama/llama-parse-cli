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

var parsingCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Parse a file by file ID or URL.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "tier",
			Usage:    "Parsing tier: 'fast' (rule-based, cheapest), 'cost_effective' (balanced), 'agentic' (AI-powered with custom prompts), or 'agentic_plus' (premium AI with highest accuracy)",
			Required: true,
			BodyPath: "tier",
		},
		&requestflag.Flag[string]{
			Name:     "version",
			Usage:    "Version for the selected tier. Use `latest`, or pin one of that tier's dated versions.\n\nCurrent `latest` by tier:\n- `fast`: `2026-06-15`\n- `cost_effective`: `2026-06-26`\n- `agentic`: `2026-06-18`\n- `agentic_plus`: `2026-07-08`\n\nFull list: `GET /api/v2/parse/versions`.",
			Required: true,
			BodyPath: "version",
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
			Name:     "agentic-options",
			Usage:    "Options for AI-powered parsing tiers (cost_effective, agentic, agentic_plus).\n\nThese options customize how the AI processes and interprets document content.\nOnly applicable when using non-fast tiers.",
			BodyPath: "agentic_options",
		},
		&requestflag.Flag[*string]{
			Name:     "client-name",
			Usage:    "Identifier for the client/application making the request. Used for analytics and debugging. Example: 'my-app-v2'",
			BodyPath: "client_name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "crop-box",
			Usage:    "Crop boundaries to process only a portion of each page. Values are ratios 0-1 from page edges",
			BodyPath: "crop_box",
		},
		&requestflag.Flag[*bool]{
			Name:     "disable-cache",
			Usage:    "Bypass result caching and force re-parsing. Use when document content may have changed or you need fresh results",
			BodyPath: "disable_cache",
		},
		&requestflag.Flag[any]{
			Name:     "fast-options",
			Usage:    "Options for fast tier parsing (rule-based, no AI).\n\nFast tier uses deterministic algorithms for text extraction without AI enhancement.\nIt's the fastest and most cost-effective option, best suited for simple documents\nwith standard layouts. Currently has no configurable options but reserved for\nfuture expansion.",
			BodyPath: "fast_options",
		},
		&requestflag.Flag[*string]{
			Name:     "file-id",
			Usage:    "ID of an existing file in the project to parse. Mutually exclusive with source_url",
			BodyPath: "file_id",
		},
		&requestflag.Flag[*string]{
			Name:     "http-proxy",
			Usage:    "HTTP/HTTPS proxy for fetching source_url. Ignored if using file_id",
			BodyPath: "http_proxy",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "input-options",
			Usage:    "Format-specific options (HTML, PDF, spreadsheet, presentation). Applied based on detected input file type",
			BodyPath: "input_options",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "output-options",
			Usage:    "Output formatting options for markdown, text, and extracted images",
			BodyPath: "output_options",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "page-ranges",
			Usage:    "Page selection: limit total pages or specify exact pages to process",
			BodyPath: "page_ranges",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "processing-control",
			Usage:    "Job execution controls including timeouts and failure thresholds",
			BodyPath: "processing_control",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "processing-options",
			Usage:    "Document processing options including OCR, table extraction, and chart parsing",
			BodyPath: "processing_options",
		},
		&requestflag.Flag[*string]{
			Name:     "source-url",
			Usage:    "Public URL of the document to parse. Mutually exclusive with file_id",
			BodyPath: "source_url",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "user-metadata",
			Usage:    "Arbitrary key/value tags to attach to this job. Returned when retrieving the job. Not searchable. Limits apply to the number of entries and the length of keys and values; oversized metadata is rejected.",
			BodyPath: "user_metadata",
		},
		&requestflag.Flag[any]{
			Name:     "webhook-configuration-id",
			Usage:    "IDs of saved webhook configurations to notify for this job.",
			BodyPath: "webhook_configuration_ids",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "webhook-configuration",
			Usage:    "Webhook endpoints for job status notifications. Multiple webhooks can be configured for different events or services",
			BodyPath: "webhook_configurations",
		},
	},
	Action:          handleParsingCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"agentic-options": {
		&requestflag.InnerFlag[*string]{
			Name:       "agentic-options.custom-prompt",
			Usage:      "Custom instructions for the AI parser. Use to guide extraction behavior, specify output formatting, or provide domain-specific context. Example: 'Extract financial tables with currency symbols. Format dates as YYYY-MM-DD.'",
			InnerField: "custom_prompt",
		},
	},
	"crop-box": {
		&requestflag.InnerFlag[*float64]{
			Name:       "crop-box.bottom",
			Usage:      "Bottom boundary as ratio (0-1). 0=top edge, 1=bottom edge. Content below this line is excluded",
			InnerField: "bottom",
		},
		&requestflag.InnerFlag[*float64]{
			Name:       "crop-box.left",
			Usage:      "Left boundary as ratio (0-1). 0=left edge, 1=right edge. Content left of this line is excluded",
			InnerField: "left",
		},
		&requestflag.InnerFlag[*float64]{
			Name:       "crop-box.right",
			Usage:      "Right boundary as ratio (0-1). 0=left edge, 1=right edge. Content right of this line is excluded",
			InnerField: "right",
		},
		&requestflag.InnerFlag[*float64]{
			Name:       "crop-box.top",
			Usage:      "Top boundary as ratio (0-1). 0=top edge, 1=bottom edge. Content above this line is excluded",
			InnerField: "top",
		},
	},
	"input-options": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "input-options.html",
			Usage:      "HTML/web page parsing options (applies to .html, .htm files)",
			InnerField: "html",
		},
		&requestflag.InnerFlag[any]{
			Name:       "input-options.pdf",
			Usage:      "PDF-specific parsing options (applies to .pdf files)",
			InnerField: "pdf",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "input-options.presentation",
			Usage:      "Presentation parsing options (applies to .pptx, .ppt, .odp, .key files)",
			InnerField: "presentation",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "input-options.spreadsheet",
			Usage:      "Spreadsheet parsing options (applies to .xlsx, .xls, .csv, .ods files)",
			InnerField: "spreadsheet",
		},
	},
	"output-options": {
		&requestflag.InnerFlag[[]string]{
			Name:       "output-options.additional-outputs",
			Usage:      "Optional additional output artifacts to save alongside the primary parse output. Each value opts in to generating and persisting one extra file; the empty list (default) saves none. The three accepted values are: 'stripped_md' — per-page markdown stripped of formatting (links, bold/italic, images, HTML), saved as JSON for full-text-search indexing; fetch via `expand=stripped_markdown_content_metadata`. 'concatenated_stripped_txt' — all stripped pages concatenated into a single plain-text file with `\\n\\n---\\n\\n` between pages, useful for feeding the document into search or embedding pipelines as one blob; fetch via `expand=concatenated_stripped_markdown_content_metadata`. 'word_bbox' — raw word-level bounding boxes (one JSON object per word, with page number and x/y/w/h coordinates) saved as JSONL, useful for highlighting or grounding extracted answers back to the source document; fetch via `expand=raw_words_content_metadata`.",
			InnerField: "additional_outputs",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "output-options.extract-printed-page-number",
			Usage:      "Extract the printed page number as it appears in the document (e.g., 'Page 5 of 10', 'v', 'A-3'). Useful for referencing original page numbers",
			InnerField: "extract_printed_page_number",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "output-options.granular-bboxes",
			Usage:      "Bounding-box granularity levels to compute for the parse. 'word' computes one bounding box per detected word; 'line' computes one per text line; 'cell' computes one per table cell. Multiple levels can be requested. Empty list (default) disables granular bboxes — only item-level layout boxes are returned on the result. When set, the computed boxes are not inlined on the result items; they are written to a separate `grounded_items` sidecar (JSONL, one row per page) and exposed as `result_content_metadata.grounded_items` (a presigned download URL) on the parse result. Each row matches the `GroundedJsonItem` shape.",
			InnerField: "granular_bboxes",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "output-options.images-to-save",
			Usage:      "Image categories to extract and save. Options: 'screenshot' (full page renders useful for visual QA), 'embedded' (images found within the document), 'layout' (cropped regions from layout detection like figures and diagrams). Empty list saves no images",
			InnerField: "images_to_save",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "output-options.markdown",
			Usage:      "Markdown formatting options including table styles and link annotations",
			InnerField: "markdown",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "output-options.spatial-text",
			Usage:      "Spatial text output options for preserving document layout structure",
			InnerField: "spatial_text",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "output-options.tables-as-spreadsheet",
			Usage:      "Options for exporting tables as XLSX spreadsheets",
			InnerField: "tables_as_spreadsheet",
		},
	},
	"page-ranges": {
		&requestflag.InnerFlag[*int64]{
			Name:       "page-ranges.max-pages",
			Usage:      "Maximum number of pages to process. Pages are processed in order starting from page 1. If both max_pages and target_pages are set, target_pages takes precedence",
			InnerField: "max_pages",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "page-ranges.target-pages",
			Usage:      "Comma-separated list of specific pages to process using 1-based indexing. Supports individual pages and ranges. Examples: '1,3,5' (pages 1, 3, 5), '1-5' (pages 1 through 5 inclusive), '1,3,5-8,10' (pages 1, 3, 5-8, and 10). Pages are sorted and deduplicated automatically. Duplicate pages cause an error",
			InnerField: "target_pages",
		},
	},
	"processing-control": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "processing-control.job-failure-conditions",
			Usage:      "Quality thresholds that determine when a job should fail vs complete with partial results",
			InnerField: "job_failure_conditions",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "processing-control.timeouts",
			Usage:      "Timeout settings for job execution. Increase for large or complex documents",
			InnerField: "timeouts",
		},
	},
	"processing-options": {
		&requestflag.InnerFlag[*bool]{
			Name:       "processing-options.aggressive-table-extraction",
			Usage:      "Use aggressive heuristics to detect table boundaries, even without visible borders. Useful for documents with borderless or complex tables",
			InnerField: "aggressive_table_extraction",
		},
		&requestflag.InnerFlag[any]{
			Name:       "processing-options.auto-mode-configuration",
			Usage:      "Conditional processing rules that apply different parsing options based on page content, document structure, or filename patterns. Each entry defines trigger conditions and the parsing configuration to apply when triggered",
			InnerField: "auto_mode_configuration",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "processing-options.cost-optimizer",
			Usage:      "Cost optimizer configuration for reducing parsing costs on simpler pages.\n\nWhen enabled, the parser analyzes each page and routes simpler pages to faster,\ncheaper processing while preserving quality for complex pages. Only works with\n'agentic' or 'agentic_plus' tiers.",
			InnerField: "cost_optimizer",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "processing-options.disable-heuristics",
			Usage:      "Disable automatic heuristics including outlined table extraction and adaptive long table handling. Use when heuristics produce incorrect results",
			InnerField: "disable_heuristics",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "processing-options.ignore",
			Usage:      "Options for ignoring specific text types (diagonal, hidden, text in images)",
			InnerField: "ignore",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "processing-options.ocr-parameters",
			Usage:      "OCR configuration including language detection settings",
			InnerField: "ocr_parameters",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "processing-options.specialized-chart-parsing",
			Usage:      "Enable AI-powered chart analysis. Modes: 'efficient' (fast, lower cost), 'agentic' (balanced), 'agentic_plus' (highest accuracy). Automatically enables extract_layout and precise_bounding_box when set",
			InnerField: "specialized_chart_parsing",
		},
	},
	"webhook-configuration": {
		&requestflag.InnerFlag[any]{
			Name:       "webhook-configuration.webhook-events",
			Usage:      "Events that trigger this webhook. Options: 'parse.success' (job completed), 'parse.error' (job failed), 'parse.partial_success' (some pages failed), 'parse.pending', 'parse.running', 'parse.cancelled'. If not specified, webhook fires for all events",
			InnerField: "webhook_events",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "webhook-configuration.webhook-headers",
			Usage:      "Custom HTTP headers to include in webhook requests. Use for authentication tokens or custom routing. Example: {'Authorization': 'Bearer xyz'}",
			InnerField: "webhook_headers",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "webhook-configuration.webhook-output-format",
			Usage:      "Format of the webhook payload body. 'string' (default) sends the payload as a JSON-encoded string; 'json' sends it as a JSON object.",
			InnerField: "webhook_output_format",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "webhook-configuration.webhook-signing-secret",
			Usage:      "Shared signing secret used to sign webhook deliveries. When set, each request includes an HMAC-SHA256 signature of the request body in the 'LC-Signature' header (value 'sha256=<hex>'). Recompute the HMAC over the raw request body with this secret to verify the delivery is authentic.",
			InnerField: "webhook_signing_secret",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "webhook-configuration.webhook-url",
			Usage:      "HTTPS URL to receive webhook POST requests. Must be publicly accessible",
			InnerField: "webhook_url",
		},
	},
})

var parsingList = cli.Command{
	Name:    "list",
	Usage:   "List parse jobs for the current project.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "created-at-on-or-after",
			Usage:     "Include items created at or after this timestamp (inclusive)",
			QueryPath: "created_at_on_or_after",
		},
		&requestflag.Flag[any]{
			Name:      "created-at-on-or-before",
			Usage:     "Include items created at or before this timestamp (inclusive)",
			QueryPath: "created_at_on_or_before",
		},
		&requestflag.Flag[any]{
			Name:      "job-id",
			Usage:     "Filter by specific job IDs",
			QueryPath: "job_ids",
		},
		&requestflag.Flag[*string]{
			Name:      "organization-id",
			QueryPath: "organization_id",
		},
		&requestflag.Flag[*int64]{
			Name:      "page-size",
			Usage:     "Number of items per page",
			QueryPath: "page_size",
		},
		&requestflag.Flag[*string]{
			Name:      "page-token",
			Usage:     "Token for pagination",
			QueryPath: "page_token",
		},
		&requestflag.Flag[*string]{
			Name:      "project-id",
			QueryPath: "project_id",
		},
		&requestflag.Flag[*string]{
			Name:      "status",
			Usage:     "Filter by job status (PENDING, RUNNING, COMPLETED, FAILED, CANCELLED)",
			QueryPath: "status",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleParsingList,
	HideHelpCommand: true,
}

var parsingGet = cli.Command{
	Name:    "get",
	Usage:   "Retrieve a parse job with optional expanded content.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "job-id",
			Required:  true,
			PathParam: "job_id",
		},
		&requestflag.Flag[[]string]{
			Name:      "expand",
			Usage:     "Fields to include: text, markdown, items, metadata, job_metadata, text_content_metadata, markdown_content_metadata, items_content_metadata, metadata_content_metadata, raw_words_content_metadata, xlsx_content_metadata, output_pdf_content_metadata, images_content_metadata. Metadata fields include presigned URLs.",
			QueryPath: "expand",
		},
		&requestflag.Flag[*string]{
			Name:      "image-filenames",
			Usage:     "Filter to specific image filenames (optional). Example: image_0.png,image_1.jpg",
			QueryPath: "image_filenames",
		},
		&requestflag.Flag[*string]{
			Name:      "organization-id",
			QueryPath: "organization_id",
		},
		&requestflag.Flag[*string]{
			Name:      "project-id",
			QueryPath: "project_id",
		},
	},
	Action:          handleParsingGet,
	HideHelpCommand: true,
}

func handleParsingCreate(ctx context.Context, cmd *cli.Command) error {
	client := llamacloudprod.NewClient(getDefaultRequestOptions(cmd)...)
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

	params := llamacloudprod.ParsingNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Parsing.New(ctx, params, options...)
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
		Title:          "parsing create",
		Transform:      transform,
	})
}

func handleParsingList(ctx context.Context, cmd *cli.Command) error {
	client := llamacloudprod.NewClient(getDefaultRequestOptions(cmd)...)
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

	params := llamacloudprod.ParsingListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Parsing.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "parsing list",
			Transform:      transform,
		})
	} else {
		iter := client.Parsing.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "parsing list",
			Transform:      transform,
		})
	}
}

func handleParsingGet(ctx context.Context, cmd *cli.Command) error {
	client := llamacloudprod.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("job-id") && len(unusedArgs) > 0 {
		cmd.Set("job-id", unusedArgs[0])
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

	params := llamacloudprod.ParsingGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Parsing.Get(
		ctx,
		cmd.Value("job-id").(string),
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
		Title:          "parsing get",
		Transform:      transform,
	})
}
