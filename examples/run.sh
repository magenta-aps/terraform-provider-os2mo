#!/bin/bash
# SPDX-FileCopyrightText: Magenta ApS
# SPDX-License-Identifier: MPL-2.0
set -e

cd ..
make install

cd examples
rm -rf .terraform.lock.hcl
terraform init
env TF_LOG=DEBUG terraform apply
