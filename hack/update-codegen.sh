#!/bin/sh

set -o errexit
set -o nounset
set -o pipefail

# Always set working directory to 'hack'
cd $(dirname "${BASH_SOURCE[0]}")

CODEGEN_DIR=../../code-generator
bash $CODEGEN_DIR/generate-groups.sh \
  "deepcopy,client,informer,lister" \
  github.com/krafton-hq/redfox/pkg/generated \
  github.com/krafton-hq/redfox/pkg/apis \
  redfox:v1alpha1 \
  --output-base ../.tmp/ \
  --go-header-file ./custom-boilerplate.go.txt

set -ex
cp -R ../.tmp/github.com/krafton-hq/redfox/pkg/ ../pkg/
rm -r ../.tmp/
