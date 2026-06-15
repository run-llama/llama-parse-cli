// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/run-llama/llama-parse-go"
	"github.com/run-llama/llama-parse-go/option"
	"github.com/stainless-sdks/llamacloud-prod-cli/internal/apiquery"
	"github.com/stainless-sdks/llamacloud-prod-cli/internal/requestflag"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var extractCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create an extraction job.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "file-input",
			Usage:    "File ID or parse job ID to extract from",
			Required: true,
			BodyPath: "file_input",
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
			Name:     "configuration",
			Usage:    "Extract configuration combining parse and extract settings.",
			BodyPath: "configuration",
		},
		&requestflag.Flag[*string]{
			Name:     "configuration-id",
			Usage:    "Saved configuration ID",
			BodyPath: "configuration_id",
		},
		&requestflag.Flag[any]{
			Name:     "webhook-configuration",
			Usage:    "Outbound webhook endpoints to notify on job status changes",
			BodyPath: "webhook_configurations",
		},
	},
	Action:          handleExtractCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"configuration": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "configuration.data-schema",
			Usage:      "JSON Schema defining the fields to extract. Validate with the /schema/validate endpoint first.",
			InnerField: "data_schema",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "configuration.cite-sources",
			Usage:      "Include citations in results",
			InnerField: "cite_sources",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "configuration.confidence-scores",
			Usage:      "Include confidence scores in results",
			InnerField: "confidence_scores",
		},
		&requestflag.InnerFlag[string]{
			Name:       "configuration.extraction-target",
			Usage:      "Granularity of extraction: per_doc returns one object per document, per_page returns one object per page, per_table_row returns one object per table row",
			InnerField: "extraction_target",
		},
		&requestflag.InnerFlag[*int64]{
			Name:       "configuration.max-pages",
			Usage:      "Maximum number of pages to process. Omit for no limit.",
			InnerField: "max_pages",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "configuration.parse-config-id",
			Usage:      "Saved parse configuration ID to control how the document is parsed before extraction",
			InnerField: "parse_config_id",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "configuration.parse-tier",
			Usage:      "Parse tier to use before extraction. Defaults to the extract tier if not specified.",
			InnerField: "parse_tier",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "configuration.system-prompt",
			Usage:      "Custom system prompt to guide extraction behavior",
			InnerField: "system_prompt",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "configuration.target-pages",
			Usage:      "Comma-separated page numbers or ranges to process (1-based). Omit to process all pages.",
			InnerField: "target_pages",
		},
		&requestflag.InnerFlag[string]{
			Name:       "configuration.tier",
			Usage:      "Extract tier: cost_effective (5 credits/page) or agentic (15 credits/page)",
			InnerField: "tier",
		},
		&requestflag.InnerFlag[string]{
			Name:       "configuration.version",
			Usage:      "Use 'latest' for the latest release for the selected tier or a date string (YYYY-MM-DD format) to pin to the nearest release at or before that date.",
			InnerField: "version",
		},
	},
	"webhook-configuration": {
		&requestflag.InnerFlag[any]{
			Name:                  "webhook-configuration.webhook-events",
			Usage:                 "Events to subscribe to (e.g. 'parse.success', 'extract.error'). If null, all events are delivered.",
			InnerField:            "webhook_events",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:                  "webhook-configuration.webhook-headers",
			Usage:                 "Custom HTTP headers sent with each webhook request (e.g. auth tokens)",
			InnerField:            "webhook_headers",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[*string]{
			Name:                  "webhook-configuration.webhook-output-format",
			Usage:                 "Response format sent to the webhook: 'string' (default) or 'json'",
			InnerField:            "webhook_output_format",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[*string]{
			Name:                  "webhook-configuration.webhook-url",
			Usage:                 "URL to receive webhook POST notifications",
			InnerField:            "webhook_url",
			OuterIsArrayOfObjects: true,
		},
	},
})

var extractList = cli.Command{
	Name:    "list",
	Usage:   "List extraction jobs with optional filtering and pagination.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[*string]{
			Name:      "configuration-id",
			Usage:     "Filter by configuration ID",
			QueryPath: "configuration_id",
		},
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
		&requestflag.Flag[*string]{
			Name:      "document-input-type",
			Usage:     "Filter by document input type (file_id or parse_job_id)",
			QueryPath: "document_input_type",
		},
		&requestflag.Flag[*string]{
			Name:      "document-input-value",
			Usage:     "Deprecated: use file_input instead",
			QueryPath: "document_input_value",
		},
		&requestflag.Flag[[]string]{
			Name:      "expand",
			Usage:     "Additional fields to include: configuration, extract_metadata",
			QueryPath: "expand",
		},
		&requestflag.Flag[*string]{
			Name:      "file-input",
			Usage:     "Filter by file input value",
			QueryPath: "file_input",
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
			Usage:     "Filter by status",
			QueryPath: "status",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleExtractList,
	HideHelpCommand: true,
}

var extractDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete an extraction job and its results.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "job-id",
			Required:  true,
			PathParam: "job_id",
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
	Action:          handleExtractDelete,
	HideHelpCommand: true,
}

var extractGenerateSchema = cli.Command{
	Name:    "generate-schema",
	Usage:   "Generate a JSON schema and return a product configuration request.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[*string]{
			Name:      "organization-id",
			QueryPath: "organization_id",
		},
		&requestflag.Flag[*string]{
			Name:      "project-id",
			QueryPath: "project_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "data-schema",
			Usage:    "Optional schema to validate, refine, or extend",
			BodyPath: "data_schema",
		},
		&requestflag.Flag[*string]{
			Name:     "file-id",
			Usage:    "Optional file ID to analyze for schema generation",
			BodyPath: "file_id",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			Usage:    "Name for the generated configuration (auto-generated if omitted)",
			BodyPath: "name",
		},
		&requestflag.Flag[*string]{
			Name:     "prompt",
			Usage:    "Natural language description of the data structure to extract",
			BodyPath: "prompt",
		},
	},
	Action:          handleExtractGenerateSchema,
	HideHelpCommand: true,
}

var extractGet = cli.Command{
	Name:    "get",
	Usage:   "Get a single extraction job by ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "job-id",
			Required:  true,
			PathParam: "job_id",
		},
		&requestflag.Flag[[]string]{
			Name:      "expand",
			Usage:     "Additional fields to include: configuration, extract_metadata",
			QueryPath: "expand",
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
	Action:          handleExtractGet,
	HideHelpCommand: true,
}

var extractValidateSchema = cli.Command{
	Name:    "validate-schema",
	Usage:   "Validate a JSON schema for extraction.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "data-schema",
			Usage:    "JSON Schema to validate for use with extract jobs",
			Required: true,
			BodyPath: "data_schema",
		},
	},
	Action:          handleExtractValidateSchema,
	HideHelpCommand: true,
}

func handleExtractCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloudprod.ExtractNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Extract.New(ctx, params, options...)
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
		Title:          "extract create",
		Transform:      transform,
	})
}

func handleExtractList(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloudprod.ExtractListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Extract.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "extract list",
			Transform:      transform,
		})
	} else {
		iter := client.Extract.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "extract list",
			Transform:      transform,
		})
	}
}

func handleExtractDelete(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloudprod.ExtractDeleteParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Extract.Delete(
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
		Title:          "extract delete",
		Transform:      transform,
	})
}

func handleExtractGenerateSchema(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloudprod.ExtractGenerateSchemaParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Extract.GenerateSchema(ctx, params, options...)
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
		Title:          "extract generate-schema",
		Transform:      transform,
	})
}

func handleExtractGet(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloudprod.ExtractGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Extract.Get(
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
		Title:          "extract get",
		Transform:      transform,
	})
}

func handleExtractValidateSchema(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloudprod.ExtractValidateSchemaParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Extract.ValidateSchema(ctx, params, options...)
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
		Title:          "extract validate-schema",
		Transform:      transform,
	})
}
