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

var betaRetrievalRetrieve = requestflag.WithInnerFlags(cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve relevant chunks via hybrid search (vector + full-text), with filtering\non built-in or user-defined metadata.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "index-id",
			Usage:    "ID of the index to retrieve against.",
			Required: true,
			BodyPath: "index_id",
		},
		&requestflag.Flag[string]{
			Name:     "query",
			Usage:    "Natural-language query to retrieve relevant chunks.",
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
		&requestflag.Flag[map[string]any]{
			Name:     "custom-filters",
			Usage:    "Filters on user-defined metadata fields.",
			BodyPath: "custom_filters",
		},
		&requestflag.Flag[*float64]{
			Name:     "full-text-pipeline-weight",
			Usage:    "Weight of the full-text search pipeline (0-1).",
			BodyPath: "full_text_pipeline_weight",
		},
		&requestflag.Flag[*int64]{
			Name:     "num-candidates",
			Usage:    "Number of candidates for approximate nearest neighbor search.",
			BodyPath: "num_candidates",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "rerank",
			Usage:    "Reranking configuration applied after hybrid search. Enabled by default.",
			BodyPath: "rerank",
		},
		&requestflag.Flag[*float64]{
			Name:     "score-threshold",
			Usage:    "Minimum score threshold for returned results.",
			BodyPath: "score_threshold",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "static-filters",
			Usage:    "Filters on built-in document fields (page range, chunk index, etc.).",
			BodyPath: "static_filters",
		},
		&requestflag.Flag[*int64]{
			Name:     "top-k",
			Usage:    "Maximum number of results to return.",
			BodyPath: "top_k",
		},
		&requestflag.Flag[*float64]{
			Name:     "vector-pipeline-weight",
			Usage:    "Weight of the vector search pipeline (0-1).",
			BodyPath: "vector_pipeline_weight",
		},
	},
	Action:          handleBetaRetrievalRetrieve,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"rerank": {
		&requestflag.InnerFlag[bool]{
			Name:       "rerank.enabled",
			Usage:      "Set to false to disable reranking.",
			InnerField: "enabled",
		},
		&requestflag.InnerFlag[*int64]{
			Name:       "rerank.top-n",
			Usage:      "Number of results to return after reranking.",
			InnerField: "top_n",
		},
	},
	"static-filters": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "static-filters.parsed-directory-file-id",
			InnerField: "parsed_directory_file_id",
		},
	},
})

var betaRetrievalFind = cli.Command{
	Name:    "find",
	Usage:   "Search for files by name.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "index-id",
			Usage:    "ID of the index to search within.",
			Required: true,
			BodyPath: "index_id",
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
			Name:     "file-name",
			Usage:    "Exact file name to match.",
			BodyPath: "file_name",
		},
		&requestflag.Flag[*string]{
			Name:     "file-name-contains",
			Usage:    "Substring match on file name (case-insensitive).",
			BodyPath: "file_name_contains",
		},
		&requestflag.Flag[*int64]{
			Name:     "page-size",
			Usage:    "The maximum number of items to return. The service may return fewer than this value. If unspecified, a default page size will be used. The maximum value is typically 1000; values above this will be coerced to the maximum.",
			BodyPath: "page_size",
		},
		&requestflag.Flag[*string]{
			Name:     "page-token",
			Usage:    "A page token, received from a previous list call. Provide this to retrieve the subsequent page.",
			BodyPath: "page_token",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleBetaRetrievalFind,
	HideHelpCommand: true,
}

var betaRetrievalGrep = cli.Command{
	Name:    "grep",
	Usage:   "Grep within a file's parsed content using a regex pattern.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "file-id",
			Usage:    "ID of the file to grep.",
			Required: true,
			BodyPath: "file_id",
		},
		&requestflag.Flag[string]{
			Name:     "index-id",
			Usage:    "ID of the index the file belongs to.",
			Required: true,
			BodyPath: "index_id",
		},
		&requestflag.Flag[string]{
			Name:     "pattern",
			Usage:    "Regex pattern to search for.",
			Required: true,
			BodyPath: "pattern",
		},
		&requestflag.Flag[*string]{
			Name:      "organization-id",
			QueryPath: "organization_id",
		},
		&requestflag.Flag[*string]{
			Name:      "project-id",
			QueryPath: "project_id",
		},
		&requestflag.Flag[*int64]{
			Name:     "context-chars",
			Usage:    "Number of characters of context to include before and after the matched pattern in the content field of the response",
			BodyPath: "context_chars",
		},
		&requestflag.Flag[*int64]{
			Name:     "page-size",
			Usage:    "The maximum number of items to return. The service may return fewer than this value. If unspecified, a default page size will be used. The maximum value is typically 1000; values above this will be coerced to the maximum.",
			BodyPath: "page_size",
		},
		&requestflag.Flag[*string]{
			Name:     "page-token",
			Usage:    "A page token, received from a previous list call. Provide this to retrieve the subsequent page.",
			BodyPath: "page_token",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleBetaRetrievalGrep,
	HideHelpCommand: true,
}

var betaRetrievalRead = cli.Command{
	Name:    "read",
	Usage:   "Read the parsed text content of a specific file.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "file-id",
			Usage:    "ID of the file to read.",
			Required: true,
			BodyPath: "file_id",
		},
		&requestflag.Flag[string]{
			Name:     "index-id",
			Usage:    "ID of the index the file belongs to.",
			Required: true,
			BodyPath: "index_id",
		},
		&requestflag.Flag[*string]{
			Name:      "organization-id",
			QueryPath: "organization_id",
		},
		&requestflag.Flag[*string]{
			Name:      "project-id",
			QueryPath: "project_id",
		},
		&requestflag.Flag[*int64]{
			Name:     "max-length",
			Usage:    "Maximum number of characters to read from the offset.",
			BodyPath: "max_length",
		},
		&requestflag.Flag[int64]{
			Name:     "offset",
			Usage:    "Starting character offset.",
			Default:  0,
			BodyPath: "offset",
		},
	},
	Action:          handleBetaRetrievalRead,
	HideHelpCommand: true,
}

func handleBetaRetrievalRetrieve(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloud.BetaRetrievalGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Beta.Retrieval.Get(ctx, params, options...)
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
		Title:          "beta:retrieval retrieve",
		Transform:      transform,
	})
}

func handleBetaRetrievalFind(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloud.BetaRetrievalFindParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Beta.Retrieval.Find(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "beta:retrieval find",
			Transform:      transform,
		})
	} else {
		iter := client.Beta.Retrieval.FindAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "beta:retrieval find",
			Transform:      transform,
		})
	}
}

func handleBetaRetrievalGrep(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloud.BetaRetrievalGrepParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Beta.Retrieval.Grep(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "beta:retrieval grep",
			Transform:      transform,
		})
	} else {
		iter := client.Beta.Retrieval.GrepAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "beta:retrieval grep",
			Transform:      transform,
		})
	}
}

func handleBetaRetrievalRead(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloud.BetaRetrievalReadParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Beta.Retrieval.Read(ctx, params, options...)
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
		Title:          "beta:retrieval read",
		Transform:      transform,
	})
}
