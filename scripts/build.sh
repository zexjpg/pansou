#!/usr/bin/env bash
#
# PanSou 跨平台构建脚本
# ----------------------------------------------------------------------------
# 由 .github/workflows/build-release.yml 的矩阵调用（每个 GOOS/GOARCH 一次），
# 也可在本地手动运行，例如：
#     GOOS=linux   GOARCH=amd64 bash scripts/build.sh
#     GOOS=darwin  GOARCH=arm64 bash scripts/build.sh
#     GOOS=windows GOARCH=amd64 bash scripts/build.sh
#
# 关键约束：
#   * bytedance/sonic v1.14.0 最高支持到 Go 1.25，必须用 GOTOOLCHAIN 锁 1.25.0
#   * CGO_ENABLED=0 做纯 Go 静态编译，才能在一个 Linux runner 上交叉编译出
#     Windows / macOS 的可执行文件（无需交叉编译器、不依赖目标平台 libc）
# ----------------------------------------------------------------------------
set -euo pipefail

GOOS="${GOOS:?GOOS is required (windows|linux|darwin)}"
GOARCH="${GOARCH:?GOARCH is required (amd64|arm64)}"

# 锁定工具链 & 静态编译
export GOTOOLCHAIN="${GOTOOLCHAIN:-go1.25.0}"
export CGO_ENABLED=0

OUT_DIR="dist"
mkdir -p "$OUT_DIR"

BIN="pansou-${GOOS}-${GOARCH}"
if [ "$GOOS" = "windows" ]; then
  BIN="${BIN}.exe"
fi
OUT="${OUT_DIR}/${BIN}"

echo ">> building ${GOOS}/${GOARCH} (GOTOOLCHAIN=${GOTOOLCHAIN}, CGO_ENABLED=${CGO_ENABLED}) -> ${OUT}"
go build -trimpath -ldflags "-s -w" -o "${OUT}" .

echo ">> done:"
ls -lh "${OUT}"
