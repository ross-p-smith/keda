#-------------------------------------------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See https://go.microsoft.com/fwlink/?linkid=2090316 for license information.
#-------------------------------------------------------------------------------------------------------------

FROM golang:1.20.4

# Avoid warnings by switching to noninteractive
ENV DEBIAN_FRONTEND=noninteractive

ENV GO111MODULE=auto

# Configure apt, install packages and tools
RUN apt-get update \
    && apt-get -y install --no-install-recommends apt-utils dialog unzip 2>&1 \
    #
    # Verify git, process tools, lsb-release (common in install instructions for CLIs) installed
    && apt-get -y install git iproute2 procps lsb-release \
    #
    # Install gocode-gomod
    && go get -x -d github.com/stamblerre/gocode 2>&1 \
    && go build -o gocode-gomod github.com/stamblerre/gocode \
    && mv gocode-gomod $GOPATH/bin/ \
    #
    # Install Go tools
    && go get -u -v \
        github.com/mdempsky/gocode \
        github.com/uudashr/gopkgs/cmd/gopkgs \
        github.com/ramya-rao-a/go-outline \
        github.com/acroca/go-symbols \
        github.com/godoctor/godoctor \
        golang.org/x/tools/cmd/guru \
        golang.org/x/tools/cmd/gorename \
        github.com/rogpeppe/godef \
        github.com/zmb3/gogetdoc \
        github.com/haya14busa/goplay/cmd/goplay \
        github.com/sqs/goreturns \
        github.com/josharian/impl \
        github.com/davidrjenni/reftools/cmd/fillstruct \
        github.com/fatih/gomodifytags \
        github.com/cweill/gotests/... \
        golang.org/x/tools/cmd/goimports \
        golang.org/x/lint/golint \
        github.com/alecthomas/gometalinter 2>&1 \
        github.com/mgechev/revive \
        github.com/derekparker/delve/cmd/dlv 2>&1 \
    && go install honnef.co/go/tools/cmd/staticcheck@latest \
    && go install golang.org/x/tools/gopls@latest \
    # Protocol Buffer Compiler
    && PROTOC_VERSION=21.9 \
    && if [ $(dpkg --print-architecture) = "amd64" ]; then PROTOC_ARCH="x86_64"; else PROTOC_ARCH="aarch_64" ; fi \
    && curl -LO "https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-$PROTOC_ARCH.zip" \
    && unzip "protoc-${PROTOC_VERSION}-linux-$PROTOC_ARCH.zip" -d $HOME/.local \
    && mv $HOME/.local/bin/protoc /usr/local/bin/protoc \
    && mv $HOME/.local/include/ /usr/local/bin/include/ \
    && protoc --version \
    # Install golangci-lint
    && curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.52.2 \
    #
    # Install pip & pre-commit
    && apt-get -y install python3-pip \
    && python3 -m pip install --no-cache-dir pre-commit \
    #
    # Clean up
    && apt-get autoremove -y \
    && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/*

# Enable go modules
ENV GO111MODULE=on

ENV OPERATOR_RELEASE_VERSION=v1.26.0
RUN ARCH=$(case $(uname -m) in x86_64) echo -n amd64 ;; aarch64) echo -n arm64 ;; *) echo -n $(uname -m) ;; esac) \
    && OS=$(uname | awk '{print tolower($0)}') \
    && OPERATOR_SDK_DL_URL=https://github.com/operator-framework/operator-sdk/releases/download/${OPERATOR_RELEASE_VERSION} \
    && curl -LO ${OPERATOR_SDK_DL_URL}/operator-sdk_${OS}_${ARCH} \
    && chmod +x operator-sdk_${OS}_${ARCH} \
    && mkdir -p /usr/local/bin/ \
    && cp operator-sdk_${OS}_${ARCH} /usr/local/bin/operator-sdk \
    && rm operator-sdk_${OS}_${ARCH}
