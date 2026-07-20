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

var betaDirectoriesFilesUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update directory-file metadata by `directory_file_id`; set `directory_id` to\nmove the file to a different directory. To resolve from `unique_id`, list with a\nfilter first.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "directory-id",
			Required:  true,
			PathParam: "directory_id",
		},
		&requestflag.Flag[string]{
			Name:      "directory-file-id",
			Required:  true,
			PathParam: "directory_file_id",
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
			Name:     "display-name",
			Usage:    "Updated display name.",
			BodyPath: "display_name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Usage:    "User-defined metadata key-value pairs. Replaces the user metadata layer.",
			BodyPath: "metadata",
		},
		&requestflag.Flag[*string]{
			Name:     "target-directory-id",
			Usage:    "Move file to a different directory.",
			BodyPath: "target_directory_id",
		},
		&requestflag.Flag[*string]{
			Name:     "unique-id",
			Usage:    "Updated unique identifier.",
			BodyPath: "unique_id",
		},
	},
	Action:          handleBetaDirectoriesFilesUpdate,
	HideHelpCommand: true,
}

var betaDirectoriesFilesList = cli.Command{
	Name:    "list",
	Usage:   "List all files within the specified directory with optional filtering and\npagination.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "directory-id",
			Required:  true,
			PathParam: "directory_id",
		},
		&requestflag.Flag[*string]{
			Name:      "display-name",
			QueryPath: "display_name",
		},
		&requestflag.Flag[*string]{
			Name:      "display-name-contains",
			QueryPath: "display_name_contains",
		},
		&requestflag.Flag[any]{
			Name:      "expand",
			Usage:     "Fields to expand on each directory file.",
			QueryPath: "expand",
		},
		&requestflag.Flag[*string]{
			Name:      "file-id",
			QueryPath: "file_id",
		},
		&requestflag.Flag[bool]{
			Name:      "include-deleted",
			Default:   false,
			QueryPath: "include_deleted",
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
		&requestflag.Flag[*string]{
			Name:      "unique-id",
			QueryPath: "unique_id",
		},
		&requestflag.Flag[any]{
			Name:      "updated-at-on-or-after",
			Usage:     "Include items updated at or after this timestamp (inclusive)",
			QueryPath: "updated_at_on_or_after",
		},
		&requestflag.Flag[any]{
			Name:      "updated-at-on-or-before",
			Usage:     "Include items updated at or before this timestamp (inclusive)",
			QueryPath: "updated_at_on_or_before",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleBetaDirectoriesFilesList,
	HideHelpCommand: true,
}

var betaDirectoriesFilesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a directory file by `directory_file_id`; to resolve from `unique_id`,\nlist with a filter first.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "directory-id",
			Required:  true,
			PathParam: "directory_id",
		},
		&requestflag.Flag[string]{
			Name:      "directory-file-id",
			Required:  true,
			PathParam: "directory_file_id",
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
	Action:          handleBetaDirectoriesFilesDelete,
	HideHelpCommand: true,
}

var betaDirectoriesFilesAdd = cli.Command{
	Name:    "add",
	Usage:   "Create a new file within the specified directory; the directory must exist in\nthe project and `file_id` must reference an existing file.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "directory-id",
			Required:  true,
			PathParam: "directory_id",
		},
		&requestflag.Flag[string]{
			Name:     "file-id",
			Usage:    "File ID for the storage location (required).",
			Required: true,
			BodyPath: "file_id",
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
			Name:     "display-name",
			Usage:    "Display name for the file. If not provided, will use the file's name.",
			BodyPath: "display_name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Usage:    "User-defined metadata key-value pairs to associate with the file.",
			BodyPath: "metadata",
		},
		&requestflag.Flag[*string]{
			Name:     "unique-id",
			Usage:    "Unique identifier for the file in the directory. If not provided, will use the file's external_file_id or name.",
			BodyPath: "unique_id",
		},
	},
	Action:          handleBetaDirectoriesFilesAdd,
	HideHelpCommand: true,
}

var betaDirectoriesFilesGet = cli.Command{
	Name:    "get",
	Usage:   "Get a directory file by `directory_file_id`; to look up by `unique_id`, use the\nlist endpoint with a filter.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "directory-id",
			Required:  true,
			PathParam: "directory_id",
		},
		&requestflag.Flag[string]{
			Name:      "directory-file-id",
			Required:  true,
			PathParam: "directory_file_id",
		},
		&requestflag.Flag[any]{
			Name:      "expand",
			Usage:     "Fields to expand.",
			QueryPath: "expand",
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
	Action:          handleBetaDirectoriesFilesGet,
	HideHelpCommand: true,
}

var betaDirectoriesFilesUpload = cli.Command{
	Name:    "upload",
	Usage:   "Upload a file and create its directory entry in one call; `unique_id` /\n`display_name` default to values derived from file metadata.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "directory-id",
			Required:  true,
			PathParam: "directory_id",
		},
		&requestflag.Flag[string]{
			Name:      "upload-file",
			Required:  true,
			BodyPath:  "upload_file",
			FileInput: true,
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
			Name:     "display-name",
			BodyPath: "display_name",
		},
		&requestflag.Flag[*string]{
			Name:     "external-file-id",
			BodyPath: "external_file_id",
		},
		&requestflag.Flag[*string]{
			Name:     "metadata",
			Usage:    "User metadata as a JSON object string.",
			BodyPath: "metadata",
		},
		&requestflag.Flag[*string]{
			Name:     "unique-id",
			BodyPath: "unique_id",
		},
	},
	Action:          handleBetaDirectoriesFilesUpload,
	HideHelpCommand: true,
}

func handleBetaDirectoriesFilesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("directory-file-id") && len(unusedArgs) > 0 {
		cmd.Set("directory-file-id", unusedArgs[0])
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

	params := llamacloud.BetaDirectoryFileUpdateParams{
		DirectoryID: cmd.Value("directory-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Beta.Directories.Files.Update(
		ctx,
		cmd.Value("directory-file-id").(string),
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
		Title:          "beta:directories:files update",
		Transform:      transform,
	})
}

func handleBetaDirectoriesFilesList(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("directory-id") && len(unusedArgs) > 0 {
		cmd.Set("directory-id", unusedArgs[0])
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

	params := llamacloud.BetaDirectoryFileListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Beta.Directories.Files.List(
			ctx,
			cmd.Value("directory-id").(string),
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
			Title:          "beta:directories:files list",
			Transform:      transform,
		})
	} else {
		iter := client.Beta.Directories.Files.ListAutoPaging(
			ctx,
			cmd.Value("directory-id").(string),
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
			Title:          "beta:directories:files list",
			Transform:      transform,
		})
	}
}

func handleBetaDirectoriesFilesDelete(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("directory-file-id") && len(unusedArgs) > 0 {
		cmd.Set("directory-file-id", unusedArgs[0])
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

	params := llamacloud.BetaDirectoryFileDeleteParams{
		DirectoryID: cmd.Value("directory-id").(string),
	}

	return client.Beta.Directories.Files.Delete(
		ctx,
		cmd.Value("directory-file-id").(string),
		params,
		options...,
	)
}

func handleBetaDirectoriesFilesAdd(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("directory-id") && len(unusedArgs) > 0 {
		cmd.Set("directory-id", unusedArgs[0])
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

	params := llamacloud.BetaDirectoryFileAddParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Beta.Directories.Files.Add(
		ctx,
		cmd.Value("directory-id").(string),
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
		Title:          "beta:directories:files add",
		Transform:      transform,
	})
}

func handleBetaDirectoriesFilesGet(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("directory-file-id") && len(unusedArgs) > 0 {
		cmd.Set("directory-file-id", unusedArgs[0])
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

	params := llamacloud.BetaDirectoryFileGetParams{
		DirectoryID: cmd.Value("directory-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Beta.Directories.Files.Get(
		ctx,
		cmd.Value("directory-file-id").(string),
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
		Title:          "beta:directories:files get",
		Transform:      transform,
	})
}

func handleBetaDirectoriesFilesUpload(ctx context.Context, cmd *cli.Command) error {
	client := llamacloud.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("directory-id") && len(unusedArgs) > 0 {
		cmd.Set("directory-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
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

	params := llamacloud.BetaDirectoryFileUploadParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Beta.Directories.Files.Upload(
		ctx,
		cmd.Value("directory-id").(string),
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
		Title:          "beta:directories:files upload",
		Transform:      transform,
	})
}
