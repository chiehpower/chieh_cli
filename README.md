# Command-Line Interface


It is my personal practice for installing some native linux packages via a CLI tool.

## ▎Usage

```
NAME:
   chieh-cli - install some basic packages for Linux Ubuntu OS.

USAGE:
   chieh-cli [global options] command [command options] [arguments...]

VERSION:
   0.1.0

DESCRIPTION:
   It is a Chieh's personal CLI tool for setting environment on Linux.

COMMANDS:
   help, h  Shows a list of commands or help for one command

   install:
     install, i, I, Install  To install some packages regarding to native Linux tools.

   list:
     list, ls  List the files in the present directory.

   pull:
     pull, p, Pull  Pull the docker image. Notice: please install Docker first.

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```


Basically we will use the subcommand, **install** and **pull**.

#### Install basic packages part

```
./chieh-cli install all
```

It will install git, docker, docker-compose, and nvidia gpu driver, etc.

## ▎Dev log

- [x] 2022/05/09 - Add the docker pull images with the specific uersname and password.
- [x] 2022/05/08 - Achieved to apply the subcommand and flag.  Also, it is able to check whether the tool exists or not. If yes, then it will not install it again.

## ▎Environment Installation

```
go env -w GO111MODULE=on 
go mod init chieh_cli
go get github.com/urfave/cli
```

## ▎Reference

- [Golang第三方命令行工具 - spf13/cobra和urfave/cli](https://strconv.com/posts/cli/)