# SPDX-FileCopyrightText: Magenta ApS
# SPDX-License-Identifier: MPL-2.0
terraform {
  required_providers {
    os2mo = {
      source = "github.com/Skeen/os2mo"
    }
  }
}

provider "os2mo" {
  url = "http://localhost:5000/graphql"
}

data "os2mo_itsystems" "all" {}

output "itoutputs" {
  value = data.os2mo_itsystems.all #.results
}

data "os2mo_itsystem" "SAP_by_userkey" {
  user_key = "SAP"
}

data "os2mo_itsystem" "SAP_by_uuid" {
  uuid = "49d91308-67b0-4b8c-b787-1cd58e3039bd"
}

/*
data "os2mo_itsystem" "NOMATCH" {
  uuid = "801ee257-57f1-4b66-ad34-cc9dae1eb343"
}

data "os2mo_itsystem" "SAP_by_both" {
  user_key = "SAP"
  uuid     = "49d91308-67b0-4b8c-b787-1cd58e3039bd"
}

data "os2mo_itsystem" "SAP_by_neither" {}
*/

output "itoutput_by_userkey" {
  value = data.os2mo_itsystem.SAP_by_userkey
}

output "itoutput_by_uuid" {
  value = data.os2mo_itsystem.SAP_by_uuid
}

resource "os2mo_organisation" "root" {
  name = "Wowzers2"
}
