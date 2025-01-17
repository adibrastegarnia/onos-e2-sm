#!/bin/sh

# SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: Apache-2.0

set -e

echo "** generate python bindings"
proto_imports=".:${GOPATH}/src/github.com/gogo/protobuf/protobuf:${GOPATH}/src/github.com/gogo/protobuf:${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate:${GOPATH}/src"

OUTPUTPATH="python/onos_e2_sm"
rm -rf "$OUTPUTPATH"
mkdir -p "$OUTPUTPATH"

cd servicemodels

# betterproto client bindings into $OUTPUTPATH
protoc "-I=$proto_imports" \
    "--python_betterproto_out=../$OUTPUTPATH" \
    e2sm_ni/v1beta1/e2sm_ni_ies.proto \
    validate/v1/validate.proto \
    e2sm_kpm_v2/v2/e2sm_kpm_v2.proto \
    e2sm_rc_pre/v2/e2sm_rc_pre_v2.proto \
    e2sm_mho/v1/e2sm_mho.proto \
    e2sm_kpm/v1beta1/e2sm_kpm_ies.proto
