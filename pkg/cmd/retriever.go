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

var retrieversCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a new Retriever.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "A name for the retriever tool. Will default to the pipeline name if not provided.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[*string]{
			Name:      "organization-id",
			QueryPath: "organization_id",
		},
		&requestflag.Flag[*string]{
			Name:      "project-id",
			QueryPath: "project_id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "pipeline",
			Usage:    "The pipelines this retriever uses.",
			BodyPath: "pipelines",
		},
	},
	Action:          handleRetrieversCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"pipeline": {
		&requestflag.InnerFlag[*string]{
			Name:       "pipeline.description",
			Usage:      "A description of the retriever tool.",
			InnerField: "description",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "pipeline.name",
			Usage:      "A name for the retriever tool. Will default to the pipeline name if not provided.",
			InnerField: "name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "pipeline.pipeline-id",
			Usage:      "The ID of the pipeline this tool uses.",
			InnerField: "pipeline_id",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "pipeline.preset-retrieval-parameters",
			Usage:      "Schema for the search params for an retrieval execution that can be preset for a pipeline.",
			InnerField: "preset_retrieval_parameters",
		},
	},
})

var retrieversUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update an existing Retriever.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "retriever-id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "pipeline",
			Usage:    "The pipelines this retriever uses.",
			Required: true,
			BodyPath: "pipelines",
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
			Name:     "name",
			Usage:    "A name for the retriever.",
			BodyPath: "name",
		},
	},
	Action:          handleRetrieversUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"pipeline": {
		&requestflag.InnerFlag[*string]{
			Name:                  "pipeline.description",
			Usage:                 "A description of the retriever tool.",
			InnerField:            "description",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[*string]{
			Name:                  "pipeline.name",
			Usage:                 "A name for the retriever tool. Will default to the pipeline name if not provided.",
			InnerField:            "name",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[string]{
			Name:                  "pipeline.pipeline-id",
			Usage:                 "The ID of the pipeline this tool uses.",
			InnerField:            "pipeline_id",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:                  "pipeline.preset-retrieval-parameters",
			Usage:                 "Schema for the search params for an retrieval execution that can be preset for a pipeline.",
			InnerField:            "preset_retrieval_parameters",
			OuterIsArrayOfObjects: true,
		},
	},
})

var retrieversList = cli.Command{
	Name:    "list",
	Usage:   "List Retrievers for a project.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[*string]{
			Name:      "name",
			QueryPath: "name",
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
	Action:          handleRetrieversList,
	HideHelpCommand: true,
}

var retrieversDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a Retriever by ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "retriever-id",
			Required: true,
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
	Action:          handleRetrieversDelete,
	HideHelpCommand: true,
}

var retrieversGet = cli.Command{
	Name:    "get",
	Usage:   "Get a Retriever by ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "retriever-id",
			Required: true,
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
	Action:          handleRetrieversGet,
	HideHelpCommand: true,
}

var retrieversSearch = requestflag.WithInnerFlags(cli.Command{
	Name:    "search",
	Usage:   "Retrieve data using specified pipelines without creating a persistent retriever.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "query",
			Usage:    "The query to retrieve against.",
			Required: true,
			BodyPath: "query",
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
			Name:     "mode",
			Usage:    "Enum for the mode of composite retrieval.",
			BodyPath: "mode",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "pipeline",
			Usage:    "The pipelines to use for retrieval.",
			BodyPath: "pipelines",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "rerank-config",
			BodyPath: "rerank_config",
		},
		&requestflag.Flag[*int64]{
			Name:     "rerank-top-n",
			Usage:    "(use rerank_config.top_n instead) The number of nodes to retrieve after reranking over retrieved nodes from all retrieval tools.",
			BodyPath: "rerank_top_n",
		},
	},
	Action:          handleRetrieversSearch,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"pipeline": {
		&requestflag.InnerFlag[*string]{
			Name:       "pipeline.description",
			Usage:      "A description of the retriever tool.",
			InnerField: "description",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "pipeline.name",
			Usage:      "A name for the retriever tool. Will default to the pipeline name if not provided.",
			InnerField: "name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "pipeline.pipeline-id",
			Usage:      "The ID of the pipeline this tool uses.",
			InnerField: "pipeline_id",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "pipeline.preset-retrieval-parameters",
			Usage:      "Schema for the search params for an retrieval execution that can be preset for a pipeline.",
			InnerField: "preset_retrieval_parameters",
		},
	},
	"rerank-config": {
		&requestflag.InnerFlag[int64]{
			Name:       "rerank-config.top-n",
			Usage:      "The number of nodes to retrieve after reranking over retrieved nodes from all retrieval tools.",
			InnerField: "top_n",
		},
		&requestflag.InnerFlag[string]{
			Name:       "rerank-config.type",
			Usage:      "The type of reranker to use.",
			InnerField: "type",
		},
	},
})

var retrieversUpsert = requestflag.WithInnerFlags(cli.Command{
	Name:    "upsert",
	Usage:   "Upsert a new Retriever.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "A name for the retriever tool. Will default to the pipeline name if not provided.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[*string]{
			Name:      "organization-id",
			QueryPath: "organization_id",
		},
		&requestflag.Flag[*string]{
			Name:      "project-id",
			QueryPath: "project_id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "pipeline",
			Usage:    "The pipelines this retriever uses.",
			BodyPath: "pipelines",
		},
	},
	Action:          handleRetrieversUpsert,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"pipeline": {
		&requestflag.InnerFlag[*string]{
			Name:       "pipeline.description",
			Usage:      "A description of the retriever tool.",
			InnerField: "description",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "pipeline.name",
			Usage:      "A name for the retriever tool. Will default to the pipeline name if not provided.",
			InnerField: "name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "pipeline.pipeline-id",
			Usage:      "The ID of the pipeline this tool uses.",
			InnerField: "pipeline_id",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "pipeline.preset-retrieval-parameters",
			Usage:      "Schema for the search params for an retrieval execution that can be preset for a pipeline.",
			InnerField: "preset_retrieval_parameters",
		},
	},
})

func handleRetrieversCreate(ctx context.Context, cmd *cli.Command) error {
	client := llamacloudprod.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := llamacloudprod.RetrieverNewParams{}

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
	_, err = client.Retrievers.New(ctx, params, options...)
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
		Title:          "retrievers create",
		Transform:      transform,
	})
}

func handleRetrieversUpdate(ctx context.Context, cmd *cli.Command) error {
	client := llamacloudprod.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("retriever-id") && len(unusedArgs) > 0 {
		cmd.Set("retriever-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := llamacloudprod.RetrieverUpdateParams{}

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
	_, err = client.Retrievers.Update(
		ctx,
		cmd.Value("retriever-id").(string),
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
		Title:          "retrievers update",
		Transform:      transform,
	})
}

func handleRetrieversList(ctx context.Context, cmd *cli.Command) error {
	client := llamacloudprod.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := llamacloudprod.RetrieverListParams{}

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
	_, err = client.Retrievers.List(ctx, params, options...)
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
		Title:          "retrievers list",
		Transform:      transform,
	})
}

func handleRetrieversDelete(ctx context.Context, cmd *cli.Command) error {
	client := llamacloudprod.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("retriever-id") && len(unusedArgs) > 0 {
		cmd.Set("retriever-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := llamacloudprod.RetrieverDeleteParams{}

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

	return client.Retrievers.Delete(
		ctx,
		cmd.Value("retriever-id").(string),
		params,
		options...,
	)
}

func handleRetrieversGet(ctx context.Context, cmd *cli.Command) error {
	client := llamacloudprod.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("retriever-id") && len(unusedArgs) > 0 {
		cmd.Set("retriever-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := llamacloudprod.RetrieverGetParams{}

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
	_, err = client.Retrievers.Get(
		ctx,
		cmd.Value("retriever-id").(string),
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
		Title:          "retrievers get",
		Transform:      transform,
	})
}

func handleRetrieversSearch(ctx context.Context, cmd *cli.Command) error {
	client := llamacloudprod.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := llamacloudprod.RetrieverSearchParams{}

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
	_, err = client.Retrievers.Search(ctx, params, options...)
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
		Title:          "retrievers search",
		Transform:      transform,
	})
}

func handleRetrieversUpsert(ctx context.Context, cmd *cli.Command) error {
	client := llamacloudprod.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := llamacloudprod.RetrieverUpsertParams{}

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
	_, err = client.Retrievers.Upsert(ctx, params, options...)
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
		Title:          "retrievers upsert",
		Transform:      transform,
	})
}
