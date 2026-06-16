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

var classifierJobsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a classify job. Experimental: not production-ready and subject to change.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "file-id",
			Usage:    "The IDs of the files to classify",
			Required: true,
			BodyPath: "file_ids",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "rule",
			Usage:    "The rules to classify the files",
			Required: true,
			BodyPath: "rules",
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
			Usage:    "The classification mode to use",
			Default:  "FAST",
			BodyPath: "mode",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "parsing-configuration",
			Usage:    "Parsing configuration for a classify job.",
			BodyPath: "parsing_configuration",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "webhook-configuration",
			Usage:    "List of webhook configurations for notifications",
			BodyPath: "webhook_configurations",
		},
	},
	Action:          handleClassifierJobsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"rule": {
		&requestflag.InnerFlag[string]{
			Name:       "rule.description",
			Usage:      "Natural language description of what to classify. Be specific about the content characteristics that identify this document type.",
			InnerField: "description",
		},
		&requestflag.InnerFlag[string]{
			Name:       "rule.type",
			Usage:      "The document type to assign when this rule matches (e.g., 'invoice', 'receipt', 'contract')",
			InnerField: "type",
		},
	},
	"parsing-configuration": {
		&requestflag.InnerFlag[string]{
			Name:       "parsing-configuration.lang",
			Usage:      "Enum for representing the languages supported by the parser.",
			InnerField: "lang",
		},
		&requestflag.InnerFlag[*int64]{
			Name:       "parsing-configuration.max-pages",
			Usage:      "The maximum number of pages to parse",
			InnerField: "max_pages",
		},
		&requestflag.InnerFlag[any]{
			Name:       "parsing-configuration.target-pages",
			Usage:      "The pages to target for parsing (0-indexed, so first page is at 0)",
			InnerField: "target_pages",
		},
	},
	"webhook-configuration": {
		&requestflag.InnerFlag[any]{
			Name:       "webhook-configuration.webhook-events",
			Usage:      "Events that trigger this webhook. Options: 'parse.success' (job completed), 'parse.error' (job failed), 'parse.partial_success' (some pages failed), 'parse.pending', 'parse.running', 'parse.cancelled'. If not specified, webhook fires for all events",
			InnerField: "webhook_events",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "webhook-configuration.webhook-headers",
			Usage:      "Custom HTTP headers to include in webhook requests. Use for authentication tokens or custom routing. Example: {'Authorization': 'Bearer xyz'}",
			InnerField: "webhook_headers",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "webhook-configuration.webhook-output-format",
			Usage:      "Format of the webhook payload body. 'string' (default) sends the payload as a JSON-encoded string; 'json' sends it as a JSON object.",
			InnerField: "webhook_output_format",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "webhook-configuration.webhook-url",
			Usage:      "HTTPS URL to receive webhook POST requests. Must be publicly accessible",
			InnerField: "webhook_url",
		},
	},
})

var classifierJobsList = cli.Command{
	Name:    "list",
	Usage:   "List classify jobs. Experimental: not production-ready and subject to change.",
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
	Action:          handleClassifierJobsList,
	HideHelpCommand: true,
}

var classifierJobsGet = cli.Command{
	Name:    "get",
	Usage:   "Get a classify job. Experimental: not production-ready and subject to change.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "classify-job-id",
			Required:  true,
			PathParam: "classify_job_id",
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
	Action:          handleClassifierJobsGet,
	HideHelpCommand: true,
}

var classifierJobsGetResults = cli.Command{
	Name:    "get-results",
	Usage:   "Get the results of a classify job. Experimental: not production-ready and\nsubject to change.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "classify-job-id",
			Required:  true,
			PathParam: "classify_job_id",
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
	Action:          handleClassifierJobsGetResults,
	HideHelpCommand: true,
}

func handleClassifierJobsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloudprod.ClassifierJobNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Classifier.Jobs.New(ctx, params, options...)
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
		Title:          "classifier:jobs create",
		Transform:      transform,
	})
}

func handleClassifierJobsList(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloudprod.ClassifierJobListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Classifier.Jobs.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "classifier:jobs list",
			Transform:      transform,
		})
	} else {
		iter := client.Classifier.Jobs.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "classifier:jobs list",
			Transform:      transform,
		})
	}
}

func handleClassifierJobsGet(ctx context.Context, cmd *cli.Command) error {
	client := llamacloudprod.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("classify-job-id") && len(unusedArgs) > 0 {
		cmd.Set("classify-job-id", unusedArgs[0])
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

	params := llamacloudprod.ClassifierJobGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Classifier.Jobs.Get(
		ctx,
		cmd.Value("classify-job-id").(string),
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
		Title:          "classifier:jobs get",
		Transform:      transform,
	})
}

func handleClassifierJobsGetResults(ctx context.Context, cmd *cli.Command) error {
	client := llamacloudprod.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("classify-job-id") && len(unusedArgs) > 0 {
		cmd.Set("classify-job-id", unusedArgs[0])
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

	params := llamacloudprod.ClassifierJobGetResultsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Classifier.Jobs.GetResults(
		ctx,
		cmd.Value("classify-job-id").(string),
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
		Title:          "classifier:jobs get-results",
		Transform:      transform,
	})
}
