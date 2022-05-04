#!/bin/bash
REGISTRY="docker.io"
USERNAME="gurken2108"
PROJECT="traefik-proxy-manager"

docker build -t ${USERNAME}/${PROJECT}:latest -f docker/Dockerfile .
docker tag ${USERNAME}/${PROJECT}:latest reg.mc8051.de/${USERNAME}/${PROJECT}:latest
#docker push ${REGISTRY}/${USERNAME}/${PROJECT}

