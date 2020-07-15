# Clevyr Scaffold

Scaffolds out our normal stack. Includes default config, Docker and Docker Compose stack, etc.

## Installing

The simplest install method is with [Homebrew](https://brew.sh/). Simply run:
```sh
brew install clevyr/tap/scaffold
```

Then, to scaffold a Laravel application, `cd` into its root and run:
```sh
scaffold
```

You will be asked some questions to determine the setup the templates, then config will be generated and output. Your responses will be saved to `.clevyr-scaffold-config` so that they can be loaded as default values if the scaffold is run in the future.

## Updating

To update to the latest version of the stack, run:
```sh
brew update
brew upgrade clevyr/tap/scaffold
```
