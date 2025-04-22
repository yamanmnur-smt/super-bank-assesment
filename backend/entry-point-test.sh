#!/bin/sh

echo "execute unit test and generate cover.out ..."
go test ./... -coverprofile ./cover.out -covermode atomic -coverpkg ./...

echo "convert coverout to cover.html"
go tool cover -html cover.out -o cover.html

echo "check total coverage.."
go tool cover -func cover.out

echo "save total coverage into file txt"
go tool cover -func cover.out | grep total: | awk '{print $2}' > coverage_percentage.txt