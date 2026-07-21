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

var retrieversRetrieverSearch = requestflag.WithInnerFlags(cli.Command{
	Name:    "search",
	Usage:   "Retrieve data using a Retriever.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "retriever-id",
			Required:  true,
			PathParam: "retriever_id",
		},
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
	Action:          handleRetrieversRetrieverSearch,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
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

func handleRetrieversRetrieverSearch(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("retriever-id") && len(unusedArgs) > 0 {
		cmd.Set("retriever-id", unusedArgs[0])
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

	params := llamacloud.RetrieverRetrieverSearchParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Retrievers.Retriever.Search(
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
		Title:          "retrievers:retriever search",
		Transform:      transform,
	})
}
