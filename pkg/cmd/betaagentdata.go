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

var betaAgentDataCreate = cli.Command{
	Name:    "create",
	Usage:   "Create new agent data.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "data",
			Required: true,
			BodyPath: "data",
		},
		&requestflag.Flag[string]{
			Name:     "deployment-name",
			Required: true,
			BodyPath: "deployment_name",
		},
		&requestflag.Flag[*string]{
			Name:      "organization-id",
			QueryPath: "organization_id",
		},
		&requestflag.Flag[*string]{
			Name:      "project-id",
			QueryPath: "project_id",
		},
		&requestflag.Flag[string]{
			Name:     "collection",
			Default:  "default",
			BodyPath: "collection",
		},
	},
	Action:          handleBetaAgentDataCreate,
	HideHelpCommand: true,
}

var betaAgentDataUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update agent data by ID (overwrites).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "item-id",
			Required:  true,
			PathParam: "item_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "data",
			Required: true,
			BodyPath: "data",
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
	Action:          handleBetaAgentDataUpdate,
	HideHelpCommand: true,
}

var betaAgentDataDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete agent data by ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "item-id",
			Required:  true,
			PathParam: "item_id",
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
	Action:          handleBetaAgentDataDelete,
	HideHelpCommand: true,
}

var betaAgentDataAggregate = cli.Command{
	Name:    "aggregate",
	Usage:   "Aggregate agent data with grouping and optional counting/first item retrieval.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "deployment-name",
			Usage:    "The agent deployment's name to aggregate data for",
			Required: true,
			BodyPath: "deployment_name",
		},
		&requestflag.Flag[*string]{
			Name:      "organization-id",
			QueryPath: "organization_id",
		},
		&requestflag.Flag[*string]{
			Name:      "project-id",
			QueryPath: "project_id",
		},
		&requestflag.Flag[string]{
			Name:     "collection",
			Usage:    "The logical agent data collection to aggregate data for",
			Default:  "default",
			BodyPath: "collection",
		},
		&requestflag.Flag[*bool]{
			Name:     "count",
			Usage:    "Whether to count the number of items in each group",
			Default:  requestflag.Ptr[bool](false),
			BodyPath: "count",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "filter",
			Usage:    "A filter object or expression that filters resources listed in the response.",
			BodyPath: "filter",
		},
		&requestflag.Flag[*bool]{
			Name:     "first",
			Usage:    "Whether to return the first item in each group (Sorted by created_at)",
			Default:  requestflag.Ptr[bool](false),
			BodyPath: "first",
		},
		&requestflag.Flag[any]{
			Name:     "group-by",
			Usage:    "The fields to group by. If empty, the entire dataset is grouped on. e.g. if left out, can be used for simple count operations",
			BodyPath: "group_by",
		},
		&requestflag.Flag[*int64]{
			Name:     "offset",
			Usage:    "The offset to start from. If not provided, the first page is returned",
			Default:  requestflag.Ptr[int64](0),
			BodyPath: "offset",
		},
		&requestflag.Flag[*string]{
			Name:     "order-by",
			Usage:    "A comma-separated list of fields to order by, sorted in ascending order. Use 'field_name desc' to specify descending order.",
			BodyPath: "order_by",
		},
		&requestflag.Flag[*int64]{
			Name:     "page-size",
			Usage:    "The maximum number of items to return. The service may return fewer than this value. If unspecified, a default page size will be used. The maximum value is typically 1000; values above this will be coerced to the maximum.",
			BodyPath: "page_size",
		},
		&requestflag.Flag[*string]{
			Name:     "page-token",
			Usage:    "A page token, received from a previous list call. Provide this to retrieve the subsequent page.",
			BodyPath: "page_token",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleBetaAgentDataAggregate,
	HideHelpCommand: true,
}

var betaAgentDataDeleteByQuery = cli.Command{
	Name:    "delete-by-query",
	Usage:   "Bulk delete agent data by query (deployment_name, collection, optional filters).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "deployment-name",
			Usage:    "The agent deployment's name to delete data for",
			Required: true,
			BodyPath: "deployment_name",
		},
		&requestflag.Flag[*string]{
			Name:      "organization-id",
			QueryPath: "organization_id",
		},
		&requestflag.Flag[*string]{
			Name:      "project-id",
			QueryPath: "project_id",
		},
		&requestflag.Flag[string]{
			Name:     "collection",
			Usage:    "The logical agent data collection to delete from",
			Default:  "default",
			BodyPath: "collection",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "filter",
			Usage:    "Optional filters to select which items to delete",
			BodyPath: "filter",
		},
	},
	Action:          handleBetaAgentDataDeleteByQuery,
	HideHelpCommand: true,
}

var betaAgentDataGet = cli.Command{
	Name:    "get",
	Usage:   "Get agent data by ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "item-id",
			Required:  true,
			PathParam: "item_id",
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
	Action:          handleBetaAgentDataGet,
	HideHelpCommand: true,
}

var betaAgentDataSearch = cli.Command{
	Name:    "search",
	Usage:   "Search agent data with filtering, sorting, and pagination.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "deployment-name",
			Usage:    "The agent deployment's name to search within",
			Required: true,
			BodyPath: "deployment_name",
		},
		&requestflag.Flag[*string]{
			Name:      "organization-id",
			QueryPath: "organization_id",
		},
		&requestflag.Flag[*string]{
			Name:      "project-id",
			QueryPath: "project_id",
		},
		&requestflag.Flag[string]{
			Name:     "collection",
			Usage:    "The logical agent data collection to search within",
			Default:  "default",
			BodyPath: "collection",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "filter",
			Usage:    "A filter object or expression that filters resources listed in the response.",
			BodyPath: "filter",
		},
		&requestflag.Flag[bool]{
			Name:     "include-total",
			Usage:    "Whether to include the total number of items in the response",
			Default:  false,
			BodyPath: "include_total",
		},
		&requestflag.Flag[*int64]{
			Name:     "offset",
			Usage:    "The offset to start from. If not provided, the first page is returned",
			Default:  requestflag.Ptr[int64](0),
			BodyPath: "offset",
		},
		&requestflag.Flag[*string]{
			Name:     "order-by",
			Usage:    "A comma-separated list of fields to order by, sorted in ascending order. Use 'field_name desc' to specify descending order.",
			BodyPath: "order_by",
		},
		&requestflag.Flag[*int64]{
			Name:     "page-size",
			Usage:    "The maximum number of items to return. The service may return fewer than this value. If unspecified, a default page size will be used. The maximum value is typically 1000; values above this will be coerced to the maximum.",
			BodyPath: "page_size",
		},
		&requestflag.Flag[*string]{
			Name:     "page-token",
			Usage:    "A page token, received from a previous list call. Provide this to retrieve the subsequent page.",
			BodyPath: "page_token",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleBetaAgentDataSearch,
	HideHelpCommand: true,
}

func handleBetaAgentDataCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloud.BetaAgentDataNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Beta.AgentData.New(ctx, params, options...)
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
		Title:          "beta:agent-data create",
		Transform:      transform,
	})
}

func handleBetaAgentDataUpdate(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("item-id") && len(unusedArgs) > 0 {
		cmd.Set("item-id", unusedArgs[0])
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

	params := llamacloud.BetaAgentDataUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Beta.AgentData.Update(
		ctx,
		cmd.Value("item-id").(string),
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
		Title:          "beta:agent-data update",
		Transform:      transform,
	})
}

func handleBetaAgentDataDelete(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("item-id") && len(unusedArgs) > 0 {
		cmd.Set("item-id", unusedArgs[0])
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

	params := llamacloud.BetaAgentDataDeleteParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Beta.AgentData.Delete(
		ctx,
		cmd.Value("item-id").(string),
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
		Title:          "beta:agent-data delete",
		Transform:      transform,
	})
}

func handleBetaAgentDataAggregate(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloud.BetaAgentDataAggregateParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Beta.AgentData.Aggregate(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "beta:agent-data aggregate",
			Transform:      transform,
		})
	} else {
		iter := client.Beta.AgentData.AggregateAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "beta:agent-data aggregate",
			Transform:      transform,
		})
	}
}

func handleBetaAgentDataDeleteByQuery(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloud.BetaAgentDataDeleteByQueryParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Beta.AgentData.DeleteByQuery(ctx, params, options...)
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
		Title:          "beta:agent-data delete-by-query",
		Transform:      transform,
	})
}

func handleBetaAgentDataGet(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("item-id") && len(unusedArgs) > 0 {
		cmd.Set("item-id", unusedArgs[0])
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

	params := llamacloud.BetaAgentDataGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Beta.AgentData.Get(
		ctx,
		cmd.Value("item-id").(string),
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
		Title:          "beta:agent-data get",
		Transform:      transform,
	})
}

func handleBetaAgentDataSearch(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloud.BetaAgentDataSearchParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Beta.AgentData.Search(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "beta:agent-data search",
			Transform:      transform,
		})
	} else {
		iter := client.Beta.AgentData.SearchAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "beta:agent-data search",
			Transform:      transform,
		})
	}
}
