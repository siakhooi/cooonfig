# cooonfig

[![License](https://img.shields.io/github/license/siakhooi/cooonfig)](https://github.com/siakhooi/cooonfig/blob/main/LICENSE)
[![Release](https://img.shields.io/github/v/release/siakhooi/cooonfig)](https://github.com/siakhooi/cooonfig/releases/latest)
[![Build](https://img.shields.io/github/actions/workflow/status/siakhooi/cooonfig/build.yaml?label=build)](https://github.com/siakhooi/cooonfig/actions/workflows/build.yaml)
[![Go Reference](https://pkg.go.dev/badge/github.com/siakhooi/cooonfig.svg)](https://pkg.go.dev/github.com/siakhooi/cooonfig)
[![Quality Gate](https://sonarcloud.io/api/project_badges/measure?project=siakhooi_cooonfig&metric=alert_status)](https://sonarcloud.io/project/overview?id=siakhooi_cooonfig)
[![Coverage](https://qlty.sh/gh/siakhooi/projects/cooonfig/coverage.svg)](https://qlty.sh/gh/siakhooi/projects/cooonfig)
[![Funding](https://img.shields.io/badge/Funding-Wise-33cb56.svg?logo=wise)](https://wise.com/pay/me/siakn3)
![visitors](https://hit-tztugwlsja-uc.a.run.app/?outputtype=badge&counter=ghmd-cooonfig)

transform config files

## Usage

```
cooonfig --help
cooonfig --version
```

`--version` prints the version, commit, and build date injected by `scripts/build.sh` and GoReleaser.

## Development

```
just ci
just build
```

`just build` writes binaries to `bin/` and snapshot packages to `dist/`. `just release` creates a GitHub release from `release.env`.
