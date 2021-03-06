#!/bin/bash -e

CURRENT_DIR=$(pwd)
GEN_DIR=$(dirname $0)
REPO_DIR="$CURRENT_DIR/$GEN_DIR/../.."

PROJECT_MODULE="github.com/flacatus/operator-tests"
IMAGE_NAME="kubernetes-codegen:latest"

CUSTOM_RESOURCE_NAME="foo"
CUSTOM_RESOURCE_VERSION="v1"

echo "Building codegen Docker image..."
docker build -f "${GEN_DIR}/Dockerfilegen" \
             -t "${IMAGE_NAME}" \
             "${REPO_DIR}"

cmd="./generate-groups.sh deepcopy,client \
    "$PROJECT_MODULE/pkg/client" \
    github.com/eclipse/che-operator/pkg/apis \
    $CUSTOM_RESOURCE_NAME:$CUSTOM_RESOURCE_VERSION"

echo "Generating client codes..."
docker run --rm \
           -v "${REPO_DIR}:/go/src/${PROJECT_MODULE}" \
           "${IMAGE_NAME}" $cmd

sudo chown $USER:$USER -R $REPO_DIR/pkg
