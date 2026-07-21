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

var betaChatCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a chat session, optionally bound to indexes (locked after the first\nmessage).",
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
		&requestflag.Flag[any]{
			Name:     "index-id",
			Usage:    "Indexes this session will retrieve from. Once set and the first message has been sent, the source set is locked for the session's lifetime. Leave null to create an unbound session.",
			BodyPath: "index_ids",
		},
	},
	Action:          handleBetaChatCreate,
	HideHelpCommand: true,
}

var betaChatRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a full session by ID, including its event history.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "session-id",
			Required:  true,
			PathParam: "session_id",
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
	Action:          handleBetaChatRetrieve,
	HideHelpCommand: true,
}

var betaChatList = cli.Command{
	Name:    "list",
	Usage:   "List all chat sessions for the current project.",
	Suggest: true,
	Flags: []cli.Flag{
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
	Action:          handleBetaChatList,
	HideHelpCommand: true,
}

var betaChatDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a session.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "session-id",
			Required:  true,
			PathParam: "session_id",
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
	Action:          handleBetaChatDelete,
	HideHelpCommand: true,
}

var betaChatGetSummary = cli.Command{
	Name:    "get-summary",
	Usage:   "Retrieve a session summary by ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "session-id",
			Required:  true,
			PathParam: "session_id",
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
	Action:          handleBetaChatGetSummary,
	HideHelpCommand: true,
}

var betaChatStream = cli.Command{
	Name:    "stream",
	Usage:   "Stream agent events for a chat turn as Server-Sent Events.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "session-id",
			Required:  true,
			PathParam: "session_id",
		},
		&requestflag.Flag[[]string]{
			Name:     "index-id",
			Usage:    "Indexes to retrieve data from.",
			Required: true,
			BodyPath: "index_ids",
		},
		&requestflag.Flag[string]{
			Name:     "prompt",
			Usage:    "User message for this chat turn.",
			Required: true,
			BodyPath: "prompt",
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
	Action:          handleBetaChatStream,
	HideHelpCommand: true,
}

func handleBetaChatCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloud.BetaChatNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Beta.Chat.New(ctx, params, options...)
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
		Title:          "beta:chat create",
		Transform:      transform,
	})
}

func handleBetaChatRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("session-id") && len(unusedArgs) > 0 {
		cmd.Set("session-id", unusedArgs[0])
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

	params := llamacloud.BetaChatGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Beta.Chat.Get(
		ctx,
		cmd.Value("session-id").(string),
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
		Title:          "beta:chat retrieve",
		Transform:      transform,
	})
}

func handleBetaChatList(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloud.BetaChatListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Beta.Chat.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "beta:chat list",
			Transform:      transform,
		})
	} else {
		iter := client.Beta.Chat.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "beta:chat list",
			Transform:      transform,
		})
	}
}

func handleBetaChatDelete(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("session-id") && len(unusedArgs) > 0 {
		cmd.Set("session-id", unusedArgs[0])
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

	params := llamacloud.BetaChatDeleteParams{}

	return client.Beta.Chat.Delete(
		ctx,
		cmd.Value("session-id").(string),
		params,
		options...,
	)
}

func handleBetaChatGetSummary(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("session-id") && len(unusedArgs) > 0 {
		cmd.Set("session-id", unusedArgs[0])
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

	params := llamacloud.BetaChatGetSummaryParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Beta.Chat.GetSummary(
		ctx,
		cmd.Value("session-id").(string),
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
		Title:          "beta:chat get-summary",
		Transform:      transform,
	})
}

func handleBetaChatStream(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("session-id") && len(unusedArgs) > 0 {
		cmd.Set("session-id", unusedArgs[0])
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

	params := llamacloud.BetaChatStreamParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Beta.Chat.Stream(
		ctx,
		cmd.Value("session-id").(string),
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
		Title:          "beta:chat stream",
		Transform:      transform,
	})
}
