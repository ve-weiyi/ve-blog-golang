#!/bin/bash

# 检查手写 Go 代码是否已格式化；生成代码与 runtime 产物不参与检查。

set -e

unformatted=$(find . -name "*.go" \
    -not -path "./.git/*" \
    -not -path "./runtime/*" \
    -not -path "*/internal/pb/*" \
    -not -name "*.pb.go" \
    -print0 | xargs -0 gofmt -l)

if [ -n "$unformatted" ]; then
    echo "以下文件未格式化，请运行 gofmt -w：" >&2
    echo "$unformatted" >&2
    exit 1
fi

echo "格式检查通过"
