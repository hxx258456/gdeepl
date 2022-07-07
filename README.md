# gdeepl deepl翻译命令行工具

[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/955/badge)](https://bestpractices.coreinfrastructure.org/projects/955) [![Go Report Card](https://goreportcard.com/badge/github.com/hyperledger/fabric)](https://goreportcard.com/report/github.com/hyperledger/fabric)

# 帮助信息
```shell
gdeepl translate's command line tools

Usage:
  gdeepl [flags]
  gdeepl [command]

Examples:
gdeepl -T "hello"

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  version     version

Flags:
  -h, --help                 help for gdeepl
  -s, --source_lang string   源语言 (default "auto")
  -t, --target_lang string   目标语言 (default "ZH")
  -T, --text string          翻译内容 (default "hello world!")

Use "gdeepl [command] --help" for more information about a command.
```

# 示例
```shell
gdeepl -h // 获取帮助信息
gdeepl version // 获取版本信息
gdeepl -T "hello world!"
```

# 安装
需要golang环境
```shell
go install github.com/hxx258456/gdeepl@latest
```

