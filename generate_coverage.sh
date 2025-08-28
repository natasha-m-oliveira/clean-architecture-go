#!/bin/sh
PKGS=$(go list ./... | grep -v /prisma/ | grep -v /test/ | grep -v /internal/infrastructure/)
go test $PKGS -buildvcs=false -coverpkg=$(echo $PKGS | tr ' ' ',') -covermode=atomic -coverprofile=coverage.out && \
go tool cover -func=coverage.out && \
go tool cover -html=coverage.out -o coverage.html