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

var betaAttachmentsList = cli.Command{
	Name:    "list",
	Usage:   "List the attachments associated with a file (e.g. per-page screenshots).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "source-id",
			Usage:     "File UUID or directory file ID (dfl-...).",
			Required:  true,
			QueryPath: "source_id",
		},
		&requestflag.Flag[*string]{
			Name:      "organization-id",
			QueryPath: "organization_id",
		},
		&requestflag.Flag[*int64]{
			Name:      "page-size",
			QueryPath: "page_size",
		},
		&requestflag.Flag[*string]{
			Name:      "page-token",
			QueryPath: "page_token",
		},
		&requestflag.Flag[*string]{
			Name:      "project-id",
			QueryPath: "project_id",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleBetaAttachmentsList,
	HideHelpCommand: true,
}

var betaAttachmentsGet = cli.Command{
	Name:    "get",
	Usage:   "Return a presigned download URL for a specific attachment.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "attachment-name",
			Required:  true,
			PathParam: "attachment_name",
		},
		&requestflag.Flag[string]{
			Name:      "source-id",
			Usage:     "File UUID or directory file ID (dfl-...).",
			Required:  true,
			QueryPath: "source_id",
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
	Action:          handleBetaAttachmentsGet,
	HideHelpCommand: true,
}

func handleBetaAttachmentsList(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
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

	params := llamacloud.BetaAttachmentListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Beta.Attachments.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "beta:attachments list",
			Transform:      transform,
		})
	} else {
		iter := client.Beta.Attachments.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "beta:attachments list",
			Transform:      transform,
		})
	}
}

func handleBetaAttachmentsGet(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("attachment-name") && len(unusedArgs) > 0 {
		cmd.Set("attachment-name", unusedArgs[0])
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

	params := llamacloud.BetaAttachmentGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Beta.Attachments.Get(
		ctx,
		cmd.Value("attachment-name").(string),
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
		Title:          "beta:attachments get",
		Transform:      transform,
	})
}
