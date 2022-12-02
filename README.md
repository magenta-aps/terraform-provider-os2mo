<!-- SPDX-FileCopyrightText: Magenta ApS -->
<!-- SPDX-License-Identifier: MPL-2.0 -->
# OS2mo Provider

[![registry.terraform.io](https://img.shields.io/badge/terraform-docs-success)](https://registry.terraform.io/providers/magenta-aps/os2mo/latest/docs)

Lifecycle management of OS2mo entities.
Maintained by Magenta ApS.

This provider enables management of OS2mo entities.

## Installation

The provider can be installed and managed automatically by Terraform. Sample `versions.tf` file:

```hcl
terraform {
  required_version = ">= 0.13"

  required_providers {
    kubectl = {
      source  = "magenta-aps/os2mo"
      version = ">= 0.0.1"
    }
  }
}
```

## Quick Start

```hcl
# Configure the Docker Hub Provider
provider "os2mo" {
  url = "http://localhost:5000/graphql"

  client_id = "terraform"
  client_secret = "3a07395c-c7cb-4529-9cb6-ff353d403229"
}

# Read out all ITSystems
data "os2mo_itsystems" "all" {}

# Read out the SAP ITSystem by user_key
data "os2mo_itsystem" "SAP_by_userkey" {
  user_key = "SAP"
}

# Read out the SAP ITSystem by uuid
data "os2mo_itsystem" "SAP_by_uuid" {
  uuid = "49d91308-67b0-4b8c-b787-1cd58e3039bd"
}
```

## Development Guide

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.12+ is *required*).
You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make terraform-provider-os2mo`. This will build the provider and put the provider binary in the local directory.

### Testing

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

### Installing

To install the provider locally (for instance for manually testing) run

```sh
$ make install
```

This installs the provider to `~/.terraform.d/`, removes `.terraform.lock.hcl` and runs `terraform init` in the local folder.
