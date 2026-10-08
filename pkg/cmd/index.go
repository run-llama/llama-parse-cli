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

var indexesCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a searchable index over a source directory.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "source-directory-id",
			Usage:    "ID of the source directory containing your documents.",
			Required: true,
			BodyPath: "source_directory_id",
		},
		&requestflag.Flag[*string]{
			Name:      "organization-id",
			QueryPath: "organization_id",
		},
		&requestflag.Flag[*string]{
			Name:      "project-id",
			QueryPath: "project_id",
		},
		&requestflag.Flag[*string]{
			Name:     "description",
			Usage:    "Optional description of the index.",
			BodyPath: "description",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			Usage:    "Optional display name for the index. If omitted, the index is named after the source directory.",
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:     "product",
			Usage:    "Product configurations for syncing. Omit to use a default parse configuration. Include an explicit entry per product type (e.g. parse, extract) to override the default.",
			BodyPath: "products",
		},
		&requestflag.Flag[any]{
			Name:     "store-attachment",
			Usage:    "Attachment kinds to store alongside parsed output. Each entry must be one of: screenshots, items. For example, ['screenshots'] renders and stores per-page screenshots; ['items'] stores structured items with bounding boxes. Omit or pass an empty list to skip attachments.",
			BodyPath: "store_attachments",
		},
		&requestflag.Flag[string]{
			Name:     "sync-frequency",
			Usage:    "How often to re-run the sync. One of: manual, daily, on_source_change. Defaults to manual.",
			Default:  "manual",
			BodyPath: "sync_frequency",
		},
		&requestflag.Flag[string]{
			Name:     "vector-target",
			Usage:    "Vector export destination for the index. 'DEFAULT' exports to the managed vector DB destination resolved from configuration. 'DISABLED' skips vector export — the export destination falls back to 'Download'.",
			Default:  "DEFAULT",
			BodyPath: "vector_target",
		},
	},
	Action:          handleIndexesCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"product": {
		&requestflag.InnerFlag[string]{
			Name:                  "product.product-config-id",
			Usage:                 "ID of the product configuration.",
			InnerField:            "product_config_id",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[string]{
			Name:                  "product.product-type",
			Usage:                 "Product type. One of: parse, extract.",
			InnerField:            "product_type",
			OuterIsArrayOfObjects: true,
		},
	},
})

var indexesList = cli.Command{
	Name:    "list",
	Usage:   "List indexes for the current project.",
	Suggest: true,
	Flags: []cli.Flag{
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
			Name:      "source-directory-id",
			QueryPath: "source_directory_id",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleIndexesList,
	HideHelpCommand: true,
}

var indexesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete an index.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "index-id",
			Required:  true,
			PathParam: "index_id",
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
	Action:          handleIndexesDelete,
	HideHelpCommand: true,
}

var indexesCancelSync = cli.Command{
	Name:    "cancel-sync",
	Usage:   "Cancel the running sync for an index. Returns 409 if no sync is running.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "index-id",
			Required:  true,
			PathParam: "index_id",
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
	Action:          handleIndexesCancelSync,
	HideHelpCommand: true,
}

var indexesGet = cli.Command{
	Name:    "get",
	Usage:   "Get an index by ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "index-id",
			Required:  true,
			PathParam: "index_id",
		},
		&requestflag.Flag[[]string]{
			Name:      "expand",
			Usage:     "Fields to expand. Supported value: sync_in_progress.",
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
	Action:          handleIndexesGet,
	HideHelpCommand: true,
}

var indexesSync = cli.Command{
	Name:    "sync",
	Usage:   "Trigger a sync and export for an existing index, re-parsing changed files and\nexporting updated chunks.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "index-id",
			Required:  true,
			PathParam: "index_id",
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
	Action:          handleIndexesSync,
	HideHelpCommand: true,
}

func handleIndexesCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloud.IndexNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Indexes.New(ctx, params, options...)
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
		Title:          "indexes create",
		Transform:      transform,
	})
}

func handleIndexesList(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloud.IndexListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Indexes.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "indexes list",
			Transform:      transform,
		})
	} else {
		iter := client.Indexes.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "indexes list",
			Transform:      transform,
		})
	}
}

func handleIndexesDelete(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("index-id") && len(unusedArgs) > 0 {
		cmd.Set("index-id", unusedArgs[0])
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

	params := llamacloud.IndexDeleteParams{}

	return client.Indexes.Delete(
		ctx,
		cmd.Value("index-id").(string),
		params,
		options...,
	)
}

func handleIndexesCancelSync(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("index-id") && len(unusedArgs) > 0 {
		cmd.Set("index-id", unusedArgs[0])
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

	params := llamacloud.IndexCancelSyncParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Indexes.CancelSync(
		ctx,
		cmd.Value("index-id").(string),
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
		Title:          "indexes cancel-sync",
		Transform:      transform,
	})
}

func handleIndexesGet(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("index-id") && len(unusedArgs) > 0 {
		cmd.Set("index-id", unusedArgs[0])
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

	params := llamacloud.IndexGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Indexes.Get(
		ctx,
		cmd.Value("index-id").(string),
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
		Title:          "indexes get",
		Transform:      transform,
	})
}

func handleIndexesSync(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("index-id") && len(unusedArgs) > 0 {
		cmd.Set("index-id", unusedArgs[0])
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

	params := llamacloud.IndexSyncParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Indexes.Sync(
		ctx,
		cmd.Value("index-id").(string),
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
		Title:          "indexes sync",
		Transform:      transform,
	})
}
