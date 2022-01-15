# deepl-cli

[![Test](https://github.com/candy12t/deepl-cli/actions/workflows/test.yml/badge.svg)](https://github.com/candy12t/deepl-cli/actions/workflows/test.yml)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/candy12t/deepl-cli)
![GitHub](https://img.shields.io/github/license/candy12t/deepl-cli)
[![Go Report Card](https://goreportcard.com/badge/github.com/candy12t/deepl-cli)](https://goreportcard.com/report/github.com/candy12t/deepl-cli)

## Installation

go install:

```shell
go install github.com/candy12t/deepl-cli/cmd/deepl-cli@v0.3.0
```

## setup

```shell
deepl-cli setup
```

See [here](https://www.deepl.com/ja/docs-api/translating-text/) for `source` language and `target` language values.

## Usage

```shell
USAGE:
	deepl-cli [global options] command [command options] [arguments...]

COMMANDS:
  setup  Setup for using this cli
  repl   Translate with REPL

GLOBAL OPTIONS:
	--help, -h     Show help for command (default: false)
	--version, -v  show deepl-cli version (default: false)
```

## LICENSE

MIT License
