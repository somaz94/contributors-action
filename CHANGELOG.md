# Changelog

All notable changes to this project will be documented in this file.

## Unreleased (2026-10-07)

### Continuous Integration

- correct stale image-seeding comments ([e412c7e](https://github.com/somaz94/contributors-action/commit/e412c7e2f7f7bf78924910418d6408b797d84359))
- trim redundant comments in gitlab-mirror workflow ([53a90ee](https://github.com/somaz94/contributors-action/commit/53a90eea94173ee906154f0be95fde10133b8416))
- correct the image-seeding comment in the release workflow ([ab0cb6d](https://github.com/somaz94/contributors-action/commit/ab0cb6d40ee6bc41679898ed58f3081e228763b2))
- retry mirror pushes on transient remote failures ([5e49b27](https://github.com/somaz94/contributors-action/commit/5e49b277ca45a0f402826bf67f4dca074e124e64))
- drop the dead issue-close trigger from changelog generation ([d026dad](https://github.com/somaz94/contributors-action/commit/d026dad7a5ac7af3dad671f8e158f42509cc23bb))
- skip release-triggered runs on the image-seeding dispatch ([479ca4c](https://github.com/somaz94/contributors-action/commit/479ca4c220a714996a323de75ce10edba23dfce4))

### Chores

- **deps:** bump golang from `cf6fca6` to `8a5910f` (#16) ([#16](https://github.com/somaz94/contributors-action/pull/16)) ([27daf3e](https://github.com/somaz94/contributors-action/commit/27daf3e8d0e63257d84b3994759ff7ea83445f4d))
- tighten golangci-lint config comments ([8fd9d1d](https://github.com/somaz94/contributors-action/commit/8fd9d1d9a1d73a196ef323fe9efd9ed7f368a15e))
- **deps:** bump alpine from `28bd5fe` to `294b683` (#15) ([#15](https://github.com/somaz94/contributors-action/pull/15)) ([fea42ba](https://github.com/somaz94/contributors-action/commit/fea42ba29ee90a0b55a793d4c1d4657755be3c41))
- **deps:** bump golang from `4c9fe60` to `cf6fca6` ([6274c59](https://github.com/somaz94/contributors-action/commit/6274c591e0430595b37a37e30af27ccd63f96967))
- **deps:** bump golang in the docker-minor group (#13) ([#13](https://github.com/somaz94/contributors-action/pull/13)) ([2c949fb](https://github.com/somaz94/contributors-action/commit/2c949fb6ff038eab0b2bd20c0b7b525377b87a7d))
- **deps:** bump golang from `0178a64` to `3889b42` ([0f2e3a7](https://github.com/somaz94/contributors-action/commit/0f2e3a72a4f3c4df7f69fd1cd2e0d2817d2257fa))

### Contributors

- somaz

<br/>

## [v1.2.0](https://github.com/somaz94/contributors-action/compare/v1.1.2...v1.2.0) (2026-08-07)

### Performance Improvements

- ship a prebuilt multi-arch image instead of building per run ([09b487e](https://github.com/somaz94/contributors-action/commit/09b487e294a0eebef61dbf0ca3a6415e3e60ea81))

### Continuous Integration

- add a golangci-lint config scoped to defect-finding linters ([cdefed2](https://github.com/somaz94/contributors-action/commit/cdefed2c5a29f95126f0df30c91e0734ffab7a94))
- remove DCO workflow ([3586d7b](https://github.com/somaz94/contributors-action/commit/3586d7bc3dac97048ef8067298531f041d145741))
- adopt semantic-pr, labels, lock-threads, PR size, and auto-assign reusables ([8237c61](https://github.com/somaz94/contributors-action/commit/8237c61808cb814fef2165f8b08828648b06bf55))
- use reusable stale-issues workflow ([dbfc850](https://github.com/somaz94/contributors-action/commit/dbfc850fe043b60e8ba0f6d107b1b569789c11c6))
- use reusable issue-greeting workflow ([cc2fd4e](https://github.com/somaz94/contributors-action/commit/cc2fd4e1adbbedd7cf24a394ee70cf594b1d18f0))
- use reusable dependabot-auto-merge workflow ([e504e39](https://github.com/somaz94/contributors-action/commit/e504e39feb1964f8d591b4fe3b7b56e9c1b69f47))
- use reusable contributors workflow ([18e051a](https://github.com/somaz94/contributors-action/commit/18e051a7ffeb35448e1ce89b2670dd3cf7d61f08))
- add ok-to-test workflow stub ([d65d205](https://github.com/somaz94/contributors-action/commit/d65d205e31d02b00daa1ca7491618935dd346773))
- add PR welcome workflow stub ([bfc1201](https://github.com/somaz94/contributors-action/commit/bfc1201f42d5a6fe0409ecba243e9b793c408a5f))

### Chores

- **deps:** bump golang from `3ad5730` to `0178a64` (#11) ([#11](https://github.com/somaz94/contributors-action/pull/11)) ([5557c4a](https://github.com/somaz94/contributors-action/commit/5557c4a97ec4114b50cac1c229c8abdf927353ab))

### Contributors

- somaz

<br/>

## [v1.1.2](https://github.com/somaz94/contributors-action/compare/v1.1.1...v1.1.2) (2026-06-24)

### Documentation

- front-load output formats in README and action.yml description ([11bedeb](https://github.com/somaz94/contributors-action/commit/11bedebb8a6ae8b12667f97947289713e1c37c96))

### Continuous Integration

- add DCO check via shared reusable workflow ([49ef48c](https://github.com/somaz94/contributors-action/commit/49ef48cf041f6f186432b4760e987a27fa853e51))

### Chores

- **deps:** bump golang from `f1ddd9f` to `3ad5730` (#9) ([#9](https://github.com/somaz94/contributors-action/pull/9)) ([523aa5e](https://github.com/somaz94/contributors-action/commit/523aa5e7bf4cf01dbdf0df920307975ead25479f))
- **deps:** bump actions/checkout from 6 to 7 (#10) ([#10](https://github.com/somaz94/contributors-action/pull/10)) ([883a3a2](https://github.com/somaz94/contributors-action/commit/883a3a2f99576d68186b91dbc472cf0320bdca26))
- **deps:** bump alpine from `f5064d3` to `28bd5fe` (#8) ([#8](https://github.com/somaz94/contributors-action/pull/8)) ([0045e83](https://github.com/somaz94/contributors-action/commit/0045e83b458b0c9992f37aeba93b4065ec209e47))
- **deps:** bump golang from `f23e8b2` to `f1ddd9f` ([6ee4c24](https://github.com/somaz94/contributors-action/commit/6ee4c24bf039c468e91ebec3797c4239a4435e57))
- **deps:** bump alpine from 3.23 to 3.24 in the docker-minor group (#6) ([#6](https://github.com/somaz94/contributors-action/pull/6)) ([cbc5c81](https://github.com/somaz94/contributors-action/commit/cbc5c81c8133103752c8399b24e3d435558e32fa))
- **deps:** bump golang from `91eda97` to `f23e8b2` ([66a2b8a](https://github.com/somaz94/contributors-action/commit/66a2b8a63fac2e0b4f198ebbbf19472b23cbc1e3))

### Contributors

- somaz

<br/>

## [v1.1.1](https://github.com/somaz94/contributors-action/compare/v1.1.0...v1.1.1) (2026-05-28)

### Code Refactoring

- harden CI workflows, Dockerfile, and Go pipeline ([6354000](https://github.com/somaz94/contributors-action/commit/6354000c61b73d2521c20650b06052b05174001f))

### Continuous Integration

- add concurrency guards to recurring workflows ([f67f6bc](https://github.com/somaz94/contributors-action/commit/f67f6bc2814f30c56210a954ae0a9e6addfc0851))
- use go-docker-action-ci-action@v1 (replace inline prelude) ([898e2fc](https://github.com/somaz94/contributors-action/commit/898e2fc08801e4e0236acb6061a4ead6a2fa70bb))

### Chores

- **deps:** bump softprops/action-gh-release from 2 to 3 ([679e9a0](https://github.com/somaz94/contributors-action/commit/679e9a00029827d324cd8521020ec5547b9f16d5))
- **deps:** bump actions/github-script from 8 to 9 ([b83a411](https://github.com/somaz94/contributors-action/commit/b83a41192af05617f6d8e16af389bed6b3016a7f))
- **deps:** bump dependabot/fetch-metadata from 2 to 3 ([616d4cd](https://github.com/somaz94/contributors-action/commit/616d4cdfae55eae61d2c5b0bc719d120f384e883))

### Contributors

- somaz

<br/>

## [v1.1.0](https://github.com/somaz94/contributors-action/compare/v1.0.0...v1.1.0) (2026-04-03)

### Bug Fixes

- rename action to Contributors Generator Action ([1dcdfd8](https://github.com/somaz94/contributors-action/commit/1dcdfd80fede8fdea7609ba9bf5b20e876ef264b))

### Code Refactoring

- improve code quality with validation, helpers, and test modernization ([d611667](https://github.com/somaz94/contributors-action/commit/d611667c0d659dc8f747b057384be597395bb5b7))

### Documentation

- remove duplicate rules covered by global CLAUDE.md ([c65d998](https://github.com/somaz94/contributors-action/commit/c65d998c511624027a4d1523dc991142a6f97b25))
- add no-push rule to CLAUDE.md ([6e533e1](https://github.com/somaz94/contributors-action/commit/6e533e17307106083703aa6bb7d1ee44498fee05))
- add CLAUDE.md project guide ([a74bf7c](https://github.com/somaz94/contributors-action/commit/a74bf7c486ae8a56878d7d4c37294bcc485f479f))

### Continuous Integration

- skip auto-generated changelog and contributors commits in release notes ([fccd17a](https://github.com/somaz94/contributors-action/commit/fccd17a41b127424f4014c2555eb5d43d0ffd8e4))
- revert to body_path RELEASE.md in release workflow ([ddd5eca](https://github.com/somaz94/contributors-action/commit/ddd5eca3f5f379984dbf6b035d7673b80204532f))
- use generate_release_notes instead of body_path in release workflow ([6d19f1e](https://github.com/somaz94/contributors-action/commit/6d19f1e124dafd62ff86c2ad315a4dd0a55854f6))
- migrate gitlab-mirror workflow to multi-git-mirror action ([67c625f](https://github.com/somaz94/contributors-action/commit/67c625f11be382acb7a4afdbf314ee53f98bd3ce))

### Chores

- remove duplicate rules from CLAUDE.md (moved to global) ([180eca3](https://github.com/somaz94/contributors-action/commit/180eca3077a6c2027e0985e50a92eed4115bf9b1))
- add git config protection to CLAUDE.md ([15857ab](https://github.com/somaz94/contributors-action/commit/15857ab701d1b6747c26a096388d218bdc6e4651))

### Contributors

- somaz

<br/>

## [v1.0.0](https://github.com/somaz94/contributors-action/releases/tag/v1.0.0) (2026-03-17)

### Features

- implement Go-based contributors action ([e24e634](https://github.com/somaz94/contributors-action/commit/e24e634000d67c8a5b8b5dd94c026174c903a58b))

### Bug Fixes

- replace real GitHub user IDs with placeholders in examples ([93091a3](https://github.com/somaz94/contributors-action/commit/93091a36901139663f92c00108f97ffa0ebcdb51))
- remove go.sum reference from Dockerfile ([0e9af5a](https://github.com/somaz94/contributors-action/commit/0e9af5a3e7c0b0a52d1c42b725e9c0e5f19c5134))
- update Go to 1.26 and Alpine to 3.23 ([1be9201](https://github.com/somaz94/contributors-action/commit/1be9201888c9154d7a4a570a6b80d471c369f3ee))

### Documentation

- create docs/ and streamline README with example output ([9d37a18](https://github.com/somaz94/contributors-action/commit/9d37a18d07c497856706d2d0875bedd034d8e52c))
- add README, .dockerignore, cliff.toml, update .gitignore ([807515b](https://github.com/somaz94/contributors-action/commit/807515bf7853c69ccaadb48c64c7904f6158f18e))

### Continuous Integration

- add operational workflows ([583550d](https://github.com/somaz94/contributors-action/commit/583550d0a0b0c35bbecf6b50290f9986300021d4))
- add CI/CD workflows and release pipeline ([11fce14](https://github.com/somaz94/contributors-action/commit/11fce1440b6fb3eb1570868f2a8d0d5fa3d46b07))

### Chores

- **deps:** bump the docker-minor group with 2 updates (#1) ([#1](https://github.com/somaz94/contributors-action/pull/1)) ([894f378](https://github.com/somaz94/contributors-action/commit/894f378edaae570fa90b6ecea346ccf432e7644d))
- remove template scaffolding and update CODEOWNERS ([5ea5ac8](https://github.com/somaz94/contributors-action/commit/5ea5ac82e891268a53c72330a2991b1bd6c53d6b))

### Contributors

- somaz

<br/>

