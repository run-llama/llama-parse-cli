// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/llamacloud-prod-cli/internal/mocktest"
	"github.com/stainless-sdks/llamacloud-prod-cli/internal/requestflag"
)

func TestRetrieversCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"retrievers", "create",
			"--name", "x",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--pipeline", "{description: description, name: x, pipeline_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, preset_retrieval_parameters: {alpha: 0, class_name: class_name, dense_similarity_cutoff: 0, dense_similarity_top_k: 1, enable_reranking: true, files_top_k: 1, rerank_top_n: 1, retrieval_mode: chunks, retrieve_image_nodes: true, retrieve_page_figure_nodes: true, retrieve_page_screenshot_nodes: true, search_filters: {filters: [{key: key, value: 0, operator: '=='}], condition: and}, search_filters_inference_schema: {foo: {foo: bar}}, sparse_similarity_top_k: 1}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(retrieversCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"retrievers", "create",
			"--name", "x",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--pipeline.description", "description",
			"--pipeline.name", "x",
			"--pipeline.pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--pipeline.preset-retrieval-parameters", "{alpha: 0, class_name: class_name, dense_similarity_cutoff: 0, dense_similarity_top_k: 1, enable_reranking: true, files_top_k: 1, rerank_top_n: 1, retrieval_mode: chunks, retrieve_image_nodes: true, retrieve_page_figure_nodes: true, retrieve_page_screenshot_nodes: true, search_filters: {filters: [{key: key, value: 0, operator: '=='}], condition: and}, search_filters_inference_schema: {foo: {foo: bar}}, sparse_similarity_top_k: 1}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: x\n" +
			"pipelines:\n" +
			"  - description: description\n" +
			"    name: x\n" +
			"    pipeline_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"    preset_retrieval_parameters:\n" +
			"      alpha: 0\n" +
			"      class_name: class_name\n" +
			"      dense_similarity_cutoff: 0\n" +
			"      dense_similarity_top_k: 1\n" +
			"      enable_reranking: true\n" +
			"      files_top_k: 1\n" +
			"      rerank_top_n: 1\n" +
			"      retrieval_mode: chunks\n" +
			"      retrieve_image_nodes: true\n" +
			"      retrieve_page_figure_nodes: true\n" +
			"      retrieve_page_screenshot_nodes: true\n" +
			"      search_filters:\n" +
			"        filters:\n" +
			"          - key: key\n" +
			"            value: 0\n" +
			"            operator: '=='\n" +
			"        condition: and\n" +
			"      search_filters_inference_schema:\n" +
			"        foo:\n" +
			"          foo: bar\n" +
			"      sparse_similarity_top_k: 1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"retrievers", "create",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestRetrieversUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"retrievers", "update",
			"--retriever-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--pipeline", "[{description: description, name: x, pipeline_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, preset_retrieval_parameters: {alpha: 0, class_name: class_name, dense_similarity_cutoff: 0, dense_similarity_top_k: 1, enable_reranking: true, files_top_k: 1, rerank_top_n: 1, retrieval_mode: chunks, retrieve_image_nodes: true, retrieve_page_figure_nodes: true, retrieve_page_screenshot_nodes: true, search_filters: {filters: [{key: key, value: 0, operator: '=='}], condition: and}, search_filters_inference_schema: {foo: {foo: bar}}, sparse_similarity_top_k: 1}}]",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--name", "name",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(retrieversUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"retrievers", "update",
			"--retriever-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--pipeline.description", "description",
			"--pipeline.name", "x",
			"--pipeline.pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--pipeline.preset-retrieval-parameters", "{alpha: 0, class_name: class_name, dense_similarity_cutoff: 0, dense_similarity_top_k: 1, enable_reranking: true, files_top_k: 1, rerank_top_n: 1, retrieval_mode: chunks, retrieve_image_nodes: true, retrieve_page_figure_nodes: true, retrieve_page_screenshot_nodes: true, search_filters: {filters: [{key: key, value: 0, operator: '=='}], condition: and}, search_filters_inference_schema: {foo: {foo: bar}}, sparse_similarity_top_k: 1}",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--name", "name",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"pipelines:\n" +
			"  - description: description\n" +
			"    name: x\n" +
			"    pipeline_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"    preset_retrieval_parameters:\n" +
			"      alpha: 0\n" +
			"      class_name: class_name\n" +
			"      dense_similarity_cutoff: 0\n" +
			"      dense_similarity_top_k: 1\n" +
			"      enable_reranking: true\n" +
			"      files_top_k: 1\n" +
			"      rerank_top_n: 1\n" +
			"      retrieval_mode: chunks\n" +
			"      retrieve_image_nodes: true\n" +
			"      retrieve_page_figure_nodes: true\n" +
			"      retrieve_page_screenshot_nodes: true\n" +
			"      search_filters:\n" +
			"        filters:\n" +
			"          - key: key\n" +
			"            value: 0\n" +
			"            operator: '=='\n" +
			"        condition: and\n" +
			"      search_filters_inference_schema:\n" +
			"        foo:\n" +
			"          foo: bar\n" +
			"      sparse_similarity_top_k: 1\n" +
			"name: name\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"retrievers", "update",
			"--retriever-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestRetrieversList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"retrievers", "list",
			"--name", "name",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestRetrieversDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"retrievers", "delete",
			"--retriever-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestRetrieversGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"retrievers", "get",
			"--retriever-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestRetrieversSearch(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"retrievers", "search",
			"--query", "x",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--mode", "routing",
			"--pipeline", "{description: description, name: x, pipeline_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, preset_retrieval_parameters: {alpha: 0, class_name: class_name, dense_similarity_cutoff: 0, dense_similarity_top_k: 1, enable_reranking: true, files_top_k: 1, rerank_top_n: 1, retrieval_mode: chunks, retrieve_image_nodes: true, retrieve_page_figure_nodes: true, retrieve_page_screenshot_nodes: true, search_filters: {filters: [{key: key, value: 0, operator: '=='}], condition: and}, search_filters_inference_schema: {foo: {foo: bar}}, sparse_similarity_top_k: 1}}",
			"--rerank-config", "{top_n: 1, type: system_default}",
			"--rerank-top-n", "0",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(retrieversSearch)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"retrievers", "search",
			"--query", "x",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--mode", "routing",
			"--pipeline.description", "description",
			"--pipeline.name", "x",
			"--pipeline.pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--pipeline.preset-retrieval-parameters", "{alpha: 0, class_name: class_name, dense_similarity_cutoff: 0, dense_similarity_top_k: 1, enable_reranking: true, files_top_k: 1, rerank_top_n: 1, retrieval_mode: chunks, retrieve_image_nodes: true, retrieve_page_figure_nodes: true, retrieve_page_screenshot_nodes: true, search_filters: {filters: [{key: key, value: 0, operator: '=='}], condition: and}, search_filters_inference_schema: {foo: {foo: bar}}, sparse_similarity_top_k: 1}",
			"--rerank-config.top-n", "1",
			"--rerank-config.type", "system_default",
			"--rerank-top-n", "0",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"query: x\n" +
			"mode: routing\n" +
			"pipelines:\n" +
			"  - description: description\n" +
			"    name: x\n" +
			"    pipeline_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"    preset_retrieval_parameters:\n" +
			"      alpha: 0\n" +
			"      class_name: class_name\n" +
			"      dense_similarity_cutoff: 0\n" +
			"      dense_similarity_top_k: 1\n" +
			"      enable_reranking: true\n" +
			"      files_top_k: 1\n" +
			"      rerank_top_n: 1\n" +
			"      retrieval_mode: chunks\n" +
			"      retrieve_image_nodes: true\n" +
			"      retrieve_page_figure_nodes: true\n" +
			"      retrieve_page_screenshot_nodes: true\n" +
			"      search_filters:\n" +
			"        filters:\n" +
			"          - key: key\n" +
			"            value: 0\n" +
			"            operator: '=='\n" +
			"        condition: and\n" +
			"      search_filters_inference_schema:\n" +
			"        foo:\n" +
			"          foo: bar\n" +
			"      sparse_similarity_top_k: 1\n" +
			"rerank_config:\n" +
			"  top_n: 1\n" +
			"  type: system_default\n" +
			"rerank_top_n: 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"retrievers", "search",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestRetrieversUpsert(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"retrievers", "upsert",
			"--name", "x",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--pipeline", "{description: description, name: x, pipeline_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, preset_retrieval_parameters: {alpha: 0, class_name: class_name, dense_similarity_cutoff: 0, dense_similarity_top_k: 1, enable_reranking: true, files_top_k: 1, rerank_top_n: 1, retrieval_mode: chunks, retrieve_image_nodes: true, retrieve_page_figure_nodes: true, retrieve_page_screenshot_nodes: true, search_filters: {filters: [{key: key, value: 0, operator: '=='}], condition: and}, search_filters_inference_schema: {foo: {foo: bar}}, sparse_similarity_top_k: 1}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(retrieversUpsert)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"retrievers", "upsert",
			"--name", "x",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--pipeline.description", "description",
			"--pipeline.name", "x",
			"--pipeline.pipeline-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--pipeline.preset-retrieval-parameters", "{alpha: 0, class_name: class_name, dense_similarity_cutoff: 0, dense_similarity_top_k: 1, enable_reranking: true, files_top_k: 1, rerank_top_n: 1, retrieval_mode: chunks, retrieve_image_nodes: true, retrieve_page_figure_nodes: true, retrieve_page_screenshot_nodes: true, search_filters: {filters: [{key: key, value: 0, operator: '=='}], condition: and}, search_filters_inference_schema: {foo: {foo: bar}}, sparse_similarity_top_k: 1}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: x\n" +
			"pipelines:\n" +
			"  - description: description\n" +
			"    name: x\n" +
			"    pipeline_id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"    preset_retrieval_parameters:\n" +
			"      alpha: 0\n" +
			"      class_name: class_name\n" +
			"      dense_similarity_cutoff: 0\n" +
			"      dense_similarity_top_k: 1\n" +
			"      enable_reranking: true\n" +
			"      files_top_k: 1\n" +
			"      rerank_top_n: 1\n" +
			"      retrieval_mode: chunks\n" +
			"      retrieve_image_nodes: true\n" +
			"      retrieve_page_figure_nodes: true\n" +
			"      retrieve_page_screenshot_nodes: true\n" +
			"      search_filters:\n" +
			"        filters:\n" +
			"          - key: key\n" +
			"            value: 0\n" +
			"            operator: '=='\n" +
			"        condition: and\n" +
			"      search_filters_inference_schema:\n" +
			"        foo:\n" +
			"          foo: bar\n" +
			"      sparse_similarity_top_k: 1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"retrievers", "upsert",
			"--organization-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--project-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
