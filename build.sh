#!/usr/bin/env sh
set -ex

export GOARCH=amd64
for GOOS in windows linux darwin ; do
    suffix="-$GOOS"
    if [ "$GOOS" = "windows" ]; then suffix=".exe"; fi
    go build -a -ldflags '-extldflags "-static"' -o "randomclass${suffix}"
done
