# Changelog

## [1.8.0](https://github.com/run-llama/llama-parse-cli/compare/v1.7.0...v1.8.0) (2026-10-07)


### Features

* add redline prompt changes to new prod version ([#27635](https://github.com/run-llama/llama-parse-cli/issues/27635)) ([30df048](https://github.com/run-llama/llama-parse-cli/commit/30df048e5a4acefa85c1679d653c733c1704b175))
* **chat:** let a chat session refuse queries from its share link ([#27012](https://github.com/run-llama/llama-parse-cli/issues/27012)) ([200379d](https://github.com/run-llama/llama-parse-cli/commit/200379d50fb14e615083b7b1c573b61376bb3598))
* **parse:** add option to include hidden PPTX slides ([#27938](https://github.com/run-llama/llama-parse-cli/issues/27938)) ([11368da](https://github.com/run-llama/llama-parse-cli/commit/11368da140781fdb9c6702896ac2b8ffccc0aee1))
* **parse:** agentic 2026-09-09 — cache-stable prompt order + Flash Lite MINIMAL thinking ([#26273](https://github.com/run-llama/llama-parse-cli/issues/26273)) ([3125104](https://github.com/run-llama/llama-parse-cli/commit/3125104e30788fdd49b306ac0d47a2d3e3b04e60))
* **parse:** apply watermark_handling to text output; add watermark e2e test ([#27932](https://github.com/run-llama/llama-parse-cli/issues/27932)) ([7ac7ed7](https://github.com/run-llama/llama-parse-cli/commit/7ac7ed7baee93bce3119a384ff0b4499dbb7b671))
* **parse:** remove_watermark output option with 2026-09-28 tier versions ([#27813](https://github.com/run-llama/llama-parse-cli/issues/27813)) ([67c5f4a](https://github.com/run-llama/llama-parse-cli/commit/67c5f4affb4f5730c3b99b2c8131549e4c77358c))
* **sdk:** publish beta.attachments list and get ([0755af8](https://github.com/run-llama/llama-parse-cli/commit/0755af8219e67064d28b2c4d3e41f3c8d7ff164b))
* **split:** target_pages page selection when splitting a parse job ([#26921](https://github.com/run-llama/llama-parse-cli/issues/26921)) ([0b72df4](https://github.com/run-llama/llama-parse-cli/commit/0b72df4926afb4dc6fa4654fae0fa24e5116f13a))


### Bug Fixes

* **chat:** report every index a chat turn could not query ([#26981](https://github.com/run-llama/llama-parse-cli/issues/26981)) ([3c1645b](https://github.com/run-llama/llama-parse-cli/commit/3c1645b8e24e702eb79c1fa39c95dc7f767e0656))
* **classifier:** drop the dead classify v1 job mappings (methods were removed from the SDKs in 2.16.0 / 1.7.0) ([20bd589](https://github.com/run-llama/llama-parse-cli/commit/20bd589c618aea25d00c9f03f2d46a946ff752e2))
* **extract:** refuse to delete a non-terminal job (LI-8700) ([#23793](https://github.com/run-llama/llama-parse-cli/issues/23793)) ([c95af7f](https://github.com/run-llama/llama-parse-cli/commit/c95af7f9a57c2a8d7cd0593d2506253735453d62))
* **split:** process target_pages in the order written, matching Extract ([#28173](https://github.com/run-llama/llama-parse-cli/issues/28173)) ([efbd797](https://github.com/run-llama/llama-parse-cli/commit/efbd7975c1f470613f930f3393d4c179cf925031))


### Chores

* **deps:** bump llama-parse-go to v1.8.0 ([0d1da10](https://github.com/run-llama/llama-parse-cli/commit/0d1da106e0c30645d9c654d31b206e55c2ad4e65))
* **sync:** resolve back-sync conflicts with production ([3fcfcbd](https://github.com/run-llama/llama-parse-cli/commit/3fcfcbd0d0d6af9160583e55a206e7429bd20f9d))


### Documentation

* **changelog:** record the classify v1 job command removal in 1.7.0 (LI-9592) ([83f5280](https://github.com/run-llama/llama-parse-cli/commit/83f5280e06ac019f6894e80ee596f67ac8d506fa))
* **changelog:** record the classify v1 job command removal in 1.7.0 (LI-9592) ([7b73541](https://github.com/run-llama/llama-parse-cli/commit/7b735418531fa4fba6e5d03a9100fcebe816855d))
* update Extract versions and pricing ([#28128](https://github.com/run-llama/llama-parse-cli/issues/28128)) ([51934da](https://github.com/run-llama/llama-parse-cli/commit/51934da3a6cea8d0c60d4e8a628f921116c6e58f))

## [1.7.0](https://github.com/run-llama/llama-parse-cli/compare/v1.6.0...v1.7.0) (2026-09-08)


### ⚠ BREAKING CHANGES

* **classifier:** the `llp classifier:jobs` commands (`create`, `list`, `get`, `get-results`) are removed. The `/api/v1/classifier/jobs*` routes were unpublished from the API surface; use `llp classify` instead.


### Features

* **api:** map DELETE /api/v2/parse/{job_id} and GET /api/v2/pipelines into the SDKs (LI-9569) ([4d53048](https://github.com/run-llama/llama-parse-cli/commit/4d53048dd54094fe50fb943064b8bf71334aab76))


### Chores

* **deps:** bump llama-parse-go to v1.7.0 ([bb97772](https://github.com/run-llama/llama-parse-cli/commit/bb97772dbda5326b11ff50a49a2a455e11dbff86))

## [1.6.0](https://github.com/run-llama/llama-parse-cli/compare/v1.5.1...v1.6.0) (2026-08-28)


### Features

* **extract:** publish the turbo tier on the public API surface (LI-8873) ([#25281](https://github.com/run-llama/llama-parse-cli/issues/25281)) ([b1cd5ba](https://github.com/run-llama/llama-parse-cli/commit/b1cd5bab214238266501ed3878c99a61ce1b5d21))

## [1.5.1](https://github.com/run-llama/llama-parse-cli/compare/v1.5.0...v1.5.1) (2026-08-20)


### Features

* **parse:** expose legal-document preset in v2 ([#24192](https://github.com/run-llama/llama-parse-cli/issues/24192)) ([0907de9](https://github.com/run-llama/llama-parse-cli/commit/0907de9633a645929c42e5ae2dfff7be2d6b240b))


### Bug Fixes

* **parse:** restore inline_images on cost_effective (new version 2026-08-11) ([#23804](https://github.com/run-llama/llama-parse-cli/issues/23804)) ([d708ab4](https://github.com/run-llama/llama-parse-cli/commit/d708ab41cd3f1ba13d5efb3ee76ea21a1e15c394))


### Reverts

* expose legal-document preset in v2 ([3b7af27](https://github.com/run-llama/llama-parse-cli/commit/3b7af27a1cf4dc08995eb3e815fc444465de29ae))

## [1.4.0](https://github.com/run-llama/llama-parse-cli/compare/v1.3.0...v1.4.0) (2026-07-22)


### Features

* **sheets:** cost_effective/agentic tiers and per-region billing ([#22508](https://github.com/run-llama/llama-parse-cli/issues/22508)) ([99079c9](https://github.com/run-llama/llama-parse-cli/commit/99079c982a787a7f9b1fd78f79cbc73eff62e957))

## [1.3.0](https://github.com/run-llama/llama-parse-cli/compare/v1.2.0...v1.3.0) (2026-07-21)


### Features

* **gdrive:** reuse-first connection picker in the data-source connect modal ([#21725](https://github.com/run-llama/llama-parse-cli/issues/21725)) ([12577ba](https://github.com/run-llama/llama-parse-cli/commit/12577ba9a2864e5d6947cf7c01dd31872e0381ec))
* **llamaparse:** agentic 2026-07-15 — Markdown-pipe table body for Gemini 3.1 Flash-Lite (EU primary) ([#22208](https://github.com/run-llama/llama-parse-cli/issues/22208)) ([7f31e38](https://github.com/run-llama/llama-parse-cli/commit/7f31e38ffc39dedbe7f646fe97f26d4f033771d3))
* **parse:** rename confidence scoring option + billing event (confidence_score_effort / confidence_score_high) ([#22290](https://github.com/run-llama/llama-parse-cli/issues/22290)) ([c228ff5](https://github.com/run-llama/llama-parse-cli/commit/c228ff5bb5917bda0db22da464b99d7caf72f72a))


### Bug Fixes

* **deps:** bump llama-parse-go to v1.3.0 for llamacloud package rename ([5eed376](https://github.com/run-llama/llama-parse-cli/commit/5eed3765750723415ed5e9c51b1fec76da2355f6))

## [1.2.0](https://github.com/run-llama/llama-parse-cli/compare/v1.1.0...v1.2.0) (2026-07-09)


### Features

* **agentic-plus:** dated version 2026-07-08 — graduate decomposed-gemini (flash-lite), fallback to 2026-06-18 ([#21738](https://github.com/run-llama/llama-parse-cli/issues/21738)) ([fce988c](https://github.com/run-llama/llama-parse-cli/commit/fce988cb14957016090e8fb65874a6bf73b6c7ca))
* rename binary from 'llamacloud-prod' to 'llp' ([7f2ec4a](https://github.com/run-llama/llama-parse-cli/commit/7f2ec4a613ae8bb8038a4d85c3480419f6239bef))
* update fast tier latest version to use liteparse + markdown ([#21669](https://github.com/run-llama/llama-parse-cli/issues/21669)) ([56691cf](https://github.com/run-llama/llama-parse-cli/commit/56691cf689deb94e8e9d9aac255afeb878a5e1b8))


### Bug Fixes

* resolve conflict markers committed by the auto-resolver ([097da65](https://github.com/run-llama/llama-parse-cli/commit/097da6565493588e483a137c932bd0d79d1a9fc7))


### Chores

* set release-please manifest to 1.1.0 to match production ([6e8cea2](https://github.com/run-llama/llama-parse-cli/commit/6e8cea2858007b928d7eeb621b5171d0f6bb6619))

## [1.1.0](https://github.com/run-llama/llama-parse-cli/compare/v1.0.0...v1.1.0) (2026-06-30)


### Chores

* correct release-please manifest to 1.0.0 (the released tag) ([7b79129](https://github.com/run-llama/llama-parse-cli/commit/7b7912940d45ebd5cf8b7a34d0e7e88a772acfe5))
* release 1.1.0 ([8c0b6a6](https://github.com/run-llama/llama-parse-cli/commit/8c0b6a661b65745080c671f8f9940ae5492ad29c))
