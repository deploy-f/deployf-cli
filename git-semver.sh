#!/bin/bash

DESCRIBE=`git describe --tags --always`

VERSION=`echo $DESCRIBE | awk '{split($0,a,"-"); print a[1]}'`
BUILD=`echo $DESCRIBE | awk '{split($0,a,"-"); print a[2]}'`
PATCH=`echo $DESCRIBE | awk '{split($0,a,"-"); print a[3]}'`

if [[ "${DESCRIBE}" =~ ^[A-Fa-f0-9]+$ ]]; then
    VERSION="0.0.0"
    BUILD=`git rev-list HEAD --count`
    PATCH=${DESCRIBE}
fi

if [[ "${BUILD}" == "" ]]; then
    BUILD=0
fi

echo ${VERSION}.${BUILD}