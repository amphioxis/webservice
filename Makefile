start:
	@echo "starting go"
	@go run ./build/goFiles/ -port=8090 -path_1=helloworld -path_2=versionz -maxReq=5 -hash=8088088aa1098050ee7761d5077d027c174a6f4a -projectURL=https://github.com/amphioxis/webservice.git .

