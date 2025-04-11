#!/bin/bash

# The 'set -e' command causes the script to immediately exit
# if any command returns a non-zero exit status (i.e., an error).
set -e

ROOT_DIR="$(pwd)"
PACKAGE_DIR="${ROOT_DIR}/packages"
GEN_DIR="${ROOT_DIR}/www/grpc/gen"
VERSION="$(echo `git -C ${ROOT_DIR} describe --abbrev=0 --tags` | sed 's/^.//')" # "v1.2.3" -> "1.2.3"

rm -rf ${PACKAGE_DIR}
mkdir -p ${PACKAGE_DIR}

echo "== Building pactus-grpc package for Python"
cp -R ${ROOT_DIR}/.github/packager/python ${PACKAGE_DIR}/python
cp ${GEN_DIR}/python/* ${PACKAGE_DIR}/python/pactus_grpc/
sed -i "s/{{ VERSION }}/$VERSION/g" ${PACKAGE_DIR}/python/setup.py
