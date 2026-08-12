# Changelog

## [1.5.0](https://github.com/run-llama/llama-parse-cli/compare/v1.4.0...v1.5.0) (2026-08-12)


### Features

* **classify:** accept webhook_configuration_ids on classify job create (LI-8138) ([#22943](https://github.com/run-llama/llama-parse-cli/issues/22943)) ([087ee56](https://github.com/run-llama/llama-parse-cli/commit/087ee565ff2f7d02f99ba93a55666cbccb15f0e1))
* **connector:** API + service layer for attaching a subscription to a directory ([#23502](https://github.com/run-llama/llama-parse-cli/issues/23502)) ([b3d53f6](https://github.com/run-llama/llama-parse-cli/commit/b3d53f6feb55174fc6fdf32c1ba9bb8d83983752))
* **extract:** accept webhook_configuration_ids on extract job create (LI-8138) ([#22907](https://github.com/run-llama/llama-parse-cli/issues/22907)) ([3705dd2](https://github.com/run-llama/llama-parse-cli/commit/3705dd27463d6d5174482204c23be6aaf1136d21))
* **extract:** pin turbo to a stable dated version; accept citations+confidence, reject only granular bboxes ([#22965](https://github.com/run-llama/llama-parse-cli/issues/22965)) ([a9b12bf](https://github.com/run-llama/llama-parse-cli/commit/a9b12bf96962672f8584ba2120ae455884840e4c))
* **extract:** reject parse_tier for parse-free tiers + pin turbo fallback to fast ([#22919](https://github.com/run-llama/llama-parse-cli/issues/22919)) ([59351c3](https://github.com/run-llama/llama-parse-cli/commit/59351c3a92c780171a58ef8c257196ba3e7bf5a9))
* **files:** rename files.get to files.content and restore files.retrieve ([2fd9092](https://github.com/run-llama/llama-parse-cli/commit/2fd90928eb45808167d518b7e8cc8723d6b6b25b))
* **parse,extract:** add expand=usage returning credits billed per job ([#23709](https://github.com/run-llama/llama-parse-cli/issues/23709)) ([1fb5977](https://github.com/run-llama/llama-parse-cli/commit/1fb59777faca00a50511e3fdcf4601cc80a2e874))
* **parse:** make the output.pdf artifact opt-in on Parse v2 (output_options.save_output_pdf) ([#23510](https://github.com/run-llama/llama-parse-cli/issues/23510)) ([1108da9](https://github.com/run-llama/llama-parse-cli/commit/1108da9935dfa061401be9919e100ed9bcfc1e14))
* **split:** accept webhook_configuration_ids on split job create (LI-8138) ([#22940](https://github.com/run-llama/llama-parse-cli/issues/22940)) ([1e7b609](https://github.com/run-llama/llama-parse-cli/commit/1e7b6098a4c766db73915c9870e8e86379510729))


### Bug Fixes

* build against llama-parse-go v1.5.0 and drop the duplicated TestClassifyCancel ([#13](https://github.com/run-llama/llama-parse-cli/issues/13)) ([f0eb259](https://github.com/run-llama/llama-parse-cli/commit/f0eb259128bea0e0bb87a9e736dce5d05634708e))
* **llamaparse:** retry qwen context-overflow 400s with a shrinking OCR anchor ([#23817](https://github.com/run-llama/llama-parse-cli/issues/23817)) ([b98043b](https://github.com/run-llama/llama-parse-cli/commit/b98043b7be875eca3c7dd0e1f88db6a0d9d063e6))


### Chores

* **api:** regenerate OpenAPI specs for new agentic parse version ([#22763](https://github.com/run-llama/llama-parse-cli/issues/22763)) ([2138578](https://github.com/run-llama/llama-parse-cli/commit/2138578c788587c4de2c0e7016d135ef3fe717bb))


### Documentation

* **parse:** shorten the images_to_save field description ([#23807](https://github.com/run-llama/llama-parse-cli/issues/23807)) ([2f78313](https://github.com/run-llama/llama-parse-cli/commit/2f78313313a0e6d93e204ecac4f3f207dee4a315))


### Refactors

* remove Depends(get_db) from permissions endpoints ([#22635](https://github.com/run-llama/llama-parse-cli/issues/22635)) ([087b803](https://github.com/run-llama/llama-parse-cli/commit/087b8037dcbe78b1abc643c38d7ce73de48241b0))

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
