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

var betaIndexesCreate = requestflag.WithInnerFlags(cli.Command{
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
	},
	Action:          handleBetaIndexesCreate,
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

var betaIndexesDelete = cli.Command{
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
	Action:          handleBetaIndexesDelete,
	HideHelpCommand: true,
}

var betaIndexesGet = cli.Command{
	Name:    "get",
	Usage:   "Get an index by ID.",
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
	Action:          handleBetaIndexesGet,
	HideHelpCommand: true,
}

var betaIndexesSync = cli.Command{
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
	Action:          handleBetaIndexesSync,
	HideHelpCommand: true,
}

func handleBetaIndexesCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloudprod.BetaIndexNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Beta.Indexes.New(ctx, params, options...)
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
		Title:          "beta:indexes create",
		Transform:      transform,
	})
}

func handleBetaIndexesDelete(ctx context.Context, cmd *cli.Command) error {
	client := llamacloudprod.NewClient(getDefaultRequestOptions(cmd)...)
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

	params := llamacloudprod.BetaIndexDeleteParams{}

	return client.Beta.Indexes.Delete(
		ctx,
		cmd.Value("index-id").(string),
		params,
		options...,
	)
}

func handleBetaIndexesGet(ctx context.Context, cmd *cli.Command) error {
	client := llamacloudprod.NewClient(getDefaultRequestOptions(cmd)...)
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

	params := llamacloudprod.BetaIndexGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Beta.Indexes.Get(
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
		Title:          "beta:indexes get",
		Transform:      transform,
	})
}

func handleBetaIndexesSync(ctx context.Context, cmd *cli.Command) error {
	client := llamacloudprod.NewClient(getDefaultRequestOptions(cmd)...)
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

	params := llamacloudprod.BetaIndexSyncParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Beta.Indexes.Sync(
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
		Title:          "beta:indexes sync",
		Transform:      transform,
	})
}
