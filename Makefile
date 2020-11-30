start:
#	HASH=$(echo test)
#	HASH=$(git log --pretty=format:'%H' -n 1)
#	git log --pretty=format:'%H' -n 1
#	@echo $(HASH)
	@echo "starting go"
#	@go run ../build/webservice.go -hash=5005 -projectURL=https://github.com/amphioxis/webservice.git #-port=${PORT} -path_1=${PATH_1} -maxReq=${MAXREQ}
#	@go run ../build/webservice.go -port=${PORT} -path_1=${PATH_1} -maxReq=${MAXREQ}
	@go run ../build/goFiles/ .

