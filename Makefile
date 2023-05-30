#!/usr/bin/env bash

#RUN REGISTRY="$(oc get route/default-route -n openshift-image-registry -o=jsonpath='{.spec.host}')/openshift"
#RUN podman login --tls-verify=false -u unused -p $(oc whoami -t) ${REGISTRY}
#RUN podman build -t ${REGISTRY}/demo .

REGISTRY=default-route-openshift-image-registry.apps-crc.testing/openshift
APPLICATION_NAME=demo
GIT_HASH=$(shell git rev-parse --short=7 HEAD)

build:
	podman login --tls-verify=false -u unused -p $(oc whoami -t) ${REGISTRY}
	podman build --tag ${REGISTRY}/${APPLICATION_NAME}:${GIT_HASH} .

push:
	podman login --tls-verify=false -u unused -p $(oc whoami -t) ${REGISTRY}
	podman push --tls-verify=false ${REGISTRY}/${APPLICATION_NAME}:${GIT_HASH}

