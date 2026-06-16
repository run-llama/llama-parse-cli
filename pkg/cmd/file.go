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

var filesCreate = cli.Command{
	Name:    "create",
	Usage:   "Upload a file using multipart/form-data.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "file",
			Usage:     "The file to upload",
			Required:  true,
			BodyPath:  "file",
			FileInput: true,
		},
		&requestflag.Flag[string]{
			Name:     "purpose",
			Usage:    "The intended purpose of the file. Valid values: 'user_data', 'parse', 'extract', 'split', 'classify', 'sheet', 'agent_app'. This determines the storage and retention policy for the file.",
			Required: true,
			BodyPath: "purpose",
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
			Name:     "external-file-id",
			Usage:    "The ID of the file in the external system",
			BodyPath: "external_file_id",
		},
	},
	Action:          handleFilesCreate,
	HideHelpCommand: true,
}

var filesList = cli.Command{
	Name:    "list",
	Usage:   "List files with optional filtering and pagination.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:      "expand",
			Usage:     "Fields to expand on each file.",
			QueryPath: "expand",
		},
		&requestflag.Flag[*string]{
			Name:      "external-file-id",
			Usage:     "Filter by external file ID.",
			QueryPath: "external_file_id",
		},
		&requestflag.Flag[any]{
			Name:      "file-id",
			Usage:     "Filter by specific file IDs.",
			QueryPath: "file_ids",
		},
		&requestflag.Flag[*string]{
			Name:      "file-name",
			Usage:     "Filter by file name (exact match).",
			QueryPath: "file_name",
		},
		&requestflag.Flag[*string]{
			Name:      "order-by",
			Usage:     "A comma-separated list of fields to order by, sorted in ascending order. Use 'field_name desc' to specify descending order.",
			QueryPath: "order_by",
		},
		&requestflag.Flag[*string]{
			Name:      "organization-id",
			QueryPath: "organization_id",
		},
		&requestflag.Flag[*int64]{
			Name:      "page-size",
			Usage:     "The maximum number of items to return. Defaults to 50, maximum is 1000.",
			QueryPath: "page_size",
		},
		&requestflag.Flag[*string]{
			Name:      "page-token",
			Usage:     "A page token received from a previous list call. Provide this to retrieve the subsequent page.",
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
	Action:          handleFilesList,
	HideHelpCommand: true,
}

var filesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a file from the project.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "file-id",
			Required:  true,
			PathParam: "file_id",
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
	Action:          handleFilesDelete,
	HideHelpCommand: true,
}

var filesGet = cli.Command{
	Name:    "get",
	Usage:   "Get a presigned URL to download the file content.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "file-id",
			Required:  true,
			PathParam: "file_id",
		},
		&requestflag.Flag[*int64]{
			Name:      "expires-at-seconds",
			QueryPath: "expires_at_seconds",
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
	Action:          handleFilesGet,
	HideHelpCommand: true,
}

var filesQuery = requestflag.WithInnerFlags(cli.Command{
	Name:    "query",
	Usage:   "Query files with filtering and pagination. Deprecated: use `GET /files`.",
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
		&requestflag.Flag[map[string]any]{
			Name:     "filter",
			Usage:    "Filter parameters for file queries.",
			BodyPath: "filter",
		},
		&requestflag.Flag[*string]{
			Name:     "order-by",
			Usage:    "A comma-separated list of fields to order by, sorted in ascending order. Use 'field_name desc' to specify descending order.",
			BodyPath: "order_by",
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
	},
	Action:          handleFilesQuery,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[*string]{
			Name:       "filter.data-source-id",
			Usage:      "Filter by data source ID",
			InnerField: "data_source_id",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "filter.external-file-id",
			Usage:      "Filter by external file ID",
			InnerField: "external_file_id",
		},
		&requestflag.InnerFlag[any]{
			Name:       "filter.file-ids",
			Usage:      "Filter by specific file IDs",
			InnerField: "file_ids",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "filter.file-name",
			Usage:      "Filter by file name",
			InnerField: "file_name",
		},
		&requestflag.InnerFlag[*bool]{
			Name:       "filter.only-manually-uploaded",
			Usage:      "Filter only manually uploaded files (data_source_id is null)",
			InnerField: "only_manually_uploaded",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "filter.project-id",
			Usage:      "Filter by project ID",
			InnerField: "project_id",
		},
	},
})

func handleFilesCreate(ctx context.Context, cmd *cli.Command) error {
	client := llamacloudprod.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	params := llamacloudprod.FileNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Files.New(ctx, params, options...)
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
		Title:          "files create",
		Transform:      transform,
	})
}

func handleFilesList(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloudprod.FileListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Files.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "files list",
			Transform:      transform,
		})
	} else {
		iter := client.Files.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "files list",
			Transform:      transform,
		})
	}
}

func handleFilesDelete(ctx context.Context, cmd *cli.Command) error {
	client := llamacloudprod.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("file-id") && len(unusedArgs) > 0 {
		cmd.Set("file-id", unusedArgs[0])
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

	params := llamacloudprod.FileDeleteParams{}

	return client.Files.Delete(
		ctx,
		cmd.Value("file-id").(string),
		params,
		options...,
	)
}

func handleFilesGet(ctx context.Context, cmd *cli.Command) error {
	client := llamacloudprod.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("file-id") && len(unusedArgs) > 0 {
		cmd.Set("file-id", unusedArgs[0])
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

	params := llamacloudprod.FileGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Files.Get(
		ctx,
		cmd.Value("file-id").(string),
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
		Title:          "files get",
		Transform:      transform,
	})
}

func handleFilesQuery(ctx context.Context, cmd *cli.Command) error {
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

	params := llamacloudprod.FileQueryParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Files.Query(ctx, params, options...)
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
		Title:          "files query",
		Transform:      transform,
	})
}
