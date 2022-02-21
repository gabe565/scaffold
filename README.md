# Clevyr Scaffold

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
