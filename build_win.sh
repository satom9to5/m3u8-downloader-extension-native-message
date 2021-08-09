#!/bin/sh

CURDIR=$(cd $(dirname $0) && pwd)

mkdir -p "${CURDIR}/bin"

GOOS=windows GOARCH=amd64 go build -o "${CURDIR}/bin/m3u8-downloader-extension-native-message.exe"
