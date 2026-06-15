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

var classifyCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a classify job.",
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
			Name:     "configuration",
			Usage:    "Configuration for a classify job.",
			BodyPath: "configuration",
		},
		&requestflag.Flag[*string]{
			Name:     "configuration-id",
			Usage:    "Saved configuration ID",
			BodyPath: "configuration_id",
		},
		&requestflag.Flag[*string]{
			Name:     "file-id",
			Usage:    "Deprecated: use file_input instead",
			BodyPath: "file_id",
		},
		&requestflag.Flag[*string]{
			Name:     "file-input",
			Usage:    "File ID or parse job ID to classify",
			BodyPath: "file_input",
		},
		&requestflag.Flag[*string]{
			Name:     "parse-job-id",
			Usage:    "Deprecated: use file_input instead",
			BodyPath: "parse_job_id",
		},
		&requestflag.Flag[*string]{
			Name:     "transaction-id",
			Usage:    "Idempotency key scoped to the project",
			BodyPath: "transaction_id",
		},
		&requestflag.Flag[any]{
			Name:     "webhook-configuration",
			Usage:    "Outbound webhook endpoints to notify on job status changes",
			BodyPath: "webhook_configurations",
		},
	},
	Action:          handleClassifyCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"configuration": {
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "configuration.rules",
			Usage:      "Classify rules to evaluate against the document (at least one required)",
			InnerField: "rules",
		},
		&requestflag.InnerFlag[string]{
			Name:       "configuration.mode",
			Usage:      "Classify execution mode",
			InnerField: "mode",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "configuration.parsing-configuration",
			Usage:      "Parsing configuration for classify jobs.",
			InnerField: "parsing_configuration",
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

var classifyList = cli.Command{
	Name:    "list",
	Usage:   "List classify jobs with optional filtering and pagination.",
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
			Usage:     "Filter by job status",
			QueryPath: "status",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleClassifyList,
	HideHelpCommand: true,
}

var classifyGet = cli.Command{
	Name:    "get",
	Usage:   "Get a classify job by ID.",
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
	Action:          handleClassifyGet,
	HideHelpCommand: true,
}

func handleClassifyCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloudprod.ClassifyNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Classify.New(ctx, params, options...)
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
		Title:          "classify create",
		Transform:      transform,
	})
}

func handleClassifyList(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloudprod.ClassifyListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Classify.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "classify list",
			Transform:      transform,
		})
	} else {
		iter := client.Classify.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "classify list",
			Transform:      transform,
		})
	}
}

func handleClassifyGet(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloudprod.ClassifyGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Classify.Get(
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
		Title:          "classify get",
		Transform:      transform,
	})
}
