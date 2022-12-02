#!/bin/bash
set -e

cd ..
make install

cd examples
rm -rf .terraform.lock.hcl
terraform init
env TF_LOG=DEBUG terraform apply
