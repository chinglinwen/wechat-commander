#!/bin/sh
# build image
set -e

suffix="$1"
suffix=${suffix:=v1}

go build

image="wechat-commander:$suffix"
echo -e "building image: $image\n"
tag="harbor.haodai.net/ops/$image"
docker build --no-cache  -t $tag .
docker push $tag
