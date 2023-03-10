# Clevyr Scaffold

[![Build](https://github.com/clevyr/scaffold/actions/workflows/build.yml/badge.svg)](https://github.com/clevyr/scaffold/actions/workflows/build.yml)

Scaffolds out our normal stack. Includes default config, Docker and Docker Compose stack, etc.

When you run the scaffold, you will be asked some questions to determine the setup the templates, then config will be
generated and output. Your responses will be saved to `.clevyr-scaffold-config` so that they can be loaded as default
values if the scaffold is run in that repository in the future.

## Running Locally

To install the command locally on your host, make sure you have [Homebrew](https://brew.sh/) first.

### Install

```sh
brew install clevyr/tap/scaffold
```

### Usage

To scaffold a Laravel application, `cd` to the desired parent directory then run:

```sh
scaffold
```

### Upgrade

```sh
brew update && brew upgrade clevyr/tap/scaffold
```

## Running in Docker

The Docker repo is located at [clevyr/scaffold](https://hub.docker.com/r/clevyr/scaffold). Docker will pull the image
the first time you run the command below.

### Usage

To scaffold a Laravel application, `cd` to the desired parent directory then run:

```sh
docker run --rm -it -v "$PWD:/data" ghcr.io/clevyr/scaffold
```

### Upgrade

```sh
docker pull ghcr.io/clevyr/scaffold
```

## Scaffold Development

To make code changes to this repository, be sure you have `go` installed, then run the following.

```sh
mkdir -p ~/go/src/github.com/clevyr
git clone https://github.com/clevyr/scaffold ~/go/src/github.com/clevyr/scaffold
```

Then go ahead and make your changes in `~/go/src/github.com/clevyr/scaffold`.

### Testing Code Changes

When the scaffold is run, it creates many files. To keep these from being committed, the `out/` directory has been
added to `.gitignore`. To test local changes, run the following command.

```sh
go run . -C out
```

Afterwards, you can look in the `out/` directory to see the generated app.

### Template Development

By default, Go's [`embed`](https://pkg.go.dev/embed) module ignores hidden files. To be able to use hidden file templates
like `.env`, a generator has been created that forces these files to be embedded. If you add any new hidden files to the
`templates/` directory, be sure to regenerate this file.

```sh
go generate
```
