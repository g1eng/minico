# About

Minico is a CLI tool to give minimal access for containers.

# Supported runtime

* Docker

# Requirements

* dockerd with modestly new version (20.0+ reccomended)
* golang >= 1.15
* MacOS or Linux recommended. Windows are not tested at now.

# Install

`go install github.com/g1eng/minico`

# Usage

run `minico help`.

```shell
$ minico help
NAME:
   minico - cli syntax sugar for container host operators

USAGE:
   minico [global options] command [command options] [command on container]

COMMANDS:
   run      run a container with the specified image
   list     list and filter containers with fuzzy name matching
   stop     stop containers with fuzzy name matching
   images   list and filter images with fuzzy name matching
   version  show version
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

# Design

WIP
