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

var pipelinesFilesCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Add files to a pipeline.",
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
	Action:          handlePipelinesFilesCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"body": {
		&requestflag.InnerFlag[string]{
			Name:       "body.file-id",
			Usage:      "The ID of the file",
			InnerField: "file_id",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "body.custom-metadata",
			Usage:      "Custom metadata for the file",
			InnerField: "custom_metadata",
		},
	},
})

var pipelinesFilesUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a file for a pipeline.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "pipeline-id",
			Required:  true,
			PathParam: "pipeline_id",
		},
		&requestflag.Flag[string]{
			Name:      "file-id",
			Required:  true,
			PathParam: "file_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "custom-metadata",
			Usage:    "Custom metadata for the file",
			BodyPath: "custom_metadata",
		},
	},
	Action:          handlePipelinesFilesUpdate,
	HideHelpCommand: true,
}

var pipelinesFilesList = cli.Command{
	Name:    "list",
	Usage:   "List files for a pipeline with optional filtering, sorting, and pagination.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "pipeline-id",
			Required:  true,
			PathParam: "pipeline_id",
		},
		&requestflag.Flag[*string]{
			Name:      "data-source-id",
			QueryPath: "data_source_id",
		},
		&requestflag.Flag[*string]{
			Name:      "file-name-contains",
			QueryPath: "file_name_contains",
		},
		&requestflag.Flag[*int64]{
			Name:      "limit",
			QueryPath: "limit",
		},
		&requestflag.Flag[*int64]{
			Name:      "offset",
			QueryPath: "offset",
		},
		&requestflag.Flag[bool]{
			Name:      "only-manually-uploaded",
			Default:   false,
			QueryPath: "only_manually_uploaded",
		},
		&requestflag.Flag[*string]{
			Name:      "order-by",
			QueryPath: "order_by",
		},
		&requestflag.Flag[any]{
			Name:      "status",
			Usage:     "Filter by file statuses",
			QueryPath: "statuses",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handlePipelinesFilesList,
	HideHelpCommand: true,
}

var pipelinesFilesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a file from a pipeline.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "pipeline-id",
			Required:  true,
			PathParam: "pipeline_id",
		},
		&requestflag.Flag[string]{
			Name:      "file-id",
			Required:  true,
			PathParam: "file_id",
		},
	},
	Action:          handlePipelinesFilesDelete,
	HideHelpCommand: true,
}

var pipelinesFilesGetStatus = cli.Command{
	Name:    "get-status",
	Usage:   "Get status of a file for a pipeline.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "pipeline-id",
			Required:  true,
			PathParam: "pipeline_id",
		},
		&requestflag.Flag[string]{
			Name:      "file-id",
			Required:  true,
			PathParam: "file_id",
		},
	},
	Action:          handlePipelinesFilesGetStatus,
	HideHelpCommand: true,
}

var pipelinesFilesGetStatusCounts = cli.Command{
	Name:    "get-status-counts",
	Usage:   "Get files for a pipeline.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "pipeline-id",
			Required:  true,
			PathParam: "pipeline_id",
		},
		&requestflag.Flag[*string]{
			Name:      "data-source-id",
			QueryPath: "data_source_id",
		},
		&requestflag.Flag[bool]{
			Name:      "only-manually-uploaded",
			Default:   false,
			QueryPath: "only_manually_uploaded",
		},
	},
	Action:          handlePipelinesFilesGetStatusCounts,
	HideHelpCommand: true,
}

func handlePipelinesFilesCreate(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
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

	params := llamacloud.PipelineFileNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Pipelines.Files.New(
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
		Title:          "pipelines:files create",
		Transform:      transform,
	})
}

func handlePipelinesFilesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("file-id") && len(unusedArgs) > 0 {
		cmd.Set("file-id", unusedArgs[0])
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

	params := llamacloud.PipelineFileUpdateParams{
		PipelineID: cmd.Value("pipeline-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Pipelines.Files.Update(
		ctx,
		cmd.Value("file-id").(string),
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
		Title:          "pipelines:files update",
		Transform:      transform,
	})
}

func handlePipelinesFilesList(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
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

	params := llamacloud.PipelineFileListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Pipelines.Files.List(
			ctx,
			cmd.Value("pipeline-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "pipelines:files list",
			Transform:      transform,
		})
	} else {
		iter := client.Pipelines.Files.ListAutoPaging(
			ctx,
			cmd.Value("pipeline-id").(string),
			params,
			options...,
		)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "pipelines:files list",
			Transform:      transform,
		})
	}
}

func handlePipelinesFilesDelete(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("file-id") && len(unusedArgs) > 0 {
		cmd.Set("file-id", unusedArgs[0])
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

	params := llamacloud.PipelineFileDeleteParams{
		PipelineID: cmd.Value("pipeline-id").(string),
	}

	return client.Pipelines.Files.Delete(
		ctx,
		cmd.Value("file-id").(string),
		params,
		options...,
	)
}

func handlePipelinesFilesGetStatus(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("file-id") && len(unusedArgs) > 0 {
		cmd.Set("file-id", unusedArgs[0])
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

	params := llamacloud.PipelineFileGetStatusParams{
		PipelineID: cmd.Value("pipeline-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Pipelines.Files.GetStatus(
		ctx,
		cmd.Value("file-id").(string),
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
		Title:          "pipelines:files get-status",
		Transform:      transform,
	})
}

func handlePipelinesFilesGetStatusCounts(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
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

	params := llamacloud.PipelineFileGetStatusCountsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Pipelines.Files.GetStatusCounts(
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
		Title:          "pipelines:files get-status-counts",
		Transform:      transform,
	})
}
