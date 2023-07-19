#!/bin/bash

NAME=go-bcrypt-service
VERSION=$(cat version)
docker build . --build-arg VERSION=${VERSION} --tag ${NAME}:${VERSION}
docker tag ${NAME}:${VERSION} ${NAME}:latest
docker tag ${NAME}:${VERSION} localhost:5000/${NAME}:${VERSION}
docker tag ${NAME}:latest localhost:5000/${NAME}:latest
