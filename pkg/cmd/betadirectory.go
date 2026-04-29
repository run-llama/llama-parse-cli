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

var betaDirectoriesCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a new directory within the specified project.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Human-readable name for the directory.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:      "organization-id",
			QueryPath: "organization_id",
		},
		&requestflag.Flag[any]{
			Name:      "project-id",
			QueryPath: "project_id",
		},
		&requestflag.Flag[any]{
			Name:     "data-source-id",
			Usage:    "Optional data source id the directory syncs from.",
			BodyPath: "data_source_id",
		},
		&requestflag.Flag[any]{
			Name:     "description",
			Usage:    "Optional description shown to users.",
			BodyPath: "description",
		},
	},
	Action:          handleBetaDirectoriesCreate,
	HideHelpCommand: true,
}

var betaDirectoriesUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update directory metadata.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "directory-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:      "organization-id",
			QueryPath: "organization_id",
		},
		&requestflag.Flag[any]{
			Name:      "project-id",
			QueryPath: "project_id",
		},
		&requestflag.Flag[any]{
			Name:     "description",
			Usage:    "Updated description for the directory.",
			BodyPath: "description",
		},
		&requestflag.Flag[any]{
			Name:     "name",
			Usage:    "Updated name for the directory.",
			BodyPath: "name",
		},
	},
	Action:          handleBetaDirectoriesUpdate,
	HideHelpCommand: true,
}

var betaDirectoriesList = cli.Command{
	Name:    "list",
	Usage:   "List Directories",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "data-source-id",
			QueryPath: "data_source_id",
		},
		&requestflag.Flag[bool]{
			Name:      "include-deleted",
			Default:   false,
			QueryPath: "include_deleted",
		},
		&requestflag.Flag[any]{
			Name:      "name",
			QueryPath: "name",
		},
		&requestflag.Flag[any]{
			Name:      "organization-id",
			QueryPath: "organization_id",
		},
		&requestflag.Flag[any]{
			Name:      "page-size",
			QueryPath: "page_size",
		},
		&requestflag.Flag[any]{
			Name:      "page-token",
			QueryPath: "page_token",
		},
		&requestflag.Flag[any]{
			Name:      "project-id",
			QueryPath: "project_id",
		},
		&requestflag.Flag[any]{
			Name:      "type",
			Usage:     `Allowed values: "user", "index".`,
			QueryPath: "type",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleBetaDirectoriesList,
	HideHelpCommand: true,
}

var betaDirectoriesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Permanently delete a directory.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "directory-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:      "organization-id",
			QueryPath: "organization_id",
		},
		&requestflag.Flag[any]{
			Name:      "project-id",
			QueryPath: "project_id",
		},
	},
	Action:          handleBetaDirectoriesDelete,
	HideHelpCommand: true,
}

var betaDirectoriesGet = cli.Command{
	Name:    "get",
	Usage:   "Retrieve a directory by its identifier.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "directory-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:      "organization-id",
			QueryPath: "organization_id",
		},
		&requestflag.Flag[any]{
			Name:      "project-id",
			QueryPath: "project_id",
		},
	},
	Action:          handleBetaDirectoriesGet,
	HideHelpCommand: true,
}

func handleBetaDirectoriesCreate(ctx context.Context, cmd *cli.Command) error {
	client := llamacloudprod.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := llamacloudprod.BetaDirectoryNewParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Beta.Directories.New(ctx, params, options...)
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
		Title:          "beta:directories create",
		Transform:      transform,
	})
}

func handleBetaDirectoriesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := llamacloudprod.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("directory-id") && len(unusedArgs) > 0 {
		cmd.Set("directory-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := llamacloudprod.BetaDirectoryUpdateParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Beta.Directories.Update(
		ctx,
		cmd.Value("directory-id").(string),
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
		Title:          "beta:directories update",
		Transform:      transform,
	})
}

func handleBetaDirectoriesList(ctx context.Context, cmd *cli.Command) error {
	client := llamacloudprod.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := llamacloudprod.BetaDirectoryListParams{}

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

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Beta.Directories.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "beta:directories list",
			Transform:      transform,
		})
	} else {
		iter := client.Beta.Directories.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "beta:directories list",
			Transform:      transform,
		})
	}
}

func handleBetaDirectoriesDelete(ctx context.Context, cmd *cli.Command) error {
	client := llamacloudprod.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("directory-id") && len(unusedArgs) > 0 {
		cmd.Set("directory-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := llamacloudprod.BetaDirectoryDeleteParams{}

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

	return client.Beta.Directories.Delete(
		ctx,
		cmd.Value("directory-id").(string),
		params,
		options...,
	)
}

func handleBetaDirectoriesGet(ctx context.Context, cmd *cli.Command) error {
	client := llamacloudprod.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("directory-id") && len(unusedArgs) > 0 {
		cmd.Set("directory-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := llamacloudprod.BetaDirectoryGetParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Beta.Directories.Get(
		ctx,
		cmd.Value("directory-id").(string),
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
		Title:          "beta:directories get",
		Transform:      transform,
	})
}
