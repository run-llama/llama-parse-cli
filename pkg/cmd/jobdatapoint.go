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

var jobDataPointsList = cli.Command{
	Name:    "list",
	Usage:   "Returns paginated job data points for the current project.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "job-type",
			Usage:     "Job type to query.",
			Required:  true,
			QueryPath: "job_type",
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
		&requestflag.Flag[int64]{
			Name:      "hours",
			Usage:     "Hours of history to include.",
			Default:   1,
			QueryPath: "hours",
		},
		&requestflag.Flag[*string]{
			Name:      "organization-id",
			QueryPath: "organization_id",
		},
		&requestflag.Flag[*int64]{
			Name:      "page-size",
			Usage:     "Number of items per page.",
			QueryPath: "page_size",
		},
		&requestflag.Flag[*string]{
			Name:      "page-token",
			Usage:     "Cursor token for the next page.",
			QueryPath: "page_token",
		},
		&requestflag.Flag[*string]{
			Name:      "project-id",
			QueryPath: "project_id",
		},
		&requestflag.Flag[any]{
			Name:      "status",
			Usage:     "Filter by status.",
			QueryPath: "status",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleJobDataPointsList,
	HideHelpCommand: true,
}

func handleJobDataPointsList(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloud.JobDataPointListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.JobDataPoints.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "job-data-points list",
			Transform:      transform,
		})
	} else {
		iter := client.JobDataPoints.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "job-data-points list",
			Transform:      transform,
		})
	}
}
