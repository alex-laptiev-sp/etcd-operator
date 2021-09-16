#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

vendor/k8s.io/code-generator/generate-groups.sh \
  "all" \
  "github.com/alex-laptiev-sp/etcd-operator/pkg/generated" \
  "github.com/alex-laptiev-sp/etcd-operator/pkg/apis" \
  "etcd:v1beta2" \
  --go-header-file "./hack/k8s/codegen/boilerplate.go.txt" \
  $@
