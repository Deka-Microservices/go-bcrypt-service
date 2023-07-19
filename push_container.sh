#!/bin/bash

NAME=go-bcrypt-service
VERSION=$(cat version)
docker push localhost:5000/${NAME}:${VERSION}
docker push localhost:5000/${NAME}:latest
