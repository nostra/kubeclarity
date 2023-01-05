#!/usr/bin/env bash

set -e
make ui backend
export DOCKER_TAG=$(date +%Y-%m-%d-%H-%M)
export DOCKER_REGISTRY=glowcr.azurecr.io
export BINARY_NAME=kubeclarity
export DOCKER_IMAGE=$DOCKER_REGISTRY/$BINARY_NAME

make docker-backend

docker image save $DOCKER_IMAGE:$DOCKER_TAG -o tmp.tar
nerdctl --namespace=k8s.io load --all-platforms -i tmp.tar
rm tmp.tar