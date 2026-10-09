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

var alphaVerifyCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a Verify job.",
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
			Usage:    "Configuration for a Verify job.",
			BodyPath: "configuration",
		},
		&requestflag.Flag[*string]{
			Name:     "file-id",
			Usage:    "Deprecated: use file_input instead",
			BodyPath: "file_id",
		},
		&requestflag.Flag[*string]{
			Name:     "file-input",
			Usage:    "File ID of the document to analyze",
			BodyPath: "file_input",
		},
		&requestflag.Flag[*string]{
			Name:     "transaction-id",
			Usage:    "Idempotency key scoped to the project. Reusing a key returns the original job; the new request body is ignored.",
			BodyPath: "transaction_id",
		},
		&requestflag.Flag[any]{
			Name:     "webhook-configuration-id",
			Usage:    "IDs of saved webhook configurations to notify for this job.",
			BodyPath: "webhook_configuration_ids",
		},
		&requestflag.Flag[any]{
			Name:     "webhook-configuration",
			Usage:    "Outbound webhook endpoints to notify on job status changes",
			BodyPath: "webhook_configurations",
		},
	},
	Action:          handleAlphaVerifyCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"configuration": {
		&requestflag.InnerFlag[*string]{
			Name:       "configuration.target-pages",
			Usage:      "Comma-separated page numbers or ranges to analyze (1-based). Omit to analyze all pages. Ignored for non-PDF inputs.",
			InnerField: "target_pages",
		},
		&requestflag.InnerFlag[string]{
			Name:       "configuration.tier",
			Usage:      "Verify tier: 'fast' runs only the quick deterministic forensic checks (metadata, content integrity, container structure, pixel statistics); 'agentic' (default) runs the full pipeline including the learned detectors and the semantic review pass.",
			InnerField: "tier",
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
			Name:                  "webhook-configuration.webhook-signing-secret",
			Usage:                 "Shared signing secret used to sign webhook deliveries. When set, each request includes an HMAC-SHA256 signature of the request body in the 'LC-Signature' header (value 'sha256=<hex>'). Recompute the HMAC over the raw request body with this secret to verify the delivery is authentic.",
			InnerField:            "webhook_signing_secret",
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

var alphaVerifyList = cli.Command{
	Name:    "list",
	Usage:   "List Verify jobs with optional filtering and pagination.",
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
		&requestflag.Flag[[]string]{
			Name:      "expand",
			Usage:     "Optional fields to include (e.g. `result`).",
			QueryPath: "expand",
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
	Action:          handleAlphaVerifyList,
	HideHelpCommand: true,
}

var alphaVerifyCancel = cli.Command{
	Name:    "cancel",
	Usage:   "Cancel a running Verify job.",
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
	Action:          handleAlphaVerifyCancel,
	HideHelpCommand: true,
}

var alphaVerifyGet = cli.Command{
	Name:    "get",
	Usage:   "Get a Verify job by ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "job-id",
			Required:  true,
			PathParam: "job_id",
		},
		&requestflag.Flag[[]string]{
			Name:      "expand",
			Usage:     "Optional fields to include (e.g. `result`).",
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
	Action:          handleAlphaVerifyGet,
	HideHelpCommand: true,
}

var alphaVerifyGetDetails = cli.Command{
	Name:    "get-details",
	Usage:   "Get the raw per-signal detail for a completed Verify job.",
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
	Action:          handleAlphaVerifyGetDetails,
	HideHelpCommand: true,
}

func handleAlphaVerifyCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloud.AlphaVerifyNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Alpha.Verify.New(ctx, params, options...)
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
		Title:          "alpha:verify create",
		Transform:      transform,
	})
}

func handleAlphaVerifyList(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloud.AlphaVerifyListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Alpha.Verify.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "alpha:verify list",
			Transform:      transform,
		})
	} else {
		iter := client.Alpha.Verify.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "alpha:verify list",
			Transform:      transform,
		})
	}
}

func handleAlphaVerifyCancel(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
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

	params := llamacloud.AlphaVerifyCancelParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Alpha.Verify.Cancel(
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
		Title:          "alpha:verify cancel",
		Transform:      transform,
	})
}

func handleAlphaVerifyGet(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
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

	params := llamacloud.AlphaVerifyGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Alpha.Verify.Get(
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
		Title:          "alpha:verify get",
		Transform:      transform,
	})
}

func handleAlphaVerifyGetDetails(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
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

	params := llamacloud.AlphaVerifyGetDetailsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Alpha.Verify.GetDetails(
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
		Title:          "alpha:verify get-details",
		Transform:      transform,
	})
}
