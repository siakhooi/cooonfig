# cooonfig

[![License](https://img.shields.io/github/license/siakhooi/cooonfig)](https://github.com/siakhooi/cooonfig/blob/main/LICENSE)
[![Release](https://img.shields.io/github/v/release/siakhooi/cooonfig)](https://github.com/siakhooi/cooonfig/releases/latest)
[![Build](https://img.shields.io/github/actions/workflow/status/siakhooi/cooonfig/build.yaml?label=build)](https://github.com/siakhooi/cooonfig/actions/workflows/build.yaml)
[![Go Reference](https://pkg.go.dev/badge/github.com/siakhooi/cooonfig.svg)](https://pkg.go.dev/github.com/siakhooi/cooonfig)
[![Quality Gate](https://sonarcloud.io/api/project_badges/measure?project=siakhooi_cooonfig&metric=alert_status)](https://sonarcloud.io/project/overview?id=siakhooi_cooonfig)
[![Coverage](https://qlty.sh/gh/siakhooi/projects/cooonfig/coverage.svg)](https://qlty.sh/gh/siakhooi/projects/cooonfig)
[![Funding](https://img.shields.io/badge/Funding-Wise-33cb56.svg?logo=wise)](https://wise.com/pay/me/siakn3)
![visitors](https://hit-tztugwlsja-uc.a.run.app/?outputtype=badge&counter=ghmd-cooonfig)

cooonfig renders configuration files from named YAML inputs. It builds one model by rendering an ordered list of model templates, then renders that model through a list of output templates.

Templates are Go `text/template` files with [Sprig](https://masterminds.github.io/sprig/) functions, plus a `required` function.

## Install

cooonfig is a single binary. Homebrew, Scoop, APT (Debian, Ubuntu, and Termux), RPM, GitHub release archives, and `go install` are in [INSTALL.md](INSTALL.md).

## Usage

```bash
cooonfig --help
cooonfig --version
cooonfig --config example/cooonfig.yaml --input-dir example --output-dir target
```

`--version` prints the version, commit, and build date injected by `scripts/build.sh` and GoReleaser.

| Flag           | Short | Default           | Meaning                    |
| -------------- | ----- | ----------------- | -------------------------- |
| `--config`     | `-c`  | `./cooonfig.yaml` | Config file                |
| `--input-dir`  | `-i`  | `./`              | Directory for input paths  |
| `--output-dir` | `-o`  | `./`              | Directory for output files |

Input paths in the config are relative to `--input-dir`. Model and output template paths are relative to the directory that contains the config file. Output paths are relative to `--output-dir`. Parent directories for output files are created.

A worked example is in [`example/`](example/).

## Config file

```yaml
inputs:
  - name: app
    path: application.yaml
  - name: env
    path: environment.yaml
models:
  - models/core.yaml
  - models/default.yaml
  - models/kubernetes.yaml
outputs:
  - template: templates/apache_config.json
    outfile: apache/config.json
  - template: templates/kafka_config.yml
    outfile: kafka/config.yaml
```

`inputs`, `models`, and `outputs` are all required.

Each input needs a `name` and a `path`. `type` defaults to `yaml`, which is the only supported type. Names must be unique. The name `model` is reserved.

Model templates run in order. Each one is rendered, parsed as YAML, and deep-merged into the model. Nested maps merge. A later value replaces an earlier value at the same key.

Output templates then render against the finished model. Each `outfile` is created or truncated.

## Templates

Model templates see each input under its name, and the model built so far under `.model`:

```yaml
app_name: { { .app.name } }
app_url: { { required ".app.source.url is required" .app.source.url } }
working_directory: "./workspace/{{ .model.app_name }}"
```

Output templates see the finished model as the root:

```yaml
Name: "{{ .app_name }}"
Host: "{{ .env_host }}"
```

`required` returns an error when its value is nil or an empty string. Sprig's `fail` and `default` work the same way as in any Sprig text template.

## Development

```bash
just ci
just build
```

`just build` writes binaries to `bin/` and snapshot packages to `dist/`. `just release` creates a GitHub release from `release.env`.

From a checkout:

```bash
go run ./cmd/cooonfig --config example/cooonfig.yaml --input-dir example --output-dir target
```
