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

var dataSinksCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a new data sink.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "component",
			Usage:    "Component that implements the data sink",
			Required: true,
			BodyPath: "component",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "The name of the data sink.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "sink-type",
			Usage:    `Allowed values: "ASTRA_DB", "AZUREAI_SEARCH", "MILVUS", "MONGODB_ATLAS", "PINECONE", "POSTGRES", "QDRANT".`,
			Required: true,
			BodyPath: "sink_type",
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
	Action:          handleDataSinksCreate,
	HideHelpCommand: true,
}

var dataSinksUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a data sink by ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "data-sink-id",
			Required:  true,
			PathParam: "data_sink_id",
		},
		&requestflag.Flag[string]{
			Name:     "sink-type",
			Usage:    `Allowed values: "ASTRA_DB", "AZUREAI_SEARCH", "MILVUS", "MONGODB_ATLAS", "PINECONE", "POSTGRES", "QDRANT".`,
			Required: true,
			BodyPath: "sink_type",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "component",
			Usage:    "Component that implements the data sink",
			BodyPath: "component",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			Usage:    "The name of the data sink.",
			BodyPath: "name",
		},
	},
	Action:          handleDataSinksUpdate,
	HideHelpCommand: true,
}

var dataSinksList = cli.Command{
	Name:    "list",
	Usage:   "List data sinks for a given project.",
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
	Action:          handleDataSinksList,
	HideHelpCommand: true,
}

var dataSinksDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a data sink by ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "data-sink-id",
			Required:  true,
			PathParam: "data_sink_id",
		},
	},
	Action:          handleDataSinksDelete,
	HideHelpCommand: true,
}

var dataSinksGet = cli.Command{
	Name:    "get",
	Usage:   "Get a data sink by ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "data-sink-id",
			Required:  true,
			PathParam: "data_sink_id",
		},
	},
	Action:          handleDataSinksGet,
	HideHelpCommand: true,
}

func handleDataSinksCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloudprod.DataSinkNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.DataSinks.New(ctx, params, options...)
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
		Title:          "data-sinks create",
		Transform:      transform,
	})
}

func handleDataSinksUpdate(ctx context.Context, cmd *cli.Command) error {
	client := llamacloudprod.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("data-sink-id") && len(unusedArgs) > 0 {
		cmd.Set("data-sink-id", unusedArgs[0])
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

	params := llamacloudprod.DataSinkUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.DataSinks.Update(
		ctx,
		cmd.Value("data-sink-id").(string),
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
		Title:          "data-sinks update",
		Transform:      transform,
	})
}

func handleDataSinksList(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloudprod.DataSinkListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.DataSinks.List(ctx, params, options...)
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
		Title:          "data-sinks list",
		Transform:      transform,
	})
}

func handleDataSinksDelete(ctx context.Context, cmd *cli.Command) error {
	client := llamacloudprod.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("data-sink-id") && len(unusedArgs) > 0 {
		cmd.Set("data-sink-id", unusedArgs[0])
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

	return client.DataSinks.Delete(ctx, cmd.Value("data-sink-id").(string), options...)
}

func handleDataSinksGet(ctx context.Context, cmd *cli.Command) error {
	client := llamacloudprod.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("data-sink-id") && len(unusedArgs) > 0 {
		cmd.Set("data-sink-id", unusedArgs[0])
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
	_, err = client.DataSinks.Get(ctx, cmd.Value("data-sink-id").(string), options...)
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
		Title:          "data-sinks get",
		Transform:      transform,
	})
}
