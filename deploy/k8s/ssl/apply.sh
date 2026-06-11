#!/bin/bash
# 将域名证书导入 Kubernetes 集群
# 证书文件位于 deploy/ssl/（已加入 .gitignore，不提交到 git）
# 使用方式：在 sparkinai-cloud 根目录执行 bash deploy/k8s/ssl/apply.sh

set -e

NAMESPACE="blog"
SSL_DIR="deploy/ssl"

echo ">>> 创建 namespace（如已存在则跳过）"
kubectl create namespace $NAMESPACE --dry-run=client -o yaml | kubectl apply -f -

echo ">>> 导入 veweiyi.cn TLS 证书"
kubectl create secret tls blog-tls \
  --cert=$SSL_DIR/veweiyi.cn_nginx/veweiyi.cn_bundle.crt \
  --key=$SSL_DIR/veweiyi.cn_nginx/veweiyi.cn.key \
  -n $NAMESPACE --dry-run=client -o yaml | kubectl apply -f -

echo ">>> 导入 app.veweiyi.cn TLS 证书"
kubectl create secret tls blog-app-tls \
  --cert=$SSL_DIR/app.veweiyi.cn_nginx/app.veweiyi.cn_bundle.crt \
  --key=$SSL_DIR/app.veweiyi.cn_nginx/app.veweiyi.cn.key \
  -n $NAMESPACE --dry-run=client -o yaml | kubectl apply -f -

echo ">>> 导入 admin.veweiyi.cn TLS 证书"
kubectl create secret tls blog-admin-tls \
  --cert=$SSL_DIR/admin.veweiyi.cn_nginx/admin.veweiyi.cn_bundle.crt \
  --key=$SSL_DIR/admin.veweiyi.cn_nginx/admin.veweiyi.cn.key \
  -n $NAMESPACE --dry-run=client -o yaml | kubectl apply -f -

echo ">>> 证书导入完成"
