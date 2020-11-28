#!/bin/sh

HASH=$(git log --pretty=format:'%H' -n 1)
PROJECTURL=$(git remote get-url origin)

go run webservice.go -port=${PORT} -path_1=${PATH_1} -path_2=${PATH_2} -maxReq=${MAXREQ} -hash=${HASH} -projectURL=${PROJECTURL}
