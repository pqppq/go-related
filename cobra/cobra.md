# cobraについてのノート

Cobra: https://cobra.dev/
> A Framework for Modern CLI Apps in Go

cmd.Execute() -> init(), rootCmd.Execute() -> rootCmd.AddCommand(subCommand)

- entry point(main)
- 設定ファイルの初期化、ルートコマンドの定義
- サブコマンドの定義、サブコマンドのaction定義

## install
```
go get -u github.com/spf13/cobra/cobra
```

import
```
import "github.com/spf13/cobra"
```

## Concept
Cobraの構成要素
- Commands(actions)
- Args
- Flags

```
▾ app
  main.go // cmd.Execute()
  ▾ cmd/
    root.go // rootCmd &cobra.Command
    foo.go // command
    bar.go // command
```

root.go
- rootCmd(&cobra.Command)を定義する
- init()をフラグと設定の初期化のために定義できる
    - initConfigも合わせて書くことが多い

```
var rootCmd = &cobra.Command{
    Use: "command",
    Short: "short description",
    Long: "long description",
    Run: func(cmd *cobra.Command, args []string) {
        ...
    },
}

// entry point(main call this)
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        ...
        os.Exit(1)
    }
}

func init() {
    cobra.OnInitialize(initConfig)
    rootCmd.PersistenFlags().Bool("foo", true, "Use foo")
    ...
}

func initConfig() {
    ...
}
```
PreRunとPostRun Hooksがあった
https://cobra.dev/#prerun-and-postrun-hooks

cmdに他のCommandを定義して、rootCmdに追加していく
```
// foo.go
func init() {
    rootCmd.AddCommand(fooCommand)
}

var fooCmd = &cobra.Command{
    ...
}
```

設定ファイルの情報は、initConfigで初期化処理などをしたあとに各コマンドで値を読み取る

Usage
```
Usage:
  cobra-cli [command]

  Available Commands:
    add         Add a command to a Cobra Application
    completion  Generate the autocompletion script for the specified shell
    help        Help about any command
    init        Initialize a Cobra Application

    Flags:
     -a, --author string    author name for copyright attribution (default "YOUR NAME")
         --config string    config file (default is $HOME/.cobra.yaml)
     -h, --help             help for cobra-cli
     -l, --license string   name of license for the project
         --viper            use Viper for configuration
```

# 初期化
```
go mod init foobar

go get -u github.com/spf13/cobra@latest
go get github.com/spf13/viper

go install github.com/spf13/cobra-cli@latest
cobra-cli init
tree
>
.
├── cmd
│   └── root.go
├── go.mod
├── go.sum
├── LICENSE
└── main.go

cobra-cli add sub
>
.
├── cmd
│   ├── root.go
│   └── sub.go
├── go.mod
├── go.sum
├── LICENSE
└── main.go
```
