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

var webhookConfigsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a reusable webhook configuration for the current project.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "webhook-url",
			Usage:    "URL to receive webhook POST notifications.",
			Required: true,
			BodyPath: "webhook_url",
		},
		&requestflag.Flag[*string]{
			Name:      "organization-id",
			QueryPath: "organization_id",
		},
		&requestflag.Flag[*string]{
			Name:      "project-id",
			QueryPath: "project_id",
		},
		&requestflag.Flag[any]{
			Name:     "webhook-event",
			Usage:    "Events to subscribe to. If null, all events are delivered.",
			BodyPath: "webhook_events",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "webhook-headers",
			Usage:    "Custom HTTP headers sent with each webhook request.",
			BodyPath: "webhook_headers",
		},
		&requestflag.Flag[*string]{
			Name:     "webhook-output-format",
			Usage:    "Response format sent to the webhook: 'string' (default) or 'json'.",
			BodyPath: "webhook_output_format",
		},
		&requestflag.Flag[*string]{
			Name:     "webhook-signing-secret",
			Usage:    "Shared secret used to sign deliveries to this endpoint. Write-only: it is never returned in responses.",
			BodyPath: "webhook_signing_secret",
		},
	},
	Action:          handleWebhookConfigsCreate,
	HideHelpCommand: true,
}

var webhookConfigsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a single webhook configuration by ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "config-id",
			Required:  true,
			PathParam: "config_id",
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
	Action:          handleWebhookConfigsRetrieve,
	HideHelpCommand: true,
}

var webhookConfigsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a webhook configuration. Only fields present in the request change.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "config-id",
			Required:  true,
			PathParam: "config_id",
		},
		&requestflag.Flag[*string]{
			Name:      "organization-id",
			QueryPath: "organization_id",
		},
		&requestflag.Flag[*string]{
			Name:      "project-id",
			QueryPath: "project_id",
		},
		&requestflag.Flag[any]{
			Name:     "webhook-event",
			Usage:    "Updated event subscriptions.",
			BodyPath: "webhook_events",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "webhook-headers",
			Usage:    "Updated headers.",
			BodyPath: "webhook_headers",
		},
		&requestflag.Flag[*string]{
			Name:     "webhook-output-format",
			Usage:    "Updated output format.",
			BodyPath: "webhook_output_format",
		},
		&requestflag.Flag[*string]{
			Name:     "webhook-signing-secret",
			Usage:    "Updated signing secret (write-only). Send to rotate the secret.",
			BodyPath: "webhook_signing_secret",
		},
		&requestflag.Flag[*string]{
			Name:     "webhook-url",
			Usage:    "Updated webhook URL.",
			BodyPath: "webhook_url",
		},
	},
	Action:          handleWebhookConfigsUpdate,
	HideHelpCommand: true,
}

var webhookConfigsList = cli.Command{
	Name:    "list",
	Usage:   "List the webhook configurations for the current project, newest first.",
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
	Action:          handleWebhookConfigsList,
	HideHelpCommand: true,
}

var webhookConfigsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a webhook configuration.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "config-id",
			Required:  true,
			PathParam: "config_id",
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
	Action:          handleWebhookConfigsDelete,
	HideHelpCommand: true,
}

func handleWebhookConfigsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloud.WebhookConfigNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.WebhookConfigs.New(ctx, params, options...)
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
		Title:          "webhook-configs create",
		Transform:      transform,
	})
}

func handleWebhookConfigsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("config-id") && len(unusedArgs) > 0 {
		cmd.Set("config-id", unusedArgs[0])
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

	params := llamacloud.WebhookConfigGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.WebhookConfigs.Get(
		ctx,
		cmd.Value("config-id").(string),
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
		Title:          "webhook-configs retrieve",
		Transform:      transform,
	})
}

func handleWebhookConfigsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("config-id") && len(unusedArgs) > 0 {
		cmd.Set("config-id", unusedArgs[0])
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

	params := llamacloud.WebhookConfigUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.WebhookConfigs.Update(
		ctx,
		cmd.Value("config-id").(string),
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
		Title:          "webhook-configs update",
		Transform:      transform,
	})
}

func handleWebhookConfigsList(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloud.WebhookConfigListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.WebhookConfigs.List(ctx, params, options...)
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
		Title:          "webhook-configs list",
		Transform:      transform,
	})
}

func handleWebhookConfigsDelete(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("config-id") && len(unusedArgs) > 0 {
		cmd.Set("config-id", unusedArgs[0])
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

	params := llamacloud.WebhookConfigDeleteParams{}

	return client.WebhookConfigs.Delete(
		ctx,
		cmd.Value("config-id").(string),
		params,
		options...,
	)
}
