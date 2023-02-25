# deepl-cli

[![Test](https://github.com/candy12t/deepl-cli/actions/workflows/test.yml/badge.svg)](https://github.com/candy12t/deepl-cli/actions/workflows/test.yml)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/candy12t/deepl-cli)
![GitHub](https://img.shields.io/github/license/candy12t/deepl-cli)
[![Go Report Card](https://goreportcard.com/badge/github.com/candy12t/deepl-cli)](https://goreportcard.com/report/github.com/candy12t/deepl-cli)

## Installation

go install:

```shell
go install github.com/candy12t/deepl-cli/cmd/deepl-cli@v0.4.2
```

## setup

```shell
deepl-cli configure
```

See [here](https://www.deepl.com/ja/docs-api/translating-text/) for `source` language and `target` language values.

## Usage

```shell
NAME:
   deepl-cli - unofficial DeepL command line tool.

USAGE:
   deepl-cli [global options] command [command options] [arguments...]

VERSION:
   v0.4.2

COMMANDS:
   repl       Translate with REPL.
   configure  Configure deepl-cli options.

GLOBAL OPTIONS:
   --help, -h     Show help for command (default: false)
   --version, -v  show deepl-cli version (default: false)
```

## LICENSE

MIT License
