# Command-Line Interface


It is my personal practice for installing some ntaive linux packages via a CLI tool.

### Environment Installation

```
go env -w GO111MODULE=on 
go mod init chieh_cli
go get github.com/urfave/cli
```

### Reference

- [Golang第三方命令行工具 - spf13/cobra和urfave/cli](https://strconv.com/posts/cli/)