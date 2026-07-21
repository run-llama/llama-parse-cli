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

var pipelinesDocumentsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Batch create documents for a pipeline.",
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
	Action:          handlePipelinesDocumentsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"body": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "body.metadata",
			InnerField: "metadata",
		},
		&requestflag.InnerFlag[string]{
			Name:       "body.text",
			InnerField: "text",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "body.id",
			InnerField: "id",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "body.excluded-embed-metadata-keys",
			InnerField: "excluded_embed_metadata_keys",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "body.excluded-llm-metadata-keys",
			InnerField: "excluded_llm_metadata_keys",
		},
		&requestflag.InnerFlag[any]{
			Name:       "body.page-positions",
			Usage:      "indices in the CloudDocument.text where a new page begins. e.g. Second page starts at index specified by page_positions[1].",
			InnerField: "page_positions",
		},
	},
})

var pipelinesDocumentsList = cli.Command{
	Name:    "list",
	Usage:   "Return a list of documents for a pipeline.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "pipeline-id",
			Required:  true,
			PathParam: "pipeline_id",
		},
		&requestflag.Flag[*string]{
			Name:      "file-id",
			QueryPath: "file_id",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Default:   10,
			QueryPath: "limit",
		},
		&requestflag.Flag[*bool]{
			Name:      "only-api-data-source-documents",
			Default:   requestflag.Ptr[bool](false),
			QueryPath: "only_api_data_source_documents",
		},
		&requestflag.Flag[*bool]{
			Name:      "only-direct-upload",
			Default:   requestflag.Ptr[bool](false),
			QueryPath: "only_direct_upload",
		},
		&requestflag.Flag[int64]{
			Name:      "skip",
			Default:   0,
			QueryPath: "skip",
		},
		&requestflag.Flag[string]{
			Name:      "status-refresh-policy",
			Usage:     `Allowed values: "cached", "ttl".`,
			Default:   "cached",
			QueryPath: "status_refresh_policy",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handlePipelinesDocumentsList,
	HideHelpCommand: true,
}

var pipelinesDocumentsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a document from a pipeline; runs async (vectors first, then MongoDB\nrecord).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "pipeline-id",
			Required:  true,
			PathParam: "pipeline_id",
		},
		&requestflag.Flag[string]{
			Name:      "document-id",
			Required:  true,
			PathParam: "document_id",
		},
	},
	Action:          handlePipelinesDocumentsDelete,
	HideHelpCommand: true,
}

var pipelinesDocumentsGet = cli.Command{
	Name:    "get",
	Usage:   "Return a single document for a pipeline.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "pipeline-id",
			Required:  true,
			PathParam: "pipeline_id",
		},
		&requestflag.Flag[string]{
			Name:      "document-id",
			Required:  true,
			PathParam: "document_id",
		},
	},
	Action:          handlePipelinesDocumentsGet,
	HideHelpCommand: true,
}

var pipelinesDocumentsGetChunks = cli.Command{
	Name:    "get-chunks",
	Usage:   "Return a list of chunks for a pipeline document.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "pipeline-id",
			Required:  true,
			PathParam: "pipeline_id",
		},
		&requestflag.Flag[string]{
			Name:      "document-id",
			Required:  true,
			PathParam: "document_id",
		},
	},
	Action:          handlePipelinesDocumentsGetChunks,
	HideHelpCommand: true,
}

var pipelinesDocumentsGetStatus = cli.Command{
	Name:    "get-status",
	Usage:   "Return a single document for a pipeline.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "pipeline-id",
			Required:  true,
			PathParam: "pipeline_id",
		},
		&requestflag.Flag[string]{
			Name:      "document-id",
			Required:  true,
			PathParam: "document_id",
		},
	},
	Action:          handlePipelinesDocumentsGetStatus,
	HideHelpCommand: true,
}

var pipelinesDocumentsSync = cli.Command{
	Name:    "sync",
	Usage:   "Sync a specific document for a pipeline.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "pipeline-id",
			Required:  true,
			PathParam: "pipeline_id",
		},
		&requestflag.Flag[string]{
			Name:      "document-id",
			Required:  true,
			PathParam: "document_id",
		},
	},
	Action:          handlePipelinesDocumentsSync,
	HideHelpCommand: true,
}

var pipelinesDocumentsUpsert = requestflag.WithInnerFlags(cli.Command{
	Name:    "upsert",
	Usage:   "Batch create or update a document for a pipeline.",
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
	Action:          handlePipelinesDocumentsUpsert,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"body": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "body.metadata",
			InnerField: "metadata",
		},
		&requestflag.InnerFlag[string]{
			Name:       "body.text",
			InnerField: "text",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "body.id",
			InnerField: "id",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "body.excluded-embed-metadata-keys",
			InnerField: "excluded_embed_metadata_keys",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "body.excluded-llm-metadata-keys",
			InnerField: "excluded_llm_metadata_keys",
		},
		&requestflag.InnerFlag[any]{
			Name:       "body.page-positions",
			Usage:      "indices in the CloudDocument.text where a new page begins. e.g. Second page starts at index specified by page_positions[1].",
			InnerField: "page_positions",
		},
	},
})

func handlePipelinesDocumentsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloud.PipelineDocumentNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Pipelines.Documents.New(
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
		Title:          "pipelines:documents create",
		Transform:      transform,
	})
}

func handlePipelinesDocumentsList(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloud.PipelineDocumentListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Pipelines.Documents.List(
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
			Title:          "pipelines:documents list",
			Transform:      transform,
		})
	} else {
		iter := client.Pipelines.Documents.ListAutoPaging(
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
			Title:          "pipelines:documents list",
			Transform:      transform,
		})
	}
}

func handlePipelinesDocumentsDelete(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("document-id") && len(unusedArgs) > 0 {
		cmd.Set("document-id", unusedArgs[0])
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

	params := llamacloud.PipelineDocumentDeleteParams{
		PipelineID: cmd.Value("pipeline-id").(string),
	}

	return client.Pipelines.Documents.Delete(
		ctx,
		cmd.Value("document-id").(string),
		params,
		options...,
	)
}

func handlePipelinesDocumentsGet(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("document-id") && len(unusedArgs) > 0 {
		cmd.Set("document-id", unusedArgs[0])
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

	params := llamacloud.PipelineDocumentGetParams{
		PipelineID: cmd.Value("pipeline-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Pipelines.Documents.Get(
		ctx,
		cmd.Value("document-id").(string),
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
		Title:          "pipelines:documents get",
		Transform:      transform,
	})
}

func handlePipelinesDocumentsGetChunks(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("document-id") && len(unusedArgs) > 0 {
		cmd.Set("document-id", unusedArgs[0])
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

	params := llamacloud.PipelineDocumentGetChunksParams{
		PipelineID: cmd.Value("pipeline-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Pipelines.Documents.GetChunks(
		ctx,
		cmd.Value("document-id").(string),
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
		Title:          "pipelines:documents get-chunks",
		Transform:      transform,
	})
}

func handlePipelinesDocumentsGetStatus(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("document-id") && len(unusedArgs) > 0 {
		cmd.Set("document-id", unusedArgs[0])
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

	params := llamacloud.PipelineDocumentGetStatusParams{
		PipelineID: cmd.Value("pipeline-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Pipelines.Documents.GetStatus(
		ctx,
		cmd.Value("document-id").(string),
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
		Title:          "pipelines:documents get-status",
		Transform:      transform,
	})
}

func handlePipelinesDocumentsSync(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("document-id") && len(unusedArgs) > 0 {
		cmd.Set("document-id", unusedArgs[0])
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

	params := llamacloud.PipelineDocumentSyncParams{
		PipelineID: cmd.Value("pipeline-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Pipelines.Documents.Sync(
		ctx,
		cmd.Value("document-id").(string),
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
		Title:          "pipelines:documents sync",
		Transform:      transform,
	})
}

func handlePipelinesDocumentsUpsert(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloud.PipelineDocumentUpsertParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Pipelines.Documents.Upsert(
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
		Title:          "pipelines:documents upsert",
		Transform:      transform,
	})
}
