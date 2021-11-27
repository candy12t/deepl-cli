# deepl-cli

## Installation

go install:

```shell
go install github.com/candy12t/deepl-cli/cmd/deepl-cli@latest
```

## setup
create config file

```shell
mkdir ~/.config/deepl-cli
touch config.yaml
```

sample

```yaml
account:
  auth_key: "xxxxxxxx"
  account_plan: "free"
default_lang:
  source_lang: "EN"
  target_lang: "JA"
```

See [here](https://www.deepl.com/ja/docs-api/translating-text/) for `source_lang` and `target_lang` values.

## Usage

```shell
USAGE:
	deepl-cli [global options] command [command options] [arguments...]

COMMANDS:
	repl  repl

GLOBAL OPTIONS:
	--help, -h     Show help for command (default: false)
	--version, -v  show deepl-cli version (default: false)
```

## LICENSE

MIT License
