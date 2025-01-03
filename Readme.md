# Token Parser  
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/bobmaertz/token-parser)](https://goreportcard.com/report/github.com/bobmaertz/token-parser)
[![Build Status](https://github.com/bobmaertz/token-parser/actions/workflows/go.yml/badge.svg)](https://github.com/bobmaertz/token-parser/actions)

A command line tool for token validation and inspection. Token Parser provides a convienant way to export the token values into a easy to parse JSON format. 

🚨 🏗️ **This project is a work in progress** 🏗️ 🚨

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Development](#development)
  - [Prerequisites](#prerequisites)
  - [Building](#building)
  - [Testing](#testing)
- [License](#license)

## Features

- Support for token inspection

## Installation

```bash
go get -u github.com/bobmaertz/token-parser
```

## Usage

```bash
token-parser [options]

Usage of token-parser:
Command Usage:
  version       Print the version of the CLI
  inspect       Inspect a token

Flag Usage:
  -t string
        the type of token being inspected (default "jwt")
  -v string
        the verifier type of token being inspected (default "none")
```

## Development

### Prerequisites

- Go 1.22 or higher

### Building

```bash
git clone https://github.com/bobmaertz/token-parser.git
cd token-parser
make build
```

### Testing

```bash
make test
```

### Running 

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.jjk
