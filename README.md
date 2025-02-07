# PageSnatcher CLI tool

## Description

The `pagesnatcher-cli` tool is a command-line interface (CLI) tool designed to help users retrieve web pages and save them locally. It provides a convenient way to download and store web content for offline access (e.g. personal archiving).

## Profiles
Use `--profile` or `-p` to use a templated configuration for `framer` or `webflow` exports.

## Examples
```sh
pagesnatcher-cli clone https://jomor-design-2019.webflow.io/ -p webflow

pagesnatcher-cli clone https://demoportfolio.framer.website/ -o /tmp/output -z zippedFolder.zip -p framer
```

## Installation
Install with go (`go install github.com/trostatik/pagesnatcher-cli@v0.0.1`), or download the latest [release](https://github.com/trostatik/pagesnatcher-cli/releases/) as an executable

### Dependencies

`pagesnatcher` requires `wget`, `find`, and `sed`.

## Version history

v0.0.1 Initial release

## Disclaimer

Please note that the use of this tool is subject to the laws and terms of service applicable to users. It is the responsibility of the user to ensure compliance with all relevant legal requirements and to respect the terms of service of the websites being accessed. The tool should be used in a responsible and ethical manner, respecting the rights and privacy of others.
