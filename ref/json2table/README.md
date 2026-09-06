# json2table

print json data in tabular format

## Usage
```
NAME:
   json2table - print json data in tabular format

USAGE:
   json2table [global options] [dataFile]

VERSION:
   v1.1.0

GLOBAL OPTIONS:
   --build                      print build info and exit
   --spec string, -s string     read spec from specFile.json, or from environment variable JSON2TABLE_SPEC or JSON2TABLE_SPEC_FILE if not provided
   --columns string, -c string  Comma separated list of columns to print, ignore -s and JSON2TABLE_SPEC or JSON2TABLE_SPEC_FILE if provided
   --help, -h                   show help
   --version, -v                print the version
```
## Examples

### Cli
```
$ json2table -s ./samples/spec1.json ./samples/data1.json

$ json2table -c 'id,desc,url,display.name' ./samples/data2.json

```
### Scripts with Spec

- [Picsum list](./examples/picsum-list.md)

## Installation

See [Installation.md](Installation.md) for Homebrew, Scoop, Linux packages, Windows winget, and manual binary installs.

## Spec Reference

- [Spec file guide](./SPEC.md)

## Quality

- https://sonarcloud.io/project/overview?id=siakhooi_json2table
- https://qlty.sh/gh/siakhooi/projects/json2table

## Deliverables

- https://pkg.go.dev/github.com/siakhooi/json2table

## Reference

- https://github.com/savioxavier/termlink
- https://github.com/fatih/color

## Badges

![GitHub](https://img.shields.io/github/license/siakhooi/json2table?logo=github)
![GitHub last commit](https://img.shields.io/github/last-commit/siakhooi/json2table?logo=github)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/siakhooi/json2table?logo=github)
![GitHub issues](https://img.shields.io/github/issues/siakhooi/json2table?logo=github)
![GitHub closed issues](https://img.shields.io/github/issues-closed/siakhooi/json2table?logo=github)
![GitHub pull requests](https://img.shields.io/github/issues-pr-raw/siakhooi/json2table?logo=github)
![GitHub closed pull requests](https://img.shields.io/github/issues-pr-closed-raw/siakhooi/json2table?logo=github)
![GitHub top language](https://img.shields.io/github/languages/top/siakhooi/json2table?logo=github)
![GitHub language count](https://img.shields.io/github/languages/count/siakhooi/json2table?logo=github)
![GitHub repo size](https://img.shields.io/github/repo-size/siakhooi/json2table?logo=github)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/siakhooi/json2table?logo=github)
![Workflow](https://img.shields.io/badge/Workflow-github-purple)
![workflow](https://github.com/siakhooi/json2table/actions/workflows/build.yaml/badge.svg)
![workflow](https://github.com/siakhooi/json2table/actions/workflows/release.yaml/badge.svg)

![Release](https://img.shields.io/badge/Release-github-purple)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/siakhooi/json2table?label=GPR%20release&logo=github)
![GitHub all releases](https://img.shields.io/github/downloads/siakhooi/json2table/total?color=33cb56&logo=github)
![GitHub Release Date](https://img.shields.io/github/release-date/siakhooi/json2table?logo=github)

![Quality-Qlty](https://img.shields.io/badge/Quality-Qlty-purple)
[![Maintainability](https://qlty.sh/gh/siakhooi/projects/json2table/maintainability.svg)](https://qlty.sh/gh/siakhooi/projects/json2table)
[![Code Coverage](https://qlty.sh/gh/siakhooi/projects/json2table/coverage.svg)](https://qlty.sh/gh/siakhooi/projects/json2table)

![Quality-Sonar](https://img.shields.io/badge/Quality-SonarCloud-purple)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=siakhooi_json2table&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=siakhooi_json2table)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=siakhooi_json2table&metric=duplicated_lines_density)](https://sonarcloud.io/summary/new_code?id=siakhooi_json2table)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=siakhooi_json2table&metric=bugs)](https://sonarcloud.io/summary/new_code?id=siakhooi_json2table)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=siakhooi_json2table&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=siakhooi_json2table)
[![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=siakhooi_json2table&metric=sqale_index)](https://sonarcloud.io/summary/new_code?id=siakhooi_json2table)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=siakhooi_json2table&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=siakhooi_json2table)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=siakhooi_json2table&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=siakhooi_json2table)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=siakhooi_json2table&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=siakhooi_json2table)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=siakhooi_json2table&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=siakhooi_json2table)
[![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=siakhooi_json2table&metric=ncloc)](https://sonarcloud.io/summary/new_code?id=siakhooi_json2table)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=siakhooi_json2table&metric=coverage)](https://sonarcloud.io/summary/new_code?id=siakhooi_json2table)
![Sonar Violations (short format)](https://img.shields.io/sonar/violations/siakhooi_json2table?server=https%3A%2F%2Fsonarcloud.io)
![Sonar Violations (short format)](https://img.shields.io/sonar/blocker_violations/siakhooi_json2table?server=https%3A%2F%2Fsonarcloud.io)
![Sonar Violations (short format)](https://img.shields.io/sonar/critical_violations/siakhooi_json2table?server=https%3A%2F%2Fsonarcloud.io)
![Sonar Violations (short format)](https://img.shields.io/sonar/major_violations/siakhooi_json2table?server=https%3A%2F%2Fsonarcloud.io)
![Sonar Violations (short format)](https://img.shields.io/sonar/minor_violations/siakhooi_json2table?server=https%3A%2F%2Fsonarcloud.io)
![Sonar Violations (short format)](https://img.shields.io/sonar/info_violations/siakhooi_json2table?server=https%3A%2F%2Fsonarcloud.io)
![Sonar Violations (long format)](https://img.shields.io/sonar/violations/siakhooi_json2table?format=long&server=http%3A%2F%2Fsonarcloud.io)

[![Wise](https://img.shields.io/badge/Funding-Wise-33cb56.svg?logo=wise)](https://wise.com/pay/me/siakn3)
![visitors](https://hit-tztugwlsja-uc.a.run.app/?outputtype=badge&counter=ghmd-json2table)
