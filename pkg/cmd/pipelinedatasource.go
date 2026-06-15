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

var pipelinesDataSourcesUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update the configuration of a data source in a pipeline.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "pipeline-id",
			Required:  true,
			PathParam: "pipeline_id",
		},
		&requestflag.Flag[string]{
			Name:      "data-source-id",
			Required:  true,
			PathParam: "data_source_id",
		},
		&requestflag.Flag[*float64]{
			Name:     "sync-interval",
			Usage:    "The interval at which the data source should be synced.",
			BodyPath: "sync_interval",
		},
	},
	Action:          handlePipelinesDataSourcesUpdate,
	HideHelpCommand: true,
}

var pipelinesDataSourcesGetDataSources = cli.Command{
	Name:    "get-data-sources",
	Usage:   "Get data sources for a pipeline.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "pipeline-id",
			Required:  true,
			PathParam: "pipeline_id",
		},
	},
	Action:          handlePipelinesDataSourcesGetDataSources,
	HideHelpCommand: true,
}

var pipelinesDataSourcesGetStatus = cli.Command{
	Name:    "get-status",
	Usage:   "Get the status of a data source for a pipeline.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "pipeline-id",
			Required:  true,
			PathParam: "pipeline_id",
		},
		&requestflag.Flag[string]{
			Name:      "data-source-id",
			Required:  true,
			PathParam: "data_source_id",
		},
	},
	Action:          handlePipelinesDataSourcesGetStatus,
	HideHelpCommand: true,
}

var pipelinesDataSourcesSync = cli.Command{
	Name:    "sync",
	Usage:   "Run incremental ingestion: pull upstream changes from the data source into the\ndata sink.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "pipeline-id",
			Required:  true,
			PathParam: "pipeline_id",
		},
		&requestflag.Flag[string]{
			Name:      "data-source-id",
			Required:  true,
			PathParam: "data_source_id",
		},
		&requestflag.Flag[any]{
			Name:     "pipeline-file-id",
			BodyPath: "pipeline_file_ids",
		},
	},
	Action:          handlePipelinesDataSourcesSync,
	HideHelpCommand: true,
}

var pipelinesDataSourcesUpdateDataSources = requestflag.WithInnerFlags(cli.Command{
	Name:    "update-data-sources",
	Usage:   "Add data sources to a pipeline.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "pipeline-id",
			Required:  true,
			PathParam: "pipeline_id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "body",
			Required: true,
			BodyRoot: true,
		},
	},
	Action:          handlePipelinesDataSourcesUpdateDataSources,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"body": {
		&requestflag.InnerFlag[string]{
			Name:       "body.data-source-id",
			Usage:      "The ID of the data source.",
			InnerField: "data_source_id",
		},
		&requestflag.InnerFlag[*float64]{
			Name:       "body.sync-interval",
			Usage:      "The interval at which the data source should be synced. Valid values are: 21600, 43200, 86400",
			InnerField: "sync_interval",
		},
	},
})

func handlePipelinesDataSourcesUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloudprod.PipelineDataSourceUpdateParams{
		PipelineID: cmd.Value("pipeline-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Pipelines.DataSources.Update(
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
		Title:          "pipelines:data-sources update",
		Transform:      transform,
	})
}

func handlePipelinesDataSourcesGetDataSources(ctx context.Context, cmd *cli.Command) error {
	client := llamacloudprod.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("pipeline-id") && len(unusedArgs) > 0 {
		cmd.Set("pipeline-id", unusedArgs[0])
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
	_, err = client.Pipelines.DataSources.GetDataSources(ctx, cmd.Value("pipeline-id").(string), options...)
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
		Title:          "pipelines:data-sources get-data-sources",
		Transform:      transform,
	})
}

func handlePipelinesDataSourcesGetStatus(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloudprod.PipelineDataSourceGetStatusParams{
		PipelineID: cmd.Value("pipeline-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Pipelines.DataSources.GetStatus(
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
		Title:          "pipelines:data-sources get-status",
		Transform:      transform,
	})
}

func handlePipelinesDataSourcesSync(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloudprod.PipelineDataSourceSyncParams{
		PipelineID: cmd.Value("pipeline-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Pipelines.DataSources.Sync(
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
		Title:          "pipelines:data-sources sync",
		Transform:      transform,
	})
}

func handlePipelinesDataSourcesUpdateDataSources(ctx context.Context, cmd *cli.Command) error {
	client := llamacloudprod.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("pipeline-id") && len(unusedArgs) > 0 {
		cmd.Set("pipeline-id", unusedArgs[0])
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

	params := llamacloudprod.PipelineDataSourceUpdateDataSourcesParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Pipelines.DataSources.UpdateDataSources(
		ctx,
		cmd.Value("pipeline-id").(string),
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
		Title:          "pipelines:data-sources update-data-sources",
		Transform:      transform,
	})
}
