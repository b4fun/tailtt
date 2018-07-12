#!/bin/bash

set -o pipefail
set -o nounset
set -o errexit

BEARYCHAT_RTM_TOKEN=${BEARYCHAT_RTM_TOKEN:-""}
BEARYCHAT_RTM_CHANNEL=${BEARYCHAT_RTM_CHANNEL:-""}

function build() {
    local ldflags=""

    if [[ "$BEARYCHAT_RTM_TOKEN" ]]
    then
        ldflags+=" -X main.EnvNotifyBearychatRTMToken=${BEARYCHAT_RTM_TOKEN}"
        echo "building with main.EnvNotifyBearychatRTMToken=${BEARYCHAT_RTM_TOKEN}"
    fi
    if [[ "$BEARYCHAT_RTM_CHANNEL" ]]
    then
        ldflags+=" -X main.EnvNotifyBearychatRTMChannel=${BEARYCHAT_RTM_CHANNEL}"
        echo "building with main.EnvNotifyBearychatRTMChannel=${BEARYCHAT_RTM_CHANNEL}"
    fi

    go build \
        -ldflags "$ldflags" \
        -o "./bin/tailtt" \
        "github.com/b4fun/tailtt/cmd/..."
}

build
