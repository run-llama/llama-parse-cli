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

var betaSheetsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a spreadsheet parsing job.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "file-id",
			Usage:    "The ID of the file to parse",
			Required: true,
			BodyPath: "file_id",
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
			Name:     "config",
			Usage:    "Configuration for spreadsheet parsing and region extraction",
			BodyPath: "config",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "configuration",
			Usage:    "Configuration for spreadsheet parsing and region extraction",
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
	Action:          handleBetaSheetsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"config": {
		&requestflag.InnerFlag[*string]{
			Name:       "config.extraction-range",
			Usage:      "A1 notation of the range to extract a single region from. If None, the entire sheet is used.",
			InnerField: "extraction_range",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "config.flatten-hierarchical-tables",
			Usage:      "Return a flattened dataframe when a detected table is recognized as hierarchical.",
			InnerField: "flatten_hierarchical_tables",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "config.generate-additional-metadata",
			Usage:      "Whether to generate additional metadata (title, description) for each extracted region.",
			InnerField: "generate_additional_metadata",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "config.include-hidden-cells",
			Usage:      "Whether to include hidden cells when extracting regions from the spreadsheet.",
			InnerField: "include_hidden_cells",
		},
		&requestflag.InnerFlag[any]{
			Name:       "config.sheet-names",
			Usage:      "The names of the sheets to extract regions from. If empty, all sheets will be processed.",
			InnerField: "sheet_names",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "config.specialization",
			Usage:      "Optional specialization mode for domain-specific extraction. Supported values: 'financial-standard', 'financial-enhanced', 'financial-precise'. Default None uses the general-purpose pipeline.",
			InnerField: "specialization",
		},
		&requestflag.InnerFlag[string]{
			Name:       "config.table-merge-sensitivity",
			Usage:      "Influences how likely similar-looking regions are merged into a single table. Useful for spreadsheets that either have sparse tables (strong merging) or many distinct tables close together (weak merging).",
			InnerField: "table_merge_sensitivity",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "config.use-experimental-processing",
			Usage:      "Enables experimental processing. Accuracy may be impacted.",
			InnerField: "use_experimental_processing",
		},
	},
	"configuration": {
		&requestflag.InnerFlag[*string]{
			Name:       "configuration.extraction-range",
			Usage:      "A1 notation of the range to extract a single region from. If None, the entire sheet is used.",
			InnerField: "extraction_range",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "configuration.flatten-hierarchical-tables",
			Usage:      "Return a flattened dataframe when a detected table is recognized as hierarchical.",
			InnerField: "flatten_hierarchical_tables",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "configuration.generate-additional-metadata",
			Usage:      "Whether to generate additional metadata (title, description) for each extracted region.",
			InnerField: "generate_additional_metadata",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "configuration.include-hidden-cells",
			Usage:      "Whether to include hidden cells when extracting regions from the spreadsheet.",
			InnerField: "include_hidden_cells",
		},
		&requestflag.InnerFlag[any]{
			Name:       "configuration.sheet-names",
			Usage:      "The names of the sheets to extract regions from. If empty, all sheets will be processed.",
			InnerField: "sheet_names",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "configuration.specialization",
			Usage:      "Optional specialization mode for domain-specific extraction. Supported values: 'financial-standard', 'financial-enhanced', 'financial-precise'. Default None uses the general-purpose pipeline.",
			InnerField: "specialization",
		},
		&requestflag.InnerFlag[string]{
			Name:       "configuration.table-merge-sensitivity",
			Usage:      "Influences how likely similar-looking regions are merged into a single table. Useful for spreadsheets that either have sparse tables (strong merging) or many distinct tables close together (weak merging).",
			InnerField: "table_merge_sensitivity",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "configuration.use-experimental-processing",
			Usage:      "Enables experimental processing. Accuracy may be impacted.",
			InnerField: "use_experimental_processing",
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

var betaSheetsList = cli.Command{
	Name:    "list",
	Usage:   "List spreadsheet parsing jobs. Experimental: not production-ready and subject to\nchange.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[*string]{
			Name:      "configuration-id",
			Usage:     "Filter by saved configuration ID",
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
		&requestflag.Flag[bool]{
			Name:      "include-results",
			Default:   false,
			QueryPath: "include_results",
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
			QueryPath: "page_size",
		},
		&requestflag.Flag[*string]{
			Name:      "page-token",
			QueryPath: "page_token",
		},
		&requestflag.Flag[*string]{
			Name:      "project-id",
			QueryPath: "project_id",
		},
		&requestflag.Flag[*string]{
			Name:      "status",
			Usage:     "Filter by job status",
			QueryPath: "status",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleBetaSheetsList,
	HideHelpCommand: true,
}

var betaSheetsDeleteJob = cli.Command{
	Name:    "delete-job",
	Usage:   "Delete a spreadsheet parsing job and its associated data. Experimental: not\nproduction-ready and subject to change.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "spreadsheet-job-id",
			Required:  true,
			PathParam: "spreadsheet_job_id",
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
	Action:          handleBetaSheetsDeleteJob,
	HideHelpCommand: true,
}

var betaSheetsGet = cli.Command{
	Name:    "get",
	Usage:   "Get a spreadsheet parsing job. When `include_results=True` (default), embeds\nextracted regions and results if complete, skipping the separate `/results`\ncall. Experimental: not production-ready and subject to change.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "spreadsheet-job-id",
			Required:  true,
			PathParam: "spreadsheet_job_id",
		},
		&requestflag.Flag[[]string]{
			Name:      "expand",
			Usage:     "Optional fields to populate on the response. Valid values: metadata_state_transitions.",
			QueryPath: "expand",
		},
		&requestflag.Flag[bool]{
			Name:      "include-results",
			Default:   true,
			QueryPath: "include_results",
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
	Action:          handleBetaSheetsGet,
	HideHelpCommand: true,
}

var betaSheetsGetResultTable = cli.Command{
	Name:    "get-result-table",
	Usage:   "Generate a presigned URL to download a specific extracted region. Experimental:\nnot production-ready and subject to change.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "spreadsheet-job-id",
			Required:  true,
			PathParam: "spreadsheet_job_id",
		},
		&requestflag.Flag[string]{
			Name:      "region-id",
			Required:  true,
			PathParam: "region_id",
		},
		&requestflag.Flag[string]{
			Name:      "region-type",
			Usage:     `Allowed values: "table", "extra", "cell_metadata".`,
			Required:  true,
			PathParam: "region_type",
		},
		&requestflag.Flag[*int64]{
			Name:      "expires-at-seconds",
			QueryPath: "expires_at_seconds",
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
	Action:          handleBetaSheetsGetResultTable,
	HideHelpCommand: true,
}

func handleBetaSheetsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloudprod.BetaSheetNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Beta.Sheets.New(ctx, params, options...)
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
		Title:          "beta:sheets create",
		Transform:      transform,
	})
}

func handleBetaSheetsList(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloudprod.BetaSheetListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Beta.Sheets.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "beta:sheets list",
			Transform:      transform,
		})
	} else {
		iter := client.Beta.Sheets.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "beta:sheets list",
			Transform:      transform,
		})
	}
}

func handleBetaSheetsDeleteJob(ctx context.Context, cmd *cli.Command) error {
	client := llamacloudprod.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("spreadsheet-job-id") && len(unusedArgs) > 0 {
		cmd.Set("spreadsheet-job-id", unusedArgs[0])
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

	params := llamacloudprod.BetaSheetDeleteJobParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Beta.Sheets.DeleteJob(
		ctx,
		cmd.Value("spreadsheet-job-id").(string),
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
		Title:          "beta:sheets delete-job",
		Transform:      transform,
	})
}

func handleBetaSheetsGet(ctx context.Context, cmd *cli.Command) error {
	client := llamacloudprod.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("spreadsheet-job-id") && len(unusedArgs) > 0 {
		cmd.Set("spreadsheet-job-id", unusedArgs[0])
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

	params := llamacloudprod.BetaSheetGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Beta.Sheets.Get(
		ctx,
		cmd.Value("spreadsheet-job-id").(string),
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
		Title:          "beta:sheets get",
		Transform:      transform,
	})
}

func handleBetaSheetsGetResultTable(ctx context.Context, cmd *cli.Command) error {
	client := llamacloudprod.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("region-type") && len(unusedArgs) > 0 {
		cmd.Set("region-type", unusedArgs[0])
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

	params := llamacloudprod.BetaSheetGetResultTableParams{
		SpreadsheetJobID: cmd.Value("spreadsheet-job-id").(string),
		RegionID:         cmd.Value("region-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Beta.Sheets.GetResultTable(
		ctx,
		llamacloudprod.BetaSheetGetResultTableParamsRegionType(cmd.Value("region-type").(string)),
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
		Title:          "beta:sheets get-result-table",
		Transform:      transform,
	})
}
