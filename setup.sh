#!/bin/sh -eux

# install or fetch dependencies
# gcloud components install --quiet app-engine-go
go get -u github.com/golang/dep/cmd/dep
dep ensure