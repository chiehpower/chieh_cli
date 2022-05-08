# Command-Line Interface


It is my personal practice for installing some ntaive linux packages via a CLI tool.

## ▎Dev log


- [x] 2022/05/08 - Achieved to apply the subcommand and flag.  Also, it is able to check whether the tool exists or not. If yes, then it will not install it again.

## ▎Environment Installation

```
go env -w GO111MODULE=on 
go mod init chieh_cli
go get github.com/urfave/cli
```

## ▎Reference

- [Golang第三方命令行工具 - spf13/cobra和urfave/cli](https://strconv.com/posts/cli/)