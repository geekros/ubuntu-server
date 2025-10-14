#!/bin/bash
# Copyright 2025 GEEKROS, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -e

lsb_release -a

architecture=$(dpkg --print-architecture)

ubuntu_code=$(lsb_release -c -s)

arch=$(uname -m)

on_init(){

    sudo apt update -y

    sudo apt install -y vim dpkg-dev gpg curl wget git gcc make cmake gcc-arm-none-eabi lsb-release libudev-dev libusb-1.0-0-dev openssl portaudio19-dev

    if [ ! -d "/usr/local/go/bin/" ]; then
        golang_version="1.24.5"
        sudo wget -q https://golang.google.cn/dl/go"${golang_version}".linux-"${architecture}".tar.gz && sudo tar -C /usr/local -xzf go"${golang_version}".linux-"${architecture}".tar.gz
        touch /etc/profile.d/geekros-golang.sh
        sudo sh -c 'echo "export PATH=$PATH:/usr/local/go/bin" >> /etc/profile.d/geekros-golang.sh'
        source /etc/profile.d/geekros-golang.sh
        sudo rm -rf go"${golang_version}".linux-"${architecture}".tar.gz
    fi

    exit 0
}

on_update(){
    
    # Clean up previous installation
    sudo cp -r ubuntu/opt/* /opt/
    find /opt -type f -name ".gitkeep" -exec rm -f {} +

    # Stop and disable the service if it exists
    sudo systemctl stop geekserver.service > /dev/null 2>&1
    sudo systemctl disable geekserver.service > /dev/null 2>&1

    # Build the main program
    /usr/local/go/bin/go env -w GOSUMDB=off
    /usr/local/go/bin/go env -w GOPATH=/tmp/golang
    /usr/local/go/bin/go env -w GOMODCACHE=/tmp/golang/pkg/mod
    /usr/local/go/bin/go env -w GOTOOLCHAIN=local
    export GO111MODULE=on && export GOPROXY=https://goproxy.io
    cd ../cmd && /usr/local/go/bin/go mod tidy && /usr/local/go/bin/go build -o ../release/geekserver main.go
    sudo cp ../release/geekserver /opt/geekros/release/
    sudo cp ../release/config.sample.yaml /opt/geekros/release/config.yaml
    sudo rm -rf ../release/geekserver && cd ../.tools

    sudo systemctl enable geekserver.service > /dev/null 2>&1
    sudo systemctl restart geekserver.service > /dev/null 2>&1
}

case "$1" in
    init)
        on_init
        ;;
    update)
        on_update
        ;;
    *)
        echo "Error: Unknown command '$1'"
        exit 1
        ;;
esac