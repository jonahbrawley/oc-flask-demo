#!/usr/bin/env bash
REGISTRY=quay.io/rh-ee-jbrawley/goresponder
GIT_HASH=$(shell git rev-parse --short=7 HEAD)

build:
	podman build --tag ${REGISTRY}:${GIT_HASH} .

push:
	podman push ${REGISTRY}:${GIT_HASH}
	echo Make sure to append the GIT_HASH to yaml!
	#echo $GIT_HASH
