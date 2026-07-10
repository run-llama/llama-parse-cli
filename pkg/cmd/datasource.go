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

var dataSourcesCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a new data source.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "component",
			Usage:    "Component that implements the data source",
			Required: true,
			BodyPath: "component",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "The name of the data source.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "source-type",
			Usage:    `Allowed values: "S3", "AZURE_STORAGE_BLOB", "GOOGLE_DRIVE", "MICROSOFT_ONEDRIVE", "MICROSOFT_SHAREPOINT", "SLACK", "NOTION_PAGE", "CONFLUENCE", "JIRA", "JIRA_V2", "BOX".`,
			Required: true,
			BodyPath: "source_type",
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
			Name:     "brokered-connection-id",
			Usage:    "Reference to a brokered managed-OAuth connection backing this source.",
			BodyPath: "brokered_connection_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "custom-metadata",
			Usage:    "Custom metadata that will be present on all data loaded from the data source",
			BodyPath: "custom_metadata",
		},
	},
	Action:          handleDataSourcesCreate,
	HideHelpCommand: true,
}

var dataSourcesUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a data source by ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "data-source-id",
			Required:  true,
			PathParam: "data_source_id",
		},
		&requestflag.Flag[string]{
			Name:     "source-type",
			Usage:    `Allowed values: "S3", "AZURE_STORAGE_BLOB", "GOOGLE_DRIVE", "MICROSOFT_ONEDRIVE", "MICROSOFT_SHAREPOINT", "SLACK", "NOTION_PAGE", "CONFLUENCE", "JIRA", "JIRA_V2", "BOX".`,
			Required: true,
			BodyPath: "source_type",
		},
		&requestflag.Flag[*string]{
			Name:     "brokered-connection-id",
			Usage:    "Reference to a brokered managed-OAuth connection backing this source.",
			BodyPath: "brokered_connection_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "component",
			Usage:    "Component that implements the data source",
			BodyPath: "component",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "custom-metadata",
			Usage:    "Custom metadata that will be present on all data loaded from the data source",
			BodyPath: "custom_metadata",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			Usage:    "The name of the data source.",
			BodyPath: "name",
		},
	},
	Action:          handleDataSourcesUpdate,
	HideHelpCommand: true,
}

var dataSourcesList = cli.Command{
	Name:    "list",
	Usage:   "List data sources for a given project. If project_id is not provided, uses the\ndefault project.",
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
	},
	Action:          handleDataSourcesList,
	HideHelpCommand: true,
}

var dataSourcesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a data source by ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "data-source-id",
			Required:  true,
			PathParam: "data_source_id",
		},
	},
	Action:          handleDataSourcesDelete,
	HideHelpCommand: true,
}

var dataSourcesGet = cli.Command{
	Name:    "get",
	Usage:   "Get a data source by ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "data-source-id",
			Required:  true,
			PathParam: "data_source_id",
		},
	},
	Action:          handleDataSourcesGet,
	HideHelpCommand: true,
}

func handleDataSourcesCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloudprod.DataSourceNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.DataSources.New(ctx, params, options...)
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
		Title:          "data-sources create",
		Transform:      transform,
	})
}

func handleDataSourcesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := llamacloudprod.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("data-source-id") && len(unusedArgs) > 0 {
		cmd.Set("data-source-id", unusedArgs[0])
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

	params := llamacloudprod.DataSourceUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.DataSources.Update(
		ctx,
		cmd.Value("data-source-id").(string),
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
		Title:          "data-sources update",
		Transform:      transform,
	})
}

func handleDataSourcesList(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloudprod.DataSourceListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.DataSources.List(ctx, params, options...)
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
		Title:          "data-sources list",
		Transform:      transform,
	})
}

func handleDataSourcesDelete(ctx context.Context, cmd *cli.Command) error {
	client := llamacloudprod.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("data-source-id") && len(unusedArgs) > 0 {
		cmd.Set("data-source-id", unusedArgs[0])
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

	return client.DataSources.Delete(ctx, cmd.Value("data-source-id").(string), options...)
}

func handleDataSourcesGet(ctx context.Context, cmd *cli.Command) error {
	client := llamacloudprod.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("data-source-id") && len(unusedArgs) > 0 {
		cmd.Set("data-source-id", unusedArgs[0])
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.DataSources.Get(ctx, cmd.Value("data-source-id").(string), options...)
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
		Title:          "data-sources get",
		Transform:      transform,
	})
}
