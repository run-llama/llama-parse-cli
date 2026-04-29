// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/stainless-sdks/llamacloud-prod-cli/internal/apiquery"
	"github.com/stainless-sdks/llamacloud-prod-cli/internal/requestflag"
	"github.com/stainless-sdks/llamacloud-prod-go"
	"github.com/stainless-sdks/llamacloud-prod-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var betaSplitCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a document split job.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "document-input",
			Usage:    "Document input specification for beta API.",
			Required: true,
			BodyPath: "document_input",
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
			Usage:    "Split configuration with categories and splitting strategy.",
			BodyPath: "configuration",
		},
		&requestflag.Flag[*string]{
			Name:     "configuration-id",
			Usage:    "Saved split configuration ID.",
			BodyPath: "configuration_id",
		},
	},
	Action:          handleBetaSplitCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"document-input": {
		&requestflag.InnerFlag[string]{
			Name:       "document-input.type",
			Usage:      "Type of document input. Valid values are: file_id",
			InnerField: "type",
		},
		&requestflag.InnerFlag[string]{
			Name:       "document-input.value",
			Usage:      "Document identifier.",
			InnerField: "value",
		},
	},
	"configuration": {
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "configuration.categories",
			Usage:      "Categories to split documents into.",
			InnerField: "categories",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "configuration.splitting-strategy",
			Usage:      "Strategy for splitting documents.",
			InnerField: "splitting_strategy",
		},
	},
})

var betaSplitList = cli.Command{
	Name:    "list",
	Usage:   "List document split jobs.",
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
			Usage:     "Filter by job status (pending, processing, completed, failed, cancelled)",
			QueryPath: "status",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleBetaSplitList,
	HideHelpCommand: true,
}

var betaSplitGet = cli.Command{
	Name:    "get",
	Usage:   "Get a document split job.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "split-job-id",
			Required:  true,
			PathParam: "split_job_id",
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
	Action:          handleBetaSplitGet,
	HideHelpCommand: true,
}

func handleBetaSplitCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloudprod.BetaSplitNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Beta.Split.New(ctx, params, options...)
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
		Title:          "beta:split create",
		Transform:      transform,
	})
}

func handleBetaSplitList(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloudprod.BetaSplitListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Beta.Split.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "beta:split list",
			Transform:      transform,
		})
	} else {
		iter := client.Beta.Split.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "beta:split list",
			Transform:      transform,
		})
	}
}

func handleBetaSplitGet(ctx context.Context, cmd *cli.Command) error {
	client := llamacloudprod.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("split-job-id") && len(unusedArgs) > 0 {
		cmd.Set("split-job-id", unusedArgs[0])
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

	params := llamacloudprod.BetaSplitGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Beta.Split.Get(
		ctx,
		cmd.Value("split-job-id").(string),
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
		Title:          "beta:split get",
		Transform:      transform,
	})
}
