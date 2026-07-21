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

var pipelinesMetadataCreate = cli.Command{
	Name:    "create",
	Usage:   "Import metadata for a pipeline.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "pipeline-id",
			Required:  true,
			PathParam: "pipeline_id",
		},
		&requestflag.Flag[string]{
			Name:      "upload-file",
			Required:  true,
			BodyPath:  "upload_file",
			FileInput: true,
		},
	},
	Action:          handlePipelinesMetadataCreate,
	HideHelpCommand: true,
}

var pipelinesMetadataDeleteAll = cli.Command{
	Name:    "delete-all",
	Usage:   "Delete metadata for all files in a pipeline.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "pipeline-id",
			Required:  true,
			PathParam: "pipeline_id",
		},
	},
	Action:          handlePipelinesMetadataDeleteAll,
	HideHelpCommand: true,
}

func handlePipelinesMetadataCreate(ctx context.Context, cmd *cli.Command) error {
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
		MultipartFormEncoded,
		false,
	)
	if err != nil {
		return err
	}

	params := llamacloud.PipelineMetadataNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Pipelines.Metadata.New(
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
		Title:          "pipelines:metadata create",
		Transform:      transform,
	})
}

func handlePipelinesMetadataDeleteAll(ctx context.Context, cmd *cli.Command) error {
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

	return client.Pipelines.Metadata.DeleteAll(ctx, cmd.Value("pipeline-id").(string), options...)
}
