#!/bin/bash

set -e

ROOT_DIR="$(pwd)"
BUILD_DIR="${ROOT_DIR}/build"

# Check the architecture
ARC="$(uname -m)"
echo "Building the binaries for Windows ${ARC} architecture"

# This fixes a bug in pkgconfig: invalid flag in pkg-config --libs: -Wl,-luuid
sed -i -e 's/-Wl,-luuid/-luuid/g' /mingw64/lib/pkgconfig/gdk-3.0.pc

CGO_ENABLED=0 go build -ldflags "-s -w" -trimpath -o ${BUILD_DIR}/pactus-daemon_unsigned.exe        ./cmd/daemon
CGO_ENABLED=0 go build -ldflags "-s -w" -trimpath -o ${BUILD_DIR}/pactus-wallet_unsigned.exe        ./cmd/wallet
CGO_ENABLED=0 go build -ldflags "-s -w" -trimpath -o ${BUILD_DIR}/pactus-shell_unsigned.exe         ./cmd/shell
go build -ldflags "-s -w -H windowsgui" -trimpath -tags gtk -o ${BUILD_DIR}/pactus-gui_unsigned.exe ./cmd/gtk
